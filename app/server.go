package app

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/netorissi/wk_api_go/app/database/nosqlstore"
	"github.com/netorissi/wk_api_go/app/database/sqlstore"
	"github.com/netorissi/wk_api_go/entities"
)

const (
	WAIT_CONNECTIONS_TO_SHUTDOWN = time.Second
	READ_TIMEOUT                 = 300
	WRITE_TIMEOUT                = 300
	PORT_SERVER                  = ":8080"
)

type Server struct {
	SqlStore   sqlstore.Store
	NoSqlStore nosqlstore.Store
	Router     *mux.Router
	Server     *http.Server
	ListenAddr *net.TCPAddr

	didFinishListen chan struct{}
}

var allowedMethods []string = []string{
	"POST",
	"GET",
	"OPTIONS",
	"PUT",
	"PATCH",
	"DELETE",
}

type CorsWrapper struct {
	config entities.ConfigFunc
	router *mux.Router
}

type RecoveryLogger struct{}

func (rl *RecoveryLogger) Println(i ...interface{}) {
	fmt.Println("[CRASH APPLICATION ERROR]\n", i)
}

func checkOrigin(r *http.Request, allowedOrigins string) bool {
	origin := r.Header.Get("Origin")
	if allowedOrigins == "*" {
		return true
	}
	for _, allowed := range strings.Split(allowedOrigins, " ") {
		if allowed == origin {
			return true
		}
	}
	return false
}

func (cw *CorsWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if checkOrigin(r, "*") {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))

		if r.Method == "OPTIONS" {
			w.Header().Set(
				"Access-Control-Allow-Methods",
				strings.Join(allowedMethods, ", "))

			w.Header().Set(
				"Access-Control-Allow-Headers",
				r.Header.Get("Access-Control-Request-Headers"))
		}
	}

	if r.Method == "OPTIONS" {
		return
	}

	cw.router.ServeHTTP(w, r)
}

func redirectHTTPToHTTPS(w http.ResponseWriter, r *http.Request) {
	if r.Host == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
	}

	url := r.URL
	url.Host = r.Host
	url.Scheme = "https"
	http.Redirect(w, r, url.String(), http.StatusFound)
}

func (a *App) StartServer() {
	var handler http.Handler = &CorsWrapper{a.Config, a.Srv.Router}

	a.Srv.Server = &http.Server{
		Handler:      handlers.RecoveryHandler(handlers.RecoveryLogger(&RecoveryLogger{}), handlers.PrintRecoveryStack(true))(handler),
		ReadTimeout:  time.Duration(READ_TIMEOUT) * time.Second,
		WriteTimeout: time.Duration(WRITE_TIMEOUT) * time.Second,
	}

	addr := PORT_SERVER
	if addr == "" {
		addr = ":http"
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	a.Srv.ListenAddr = listener.Addr().(*net.TCPAddr)
	fmt.Println("[INFO] Listen server on:", listener.Addr().String())

	a.Srv.didFinishListen = make(chan struct{})

	go func() {
		var err error
		err = a.Srv.Server.Serve(listener)

		if err != nil && err != http.ErrServerClosed {
			fmt.Println("Listen server off:", err)
			time.Sleep(time.Second)
		}
		close(a.Srv.didFinishListen)
	}()
}

func (a *App) StopServer() {
	if a.Srv.Server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), WAIT_CONNECTIONS_TO_SHUTDOWN)
		defer cancel()
		didShutdown := false
		for a.Srv.didFinishListen != nil && !didShutdown {
			a.Srv.Server.Shutdown(ctx)
			timer := time.NewTimer(time.Millisecond * 50)
			select {
			case <-a.Srv.didFinishListen:
				didShutdown = true
			case <-timer.C:
			}
			timer.Stop()
		}
		a.Srv.Server.Close()
		a.Srv.Server = nil
	}
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tcpConn, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	tcpConn.SetKeepAlive(true)
	tcpConn.SetKeepAlivePeriod(3 * time.Minute)
	return tcpConn, nil
}

func (a *App) Listen(addr string) (net.Listener, error) {
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return tcpKeepAliveListener{ln.(*net.TCPListener)}, nil
}

func consumeAndClose(r *http.Response) {
	if r.Body != nil {
		io.Copy(ioutil.Discard, r.Body)
		r.Body.Close()
	}
}
