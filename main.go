package main

import (
	"os"

	flag "github.com/spf13/pflag"
	xservice "github.com/xwy27/CloudGo-IO/service"
)

const (
	// PORT is the default port number
	PORT string = "9000"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	customPort := flag.StringP("port", "p", PORT, "server listening port")
	flag.Parse()
	if len(*customPort) != 0 {
		port = *customPort
	}

	server := xservice.NewServer()
	server.Run(":" + port)
}
