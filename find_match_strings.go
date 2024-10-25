package main

import (
	"fmt"
	"strconv"
	"strings"
)

func printFindMatchStringHelp() {
	fmt.Println("Usage:")
	fmt.Println("  mnc-go-test find-match-string [total] [str-1] [str-2] ... [str-n]")
	fmt.Println("")
	fmt.Println("Description:")
	fmt.Println("  Finds and outputs the indexes of matching strings from the provided list.")
	fmt.Println("  The comparison is case-insensitive, and only the first matching set is displayed.")
	fmt.Println("")
	fmt.Println("Arguments:")
	fmt.Println("  total               Number of strings to compare")
	fmt.Println("  str-1, str-2, ...   List of strings to be compared")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("  mnc-go-test find-match-string 4 abcd acbd aaab acbd")
	fmt.Println("  Output: 2 4")
}

// findMatchStrings
// Use map to find index of first same string rather than use nested array
// The process time can be shortened to O(n) rather than O(n^2)
func findMatchStrings(n int, stringsArr []string) []int {
	// Map to find first string that has same value case-insensitive
	strMap := make(map[string]int)

	// First find string which has same value
	var sameStr string
	sameStrIdx := make([]int, 0)

	// Loop to check each string
	for i := 0; i < n; i++ {
		// Convert string to lowercase for case-insensitive comparison\
		idxPlusOne := i + 1
		lowerStr := strings.ToLower(stringsArr[i])

		// If same string already found,
		// then compare it and append the index
		if sameStr != "" && sameStr == lowerStr {
			sameStrIdx = append(sameStrIdx, idxPlusOne)
			continue
		}

		// If the string is already in the map,
		// assign to sameStr and sameStrIdx
		if idx, found := strMap[lowerStr]; sameStr == "" && found {
			sameStr = lowerStr
			sameStrIdx = append(sameStrIdx, idx)
			sameStrIdx = append(sameStrIdx, idxPlusOne)
			continue
		}

		// If not found, add the string to the map
		strMap[lowerStr] = idxPlusOne
	}

	return sameStrIdx
}

func execFindMatchString(args []string) {
	if !validateHelpInCommand(args, printFindMatchStringHelp) {
		return
	}

	// Try convert arg1 to integer
	arg1 := args[2]
	total, err := strconv.Atoi(arg1)
	if err != nil {
		fmt.Println("Total must be a valid number:", arg1)
		return
	}

	// Check array of string length
	// if length < given total, reject
	if len(args) < 3+int(total) {
		fmt.Printf("Length of given string array must be equals or greater than given total\n"+
			"total=%v, arrayLength=%v", total, len(args)-3)
		return
	}

	indexes := findMatchStrings(total, args[3:])
	fmt.Printf("Input: %v\n", args[2:])
	if len(indexes) > 0 {
		fmt.Printf("Output: %v", indexes)
	} else {
		fmt.Printf("Output: %v", false)
	}
}
