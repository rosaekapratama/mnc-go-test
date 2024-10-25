package main

import (
	"fmt"
)

func printValidateStringHelp() {
	fmt.Println("Usage:")
	fmt.Println("  mnc-go-test validate-string [input]")
	fmt.Println("")
	fmt.Println("Description:")
	fmt.Println("  Validates the given input string based on the following rules:")
	fmt.Println("  - The string can only contain the characters <>{}[]")
	fmt.Println("  - Every opening character must have a matching closing character.")
	fmt.Println("  - No closing character should appear before its corresponding opening character (e.g., ']<>' is invalid).")
	fmt.Println("  - Brackets cannot enclose mismatched bracket types (e.g., '<[>]' is invalid).")
	fmt.Println("  - Brackets can be nested inside others as long as they are fully enclosed (e.g., '<[]{<>}>' is valid).")
	fmt.Println("  - The length of the string must be between 1 and 4096 characters.")
	fmt.Println("")
	fmt.Println("Arguments:")
	fmt.Println("  input    The string to be validated.")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("  mnc-go-test validate-string \"{{[<>[{{}}]]}}\"")
	fmt.Println("  Output: Valid")
	fmt.Println("")
	fmt.Println("  mnc-go-test validate-string \"][\"")
	fmt.Println("  Output: Invalid")
}

func validateString(input string) bool {
	// Check the length of the input
	if len(input) < 1 || len(input) > 4096 {
		return false
	}

	// Create a stack to store opening brackets
	var stack []rune

	// Iterate over the input
	for _, char := range input {
		switch char {
		case '<', '{', '[':
			// Push opening brackets onto the stack
			stack = append(stack, char)
		case '>':
			// Check if the stack is empty or if the top of the stack doesn't match
			if len(stack) == 0 || stack[len(stack)-1] != '<' {
				return false
			}
			// If open bracket match the close bracket
			// then pop the top of the stack
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			// Invalid character
			return false
		}
	}

	// Ensure the stack is empty at the end (all opening brackets have been closed)
	return len(stack) == 0
}

func execValidateString(args []string) {
	if !validateHelpInCommand(args, printValidateStringHelp) {
		return
	}

	arg := args[2]
	fmt.Printf("Input: \"%s\"\n", arg)
	if validateString(arg) {
		fmt.Println("Output: Valid")
	} else {
		fmt.Println("Output: Invalid")
	}
}
