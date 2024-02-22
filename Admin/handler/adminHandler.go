package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	adminpb "github.com/grpc/gobus/apigw/Admin/pb"
	"github.com/grpc/gobus/apigw/models"
)

func AdminLoginHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	var admin *models.AdminLogin
	if err := c.ShouldBindJSON(&admin); err != nil {
		log.Printf("error binding JSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	ctx := context.Background()
	logReq := &adminpb.LoginRequest{
		Username: admin.Username,
		Password: admin.Password,
	}
	response, err := client.AdminLogin(ctx, logReq)
	if err != nil {
		log.Printf("error logging in admin %v err: %v", admin.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v logged in succesfully", admin.Username),
		"data":    response,
	})
}
func EditUserHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	var user *models.UserSignUp
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding user")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}
	ctx := context.Background()
	response, err := client.EditUser(ctx, &adminpb.UserDataRequest{
		UserName: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		log.Printf("error editing user %v err: %v", user.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v user edited succesfully", user.Username),
		"data":    response,
	})
}
func DeleteUserByIdHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "id param missing",
		})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}
	ctx := context.Background()
	response, err := client.DeleteUserById(ctx, &adminpb.IdRequest{
		Id: uint32(id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error deleting user",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "user deleted successfully",
		"data":    response,
	})
}
func FindUserByIdHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "id param missing",
		})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}
	ctx := context.Background()
	response, err := client.FindUserById(ctx, &adminpb.IdRequest{
		Id: uint32(id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error finding user",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  response.Status,
		"message": response.Message,
		"data":    response.Status,
	})
}
func CreateUserHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	var user *models.UserSignUp
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding user")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}
	ctx := context.Background()
	response, err := client.CreateUser(ctx, &adminpb.UserDataRequest{
		UserName: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		log.Printf("error editing user %v err: %v", user.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v user created succesfully", user.Username),
		"data":    response,
	})
}
func EditProviderHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	var provider *models.ProviderSignUp
	if err := c.ShouldBindJSON(&provider); err != nil {
		log.Printf("Error binding provider")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}
	ctx := context.Background()
	response, err := client.EditProvider(ctx, &adminpb.ProviderDataRequest{
		UserName: provider.Username,
		Email:    provider.Email,
		Password: provider.Password,
	})
	if err != nil {
		log.Printf("error editing provider %v err: %v", provider.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v provider edited succesfully", provider.Username),
		"data":    response,
	})
}
func DeleteProviderByIdHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "id param missing",
		})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}
	ctx := context.Background()
	response, err := client.DeleteProviderById(ctx, &adminpb.IdRequest{
		Id: uint32(id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error deleting provider",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "provider deleted successfully",
		"data":    response,
	})
}
func FindProviderByIdHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "id param missing",
		})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}
	ctx := context.Background()
	response, err := client.FindProviderById(ctx, &adminpb.IdRequest{
		Id: uint32(id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error finding provider",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  response.Status,
		"message": response.Message,
		"data":    response.Status,
	})
}
func CreateProviderHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	var provider *models.ProviderSignUp
	if err := c.ShouldBindJSON(&provider); err != nil {
		log.Printf("Error binding provider")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}
	ctx := context.Background()
	response, err := client.CreateProvider(ctx, &adminpb.ProviderDataRequest{
		UserName: provider.Username,
		Email:    provider.Email,
		Password: provider.Password,
	})
	if err != nil {
		log.Printf("error editing provider %v err: %v", provider.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v provider created succesfully", provider.Username),
		"data":    response,
	})
}
