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

// CategoriesAvg returns average sun of payments in the category
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	categoriesCount := map[types.Category]int{}

	for _, p := range payments {
		if p.Status == types.StatusFail {
			continue
		}
		categories[p.Category] += p.Amount
		categoriesCount[p.Category] += 1
	}

	for k := range categories {
		categories[k] = categories[k] / types.Money(categoriesCount[k])
	}

	return categories
}

// PeriodsDynamic compares expenses in categories for 2 periods
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	for k := range second {
		if _, ok := first[k]; ok {
			result[k] = second[k] - first[k]
		} else {
			result[k] = second[k]
		}
	}
	for k := range first {
		if _, ok := second[k]; !ok {
			result[k] = -first[k]
		}
	}
	return result
}
