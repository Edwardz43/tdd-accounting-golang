package accounting_test

import (
	accounting "Edwardz43/tdd-accounting-golang"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)
var account *accounting.Accounting

func beforeEach() {
	account = &accounting.Accounting{BudgetRepo: &accounting.BudgetRepo{}}
}

func createPeriod(startStr string, endStr string) (time.Time, time.Time) {
	start, _ := time.Parse("2006-01-02", startStr)
	end, _ := time.Parse("2006-01-02", endStr)
	return start, end
}

func TestNoBudgets(t *testing.T) {
	beforeEach()
	start, end := createPeriod("2021-04-02", "2021-04-02")
	assert.Equal(t, .0, account.TotalAmount(start, end))
}

type Budget struct {
	YearMonth string
	Amount int
}

func TestPeriodInWholeMonth(t *testing.T) {
	beforeEach()
	//account.
	start, end := createPeriod("2021-04-01", "2021-04-30")
	assert.Equal(t, 30.0, account.TotalAmount(start, end))
}