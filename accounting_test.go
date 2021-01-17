package account_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Accounting struct {

}

func (a Accounting) totalAmount() float64 {
	return 0
}

func TestNoBudget(t *testing.T) {
	account := new(Accounting)
	assert.Equal(t, .0, account.totalAmount())
}
