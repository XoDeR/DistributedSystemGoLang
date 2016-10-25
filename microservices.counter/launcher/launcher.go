package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"microservices.counter/common"
)

var addr = flag.String("addr", ":8888", "http service address")

func createItemRecord(item common.Item) {
	// POST request
	itemJson, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := strings.NewReader(fmt.Sprintf("%s", string(itemJson)))
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
	var items = make([]common.Item, itemsCount)

	items[0] = common.Item{"stick01", "bundle"}
	items[1] = common.Item{"stick02", "bundle"}
	items[2] = common.Item{"stick03", "bundle"}

	items[3] = common.Item{"key01", "bunch"}
	items[4] = common.Item{"key02", "bunch"}
	items[5] = common.Item{"key03", "bunch"}

	items[6] = common.Item{"fly01", "swarm"}
	items[7] = common.Item{"fly02", "swarm"}
	items[8] = common.Item{"fly03", "swarm"}

	items[9] = common.Item{"bee01", "hive"}
	items[10] = common.Item{"bee02", "hive"}
	items[11] = common.Item{"bee03", "hive"}

	items[12] = common.Item{"star01", "galaxy"}
	items[13] = common.Item{"star02", "galaxy"}
	items[14] = common.Item{"star03", "galaxy"}

	items[15] = common.Item{"stone01", "heap"}
	items[16] = common.Item{"stone02", "heap"}
	items[17] = common.Item{"stone03", "heap"}

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

func main() {

	fmt.Println("Test Launcher for MicroServices Counter")

	// create coordinator
	exec.Command("f:\\GoLang\\src\\microservices.counter\\coordinator\\coordinator.exe")

	// create hub
	// "rpc_addr", "127.0.0.1:7777"
	exec.Command("f:\\GoLang\\src\\microservices.counter\\hub\\hub.exe", "-rpc_addr 127.0.0.1:7777")

	createItems()

	var itemsByTenantCount = getItemsCount("Yyyyy")
	fmt.Println(itemsByTenantCount)

	// pause
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err1 := bio.ReadLine()
	fmt.Println(err1)
	fmt.Println(line)
	fmt.Println(hasMoreInLine)
}
