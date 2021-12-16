package shopping

import (
	"fmt"
	"testing"
	"time"
)

func TestToEbayDateTime(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "2021/11/27T00/28/30.123",
			want:  "2021-11-27T00:28:30.123Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			dt, _ := time.Parse("2006/01/02T15/04/05.000", tt.input)
			fmt.Println(dt)
			fmt.Println(ToEbayDateTime(dt))
			if got := ToEbayDateTime(dt); got != tt.want {
				t.Errorf("ToEbayDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromEbayDateTime(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "2021-11-27T00:28:30.123Z",
			want:  "2021/11/27T00/28/30.123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			dt, _ := FromEbayDateTime(tt.input)
			datetime := dt.Format("2006/01/02T15/04/05.000")
			if got := datetime; got != tt.want {
				t.Errorf("FromEbayDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromEbayDuration(t *testing.T) {
	tests := []struct {
		d     string
		wantD time.Duration
	}{
		{
			d:     "P2DT23H32M51S",
			wantD: time.Hour*24*2 + time.Hour*23 + time.Minute*32 + time.Second*51,
		},
		{
			d:     "P0DT0H0M1S",
			wantD: time.Second,
		},
		{
			d:     "P0DT0H0M0S",
			wantD: time.Second * 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.d, func(t *testing.T) {
			curDuration := FromEbayDuration(tt.d)
			if tt.wantD != curDuration {
				t.Errorf("wanted: %d , but got %d", tt.wantD, curDuration)
			}

		})
	}
}
