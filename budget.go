package accounting

import "time"

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