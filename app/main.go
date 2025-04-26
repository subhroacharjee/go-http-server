package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/internal/servercore"
)

func main() {
	httpServer := servercore.NewHttpServer()
	if err := httpServer.Listen(4221); err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		os.Exit(1)
	}
}
