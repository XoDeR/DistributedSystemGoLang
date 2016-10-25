package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/NYTimes/gizmo/server"
	"microservices.counter/service"
)

var (
	// RPC server (hub) port 7777
	serverAddr = flag.String("server_addr", "127.0.0.1:7777", "The server address in the format of host:port")

	mostPopularFlag = flag.Bool("most-popular", true, "flag make the GetMostPopular call")
)

func main() {
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

	nytClient := service.NewNYTProxyServiceClient(conn)

	if *mostPopularFlag {
		mostPop, err := nytClient.GetMostPopular(context.Background(), &service.MostPopularRequest{
			ResourceType:   "mostviewed",
			Section:        "all-sections",
			TimePeriodDays: uint32(1),
		})
		if err != nil {
			//log.Fatal("get most popular list error: ", err)
		}

		fmt.Println("Most Popular Results:")
		out, _ := json.MarshalIndent(mostPop, "", "    ")
		fmt.Fprint(os.Stdout, string(out))
		fmt.Println("")
	}

	// pause
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	fmt.Println(err)
	fmt.Println(line)
	fmt.Println(hasMoreInLine)
}
