package main

import (
	version "cicd/verison"
	"fmt"
	"log"
	"net/http"
)

var logger = log.Default()

func hello(w http.ResponseWriter, req *http.Request) {
	logger.Printf("get request from host %s\n", req.Host)
	fmt.Fprintf(w, "hello, this is version: %s\n", version.Version)
}

func main() {
	logger.Printf("success deploy %s\n", version.Version)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", nil)
}
