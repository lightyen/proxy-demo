package main

import "flag"

var port = flag.String("port", "8080", "port")

func main() {
	_ = runFiber()
}
