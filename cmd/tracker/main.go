package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var isConfigTest = flag.Bool("test", false, "config test")

	flag.Parse()

	testConfig(*isConfigTest)
}

func testConfig(configtest bool) {
	if configtest {
		log.Println("Config Test Successfull, exiting")
		os.Exit(0)
	}
}
