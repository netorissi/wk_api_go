package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/netorissi/wk_api_go/app"
	"github.com/netorissi/wk_api_go/app/infra"
	"github.com/netorissi/wk_api_go/routes"
)

func main() {
	if err := runServer(); err != nil {
		os.Exit(1)
	}
}

func runServer() error {
	if err := infra.LoadGlobalConfig(); err != nil {
		fmt.Println("[ERROR] Config fail loaded!")
		return err
	}

	a := app.New()
	defer a.Shutdown()

	a.StartServer()
	// nats.InitNats(a)
	routes.Init(a, a.Srv.Router)

	fmt.Println("Server running")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	return nil
}
