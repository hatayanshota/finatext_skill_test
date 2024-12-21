package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/dao"
	"main/handler"
)

func main() {
	r := gin.Default()

	container := NewContainer()

	g := r.Group("/api")
	{
		g.GET("/candle", container.Handler.Candle)

		g.PUT("/login", container.Handler.Login)
		g.PUT("/flag", container.Handler.Flag)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

type Container struct {
	Handler *handler.Handler
}

func NewContainer() Container {
	orderBookDao := dao.NewOrderBookDao()

	return Container{
		Handler: handler.NewHandler(orderBookDao),
	}
}
