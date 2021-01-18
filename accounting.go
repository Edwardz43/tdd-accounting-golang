package accounting

import (
	"time"
)

type Accounting struct {
	*BudgetRepo
}

func (a *Accounting) TotalAmount(start time.Time, end time.Time) float64 {
	budgets := a.GetAll()
	if len(budgets) > 0 {
		budget := budgets[0]
		if start.After(budget.lastDay()) || end.Before(budget.firstDay()) {
			return 0
		}
		var overlappingStart time.Time
		if start.After(budget.firstDay()) {
			overlappingStart = start
		} else {
			overlappingStart = budget.firstDay()
		}
		return end.Sub(overlappingStart).Hours()/24 + 1
	}
	return 0
}

