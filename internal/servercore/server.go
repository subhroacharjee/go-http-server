package servercore

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"unicode"

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

	acceptedEncoding, existsAcceptedEncoding := request.Headers["accept-encoding"]

	handlers, pathParams := h.router.GetHandlers(request.Method, request.Path)
	// fmt.Println(handlers == nil, pathParams)

	if handlers == nil {
		errorResult := httpcore.NewHttpResponseWriter()
		errorResult.SetStatus(httpcore.StatusNotFound)
		if existsAcceptedEncoding && acceptedEncoding == "gzip" {
			errorResult.SetHeader("Content-Encoding", acceptedEncoding)
		}
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

	handleEncoding(*request, &response)

	if _, err := conn.Write(response.ToResponseByte()); err != nil {
		fmt.Printf("Error writing the response %v", err)
	}
}

func handleEncoding(r httpcore.Request, w *httpcore.HttpResponseWriter) {
	fmt.Println("Called in handlers")
	accepted, exists := r.Headers["accept-encoding"]
	if !exists {
		return
	}

	encodings := strings.Split(accepted, ", ")
	canCompress := false
	for _, encoding := range encodings {
		if strings.TrimFunc(encoding, unicode.IsSpace) == "gzip" {
			canCompress = true
			break
		}
	}

	if canCompress {
		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)

		if _, err := zw.Write(w.Body); err != nil {
			return
		}

		if err := zw.Close(); err != nil {
			fmt.Println("Error >>>", err)
			return
		}

		compressedBody := buf.Bytes()
		w.SetHeader("Content-Encoding", "gzip")
		w.Write(compressedBody)

	}
}
