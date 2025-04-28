package servercore

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/codecrafters-io/http-server-starter-go/internal/httpcore"
	"github.com/codecrafters-io/http-server-starter-go/internal/router"
)

type HttpServer struct {
	router router.ReadOnlyRouter
}

func NewHttpServer(appRouter router.IRouter) HttpServer {
	router := router.Router{}
	router.CopyPath(appRouter)
	return HttpServer{
		router: &router,
	}
}

func (h *HttpServer) Listen(port uint) error {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return err
	}

	connChan := make(chan net.Conn, 3) // async data passing
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		h.listenOverLoop(l, connChan)
	}()

	var wg sync.WaitGroup

	go func() {
		for conn := range connChan {
			wg.Add(1)
			go func() {
				h.handleRequests(conn)
				wg.Done()
			}()
		}
	}()

	<-stop
	fmt.Println("Shutting down gracefully")
	l.Close()
	close(connChan)
	wg.Wait()
	fmt.Println("graceful Shutdown complete")
	return nil
}

func (h *HttpServer) listenOverLoop(ln net.Listener, connChan chan<- net.Conn) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Stopped listening")
			return
		}

		connChan <- conn
	}
}

func (h *HttpServer) handleRequests(conn net.Conn) {
	defer conn.Close()

	request, err := httpcore.ParseRequest(bufio.NewReader(conn))
	if err != nil {
		errorResult := httpcore.NewHttpResponseWriter()
		errorResult.SetStatus(httpcore.StatusBadRequest)
		if _, err := conn.Write(errorResult.ToResponseByte()); err != nil {
			fmt.Printf("Error writing to the connection %v", err)
		}
		return

	}

	handlers, pathParams := h.router.GetHandlers(request.Method, request.Path)
	fmt.Println(handlers == nil, pathParams)

	if handlers == nil {
		errorResult := httpcore.NewHttpResponseWriter()
		errorResult.SetStatus(httpcore.StatusNotFound)
		if _, err := conn.Write(errorResult.ToResponseByte()); err != nil {
			fmt.Printf("Error writing to the connection %v", err)
		}
		return
	}

	request.PathParams = pathParams
	response := httpcore.NewHttpResponseWriter()

	for _, handler := range handlers {
		handler(*request, &response)
		if response.IsReadyForResponse() {
			break
		}
	}

	if !response.IsReadyForResponse() || !response.IsStatusSet() {
		response.SetStatus(httpcore.StatusOK)
	}

	if _, err := conn.Write(response.ToResponseByte()); err != nil {
		fmt.Printf("Error writing the response %v", err)
	}
}
