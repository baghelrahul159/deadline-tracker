package main

import (
	"flag"
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
	router.HandleFunc("/", testHandler)
	log.Fatal(http.ListenAndServe(":9000", router))
}

func testConfig(configtest bool) {
	if configtest {
		log.Println("Config Test Successfull, exiting")
		os.Exit(0)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test!\n"))
}
