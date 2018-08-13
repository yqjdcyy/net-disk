package main

import (
	"config"
	"flag"
	"server"
)

func main() {

	flag.Parse()

	// log.*
	config.Init()

	// server
	server.Start()
}
