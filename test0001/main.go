// Client0001 project main.go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	//"io/ioutil"
	//"log"
	"net/http"
	//"strconv"
	"strings"
)

type Item struct {
	Id     string
	Tenant string
}

func addItemList(itemList []Item) {

	// POST request
	jsonList, err := json.Marshal(itemList)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := strings.NewReader(fmt.Sprintf("%s", string(jsonList)))
	fmt.Println(reader)
	request, err := http.NewRequest("POST", "http://localhost:7777/api", reader)
	fmt.Println(err)
	client := &http.Client{}
	resp, err := client.Do(request)
	fmt.Println(err)
	fmt.Println(resp)
}

func createItems() {

	var itemsCount int = 18
	var items = make([]string, itemsCount)

	items[0] = Item{"stick01", "bundle"}
	items[1] = Item{"stick02", "bundle"}
	items[2] = Item{"stick03", "bundle"}

	items[3] = Item{"key01", "bunch"}
	items[4] = Item{"key02", "bunch"}
	items[5] = Item{"key03", "bunch"}

	items[6] = Item{"fly01", "swarm"}
	items[7] = Item{"fly02", "swarm"}
	items[8] = Item{"fly03", "swarm"}

	items[9] = Item{"bee01", "hive"}
	items[10] = Item{"bee02", "hive"}
	items[11] = Item{"bee03", "hive"}

	items[12] = Item{"star01", "galaxy"}
	items[13] = Item{"star02", "galaxy"}
	items[14] = Item{"star03", "galaxy"}

	items[15] = Item{"stone01", "heap"}
	items[16] = Item{"stone02", "heap"}
	items[17] = Item{"stone03", "heap"}

	addItemList(items)
}

func queryItemList(tenant string, uint32 count) itemList []Item {
	// GET request
	reader := strings.NewReader(`{"body":123}`)
	request, err := http.NewRequest("GET", "http://localhost:7777/api", reader)
	fmt.Println(err)
	client := &http.Client{}
	resp, err := client.Do(request)
	fmt.Println(err)
	fmt.Println(resp)
}

func main() {
	fmt.Println("Test Client for Coordinator Server")

	createItems()

	itemList []Item

	// pause
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	fmt.Println(err)
	fmt.Println(line)
	fmt.Println(hasMoreInLine)
}
