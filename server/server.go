package server

import (
	"github.com/gin-gonic/gin"
	"github.com/grpc/gobus/apigw/config"
)

type ServerStruct struct {
	R *gin.Engine
}

func (s *ServerStruct) StartServer(cfg *config.Config) {
	s.R.Run(":" + cfg.APIGATEWAYPORT)
}

func Server() *ServerStruct {
	engine := gin.Default()
	return &ServerStruct{
		R: engine,
	}
}
