package routes

import (
	"net/http"
	"strings"

	"github.com/netorissi/wk_api_go/app"
	"github.com/netorissi/wk_api_go/entities"
	"github.com/netorissi/wk_api_go/utils"
)

type Context struct {
	App       *app.App
	Session   entities.Session
	Params    *Params
	Err       *entities.AppError
	Path      string
	IpAddress string
}

type handler struct {
	app            *app.App
	handleFunc     func(*Context, http.ResponseWriter, *http.Request)
	requireSession bool
	requireAdmin   bool
}

func (routes *Routes) Public(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		app:            routes.App,
		handleFunc:     h,
		requireSession: false,
		requireAdmin:   false,
	}
}

func (routes *Routes) Private(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		app:            routes.App,
		handleFunc:     h,
		requireSession: true,
		requireAdmin:   false,
	}
}

func (routes *Routes) Admin(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		app:            routes.App,
		handleFunc:     h,
		requireSession: true,
		requireAdmin:   true,
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{}
	c.App = h.app
	c.Params = GetParamsFromRequest(r)
	c.IpAddress = utils.GetIpAddress(r)

	token := ""

	authHeader := r.Header.Get("Authorization")
	if len(authHeader) > 6 && strings.ToUpper(authHeader[0:6]) == "BEARER" {
		token = authHeader[7:]

	} else if len(authHeader) > 5 && strings.ToLower(authHeader[0:5]) == "BEARER" {
		token = authHeader[6:]
	}

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		w.Header().Set("Expires", "0")
	}

	if h.requireSession {
		if ok, err := h.app.ValidToken(token); ok != true || err != nil {
			c.Err = err
			w.WriteHeader(c.Err.StatusCode)
			w.Write([]byte(c.Err.ToJson()))
			return
		}

		if c.Session.UserID == 0 {
			c.Err = entities.NewAppError("", "Inválido ou sessão expirada, por favor faça login novamente.", nil, "UserRequired", http.StatusUnauthorized)

			w.WriteHeader(c.Err.StatusCode)
			w.Write([]byte(c.Err.ToJson()))

			return
		}

		// if h.requireAdmin && !c.Session.IsAdmin() {
		// 	c.Err = entities.NewAppError("", "Você não possui privilégios suficientes para realizar essa operação.", nil, "AdminRequired", http.StatusUnauthorized)

		// 	w.WriteHeader(c.Err.StatusCode)
		// 	w.Write([]byte(c.Err.ToJson()))

		// 	return
		// }
	}

	h.handleFunc(c, w, r)
}
