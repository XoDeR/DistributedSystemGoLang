package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/golang/protobuf/proto"
	"microservices.counter/common"
)

var WorkerId uint32
var HubAddr string
var ItemsMap map[uint32][]uint32 // tenantId : itemId slice

func testSendResponse(hubAddr string) {
	var tenantId uint32 = 1

	data, err := createMessage(tenantId)
	checkError(err)
	sendDataToDest(data, &hubAddr)
}

func main() {
	WorkerId = uint32(*flag.Uint("id", 0, "Worker Id"))
	HubAddr = *flag.String("hub_addr", "127.0.0.1:2110", "Hub socket address")
	flag.Parse()

	ItemsMap = make(map[uint32][]uint32)

	testSendResponse(HubAddr)
}

func createMessage(tenantId uint32) ([]byte, error) {
	ProtoMessage := new(common.WorkerGetItemListResponse)
	ProtoMessage.WorkerId = WorkerId

	itemIdList := ItemsMap[tenantId]

	for _, v := range itemIdList {
		ProtoMessage.Items = append(ProtoMessage.Items, v)
	}

	return proto.Marshal(ProtoMessage)
}

func sendDataToDest(data []byte, dst *string) {
	conn, err := net.Dial("tcp", *dst)
	checkError(err)
	n, err := conn.Write(data)
	checkError(err)
	fmt.Println("Sent " + strconv.Itoa(n) + " bytes")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
