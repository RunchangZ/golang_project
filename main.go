package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/RunchangZ/golang_project/api"
	"os"
)



func main(){

	listenAddr := flag.String("listenAddr", ":8000", "the server address")
	flag.Parse()

	//Call server.go to run it at port :8000 
	server := api.NewServer(*listenAddr)

	//get the filename from the env 
	filename := os.Getenv("INPUT_FILENAME")


	// Default filename 
	if filename == "" {
		filename = "M&M.json"
	}

	log.Printf("Input filename: %s", filename)

	filepath := "example/" + filename

	fmt.Println("Server is running on: ", *listenAddr)
	log.Fatal(server.Start(filepath))

}


