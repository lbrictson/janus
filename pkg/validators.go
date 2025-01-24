package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

func isJobOrSecretNameValid(name string) error {
	if len(name) > 255 {
		return fmt.Errorf("Name is too long, must be less than 255 characters")
	}
	// It cannot contain any spaces
	if strings.Contains(name, " ") {
		return fmt.Errorf("Name cannot contain spaces")
	}
	// it cannot contain any special characters
	if strings.ContainsAny(name, "!@#$%^&*()+={}[]|\\:;\"'<>,.?/") {
		return fmt.Errorf("Name cannot contain special characters")
	}
	// it cannot start with a number
	if _, err := strconv.Atoi(string(name[0])); err == nil {
		return fmt.Errorf("Name cannot start with a number")
	}
	return nil
}
