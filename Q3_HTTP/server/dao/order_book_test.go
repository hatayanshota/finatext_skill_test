package dao

import (
	"testing"
	"time"
)

func TestOrderBookDao_isTargetRecord(t *testing.T) {
	t.Parallel()

	type args struct {
		code      string
		startTime time.Time
		endTime   time.Time
		record    []string
	}

	type res struct {
		isTarget   bool
		targetCode string
		targetTime time.Time
	}

	location, _ := time.LoadLocation("Asia/Tokyo")

	tests := map[string]struct {
		in   args
		want res
	}{
		"startTimeと同じ時間は対象となる": {
			in: args{
				code:      "code1",
				startTime: time.Date(2021, 1, 1, 0, 0, 0, 0, location),
				endTime:   time.Date(2021, 1, 1, 1, 0, 0, 0, location),
				record: []string{
					"2021-01-01 00:00:00 +0900 JST", "code1", "100",
				},
			},
			want: res{
				isTarget:   true,
				targetCode: "code1",
				targetTime: time.Date(2021, 1, 1, 0, 0, 0, 0, location),
			},
		},
		"startTimeより前の時間は対象とならない": {
			in: args{
				code:      "code1",
				startTime: time.Date(2021, 1, 1, 1, 0, 0, 0, location),
				endTime:   time.Date(2021, 1, 1, 2, 0, 0, 0, location),
				record: []string{
					"2021-01-01 00:59:59 +0900 JST", "code1", "100",
				},
			},
			want: res{
				isTarget:   false,
				targetCode: "",
				targetTime: time.Time{},
			},
		},
		"endTimeと同じ時間は対象とならない": {
			in: args{
				code:      "code1",
				startTime: time.Date(2021, 1, 1, 1, 0, 0, 0, location),
				endTime:   time.Date(2021, 1, 1, 2, 0, 0, 0, location),
				record: []string{
					"2021-01-01 02:00:00 +0900 JST", "code1", "100",
				},
			},
			want: res{
				isTarget:   false,
				targetCode: "",
				targetTime: time.Time{},
			},
		},
		"codeが違う場合は対象とならない": {
			in: args{
				code:      "code1",
				startTime: time.Date(2021, 1, 1, 1, 0, 0, 0, location),
				endTime:   time.Date(2021, 1, 1, 2, 0, 0, 0, location),
				record: []string{
					"2021-01-01 01:00:00 +0900 JST", "code2", "100",
				},
			},
			want: res{
				isTarget:   false,
				targetCode: "",
				targetTime: time.Time{},
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gotIsTarget, gotTargetCode, gotTargetTime, err := isTargetRecord(tt.in.code, tt.in.startTime, tt.in.endTime, tt.in.record)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if gotIsTarget != tt.want.isTarget {
				t.Errorf("isTargetRecord() isTarget = %v, want %v", gotIsTarget, tt.want.isTarget)
			}

			if gotTargetCode != tt.want.targetCode {
				t.Errorf("isTargetRecord() targetCode = %v, want %v", gotTargetCode, tt.want.targetCode)
			}

			if gotTargetTime.Unix() != tt.want.targetTime.Unix() {
				t.Errorf("isTargetRecord() targetTime = %v, want %v", gotTargetTime, tt.want.targetTime)
			}
		})
	}
}
