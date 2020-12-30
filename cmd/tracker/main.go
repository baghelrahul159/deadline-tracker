package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var isConfigTest bool

func init() {
	flag.BoolVar(&isConfigTest, "test", false, "config test")
}

func main() {
	flag.Parse()

	testConfig(isConfigTest)

	router := mux.NewRouter()
	router.HandleFunc("/health", healthCheckHandler)
	log.Fatal(http.ListenAndServe(":9000", router))
}

func testConfig(configtest bool) {
	if configtest {
		log.Println("Config Test Successfull, exiting")
		os.Exit(0)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
