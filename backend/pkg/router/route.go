package route

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Route is an HTTP route with optional children.
type Route struct {
	Method   string
	Path     string
	Handler  http.Handler
	Children []*Route
}

// Simple creates a simple route.
func Simple(method, path string, handler http.Handler) *Route {
	return &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

// Group creates a route group without an own handler.
func Group(path string, children []*Route) *Route {
	return &Route{
		Path:     path,
		Children: children,
	}
}

// Build builds the route into an HTTP handler.
func (r *Route) Build() (http.Handler, error) {
	router := httprouter.New()
	err := r.register(router, "/")
	if err != nil {
		return nil, err
	}
	return router, nil
}

// register registers the route and all its children.
func (r *Route) register(router *httprouter.Router, prefix string) error {
	if r.Method != "" {
		router.Handler(r.Method, prefix+r.Path, r.Handler)
	}
	for _, child := range r.Children {
		err := child.register(router, prefix+r.Path)
		if err != nil {
			return err
		}
	}
	return nil
}
