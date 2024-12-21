package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/usecase"
	"net/http"
	"strconv"
)

type CandleRequest struct {
	Code  string `json:"code"`
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
	Hour  int    `json:"hour"`
}

type CandleResponse struct {
	Open  int `json:"open"`
	Close int `json:"close"`
	High  int `json:"high"`
	Low   int `json:"low"`
}

func (h *Handler) Candle(c *gin.Context) {
	year := c.Query("year")
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		log.Printf("error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	month := c.Query("month")
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		log.Printf("error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	day := c.Query("day")
	dayInt, err := strconv.Atoi(day)
	if err != nil {
		log.Printf("error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hour := c.Query("hour")
	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		log.Printf("error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reqBody := CandleRequest{
		c.Query("code"),
		yearInt,
		monthInt,
		dayInt,
		hourInt,
	}

	log.Printf("request: %+v\n", reqBody)

	result, err := usecase.NewCandleUseCase(h.OrderBookDao).Run(usecase.CandleParam{
		Code:  reqBody.Code,
		Year:  reqBody.Year,
		Month: reqBody.Month,
		Day:   reqBody.Day,
		Hour:  reqBody.Hour,
	})
	if err != nil {
		log.Printf("error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, CandleResponse{
		Open:  result.Open,
		Close: result.Close,
		High:  result.High,
		Low:   result.Low,
	})
}
