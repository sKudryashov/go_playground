package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/sKudryashov/graphsrv/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	var method, port, body string
	flag.StringVar(&method, "method", "put/get/delete", "call method")
	flag.StringVar(&body, "body", "", "call body which conforms to appropriate method .proto specification")
	flag.StringVar(&port, "port", "50051", "server grpc port")
	flag.Parse()
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("unable to reach the server %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	cli := pb.NewGraphServiceClient(conn)
	ctx := context.Background()
	switch method {
	case "put":
		graph := &pb.Graph{}
		if err := json.Unmarshal([]byte(body), graph); err != nil {
			fmt.Printf("unable to parse req body %v", err)
		}
		rsp, err := cli.Put(ctx, graph)
		if err != nil {
			fmt.Printf("got response error %v", err)
			os.Exit(1)
		}
		printResponse(rsp, "added")
	case "get":
		graph := &pb.Search{}
		if err := json.Unmarshal([]byte(body), graph); err != nil {
			fmt.Printf("unable to parse req body %v", err)
		}
		rsp, err := cli.GetLast(ctx, graph)
		if err != nil {
			fmt.Printf("got response error %v", err)
			os.Exit(1)
		}
		printResponse(rsp, "got last record")
	case "delete":
		graph := &pb.Graph{}
		if err := json.Unmarshal([]byte(body), graph); err != nil {
			fmt.Printf("unable to parse req body %v", err)
		}
		rsp, err := cli.Delete(ctx, graph)
		if err != nil {
			fmt.Printf("got response error %v", err)
			os.Exit(1)
		}
		printResponse(rsp, "deleted")
	default:
		fmt.Println("unregistered action")
		os.Exit(1)
	}
}

func printResponse(rsp interface{}, action string) {
	rspByte, err := json.Marshal(rsp)
	if err != nil {
		fmt.Printf("got marshal error %v", err)
		os.Exit(1)
	}
	fmt.Println("action: ", action)
	fmt.Println("server response: ")
	fmt.Println(string(rspByte))
	os.Exit(0)
}
