package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/netorissi/wk_api_go/app"
	"github.com/netorissi/wk_api_go/entities"
)

type RoutesMux struct {
	Root     *mux.Router
	ApiRoot  *mux.Router
	ApiUsers *mux.Router
	ApiAuth  *mux.Router
}

type Routes struct {
	App        *app.App
	BaseRoutes *RoutesMux
}

func Init(a *app.App, root *mux.Router) *Routes {
	routes := &Routes{
		App:        a,
		BaseRoutes: &RoutesMux{},
	}

	routes.BaseRoutes.Root = root
	routes.BaseRoutes.ApiRoot = root.PathPrefix("/api/v1/").Subrouter()
	routes.BaseRoutes.ApiAuth = routes.BaseRoutes.ApiRoot.PathPrefix("/auth/").Subrouter()
	routes.BaseRoutes.ApiUsers = routes.BaseRoutes.ApiRoot.PathPrefix("/users/").Subrouter()

	// INJECTOR-INIT

	root.Handle("/status", http.HandlerFunc(ReturnStatusOK))
	root.Handle("/{anything:.*}", http.HandlerFunc(Handle404))

	return routes
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	err := entities.NewAppError("Handle404", "Erro", nil, "", http.StatusNotFound)

	w.WriteHeader(err.StatusCode)
	err.DetailedError = "No exist url='" + r.URL.Path + "'."
	w.Write([]byte(err.ToJson()))
}

func ReturnStatusOK(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["status"] = "OK"
	w.Write([]byte(entities.MapToJson(m)))
}

func ReturnStatus(w http.ResponseWriter, sc int, resp []byte) {
	w.WriteHeader(sc)
	w.Write(resp)
}
