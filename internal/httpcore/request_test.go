package httpcore_test

import (
	"bytes"
	"testing"

	"github.com/codecrafters-io/http-server-starter-go/internal/httpcore"
)

func TestRequest(t *testing.T) {
	testCases := []struct {
		Name        string
		RawRequest  []byte
		Path        string
		Method      httpcore.Method
		QueryMap    map[string]string
		Headers     map[string]string
		Body        []byte
		ExpectError bool
	}{
		{
			Name:        "Should return error when malformed request is incoming",
			RawRequest:  []byte("GET / HTTP/1.1\r\n"),
			ExpectError: true,
		},
		{
			Name:        "Should return error when malformed request line is incoming",
			RawRequest:  []byte("GET /HTTP/1.1\r\n\r\n"),
			ExpectError: true,
		},
		{
			Name:        "Empty request with only request line",
			RawRequest:  []byte("GET / HTTP/1.1\r\n\r\n"),
			Path:        "/",
			Method:      httpcore.GET,
			ExpectError: false,
		},
		{
			Name:       "Empty request with only request line and query",
			RawRequest: []byte("GET /some-query?q=1&y=2 HTTP/1.1\r\n\r\n"),
			Path:       "/some-query",
			Method:     httpcore.GET,
			QueryMap: map[string]string{
				"q": "1",
				"y": "2",
			},
			ExpectError: false,
		},
		{
			Name:       "Request with headers",
			RawRequest: []byte("GET / HTTP/1.1\r\nContent-Type: text/plain\r\nServer: Go-server\r\n\r\n"),
			Path:       "/",
			Method:     httpcore.GET,
			Headers: map[string]string{
				"Content-Type": "text/plain",
				"Server":       "Go-server",
			},
			ExpectError: false,
		},
		{
			Name:       "Request with headers and body",
			RawRequest: []byte("GET / HTTP/1.1\r\nContent-Type: text/plain\r\nServer: Go-server\r\n\r\nHello World!"),
			Path:       "/",
			Method:     httpcore.GET,
			Headers: map[string]string{
				"Content-Type": "text/plain",
				"Server":       "Go-server",
			},
			Body:        []byte("Hello World!"),
			ExpectError: false,
		},
		{
			Name:       "Request with headers, query and body",
			RawRequest: []byte("GET /path?long=10&lat=20.3 HTTP/1.1\r\nContent-Type: text/plain\r\nServer: Go-server\r\n\r\nHello World!"),
			Path:       "/path",
			Method:     httpcore.GET,
			Headers: map[string]string{
				"Content-Type": "text/plain",
				"Server":       "Go-server",
			},
			Body: []byte("Hello World!"),
			QueryMap: map[string]string{
				"long": "10",
				"lat":  "20.3",
			},
			ExpectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc := tc
			response, err := httpcore.ParseRequest(tc.RawRequest)
			if tc.ExpectError && err == nil {
				t.Errorf("[ %s ]Was expecting error but no error was returned", tc.Name)
				t.Fail()
			} else if !tc.ExpectError && err != nil {
				t.Errorf("[ %s ]Was not expecting error but error (%v) was returned", tc.Name, err)
				t.Fail()
			} else if !tc.ExpectError && err == nil {
				if tc.Body != nil && !bytes.Equal(tc.Body, response.Body) {
					t.Errorf("[ %s ]actual body and expected body are separate", tc.Name)
					t.Fail()
				}
				if tc.Headers != nil && !mapsAreEqual(tc.Headers, response.Headers) {
					t.Errorf("[ %s ]actual headers and expected headers are separate", tc.Name)
					t.Fail()
				}
				if tc.QueryMap != nil && !mapsAreEqual(tc.QueryMap, response.Query) {
					t.Errorf("[ %s ]actual query and expected query are separate", tc.Name)
					t.Fail()
				}
				if tc.Method != response.Method {
					t.Errorf("[ %s ]actual method and expected method are different", tc.Name)
					t.Fail()
				}
				if tc.Path != response.Path {
					t.Errorf("[ %s ]actual path and expected path are different", tc.Name)
					t.Fail()
				}
			}
		})
	}
}

func mapsAreEqual(map1, map2 map[string]string) bool {
	// First check if the lengths are the same
	if len(map1) != len(map2) {
		return false
	}

	// Then, check if each key-value pair is the same
	for key, value := range map1 {
		if map2[key] != value {
			return false
		}
	}

	return true
}
