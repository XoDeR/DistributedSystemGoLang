package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/golang/protobuf/proto"
	"microservices.counter/common"
)

type Headers []string

const CLIENT_ID = 2

func main() {
	dest := flag.String("d", "127.0.0.1:2110", "Enter the destination socket address")
	flag.Parse()

	data, err := createMessage(tenantId uint32)
	checkError(err)
	sendDataToDest(data, dest)
}

func createMessage(tenantId uint32) ([]byte, error) 
{
	ProtoMessage := new(ProtobufTest.TestMessage)
	ProtoMessage.ClientId = proto.Int32(CLIENT_ID)

	//loop through the records
	for {
		record, err := csvreader.Read()
		if err != io.EOF {
			checkError(err)
		} else {

			break
		}
		//Populate items
		testMessageItem := new(ProtobufTest.TestMessage_MsgItem)
		itemid, err := strconv.Atoi(record[ITEMIDINDEX])
		checkError(err)
		testMessageItem.Id = proto.Int32(int32(itemid))
		testMessageItem.ItemName = &record[ITEMNAMEINDEX]
		itemvalue, err := strconv.Atoi(record[ITEMVALUEINDEX])
		checkError(err)
		testMessageItem.ItemValue = proto.Int32(int32(itemvalue))
		itemtype, err := strconv.Atoi(record[ITEMTYPEINDEX])
		checkError(err)
		iType := ProtobufTest.TestMessage_ItemType(itemtype)
		testMessageItem.ItemType = &iType

		ProtoMessage.Messageitems = append(ProtoMessage.Messageitems, testMessageItem)

		fmt.Println(record)
	}

	//fmt.Println(ProtoMessage.Messageitems)
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

func (h Headers) getHeaderIndex(headername string) int {
	if len(headername) >= 2 {
		for index, s := range h {
			if s == headername {
				return index
			}
		}
	}
	return -1
}
