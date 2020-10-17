package routes

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	PAGE_DEFAULT  int = 0
	LIMIT_DEFAULT int = 10
)

type Params struct {
	UserID string
	Page   int
	Limit  int
}

func GetParamsFromRequest(r *http.Request) *Params {
	params := &Params{}

	props := mux.Vars(r)
	query := r.URL.Query()

	if val, ok := props["user_id"]; ok {
		params.UserID = val
	}

	page := query.Get("page")
	if num, err := strconv.Atoi(page); len(page) == 0 || err != nil {
		params.Page = PAGE_DEFAULT
	} else {
		params.Page = num
	}

	limit := query.Get("limit")
	if num, err := strconv.Atoi(limit); len(limit) == 0 || err != nil {
		params.Limit = LIMIT_DEFAULT
	} else {
		params.Limit = num
	}

	return params
}
