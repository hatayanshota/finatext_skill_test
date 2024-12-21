package dao

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const (
	csvFilePath   = "resource/csv/order_books.csv"
	csvTimeLayout = "2006-01-02 15:04:05 -0700 MST"
)

type OrderBookDao struct{}

func NewOrderBookDao() *OrderBookDao {
	return &OrderBookDao{}
}

type OrderBook struct {
	Time  time.Time
	Code  string
	Price int
}

func (dao *OrderBookDao) FindByCodeDuration(code string, startTime, endTime time.Time) ([]OrderBook, error) {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open csv file: %w", err)
	}

	reader := csv.NewReader(csvFile)

	// ヘッダー行を読み飛ばす
	_, err = reader.Read()
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	results := make([]OrderBook, 0)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("failed to read record: %w", err)
		}

		isTarget, targetCode, targetTime, err := isTargetRecord(code, startTime, endTime, record)
		if err != nil {
			return nil, fmt.Errorf("failed to check record: %w", err)
		}

		if !isTarget {
			continue
		}

		targetPrice, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, fmt.Errorf("failed to parse price: %w", err)
		}

		results = append(results, OrderBook{
			Time:  targetTime,
			Code:  targetCode,
			Price: targetPrice,
		})
	}

	return results, nil
}

// isTargetRecord レコードが対象のデータかどうかを判定する
// 対象の場合はtrueを返し、コードと時間を返す
func isTargetRecord(code string, startTime, endTime time.Time, record []string) (bool, string, time.Time, error) {
	recordTime, err := time.Parse(csvTimeLayout, record[0])
	if err != nil {
		return false, "", time.Time{}, fmt.Errorf("failed to parse time: %w", err)
	}

	if recordTime.Before(startTime) || recordTime.After(endTime) || recordTime.Equal(endTime) {
		// 時間が範囲外の場合は無効なデータ
		return false, "", time.Time{}, nil
	}

	recordCode := record[1]
	if recordCode != code {
		// コードが一致しない場合は無効なデータ
		return false, "", time.Time{}, nil
	}

	return true, recordCode, recordTime, nil
}
