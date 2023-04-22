package main

import (
	"demo/internal/ctrler/http"

	"github.com/why2go/gostarter/boot"
	"github.com/why2go/gostarter/ginstarter"
)

func main() {
	boot.AddStarters(
		http.RegisterHandlers,
		ginstarter.StartHttpServer,
	)

	boot.AddSweepers(
		ginstarter.StopHttpServer,
	)

	boot.Run()
}
