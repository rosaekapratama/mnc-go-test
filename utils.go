package main

import (
	"fmt"
	"strings"
)

func validateHelpInCommand(args []string, printHelpFunc func()) (valid bool) {
	if len(args) < 3 {
		printHelpFunc()
		return
	}

	arg1 := args[2]
	switch arg1 {
	case flagHelpShort:
		printHelpFunc()
		return
	case flagHelpLong:
		printHelpFunc()
		return
	}

	// Validate flag help
	if strings.HasPrefix(arg1, "-") {
		fmt.Println("Unknown flag:", arg1)
		printHelpFunc()
		return
	}

	return true
}
