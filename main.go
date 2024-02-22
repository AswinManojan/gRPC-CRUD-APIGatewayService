package main

import (
	admin "github.com/grpc/gobus/apigw/Admin"
	provider "github.com/grpc/gobus/apigw/Provider"
	user "github.com/grpc/gobus/apigw/User"
	"github.com/grpc/gobus/apigw/config"
	"github.com/grpc/gobus/apigw/server"
)

func main() {
	cfg := config.LoadConfig()
	srv := server.Server()
	user.NewUserRoute(srv.R, cfg)
	admin.NewAdminRoute(srv.R, cfg)
	provider.NewProviderRoute(srv.R, cfg)
	srv.StartServer(cfg)
}
