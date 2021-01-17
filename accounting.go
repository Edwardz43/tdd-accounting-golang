package accounting

import (
	"github.com/jinzhu/now"
	"time"
)

type Accounting struct {
	*BudgetRepo
}

func (a *Accounting) TotalAmount(start time.Time, end time.Time) float64 {
	budgets := a.BudgetRepo.GetAll()
	if len(budgets) > 0 {
		if start.After(a.lastDay(budgets)) || end.Before(a.firstDay(budgets)) {
			return 0
		}
		return end.Sub(start).Hours()/24 + 1
	}
	return 0
}

func (a *Accounting) lastDay(budgets []Budget) time.Time {
	budgetLastDate, _ := time.Parse("200601", budgets[0].YearMonth)
	return now.With(budgetLastDate).EndOfMonth()
}

func (a *Accounting) firstDay(budgets []Budget) time.Time {
	budgetFirstDay, _ := time.Parse("200601", budgets[0].YearMonth)
	return budgetFirstDay
}
