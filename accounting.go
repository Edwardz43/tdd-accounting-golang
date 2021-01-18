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
		firstDay := budget.firstDay()
		lastDay := budget.lastDay()
		if start.After(lastDay) || end.Before(firstDay) {
			return 0
		}
		var overlappingStart time.Time
		if start.After(firstDay) {
			overlappingStart = start
		} else {
			overlappingStart = firstDay
		}
		return end.Sub(overlappingStart).Hours()/24 + 1
	}
	return 0
}

