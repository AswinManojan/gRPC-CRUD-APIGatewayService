package provider

import (
	"log"

	"github.com/gin-gonic/gin"
	handler "github.com/grpc/gobus/apigw/Provider/Handler"
	pb "github.com/grpc/gobus/apigw/Provider/PB"
	"github.com/grpc/gobus/apigw/config"
)

type Provider struct {
	client pb.ProviderServicesClient
}

func (u *Provider) Login(c *gin.Context) {
	handler.ProviderLoginHandler(c, u.client)
}

func (u *Provider) SignUp(c *gin.Context) {
	handler.ProviderSignUpHandler(c, u.client)
}

func NewProviderRoute(c *gin.Engine, cfg *config.Config) {
	client, err := ProviderClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with gRPC server, %v", err.Error())
	}
	providerHandler := Provider{
		client: client,
	}
	apiProvider := c.Group("/api/provider")
	{
		apiProvider.POST("/login", providerHandler.Login)
		apiProvider.POST("/signup", providerHandler.SignUp)
	}
}
