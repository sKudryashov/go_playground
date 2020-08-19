package main

import (
	"fmt"
	"os"

	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
)

func main() {
	report, err := runner.Run("pb.GraphService/Put", ":50051",
		runner.WithProtoFile("graphsrv.proto", []string{"../../pkg/pb"}),
		runner.WithDataFromFile("./data.json"),
		runner.WithInsecure(true),
	)
	if err != nil {
		fmt.Println("unable to configure Runner", err.Error())
		os.Exit(1)
	}
	pp := printer.ReportPrinter{
		Out:    os.Stdout,
		Report: report,
	}
	pp.Print("pretty")
}
