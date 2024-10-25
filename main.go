package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	cmdFindMatchString = "find-match-string"
	cmdCalculateChange = "calculate-change"
	cmdValidateString  = "validate-string"
	cmdLeavePermit     = "leave-permit"
	flagHelpShort      = "-h"
	flagHelpLong       = "--help"
)

func printCmdHelp() {
	fmt.Println("Usage:")
	fmt.Println("  mnc-go-test [command] [arguments]")
	fmt.Println("")
	fmt.Println("Available Commands:")
	fmt.Println("  find-match-string   Find first matching strings index from the given list")
	fmt.Println("  calculate-change    Calculate and output change amount from the given total purchase and paid amount")
	fmt.Println("  validate-string     Validate strings without using regex")
	fmt.Println("  leave-permit        Check employee leave eligibility")
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("  -h, --help          Show help for a command")
	fmt.Println("")
	fmt.Println("Use \"mnc-go-test [command] --help\" for more information about a command.")
}

func main() {
	args := os.Args

	// Print help if needed
	if len(args) < 2 {
		printCmdHelp()
		return
	}
	cmd := args[1]

	switch cmd {
	case flagHelpShort:
		printCmdHelp()
	case flagHelpLong:
		printCmdHelp()
	case cmdFindMatchString:
		execFindMatchString(args)
	case cmdCalculateChange:
		execCalculateChange(args)
	case cmdValidateString:
		execValidateString(args)
	case cmdLeavePermit:
		execLeavePermit(args)
	default:
		if strings.HasPrefix(cmd, "-") {
			fmt.Println("Unknown flag:", cmd)
		} else {
			fmt.Println("Unknown command:", cmd)
		}
		printCmdHelp()
	}
}
