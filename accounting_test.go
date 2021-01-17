package account_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Accounting struct {

}

func (a Accounting) totalAmount(start time.Time, end time.Time) float64 {
	return 0
}

func createPeriod(startStr string, endStr string) (time.Time, time.Time) {
	start, _ := time.Parse("2006-01-02", startStr)
	end, _ := time.Parse("2006-01-02", endStr)
	return start, end
}

func TestNoBudgets(t *testing.T) {
	account := new(Accounting)
	start, end := createPeriod("2021-04-02", "2021-04-02")
	assert.Equal(t, .0, account.totalAmount(start, end))
}

