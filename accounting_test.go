package accounting_test

import (
	accounting "Edwardz43/tdd-accounting-golang"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createPeriod(startStr string, endStr string) (time.Time, time.Time) {
	start, _ := time.Parse("2006-01-02", startStr)
	end, _ := time.Parse("2006-01-02", endStr)
	return start, end
}

func TestNoBudgets(t *testing.T) {
	account := new(accounting.Accounting)
	start, end := createPeriod("2021-04-02", "2021-04-02")
	assert.Equal(t, .0, account.TotalAmount(start, end))
}

