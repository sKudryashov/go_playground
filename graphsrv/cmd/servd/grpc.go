package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/labstack/gommon/log"
	handler "github.com/sKudryashov/graphsrv/internal/handler/grpc"
	"github.com/sKudryashov/graphsrv/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	var addr string
	flag.StringVar(&addr, "port", "50051", "service address endpoint")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))
	if err != nil {
		fmt.Println("unable to start the server", err.Error())
		os.Exit(1)
	}
	lgr := log.New("graph_server")
	h := handler.NewHandler(lgr)
	srv := grpc.NewServer()
	pb.RegisterGraphServiceServer(srv, h)
	lgr.Infof("Server started on port %s", addr)
	srv.Serve(listener)
}
