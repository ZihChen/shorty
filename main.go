package main

import (
	"shorty/internal/mongodb"
	"shorty/internal/server"
)

func init() {
	mongodb.NewInstance()
}

func main() {
	server.Run()
}
