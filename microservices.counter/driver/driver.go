package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"microservices.counter/common/item"
)

var addr = flag.String("addr", ":8888", "http service address")
var homeTemplate = template.Must(template.ParseFiles("home.html"))

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

func addItemList(itemList []item.Item) {

	// POST request
	jsonList, err := json.Marshal(itemList)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := strings.NewReader(fmt.Sprintf("%s", string(jsonList)))
	fmt.Println(reader)
	request, err := http.NewRequest("POST", "http://localhost:8888/api", reader)
	fmt.Println(err)
	client := &http.Client{}
	resp, err := client.Do(request)
	fmt.Println(err)
	fmt.Println(resp)
}

func createItems() {

	var itemsCount int = 18
	var items = make([]item.Item, itemsCount)

	items[0] = item.Item{"stick01", "bundle"}
	items[1] = item.Item{"stick02", "bundle"}
	items[2] = item.Item{"stick03", "bundle"}

	items[3] = item.Item{"key01", "bunch"}
	items[4] = item.Item{"key02", "bunch"}
	items[5] = item.Item{"key03", "bunch"}

	items[6] = item.Item{"fly01", "swarm"}
	items[7] = item.Item{"fly02", "swarm"}
	items[8] = item.Item{"fly03", "swarm"}

	items[9] = item.Item{"bee01", "hive"}
	items[10] = item.Item{"bee02", "hive"}
	items[11] = item.Item{"bee03", "hive"}

	items[12] = item.Item{"star01", "galaxy"}
	items[13] = item.Item{"star02", "galaxy"}
	items[14] = item.Item{"star03", "galaxy"}

	items[15] = item.Item{"stone01", "heap"}
	items[16] = item.Item{"stone02", "heap"}
	items[17] = item.Item{"stone03", "heap"}

	addItemList(items)
}

func getItemsCount(tenant string) uint32 {
	var result uint32 = 0

	// GET request
	reader := strings.NewReader(`{"body":123}`)
	request, err := http.NewRequest("GET", "http://localhost:8888//cats", reader)
	fmt.Println(err)
	client := &http.Client{}
	resp, err := client.Do(request)
	fmt.Println(err)
	fmt.Println(resp)

	return result
}

func main() {

	fmt.Println("Test Driver for MicroServices Counter Hub (Main Server)")

	//flag.Parse()
	//hub := newHub()
	//go hub.run()

	//http.HandleFunc("/", serveHome)
	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//serveWs(hub, w, r)
	//})
	//err := http.ListenAndServe(*addr, nil)
	//if err != nil {
	//log.Fatal("ListenAndServe: ", err)
	//}

	createItems()

	var itemsByTenantCount = getItemsCount("Yyyyy")
	fmt.Println(itemsByTenantCount)

	// pause
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	fmt.Println(err)
	fmt.Println(line)
	fmt.Println(hasMoreInLine)
}
