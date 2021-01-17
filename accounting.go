package accounting

import (
	"time"
)

type Accounting struct {
	*BudgetRepo
}

func (a *Accounting) TotalAmount(start time.Time, end time.Time) float64 {
	budgets := a.BudgetRepo.GetAll()
	if len(budgets) > 0 {
		if end.Before(a.firstDay(budgets)) {
			return 0
		}
		days := end.Sub(start).Hours()/24 + 1
		return days
	}
	return 0
}

func (a *Accounting) firstDay(budgets []Budget) time.Time {
	budgetFirstDay, _ := time.Parse("200601", budgets[0].YearMonth)
	return budgetFirstDay
}
