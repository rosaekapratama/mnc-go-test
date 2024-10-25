package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	dateLayout = "2006-01-02"
)

func printLeavePermitHelp() {
	fmt.Println("Usage:")
	fmt.Println("  mnc-go-test leave-permit [total-holidays] [join-date] [leave-date] [leave-duration]")
	fmt.Println("")
	fmt.Println("Description:")
	fmt.Println("  Validates if an employee can take personal leave based on company rules.")
	fmt.Println("  The function checks for eligibility based on joining date, leave date,")
	fmt.Println("  total holidays, and the duration of leave requested.")
	fmt.Println("")
	fmt.Println("Arguments:")
	fmt.Println("  total-holidays     Total number of holidays leave (int).")
	fmt.Println("  join-date          Employee's joining date (yyyy-MM-dd).")
	fmt.Println("  leave-date         Planned leave date (yyyy-MM-dd).")
	fmt.Println("  leave-duration     Number of days the employee wishes to take leave (int).")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("  mnc-go-test leave-permit 7 2021-05-01 2021-11-05 3")
	fmt.Println("  Output: False")
	fmt.Println("  Reason: Because personal leave can only be taken for a maximum of 3 consecutive days.")
}

// leavePermit Function to determine if the employee can take personal leave
func leavePermit(totalHolidays int, joinDate, leaveDate time.Time, leaveDuration int) (bool, string) {
	// Check if leave date < join date
	if leaveDate.Before(joinDate) {
		return false, "Leave date must be greater than join date"
	}

	// Calculate the eligible leave period
	startEligibleLeaveDate := joinDate.Add(180 * 24 * time.Hour)
	endEligibleLeaveDate := time.Date(joinDate.Year(), time.December, 31, 0, 0, 0, 0, time.UTC)

	// Check if it's within 180 days from the join date
	if leaveDate.Before(startEligibleLeaveDate) {
		return false, "Because it has not been 180 days since the employee joined."
	}

	// Calculate total days eligible for personal leave
	totalEligibleLeaveDays := int(endEligibleLeaveDate.Sub(startEligibleLeaveDate).Hours()/24) + 1 // Including the start date

	// Calculate personal leave quota
	personalLeaveQuota := totalEligibleLeaveDays * totalHolidays / 365 // Rounded down

	// Check leave duration
	if leaveDuration > personalLeaveQuota {
		return false, fmt.Sprintf("Because the leave duration exceeds the personal leave quota, quota=%d.", personalLeaveQuota)
	}

	if leaveDuration > 3 {
		return false, "Because personal leave can only be taken for a maximum of 3 consecutive days."
	}

	return true, ""
}

func execLeavePermit(args []string) {
	if !validateHelpInCommand(args, printLeavePermitHelp) {
		return
	}

	if len(args) < 6 {
		printLeavePermitHelp()
		return
	}

	totalHolidaysStr := args[2]
	joinDateStr := args[3]
	leaveDateStr := args[4]
	leaveDurationStr := args[5]

	fmt.Println("Total holidays: ", totalHolidaysStr)
	fmt.Println("Join date: ", joinDateStr)
	fmt.Println("Leave date: ", leaveDateStr)
	fmt.Println("Leave duration: ", leaveDurationStr)

	// Parse total holidays and leave duration
	totalHolidays, err := strconv.Atoi(totalHolidaysStr)
	if err != nil {
		fmt.Printf("Total holidays must be an integer")
		return
	}
	leaveDuration, err := strconv.Atoi(leaveDurationStr)
	if err != nil {
		fmt.Printf("Leave duration must be an integer")
		return
	}

	// Parse the join date and leave date
	joinDate, err := time.Parse(dateLayout, joinDateStr)
	if err != nil {
		fmt.Printf("Invalid join date format '%s', valid format is yyyy-MM-dd", joinDateStr)
		return
	}
	leaveDate, err := time.Parse(dateLayout, leaveDateStr)
	if err != nil {
		fmt.Printf("Invalid leave date format '%s', valid format is yyyy-MM-dd", leaveDateStr)
		return
	}

	isValid, reason := leavePermit(totalHolidays, joinDate, leaveDate, leaveDuration)
	fmt.Println("Eligibility: ", isValid)
	if reason != "" {
		fmt.Println("Reason: ", reason)
	}
}
