package stats

import (
	"fmt"

	"github.com/root5427/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   1000,
			Category: "Mobile",
			Status:   "OK",
		},
		{
			ID:       2,
			Amount:   2000,
			Category: "Mobile",
			Status:   "OK",
		},
		{
			ID:       3,
			Amount:   3000,
			Category: "Withdraw",
			Status:   "FAIL",
		},
	}

	fmt.Println(Avg(payments))

	// Output:
	// 1500
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   1000,
			Category: "Mobile",
			Status:   "OK",
		},
		{
			ID:       2,
			Amount:   2000,
			Category: "Mobile",
			Status:   "FAIL",
		},
		{
			ID:       3,
			Amount:   3000,
			Category: "Withdraw",
			Status:   "OK",
		},
	}

	fmt.Println(TotalInCategory(payments, "Mobile"))

	// Output:
	// 1000
}
