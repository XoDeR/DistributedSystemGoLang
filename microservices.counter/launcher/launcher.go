package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"

	"microservices.counter/common/item"
	"microservices.counter/hub"
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

func createItemRecord(item item.Item) {
	// POST request
	jsonList, err := json.Marshal(itemList)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := strings.NewReader(fmt.Sprintf("%s", string(jsonList)))
	fmt.Println(reader)
	request, err := http.NewRequest("POST", "http://localhost:8888/items", reader)
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

	// TODO
	// for 0..size createItemRecord
}

func getItemsCount(tenant string) uint32 {
	var result uint32 = 0

	// GET request
	reader := strings.NewReader(`{"body":123}`)
	request, err := http.NewRequest("GET", "http://localhost:8888/items/{tenantId}/count", reader)
	fmt.Println(err)
	client := &http.Client{}
	resp, err := client.Do(request)
	fmt.Println(err)
	fmt.Println(resp)

	return result
}

var nextUniqueId uint32 = 0

func getNextUniqueId() uint32 {
	var currentUniqueId uint32 = nextUniqueId
	nextUniqueId = nextUniqueId + 1
	return currentUniqueId
}

func Handler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	webpage, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(response, fmt.Sprintf("home.html file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage))
}

func APIHandler(response http.ResponseWriter, request *http.Request) {

	//set mime type to JSON
	response.Header().Set("Content-type", "application/json")

	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}

	// fixed size 1000
	var result = make([]string, 1000)

	switch request.Method {
	case "GET":

		// query items
		tenant := request.PostFormValue("tenant")

		var itemList []Item = getItemListByTenant(tenant, 10)

		//	b, err := json.Marshal(itemList)
		//	if err != nil {
		//	fmt.Println(err)
		//return
		//	}
		//result[i] = fmt.Sprintf("%s", string(b))

		result = result[:1]

	case "POST":
		id := request.PostFormValue("id")

		// insert new entry with id and tenant
		var res bool = true
		if res == true {
			result[0] = "true"
		}
		result = result[:1]

	case "PUT":
		id := request.PostFormValue("id")
		tenant := request.PostFormValue("tenant")

		// update existing entry
		var res bool
		if res == true {
			result[0] = "true"
		}
		result = result[:1]
	case "DELETE":
		id := strings.Replace(request.URL.Path, "/api/", "", -1)

		// delete item with id
		var res bool
		if res == true {
			result[0] = "true"
		}
		result = result[:1]

	default:
	}

	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the text diagnostics to the client.
	fmt.Fprintf(response, "%v", string(json))
}

func main() {

	fmt.Println("Test Launcher for MicroServices Counter")

	// create hub (coordinator)

	// create 5 counters

	createItems()

	var itemsByTenantCount = getItemsCount("Yyyyy")
	fmt.Println(itemsByTenantCount)

	port := 7777
	var err string
	portstring := strconv.Itoa(port)

	mux := http.NewServeMux()
	mux.Handle("/api/", http.HandlerFunc(APIHandler))
	mux.Handle("/", http.HandlerFunc(Handler))

	log.Print("Listening on port " + portstring + " ... ")
	errs := http.ListenAndServe(":"+portstring, mux)
	if errs != nil {
		log.Fatal("ListenAndServe error: ", err)
	}

	// pause
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	fmt.Println(err)
	fmt.Println(line)
	fmt.Println(hasMoreInLine)
}
