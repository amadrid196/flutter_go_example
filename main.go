package main

import (
	"github.com/Goinitial/server"
)

func main() {

	server := server.New(":8080")
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
