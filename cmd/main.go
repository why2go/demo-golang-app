package main

import (
	"demo/internal/ctrler/grpc"
	"demo/internal/ctrler/http"

	"github.com/why2go/gostarter/boot"
	"github.com/why2go/gostarter/ginstarter"
	"github.com/why2go/gostarter/grpcstarter"
)

func main() {
	boot.AddStarters(
		http.RegisterHandlers,
		ginstarter.StartHttpServer,
		grpc.RegisterGrpcService,
		grpcstarter.StartGrpcServer,
	)

	boot.AddSweepers(
		ginstarter.StopHttpServer,
		grpcstarter.StopGrpcServer,
	)

	boot.Run()
}
