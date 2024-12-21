package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FlagRequest struct {
	Flag string `json:"flag"`
}

func (h *Handler) Flag(c *gin.Context) {
	var reqBody FlagRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Printf("error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("flag: %s\n", reqBody.Flag)
}
