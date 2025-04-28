package httpcore

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/internal/common"
)

type HandlerFunc func(w Request, r *HttpResponseWriter)

type Request struct {
	Method     common.Method
	Path       string
	Headers    HeaderMap
	Body       []byte
	Query      map[string]string
	PathParams map[string]string
}

func ParseRequest(reader *bufio.Reader) (*Request, error) {
	// Read request line
	requestLineBytes, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read request line: %w", err)
	}
	requestLineBytes = bytes.TrimRight(requestLineBytes, "\r\n")

	requestLineItems := bytes.Split(requestLineBytes, []byte(" "))
	if len(requestLineItems) != 3 {
		return nil, fmt.Errorf("invalid request line structure")
	}

	method := common.Method(string(requestLineItems[0]))
	requestPath, queryMap := getQueryMapFromPath(string(requestLineItems[1]))
	headerMap := make(HeaderMap)

	// Read headers
	for {
		headerLineBytes, err := reader.ReadBytes('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read header line: %w", err)
		}
		headerLineBytes = bytes.TrimRight(headerLineBytes, "\r\n")

		if len(headerLineBytes) == 0 {
			// Empty line signals end of headers
			break
		}

		key, value, found := bytes.Cut(headerLineBytes, []byte(": "))
		if found {
			headerMap[strings.ToLower(string(key))] = string(value)
		}
	}

	// Read body if Content-Length exists
	var body []byte
	if contentLengthStr, ok := headerMap["content-length"]; ok {
		contentLength, err := strconv.Atoi(contentLengthStr)
		if err != nil {
			return nil, fmt.Errorf("invalid Content-Length: %w", err)
		}

		body = make([]byte, contentLength)
		_, err = io.ReadFull(reader, body)
		if err != nil {
			return nil, fmt.Errorf("failed to read body: %w", err)
		}
	}

	return &Request{
		Method:  method,
		Path:    requestPath,
		Headers: headerMap,
		Body:    body,
		Query:   queryMap,
	}, nil
}

func getQueryMapFromPath(urlPath string) (string, map[string]string) {
	queryMap := make(map[string]string)
	queryLineIdx := strings.Index(urlPath, "?")

	if queryLineIdx == -1 {
		return urlPath, nil
	}

	path, query, _ := strings.Cut(urlPath, "?")
	for _, substr := range strings.Split(query, "&") {
		key, value, found := strings.Cut(substr, "=")
		if !found {
			queryMap[substr] = ""
		} else {
			queryMap[key] = value
		}

	}

	return path, queryMap
}
