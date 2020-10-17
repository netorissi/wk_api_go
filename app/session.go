package app

import (
	"net/http"

	"github.com/netorissi/wk_api_go/entities"
	"github.com/netorissi/wk_api_go/utils"
)

func (a *App) CreateSession(r *http.Request, session *entities.Session) (*entities.Session, *entities.AppError) {

	session.IPAddress = utils.GetIpAddress(r)

	result := <-a.Srv.SqlStore.Sessions().Create(session)

	if result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*entities.Session), nil
}
