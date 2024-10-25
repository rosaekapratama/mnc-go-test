# MNC Go Test CLI

This command-line application provides various utilities for string validation, leave management, and change calculation.

Windows:

```mnc-go-test.exe [command] [arguments]```

Linux:

```mnc-go-test [command] [arguments]```

## Available Commands:
- **find-match-string**: Find first matching strings index from the given list
- **calculate-change**: Calculate and output change amount from the given total purchase and paid amount
- **validate-string**: Validate strings without using regex
- **leave-permit**: Check employee leave eligibility

## Flags:
- `-h, --help`: Show help for a command

## Additional Info:
Use `mnc-go-test [command] --help` for more information about a command.

---

## Commands

### 1. Find Match String

#### Usage:
Windows:

```mnc-go-test.exe find-match-string [total] [str-1] [str-2] ... [str-n]```

Linux:

```mnc-go-test find-match-string [total] [str-1] [str-2] ... [str-n]```


#### Description:
Finds and outputs the indexes of matching strings from the provided list. The comparison is case-insensitive, and only the first matching set is displayed.

#### Arguments:
- `total`: Number of strings to compare.
- `str-1, str-2, ...`: List of strings to be compared.

#### Example:
```
.\mnc-go-test.exe find-match-string 4 abcd acbd aaab acbd 
Output: 2 4
```

---

### 2. Calculate Change

#### Usage:
Windows:

```mnc-go-test.exe calculate-change [total-purchase] [amount-paid]```

Linux:

```mnc-go-test calculate-change [total-purchase] [amount-paid]```


#### Description:
Calculates and outputs the change to be returned to the customer. The change is rounded down to the nearest hundred. If the amount paid is less than the total purchase, an error message is returned.

#### Arguments:
- `total-purchase`: Total purchase amount (in Rp).
- `amount-paid`: Amount paid by the customer (in Rp).

#### Example:
```
.\mnc-go-test.exe calculate-change 700649 800000
Output: 
Change to be given by the cashier: Rp 99.351 
Rounded down to: Rp 99.300 
Denomination breakdown: 
1 sheet of Rp 50.000 
2 sheets of Rp 20.000 
1 sheet of Rp 5.000 
2 coins of Rp 2.000 
1 coin of Rp 200 
1 coin of Rp 100
```

---

### 3. Validate String

#### Usage:
Windows:

```mnc-go-test.exe validate-string [input]```

Linux:

```mnc-go-test validate-string [input]```


#### Description:
Validates the given input string based on the following rules:
- The string can only contain the characters `<>{}[]`.
- Every opening character must have a matching closing character.
- No closing character should appear before its corresponding opening character (e.g., ']<>' is invalid).
- Brackets cannot enclose mismatched bracket types (e.g., '<[>]' is invalid).
- Brackets can be nested inside others as long as they are fully enclosed (e.g., '<[]{<>}>' is valid).
- The length of the string must be between 1 and 4096 characters.

#### Arguments:
- `input`: The string to be validated.

#### Example:
```
.\mnc-go-test.exe validate-string "{{[<>[{{}}]]}}"
Output: Invalid
```

---

### 4. Leave Permit

#### Usage:

Windows:

```mnc-go-test.exe leave-permit [total-holidays] [join-date] [leave-date] [leave-duration]```

Linux:

```mnc-go-test leave-permit [total-holidays] [join-date] [leave-date] [leave-duration]```


#### Description:
Validates if an employee can take personal leave based on company rules. The function checks for eligibility based on joining date, leave date, total holidays, and the duration of leave requested.

#### Arguments:
- `total-holidays`: Total number of holidays leave (int).
- `join-date`: Employee's joining date (yyyy-MM-dd).
- `leave-date`: Planned leave date (yyyy-MM-dd).
- `leave-duration`: Number of days the employee wishes to take leave (int).

#### Example:
```
.\mnc-go-test.exe leave-permit 7 2021-05-01 2021-11-05 3
Output: False 
Reason: Because personal leave can only be taken for a maximum of 3 consecutive days.
```