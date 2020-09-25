package stats

import (
	"reflect"
	"testing"

	"github.com/root5427/bank/v2/pkg/types"
)

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "mobile"},
		{ID: 3, Category: "food"},
		{ID: 4, Category: "withdraw"},
		{ID: 5, Category: "transfer"},
	}
	result := FilterByCategory(payments, "epay")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "mobile"},
		{ID: 3, Category: "food"},
		{ID: 4, Category: "withdraw"},
		{ID: 5, Category: "transfer"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
	}

	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "mobile"},
		{ID: 3, Category: "mobile"},
		{ID: 4, Category: "withdraw"},
		{ID: 5, Category: "transfer"},
	}
	expected := []types.Payment{
		{ID: 2, Category: "mobile"},
		{ID: 3, Category: "mobile"},
	}

	result := FilterByCategory(payments, "mobile")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 10000},
		{ID: 2, Category: "mobile", Amount: 10000},
		{ID: 3, Category: "mobile", Amount: 10000},
		{ID: 4, Category: "withdraw", Amount: 10000},
		{ID: 5, Category: "transfer", Amount: 10000},
	}
	expected := map[types.Category]types.Money{
		"auto":     10000,
		"mobile":   20000,
		"withdraw": 10000,
		"transfer": 10000,
	}

	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 10000, Status: types.StatusOk},
		{ID: 2, Category: "mobile", Amount: 10000, Status: types.StatusOk},
		{ID: 3, Category: "mobile", Amount: 10000, Status: types.StatusOk},
		{ID: 4, Category: "withdraw", Amount: 10000, Status: types.StatusOk},
		{ID: 5, Category: "transfer", Amount: 10000, Status: types.StatusOk},
		{ID: 6, Category: "transfer", Amount: 10000, Status: types.StatusFail},
		{ID: 7, Category: "transfer", Amount: 20000, Status: types.StatusOk},
	}
	expected := map[types.Category]types.Money{
		"auto":     10000,
		"mobile":   10000,
		"withdraw": 10000,
		"transfer": 15000,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto":     10,
		"food":     20,
		"withdraw": 30,
	}
	second := map[types.Category]types.Money{
		"auto":   10,
		"food":   25,
		"mobile": 5,
	}
	want := map[types.Category]types.Money{
		"auto":     0,
		"food":     5,
		"mobile":   5,
		"withdraw": 30,
	}
	got := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
