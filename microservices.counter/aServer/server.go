//package server
package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/websocket"
)

var SERVER_PORT = "8000"

var CLIENT_CONFIG_DICT = map[string]string{
	"C1": "aa",
	"C2": "bb",
	"C3": "cc",
	"C4": "aa",
}

var CLIENT_RESP string

//this to be migrated to redis
var CMAP = map[string][][]byte{}

func readconfig(ws *websocket.Conn) {
	log.Print("reading config..")
	var incoming []byte
	if err := websocket.Message.Receive(ws, &incoming); err != nil {
		log.Print("client stream ended..") //break
	}
	incoming_str := strings.Trim(string(incoming), string(10))
	log.Print("data received from client: ", string(incoming_str))

	cumu_v := []byte{}
	for _, v := range CMAP[string(incoming_str)] {
		if string(v) != "none" {
			cumu_v = append(append(cumu_v, v...), []byte("##")...)
		}
	}
	io.Copy(ws, bytes.NewBuffer(cumu_v))
	//clean any re-init
	delete(CMAP, string(incoming_str))
	CMAP[string(incoming_str)] = append(CMAP[string(incoming_str)], []byte("none"))
	log.Print("config sent..")
}

func storeconfig(ws *websocket.Conn) {
	log.Print("processing client's response..")

	var incoming []byte
	if err := websocket.Message.Receive(ws, &incoming); err != nil {
		log.Print("client stream ended..") //break
	}
	res := strings.Split(string(incoming), "##")
	res = res[:len(res)-1]
	for _, resp := range res {
		CLIENT_RESP = strconv.Itoa(rand.Int()) + resp
	}
	io.Copy(ws, bytes.NewBuffer([]byte("response stored at server..")))
}

func heartbeat(ws *websocket.Conn) {

	var incoming []byte
	if err := websocket.Message.Receive(ws, &incoming); err != nil {
		log.Print("client stream ended..") //break
	}
	//incoming_str at the moment is hostname and is the key in kv store
	incoming_str := strings.Trim(string(incoming), string(10))
	log.Print("heartbeat from client: ", incoming_str)

	_, present := CMAP[incoming_str]
	if !present {

		//client not present - initialize with none - status code #001
		CMAP[incoming_str] = append(CMAP[incoming_str], []byte("none"))
		io.Copy(ws, bytes.NewBuffer([]byte("001")))
	} else if present {

		//client present in kv store
		if len(CMAP[incoming_str]) == 1 { //string(val) == "none" { //client initialized but no jobs in queue - status code #002
			for _, conf := range CLIENT_CONFIG_DICT {
				c, _ := ioutil.ReadFile(conf)
				CMAP[incoming_str] = append(CMAP[incoming_str], c)
			}
			io.Copy(ws, bytes.NewBuffer([]byte("002")))
		} else if len(CMAP[incoming_str]) > 1 { //string(val) != "none" { //client initialized and jobs in queue - status code #003

			io.Copy(ws, bytes.NewBuffer([]byte("003")))
		}
	} else {
		log.Print("nothing found.. ")
	}
}

func main() {
	http.Handle("/readconfig", websocket.Handler(readconfig))
	http.Handle("/storeconfig", websocket.Handler(storeconfig))
	http.Handle("/", websocket.Handler(heartbeat))
	log.Print("server running on port " + SERVER_PORT + "..")
	err := http.ListenAndServe(":"+SERVER_PORT, nil)
	if err != nil {
		panic("unable to serve.. " + err.Error())
	}

}
