package main

import (
	"shorty/internal/mongodb"
	"shorty/internal/server"
)

func init() {
	db := mongodb.NewInstance()
	db.GetConnection()
}

func main() {
	server.Run()
}
