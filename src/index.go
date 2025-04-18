package main

import (
	"flag"

	"github.com/deanxv/CycleTLS/cycletls"
	"log"
	nhttp "net/http"
	"os"
	"runtime"
)

func setupRoutes() {
	nhttp.HandleFunc("/", cycletls.WSEndpoint)
}

func main() {
	port, exists := os.LookupEnv("WS_PORT")
	var addr *string
	if exists {
		addr = flag.String("addr", ":"+port, "http service address")
	} else {
		addr = flag.String("addr", ":9112", "http service address")
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	setupRoutes()
	log.Fatal(nhttp.ListenAndServe(*addr, nil))
}
