package main

import (
	"github.com/s21toolkit/s21adapter/internal/server"
)

func main() {
	root := server.InitServer()

	server.Serve(root, ":1323")
}
