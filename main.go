package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Pinkal777/vmtsp/mtsp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide problem file Path as commandline arguments.")
	}
	problemFilePath := os.Args[1]
	problems, err := mtsp.LoadProblems(problemFilePath)
	if err != nil {
		log.Fatal("Not able to read loads/problems from file", err)
	}
	driversRoute := mtsp.AssignLoads(&problems)
	for _, dr := range driversRoute {
		fmt.Printf("[%s]\n", strings.Join(dr.RouteLoads, ","))
	}
}
