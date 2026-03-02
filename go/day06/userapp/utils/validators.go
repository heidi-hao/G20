package utils

import (
	"strings"
)

func ValidateEmail(email string) bool {
	return strings.Contains(email, "@")
}

func ValidateAge(age int) bool {
	return age > 0 && age < 150
}
