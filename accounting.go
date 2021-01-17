package accounting

import (
	"github.com/jinzhu/now"
	"time"
)

type Accounting struct {
	*BudgetRepo
}

func (a Accounting) TotalAmount(start time.Time, end time.Time) float64 {
	budgets := a.BudgetRepo.GetAll()
	if len(budgets) > 0 {
		yearMonth, _ := time.Parse("200601", budgets[0].YearMonth)
		return float64(now.With(yearMonth).EndOfMonth().Day())
	}
	return 0
}
