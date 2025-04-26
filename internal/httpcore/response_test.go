package httpcore_test

import (
	"bytes"
	"testing"

	"github.com/codecrafters-io/http-server-starter-go/internal/httpcore"
)

func TestResponse(t *testing.T) {
	testCases := []struct {
		Name     string
		Headers  map[string]string
		Status   httpcore.HttpStatus
		Body     []byte
		Expected []byte
	}{
		{
			Name:     "Empty 200 response",
			Status:   httpcore.StatusAccepted,
			Expected: []byte("HTTP/1.1 200 OK\r\n\r\n"),
		},
		{
			Name: "Response with header",
			Headers: map[string]string{
				"Content-Type": "text/plain",
				"Server":       "Go-server",
			},
			Status:   httpcore.StatusAccepted,
			Expected: []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nServer: Go-server\r\n\r\n"),
		},
		{
			Name: "Response with Header and body",
			Headers: map[string]string{
				"Server":       "Go-server",
				"Content-Type": "text/plain",
			},
			Body:     []byte("Hello world"),
			Status:   httpcore.StatusAccepted,
			Expected: []byte("HTTP/1.1 200 OK\r\nServer: Go-server\r\nContent-Type: text/plain\r\n\r\nHello world"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc := tc
			writer := httpcore.NewHttpResponseWriter()
			if tc.Headers != nil {
				for key, value := range tc.Headers {
					writer.SetHeader(key, value)
				}
			}
			if tc.Body != nil {
				writer.Write(tc.Body)
			}

			if !bytes.Equal(tc.Expected, writer.ToResponseByte()) {
				t.Errorf("Actual: %q", writer.ToResponseByte())
				t.Fail()
			}
		})
	}
}
