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

func getBudgets(budgets []accounting.Budget) {
	account.GetBudgets(budgets)
}

func TestNoBudgets(t *testing.T) {
	beforeEach()
	start, end := createPeriod("2021-04-02", "2021-04-02")
	assert.Equal(t, .0, account.TotalAmount(start, end))
}

func TestPeriodInWholeMonth(t *testing.T) {
	beforeEach()
	var budgets []accounting.Budget
	budgets = append(budgets, accounting.Budget{YearMonth: "202104", Amount: 30})
	getBudgets(budgets)
	start, end := createPeriod("2021-04-01", "2021-04-30")
	assert.Equal(t, 30.0, account.TotalAmount(start, end))
}

func TestPeriodInsideBudgetMonth(t *testing.T) {
	beforeEach()
	var budgets []accounting.Budget
	budgets = append(budgets, accounting.Budget{YearMonth: "202104", Amount: 30})
	getBudgets(budgets)
	start, end := createPeriod("2021-04-02", "2021-04-02")
	assert.Equal(t, 1.0, account.TotalAmount(start, end))
}

func TestPeriodNoOverlapBudgetFirstDay(t *testing.T) {
	beforeEach()
	var budgets []accounting.Budget
	budgets = append(budgets, accounting.Budget{YearMonth: "202104", Amount: 30})
	getBudgets(budgets)
	start, end := createPeriod("2021-03-30", "2021-03-30")
	assert.Equal(t, 0.0, account.TotalAmount(start, end))
}