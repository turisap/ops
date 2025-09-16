package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(
		formatRussianName("Шакиров Кирилл Сергеевич"),
	)
}
func formatRussianName(fullName string) string {
	parts := strings.Fields(fullName)
	if len(parts) < 3 {
		return fullName // Return original if not in expected format
	}

	lastName := parts[0]
	firstName := string([]rune(parts[1])[0]) + "."
	middleName := string([]rune(parts[2])[0]) + "."
	
	return lastName + " " + firstName + middleName
}
