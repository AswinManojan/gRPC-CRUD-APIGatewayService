package user

import (
	"log"

	"github.com/gin-gonic/gin"
	handler "github.com/grpc/gobus/apigw/User/Handler"
	pb "github.com/grpc/gobus/apigw/User/PB"
	"github.com/grpc/gobus/apigw/config"
)

type User struct {
	client pb.UserServicesClient
}

func (u *User) Login(c *gin.Context) {
	handler.UserLoginHandler(c, u.client)
}

func (u *User) SignUp(c *gin.Context) {
	handler.UserSignUpHandler(c, u.client)
}

func NewUserRoute(c *gin.Engine, cfg *config.Config) {
	client, err := UserClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with gRPC server, %v", err.Error())
	}
	userHandler := User{
		client: client,
	}
	apiUser := c.Group("/api/user")
	{
		apiUser.POST("/login", userHandler.Login)
		apiUser.POST("/signup", userHandler.SignUp)
	}
}
