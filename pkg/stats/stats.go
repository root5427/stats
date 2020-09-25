package stats

import (
	"github.com/root5427/bank/v2/pkg/types"
)

// Avg calcs avg payment sum
func Avg(payments []types.Payment) types.Money {
	var avg, sum types.Money = 0, 0
	var exc = 0
	for _, p := range payments {
		if p.Status == types.StatusFail {
			exc++
			continue
		}
		sum += p.Amount
	}
	avg = sum / types.Money(len(payments)-exc)
	return avg
}

// TotalInCategory calcs sum of payments in a given category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money = 0
	for _, p := range payments {
		if p.Status == types.StatusFail {
			continue
		}
		if p.Category == category {
			sum += p.Amount
		}
	}
	return sum
}

// FilterByCategory returns payments in a given category
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, p := range payments {
		if p.Category == category {
			filtered = append(filtered, p)
		}
	}

	return filtered
}

// CategoriesTotal returns sum of payments in a category
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, p := range payments {
		categories[p.Category] += p.Amount
	}

	return categories
}
