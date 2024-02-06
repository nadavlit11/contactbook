package main

import (
	"flag"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
)

func main() {
	InitControllers(port)
}
