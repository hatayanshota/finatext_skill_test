package usecase

import (
	"fmt"
	"github.com/samber/lo"
	"main/dao"
	"sort"
	"time"
)

type CandleUseCase struct {
	OrderBookDao *dao.OrderBookDao
}

func NewCandleUseCase(orderBookDao *dao.OrderBookDao) *CandleUseCase {
	return &CandleUseCase{
		OrderBookDao: orderBookDao,
	}
}

type CandleParam struct {
	Code  string
	Year  int
	Month int
	Day   int
	Hour  int
}

type CandleResult struct {
	Open  int
	Close int
	High  int
	Low   int
}

func (u *CandleUseCase) Run(param CandleParam) (CandleResult, error) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	startTime := time.Date(param.Year, time.Month(param.Month), param.Day, param.Hour, 0, 0, 0, location)
	endTime := startTime.Add(1 * time.Hour)

	targetRecords, err := u.OrderBookDao.FindByCodeDuration(param.Code, startTime, endTime)
	if err != nil {
		return CandleResult{}, fmt.Errorf("failed to FindByCodeDuration: %w", err)
	}

	if len(targetRecords) == 0 {
		return CandleResult{}, fmt.Errorf("no records found")
	}

	// 時間の昇順でソート
	sort.Slice(targetRecords, func(i, j int) bool {
		return targetRecords[i].Time.Before(targetRecords[j].Time)
	})

	prices := lo.Map(targetRecords, func(i dao.OrderBook, _ int) int {
		return i.Price
	})

	return CandleResult{
		Open:  targetRecords[0].Price,
		Close: targetRecords[len(targetRecords)-1].Price,
		High:  lo.Max(prices),
		Low:   lo.Min(prices),
	}, nil
}
