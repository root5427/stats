package stats

import (
	"github.com/root5427/bank/pkg/bank/types"
)

// Avg calcs avg payment sum
func Avg(payments []types.Payment) types.Money {
	var avg, sum types.Money = 0, 0
	for _, p := range payments {
		sum += p.Amount
	}
	avg = sum / types.Money(len(payments))
	return avg
}

// TotalInCategory calcs sum of payments in a given category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money = 0
	for _, p := range payments {
		if p.Category == category {
			sum += p.Amount
		}
	}
	return sum
}
