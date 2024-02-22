package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/grpc/gobus/apigw/Provider/PB"
	"github.com/grpc/gobus/apigw/models"
)

func ProviderLoginHandler(c *gin.Context, client pb.ProviderServicesClient) {
	ctx := context.Background()
	var login models.Provider
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Printf("error binding JSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	req := &pb.LoginReq{
		Username: login.Username,
		Password: login.Password,
	}
	res, err := client.ProviderLogin(ctx, req)
	if err != nil {
		log.Printf("error logging in user %v err: %v", login.Username, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v logged in successfully", login.Username),
		"data":    res.Message,
	})
}

func ProviderSignUpHandler(c *gin.Context, client pb.ProviderServicesClient) {
	ctx := context.Background()
	var signup models.ProviderSignUp
	if err := c.ShouldBindJSON(&signup); err != nil {
		log.Printf("error binding JSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	req := &pb.ProviderDataReq{
		UserName: signup.Username,
		Password: signup.Password,
		Email:    signup.Email,
	}
	res, err := client.ProviderSignUp(ctx, req)
	if err != nil {
		log.Printf("error logging in user %v err: %v", signup.Username, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v signed up successfully", signup.Username),
		"data":    res.Message,
	})
}
