// REST APIs with Go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Item struct {
	Id     string
	Tenant string
}

type Node struct {
}

type Coordinator struct {
	nodesCount uint32
	nodes      []Node
}

func createCoordinator() Coordinator {
	var nodesCount uint32 = 6
	var nodes = make([]Node, nodesCount)
	var coordinator Coordinator // = {nodesCount, nodes}
	return coordinator
}

var coordinator Coordinator = createCoordinator()

func getItemListByTenant(tenant string, count uint32) []Item {
	// query each node, stop querying if [count] is already filled
	return nil
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

		/*


			b, err := json.Marshal(itemList)
				if err != nil {
					fmt.Println(err)
					return
				}
				result[i] = fmt.Sprintf("%s", string(b))
		*/

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
}
