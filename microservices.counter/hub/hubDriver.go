package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"

	"net"
	"strconv"

	"github.com/golang/protobuf/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/NYTimes/gizmo/server"
	"microservices.counter/common"
)

var (
	// RPC server (coordinator-hub) port 7777
	serverAddr = flag.String("rpc_addr", "127.0.0.1:7777", "The server address in the format of host:port")

	hubClient *common.NYTProxyServiceClient
)

func createWorker(index uint32, hubIp *string, hubPort *string) {
	var addr string = *hubIp + ":" + *hubPort
	var param1 string = "-id=" + strconv.Itoa(int(index))
	var param2 string = "-hub_addr=" + addr
	exec.Command("f:\\GoLang\\src\\microservices.counter\\worker\\worker.exe", param1, param2)
}

func testAddNewItem() {

}

func main() {
	// Connect to Coordinator
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			server.Log.Warn("unable to close gRPC connection: ", err)
		}
	}()

	hubClient := common.NewNYTProxyServiceClient(conn)

	// testAddNewItem

	addNewItem, err := hubClient.AddNewItemWithTenant(context.Background(), &common.AddNewItemWithTenantRequest{
		ItemId:   uint32(1),
		TenantId: uint32(1),
	})
	if err != nil {
		//log.Fatal("get most popular list error: ", err)
	}

	fmt.Println("Add New Item Results:")
	out, _ := json.MarshalIndent(addNewItem, "", "    ")
	fmt.Fprint(os.Stdout, string(out))
	fmt.Println("")

	// create 5 workers
	var hubIp string = "127.0.0.1"
	var workersCount int = 5

	var hubPortList []string = make([]string, workersCount)
	hubPortList[0] = "7771"
	hubPortList[1] = "7772"
	hubPortList[2] = "7773"
	hubPortList[3] = "7774"
	hubPortList[4] = "7775"

	for i := 1; i <= workersCount; i++ {
		createWorker(uint32(i), &hubIp, &hubPortList[int(i-1)])
	}

	fmt.Printf("Started ProtoBuf Server")
	c := make(chan *common.WorkerGetItemListResponse)
	go func() {
		for {
			message := <-c
			mergeValues(message)

		}
	}()

	//Listen to the TCP port
	listener, err := net.Listen("tcp", "127.0.0.1:2110")
	checkError(err)
	for {
		if conn, err := listener.Accept(); err == nil {
			//If err is nil then that means that data is available for us so we take up this data and pass it to a new goroutine
			go handleProtoClient(conn, c)
		} else {
			continue
		}
	}

	// pause
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	fmt.Println(err)
	fmt.Println(line)
	fmt.Println(hasMoreInLine)
}

func handleProtoClient(conn net.Conn, c chan *common.WorkerGetItemListResponse) {
	fmt.Println("Connection established")
	//Close the connection when the function exits
	defer conn.Close()
	//Create a data buffer of type byte slice with capacity of 4096
	data := make([]byte, 4096)
	//Read the data waiting on the connection and put it in the data buffer
	n, err := conn.Read(data)
	checkError(err)
	fmt.Println("Decoding Protobuf message")
	//Create an struct pointer of type WorkerResponse struct
	protodata := new(common.WorkerGetItemListResponse)
	//Convert all the data retrieved into the ProtobufTest.TestMessage struct type
	err = proto.Unmarshal(data[0:n], protodata)
	checkError(err)
	//Push the protobuf message into a channel
	c <- protodata
}

func mergeValues(datatowrite *common.WorkerGetItemListResponse) {

	//Retreive client information from the protobuf message
	WorkerId := strconv.Itoa(int(datatowrite.WorkerId))

	// retrieve the message items list
	items := datatowrite.Items

	fmt.Println("Merging results from worker")
	//Go through the list of message items
	for _, item := range items {
		record := []string{WorkerId, strconv.Itoa(int(item))}
		fmt.Println(record)
	}
	fmt.Println("Finished merging results from worker")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
