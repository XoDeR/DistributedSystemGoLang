package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/websocket"

	"microservices.counter/testCommon"
)

var addr = flag.String("addr", ":8080", "http service address")
var homeTemplate = template.Must(template.ParseFiles("home.html"))

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *testCommon.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// creates new client
	client := &testCommon.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client
	go client.WritePump()
	client.ReadPump()
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}

func testSendMessage() {
	// GET request
	reader := strings.NewReader(`{"body":123}`)
	request, err := http.NewRequest("GET", "http://localhost:8080/ws", reader)
	fmt.Println(err)
	client := &http.Client{}
	resp, err := client.Do(request)
	fmt.Println(err)
	fmt.Println(resp)
}

func main() {
	flag.Parse()
	hub := testCommon.NewHub()
	go hub.Run()

	http.HandleFunc("/", serveHome)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	testSendMessage()
}
