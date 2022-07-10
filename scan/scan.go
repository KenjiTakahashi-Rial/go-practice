package scan

import (
	"fmt"
	"strconv"
)

func ScanInt(prompt string, min, max int) int {
	hasRange := min < max
	fmt.Print(prompt)

	for {
		var response string
		fmt.Scanln(&response)
		n, err := strconv.Atoi(response)

		if err != nil {
			fmt.Printf("Invalid number. %s", prompt)
		} else if hasRange && (n < min || n > max) {
			fmt.Printf("Number must be between %d and %d. %s", min, max, prompt)
		} else {
			return n
		}
	}
}
