package app

import (
	"fmt"
	"sync/atomic"

	"github.com/gorilla/mux"
	"github.com/netorissi/wk_api_go/app/database/nosqlstore"
	"github.com/netorissi/wk_api_go/app/database/sqlstore"
	"github.com/netorissi/wk_api_go/app/infra"
	"github.com/netorissi/wk_api_go/entities"
)

// App default config
type App struct {
	goroutineCount      int32
	NatsID              string
	goroutineExitSignal chan struct{}

	Srv           *Server
	newSQLStore   func() sqlstore.Store
	newNoSQLStore func() nosqlstore.Store

	// SessionCache     *utils.Cache
	configListenerID string

	// NC                  *nats.Conn
}

// Option struct at App
type Option func(a *App)

// ConfigPackage default at package
func ConfigPackage() Option {
	return func(a *App) {}
}

func (a *App) Config() *entities.Config {
	return infra.Configurations
}

var instanceCount = 0

func New(options ...Option) *App {
	instanceCount++
	if instanceCount > 1 {
		panic("Only one App should exist at a time. Did you forget to call Shutdown()?")
	}

	app := &App{
		goroutineExitSignal: make(chan struct{}, 1),
		Srv: &Server{
			Router: mux.NewRouter(),
		},
	}

	for _, option := range options {
		option(app)
	}

	if app.newSQLStore == nil {
		app.newSQLStore = func() sqlstore.Store {
			return sqlstore.NewLayeredStore(sqlstore.NewSqlSupplier(app.Config()))
		}
	}

	if app.newNoSQLStore == nil {
		app.newNoSQLStore = func() nosqlstore.Store {
			return nosqlstore.NewLayeredStore(nosqlstore.NewNoSqlSupplier(app.Config()))
		}
	}

	app.Srv.SqlStore = app.newSQLStore()
	app.Srv.NoSqlStore = app.newNoSQLStore()

	fmt.Println("[INFO] Server is initializing...")

	return app
}

func (a *App) WaitForGoroutines() {
	for atomic.LoadInt32(&a.goroutineCount) != 0 {
		<-a.goroutineExitSignal
	}
}

func (a *App) Shutdown() {
	instanceCount--

	a.StopServer()

	a.WaitForGoroutines()

	a.Srv.SqlStore.Close()
	a.Srv.NoSqlStore.Close()
	a.Srv = nil
}
