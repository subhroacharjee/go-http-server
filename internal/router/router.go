package router

import (
	"strings"
	"unicode"

	"github.com/codecrafters-io/http-server-starter-go/internal/common"
	"github.com/codecrafters-io/http-server-starter-go/internal/httpcore"
)

type Route struct {
	children   map[string]*Route
	middleware []httpcore.HandlerFunc
	hasParam   bool
}

func NewRoute() *Route {
	return &Route{children: make(map[string]*Route), middleware: make([]httpcore.HandlerFunc, 0), hasParam: false}
}

type ReadOnlyRouter interface {
	GetHandlers(method common.Method, path string) ([]httpcore.HandlerFunc, map[string]string)
	CopyPath(router IRouter)
}

type IRouter interface {
	Get(path string, handlers ...httpcore.HandlerFunc)
	Post(path string, handlers ...httpcore.HandlerFunc)
	Put(path string, handlers ...httpcore.HandlerFunc)
	Patch(path string, handlers ...httpcore.HandlerFunc)
	Head(path string, handlers ...httpcore.HandlerFunc)
	Delete(path string, handlers ...httpcore.HandlerFunc)
}

type Router struct {
	root *Route
}

func (r *Router) Get(path string, handlers ...httpcore.HandlerFunc) {
	r.addRoute(common.GET, path, handlers...)
}

func (r *Router) Post(path string, handlers ...httpcore.HandlerFunc) {
	r.addRoute(common.POST, path, handlers...)
}

func (r *Router) Head(path string, handlers ...httpcore.HandlerFunc) {
	r.addRoute(common.HEAD, path, handlers...)
}

func (r *Router) Put(path string, handlers ...httpcore.HandlerFunc) {
	r.addRoute(common.PUT, path, handlers...)
}

func (r *Router) Patch(path string, handlers ...httpcore.HandlerFunc) {
	r.addRoute(common.PATCH, path, handlers...)
}

func (r *Router) Delete(path string, handlers ...httpcore.HandlerFunc) {
	r.addRoute(common.DELETE, path, handlers...)
}

func NewRouter() IRouter {
	return &Router{
		root: NewRoute(),
	}
}

func (r *Router) CopyPath(router IRouter) {
	r.root = router.(*Router).root
}

func (r Router) GetHandlers(method common.Method, path string) ([]httpcore.HandlerFunc, map[string]string) {
	routeSegments := strings.Split(strings.TrimRightFunc(strings.ReplaceAll(path, "/", " "), unicode.IsSpace), " ")
	routeSegments = append([]string{string(method)}, routeSegments...)

	handlers, pathParam := make([]httpcore.HandlerFunc, 0), make(map[string]string)
	current := r.root
	for _, segment := range routeSegments {
		child, exists := current.children[segment]
		if !exists {

			for key, value := range current.children {
				if value.hasParam {
					child = value
					pathParam[strings.ReplaceAll(key, ":", "")] = segment
					exists = true
					break
				}
			}
			// in case no path param present
			if !exists {
				return nil, pathParam
			}

		}
		current = child
	}
	handlers = current.middleware

	return handlers, pathParam
}

func (r *Router) addRoute(method common.Method, path string, handlers ...httpcore.HandlerFunc) {
	routeSegments := strings.Split(strings.TrimRightFunc(strings.ReplaceAll(path, "/", " "), unicode.IsSpace), " ")
	routeSegments = append([]string{string(method)}, routeSegments...)
	current := r.root

	for _, segment := range routeSegments {
		// fmt.Println(segment)
		child, exists := current.children[segment]
		if !exists {
			child = NewRoute()

			if len(segment) > 0 && segment[0] == ':' {
				child.hasParam = true
			}
			current.children[segment] = child
		}
		current = child
	}

	current.middleware = handlers
}
