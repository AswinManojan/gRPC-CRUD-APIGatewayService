package admin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/grpc/gobus/apigw/Admin/handler"
	adminpb "github.com/grpc/gobus/apigw/Admin/pb"
	"github.com/grpc/gobus/apigw/config"
)

type AdminRoutes struct {
	client adminpb.AdminServiceClient
}

func NewAdminRoute(c *gin.Engine, cfg *config.Config) {
	client, err := AdminClientDial(cfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}
	adminHandler := AdminRoutes{
		client: client,
	}
	apiAdmin := c.Group("/api/admin")
	{
		apiAdmin.POST("/login", adminHandler.Login)
		apiAdmin.PATCH("/edit/user", adminHandler.EditUser)
		apiAdmin.DELETE("delete/user", adminHandler.DeleteUser)
		apiAdmin.GET("/user/id", adminHandler.FindUserByID)
		apiAdmin.POST("/create/user", adminHandler.CreateUser)
		apiAdmin.PATCH("/edit/provider", adminHandler.EditProvider)
		apiAdmin.DELETE("delete/provider", adminHandler.DeleteProviderById)
		apiAdmin.GET("/provider/id", adminHandler.FindProviderByID)
		apiAdmin.POST("/create/provider", adminHandler.CreateProvider)
	}
}

func (a *AdminRoutes) Login(c *gin.Context) {
	handler.AdminLoginHandler(c, a.client)
}

func (a *AdminRoutes) EditUser(c *gin.Context) {
	handler.EditUserHandler(c, a.client)
}

func (a *AdminRoutes) DeleteUser(c *gin.Context) {
	handler.DeleteUserByIdHandler(c, a.client)
}

func (a *AdminRoutes) FindUserByID(c *gin.Context) {
	handler.FindUserByIdHandler(c, a.client)
}
func (a *AdminRoutes) CreateUser(c *gin.Context) {
	handler.CreateUserHandler(c, a.client)
}
func (a *AdminRoutes) EditProvider(c *gin.Context) {
	handler.EditProviderHandler(c, a.client)
}

func (a *AdminRoutes) DeleteProviderById(c *gin.Context) {
	handler.DeleteProviderByIdHandler(c, a.client)
}

func (a *AdminRoutes) FindProviderByID(c *gin.Context) {
	handler.FindProviderByIdHandler(c, a.client)
}
func (a *AdminRoutes) CreateProvider(c *gin.Context) {
	handler.CreateProviderHandler(c, a.client)
}
