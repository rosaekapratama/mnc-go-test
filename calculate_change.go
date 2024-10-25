package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	unitBill = "bill"
	unitCoin = "coin"
)

var (
	availableDenoms = []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
)

func printCalculateChangeHelp() {
	fmt.Println("Usage:")
	fmt.Println("  mnc-go-test calculate-change [total-purchase] [amount-paid]")
	fmt.Println("")
	fmt.Println("Description:")
	fmt.Println("  Calculates and outputs the change to be returned to the customer.")
	fmt.Println("  The change is rounded down to the nearest hundred.")
	fmt.Println("  If the amount paid is less than the total purchase, an error message is returned.")
	fmt.Println("")
	fmt.Println("Arguments:")
	fmt.Println("  total-purchase    Total purchase amount (in Rp).")
	fmt.Println("  amount-paid       Amount paid by the customer (in Rp).")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("  mnc-go-test calculate-change 700649 800000")
	fmt.Println("  Output: Change to be given by the cashier: Rp 99.351")
	fmt.Println("          Rounded down to: Rp 99.300")
	fmt.Println("          Denomination breakdown:")
	fmt.Println("          1 sheet of Rp 50.000")
	fmt.Println("          2 sheets of Rp 20.000")
	fmt.Println("          1 sheet of Rp 5.000")
	fmt.Println("          2 coins of Rp 2.000")
	fmt.Println("          1 coin of Rp 200")
	fmt.Println("          1 coin of Rp 100")
}

// calculateChange Function to calculate change and provide the breakdown of denomination
func calculateChange(totalPurchase, amountPaid int) {
	if amountPaid < totalPurchase {
		fmt.Println("False, insufficient payment")
		return
	}

	// Calculate change
	change := amountPaid - totalPurchase
	// Round down to the nearest hundred
	roundedChange := (change / 100) * 100

	// Calculate the breakdown of availableDenoms
	changeDetails := make(map[int]int)
	for _, denom := range availableDenoms {
		for roundedChange >= denom {
			changeDetails[denom]++
			roundedChange -= denom
		}
	}

	// Get the keys (changeDenoms) from the map
	changeDenoms := make([]int, 0, len(changeDetails))
	for denom := range changeDetails {
		changeDenoms = append(changeDenoms, denom)
	}

	// Sort the keys in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(changeDenoms)))

	// Output the results
	fmt.Printf("Change to be given by the cashier: Rp %s,\n", formatDenom(change))
	fmt.Printf("Rounded down to: Rp %s\n", formatDenom((change/100)*100))
	fmt.Println("Denomination breakdown:")
	for _, changeDenom := range changeDenoms {
		count := changeDetails[changeDenom]
		if count > 0 {
			// Decide unit bill or coin
			unit := unitBill
			if changeDenom < 1000 {
				unit = unitCoin
			}

			// Add plural or not
			if count > 1 {
				unit += "s"
			}

			// Format changeDenom with period for thousands
			formattedDenom := formatDenom(changeDenom)
			fmt.Printf("%d %s of Rp %s\n", count, unit, formattedDenom)
		}
	}
}

// formatDenom Helper function to format the denomination with dot separators
func formatDenom(denom int) string {
	s := fmt.Sprintf("%d", denom)
	// Insert dot as denom separator
	if len(s) > 3 {
		// Split the string and insert denom separator
		return strings.Join([]string{s[:len(s)-3], s[len(s)-3:]}, ".")
	}
	return s
}

func execCalculateChange(args []string) {
	if !validateHelpInCommand(args, printCalculateChangeHelp) {
		return
	}

	if len(args) < 4 {
		printCalculateChangeHelp()
		return
	}

	// Try convert arg1 and arg2 to integer
	arg1 := args[2]
	totalPurchase, err := strconv.Atoi(arg1)
	if err != nil {
		fmt.Println("Total purchase must be a valid number:", arg1)
		return
	}
	arg2 := args[3]
	paidAmount, err := strconv.Atoi(arg2)
	if err != nil {
		fmt.Println("Paid amount must be a valid number:", arg2)
		return
	}

	fmt.Printf("Total purchase: Rp %s\n", formatDenom(totalPurchase))
	fmt.Printf("Paid amount: Rp %s\n", formatDenom(paidAmount))
	calculateChange(totalPurchase, paidAmount)
}
