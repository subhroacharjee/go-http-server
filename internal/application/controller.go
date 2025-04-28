package application

import (
	"fmt"
	"io"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/internal/httpcore"
	"github.com/codecrafters-io/http-server-starter-go/internal/router"
)

func RegisterControllers(appRouter router.IRouter, directory *string) {
	appRouter.Get("/", func(r httpcore.Request, w *httpcore.HttpResponseWriter) {
		w.SetStatus(httpcore.StatusOK)
	})
	appRouter.Get("/index.html", func(r httpcore.Request, w *httpcore.HttpResponseWriter) {
		w.SetStatus(httpcore.StatusOK)
	})
	appRouter.Get("/echo/:str", func(r httpcore.Request, w *httpcore.HttpResponseWriter) {
		value := r.PathParams["str"]
		w.SetHeader("Content-Type", "text/plain")
		w.Write([]byte(value))
	})
	appRouter.Get("/user-agent", func(r httpcore.Request, w *httpcore.HttpResponseWriter) {
		value := r.Headers["user-agent"]
		w.SetHeader("Content-Type", "text/plain")
		w.Write([]byte(value))
	})

	appRouter.Get("/files/:filename", func(r httpcore.Request, w *httpcore.HttpResponseWriter) {
		filename, exists := r.PathParams["filename"]
		if !exists {
			w.SetStatus(httpcore.StatusNotFound)
			return
		}
		absolutePath := fmt.Sprintf("%s%s", *directory, filename)

		file, err := os.Open(absolutePath)
		if err != nil {
			if os.IsNotExist(err) {
				w.SetStatus(httpcore.StatusNotFound)
				return
			}

			fmt.Println(err)
			w.SetStatus(httpcore.StatusInternalServerError)
			return
		}

		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			w.SetStatus(httpcore.StatusInternalServerError)
			return
		}

		w.SetHeader("Content-Type", "application/octet-stream")
		w.Write(content)
	})

	appRouter.Post("/files/:filename", func(r httpcore.Request, w *httpcore.HttpResponseWriter) {
		filename, exists := r.PathParams["filename"]
		if !exists {
			w.SetStatus(httpcore.StatusNotFound)
			return
		}
		absolutePath := fmt.Sprintf("%s%s", *directory, filename)

		if err := os.WriteFile(absolutePath, r.Body, 0644); err != nil {
			w.SetStatus(httpcore.StatusInternalServerError)
			return
		}

		w.SetStatus(httpcore.StatusAccepted)
	})
}
