package routes

import (
	"net/http"

	"github.com/netorissi/wk_api_go/entities"
)

// InitRouteUsers - init routes to users
func (routes *Routes) InitRouteUsers() {

	apiUsers := routes.BaseRoutes.ApiUsers
	public := routes.Public

	apiUsers.Handle("", public(createUser)).Methods("POST")
}

func createUser(c *Context, w http.ResponseWriter, r *http.Request) {
	body := entities.UserFromJSON(r.Body)

	user, err := c.App.CreateUser(body)

	if err != nil {
		c.Err = err
		return
	}

	ReturnStatus(w, http.StatusCreated, []byte(user.UserToJSON()))
}
