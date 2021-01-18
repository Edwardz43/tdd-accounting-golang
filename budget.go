package accounting

import (
	"github.com/jinzhu/now"
	"time"
)

type BudgetRepo struct {
	budgets []Budget
}

func (r *BudgetRepo) GetBudgets(budgets []Budget) {
	r.budgets = budgets
}

func (r *BudgetRepo) GetAll() []Budget {
	return r.budgets
}

type Budget struct {
	YearMonth string
	Amount    int
}

func (b *Budget) firstDay() time.Time {
	budgetFirstDay, _ := time.Parse("200601", b.YearMonth)
	return budgetFirstDay
}

func (b *Budget) lastDay() time.Time {
	budgetLastDate, _ := time.Parse("200601", b.YearMonth)
	return now.With(budgetLastDate).EndOfMonth().Add(time.Hour * -24).Round(time.Hour)
}
