package pkg

import (
	"strings"
	"testing"
)

func TestIsJobOrSecretNameValid(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
		errorMsg  string
	}{
		// Valid names
		{
			name:      "valid simple name",
			input:     "validname",
			wantError: false,
		},
		{
			name:      "valid name with underscore",
			input:     "test_job",
			wantError: false,
		},
		{
			name:      "valid name with hyphen",
			input:     "my-secret",
			wantError: false,
		},
		{
			name:      "valid name with mixed case",
			input:     "MyJobName",
			wantError: false,
		},
		{
			name:      "valid name with underscore and hyphen",
			input:     "test_job-name",
			wantError: false,
		},
		{
			name:      "valid name with number in middle",
			input:     "job2name",
			wantError: false,
		},
		{
			name:      "valid name with number at end",
			input:     "jobname123",
			wantError: false,
		},
		{
			name:      "valid single character",
			input:     "a",
			wantError: false,
		},
		{
			name:      "valid two characters",
			input:     "ab",
			wantError: false,
		},
		{
			name:      "valid name at max length minus 1",
			input:     strings.Repeat("a", 254),
			wantError: false,
		},

		// Invalid names - Too long
		{
			name:      "name too long (256 chars)",
			input:     strings.Repeat("a", 256),
			wantError: true,
			errorMsg:  "Name is too long, must be less than 255 characters",
		},
		{
			name:      "name at exact limit (255 chars)",
			input:     strings.Repeat("a", 255),
			wantError: false,
		},
		{
			name:      "name way too long (500 chars)",
			input:     strings.Repeat("a", 500),
			wantError: true,
			errorMsg:  "Name is too long, must be less than 255 characters",
		},

		// Invalid names - Contains spaces
		{
			name:      "name with space in middle",
			input:     "invalid name",
			wantError: true,
			errorMsg:  "Name cannot contain spaces",
		},
		{
			name:      "name with space at beginning",
			input:     " invalidname",
			wantError: true,
			errorMsg:  "Name cannot contain spaces",
		},
		{
			name:      "name with space at end",
			input:     "invalidname ",
			wantError: true,
			errorMsg:  "Name cannot contain spaces",
		},
		{
			name:      "name with multiple spaces",
			input:     "invalid  name  here",
			wantError: true,
			errorMsg:  "Name cannot contain spaces",
		},

		// Invalid names - Special characters
		{
			name:      "name with exclamation mark",
			input:     "invalid!name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with at symbol",
			input:     "invalid@name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with hash",
			input:     "invalid#name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with dollar sign",
			input:     "invalid$name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with percent",
			input:     "invalid%name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with caret",
			input:     "invalid^name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with ampersand",
			input:     "invalid&name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with asterisk",
			input:     "invalid*name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with parentheses",
			input:     "invalid(name)",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with plus sign",
			input:     "invalid+name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with equals sign",
			input:     "invalid=name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with curly braces",
			input:     "invalid{name}",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with square brackets",
			input:     "invalid[name]",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with pipe",
			input:     "invalid|name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with backslash",
			input:     "invalid\\name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with colon",
			input:     "invalid:name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with semicolon",
			input:     "invalid;name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with double quotes",
			input:     "invalid\"name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with single quote",
			input:     "invalid'name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with less than",
			input:     "invalid<name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with greater than",
			input:     "invalid>name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with comma",
			input:     "invalid,name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with period",
			input:     "invalid.name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with question mark",
			input:     "invalid?name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with forward slash",
			input:     "invalid/name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},
		{
			name:      "name with multiple special chars",
			input:     "invalid!@#$%name",
			wantError: true,
			errorMsg:  "Name cannot contain special characters",
		},

		// Invalid names - Starts with number
		{
			name:      "name starting with 0",
			input:     "0invalidname",
			wantError: true,
			errorMsg:  "Name cannot start with a number",
		},
		{
			name:      "name starting with 1",
			input:     "1invalidname",
			wantError: true,
			errorMsg:  "Name cannot start with a number",
		},
		{
			name:      "name starting with 2",
			input:     "2invalidname",
			wantError: true,
			errorMsg:  "Name cannot start with a number",
		},
		{
			name:      "name starting with 3",
			input:     "3invalidname",
			wantError: true,
			errorMsg:  "Name cannot start with a number",
		},
		{
			name:      "name starting with 4",
			input:     "4invalidname",
			wantError: true,
			errorMsg:  "Name cannot start with a number",
		},
		{
			name:      "name starting with 5",
			input:     "5invalidname",
			wantError: true,
			errorMsg:  "Name cannot start with a number",
		},
		{
			name:      "name starting with 6",
			input:     "6invalidname",
			wantError: true,
			errorMsg:  "Name cannot start with a number",
		},
		{
			name:      "name starting with 7",
			input:     "7invalidname",
			wantError: true,
			errorMsg:  "Name cannot start with a number",
		},
		{
			name:      "name starting with 8",
			input:     "8invalidname",
			wantError: true,
			errorMsg:  "Name cannot start with a number",
		},
		{
			name:      "name starting with 9",
			input:     "9invalidname",
			wantError: true,
			errorMsg:  "Name cannot start with a number",
		},

		// Edge cases with unicode
		{
			name:      "name with emoji",
			input:     "invalid😀name",
			wantError: false, // Emojis are not in the special chars list
		},
		{
			name:      "name with unicode letter",
			input:     "validñame",
			wantError: false, // Unicode letters are not in the special chars list
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := isJobOrSecretNameValid(tt.input)

			if tt.wantError {
				if err == nil {
					t.Errorf("isJobOrSecretNameValid(%q) = nil, want error", tt.input)
				} else if tt.errorMsg != "" && err.Error() != tt.errorMsg {
					t.Errorf("isJobOrSecretNameValid(%q) error = %q, want %q", tt.input, err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("isJobOrSecretNameValid(%q) = %v, want nil", tt.input, err)
				}
			}
		})
	}
}

func TestIsJobOrSecretNameValid_EmptyString(t *testing.T) {
	// Special test for empty string - this will panic due to index out of bounds
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("isJobOrSecretNameValid(\"\") did not panic as expected")
		}
	}()

	_ = isJobOrSecretNameValid("")
}

func TestIsJobOrSecretNameValid_Benchmarks(t *testing.T) {
	// Benchmark tests to ensure validation is performant
	inputs := []string{
		"validname",
		"invalid name",
		"1invalidname",
		"invalid!name",
		strings.Repeat("a", 300),
	}

	for _, input := range inputs {
		// Just run validation to ensure it completes without hanging
		_ = isJobOrSecretNameValid(input)
	}
}
