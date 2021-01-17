package accounting

type BudgetRepo struct {
	budgets []Budget
}

func (r *BudgetRepo) GetBudgets(budgets []Budget) {
	r.budgets = budgets
}

type Budget struct {
	YearMonth string
	Amount    int
}