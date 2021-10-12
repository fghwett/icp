package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/fghwett/icp/api"
)

var port = flag.Int("port", 2080, "api端口")

func main() {
	flag.Parse()

	http.HandleFunc("/query", api.Query)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
