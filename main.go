package main

import (
	"github.com/ishanshre/gotodo/pkg/middleware"
)

func main() {
	server := middleware.NewAPIServer(":8000")
	server.Run()
}
