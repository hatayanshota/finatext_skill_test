package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/usecase"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *Handler) Login(c *gin.Context) {
	var reqBody LoginRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Printf("error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	param := usecase.LoginParam{
		Username: reqBody.Username,
		Password: reqBody.Password,
	}

	result := usecase.Login(param)

	response := LoginResponse{
		Token: result.Token,
	}

	c.JSON(http.StatusOK, response)
}
