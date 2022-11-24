package main

import (
	"fmt"
	"shorty/internal/mongodb"
)

func init() {

}

func main() {

	dbIns := mongodb.NewInstance()

	dbIns.GetConnection()

	fmt.Println("hello world!")
}
