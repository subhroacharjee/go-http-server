package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/internal/application"
	"github.com/codecrafters-io/http-server-starter-go/internal/router"
	"github.com/codecrafters-io/http-server-starter-go/internal/servercore"
)

func main() {
	directory := flag.String("directory", "/tmp", "Directory where files are stored")
	flag.Parse()
	router := router.NewRouter()
	application.RegisterControllers(router, directory)
	httpServer := servercore.NewHttpServer(router)
	if err := httpServer.Listen(4221); err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		os.Exit(1)
	}
}
