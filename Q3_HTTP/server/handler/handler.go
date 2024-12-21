package handler

import "main/dao"

type Handler struct {
	OrderBookDao *dao.OrderBookDao
}

func NewHandler(orderBookDao *dao.OrderBookDao) *Handler {
	return &Handler{
		OrderBookDao: orderBookDao,
	}
}
