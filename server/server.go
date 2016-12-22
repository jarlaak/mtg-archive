package server

import (
	"github.com/gorilla/mux"
)

type AdvancedRouter struct {
	*mux.Router
}

type AdvancedRoute struct {
	*mux.Route
}

func NewRouter() *AdvancedRouter {
	return NewAdvancedRouter(mux.NewRouter())
}

func NewAdvancedRouter(r *mux.Router) *AdvancedRouter {
	return &AdvancedRouter{r}

}
func NewAdvancedRoute(r *mux.Route) *AdvancedRoute {
	return &AdvancedRoute{r}
}

func (r *AdvancedRouter) HandleFunc(path string, f requestHandler) {
	handler := LogRequest(f)
	r.Router.HandleFunc(path, handler)
}

func (r *AdvancedRoute) Subrouter() *AdvancedRouter {
	sr := r.Route.Subrouter()
	return NewAdvancedRouter(sr)
}

func (r *AdvancedRouter) PathPrefix(path string) *AdvancedRoute {
	return NewAdvancedRoute(r.Router.PathPrefix(path))
}
