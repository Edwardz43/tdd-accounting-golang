package accounting

import "time"

type Accounting struct {
	*BudgetRepo
}

func (a Accounting) TotalAmount(start time.Time, end time.Time) float64 {
	return 0
}
