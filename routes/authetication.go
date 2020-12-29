package routes

import (
	"fmt"
	"net/http"

	"github.com/netorissi/wk_api_go/entities"
)

func (routes *Routes) InitRouteAuth() {
	fmt.Println("[START] - InitRouteAuth")

	apiAuth := routes.BaseRoutes.ApiAuth
	public := routes.Public

	apiAuth.Handle("/login", public(login)).Methods("POST")
}

func login(c *Context, w http.ResponseWriter, r *http.Request) {
	body := entities.AuthenticationFromJSON(r.Body)

	auth, err := c.App.Login(body, r)

	if err != nil {
		c.Err = err
		ReturnStatus(w, err.StatusCode, []byte(err.ToJson()))
		return
	}

	ReturnStatus(w, http.StatusCreated, []byte(auth.ResponseAuthToJSON()))
}
