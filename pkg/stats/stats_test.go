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
