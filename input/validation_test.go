package input

import (
	"testing"
)

func TestValidateAlphanumeric(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"letters only", "hello", false},
		{"digits only", "123", false},
		{"mixed", "hello123", false},
		{"with spaces", "hello 123", false},
		{"empty string", "", false},
		{"special chars", "hello!", true},
		{"hyphen", "hello-world", true},
		{"underscore", "hello_world", true},
		{"unicode", "caf\u00e9", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAlphanumeric(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAlphanumeric(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if err != nil && err.Error() != ruleErrorMsgMap[Alphanumeric] {
				t.Errorf("ValidateAlphanumeric(%q) error = %q, want %q", tt.input, err.Error(), ruleErrorMsgMap[Alphanumeric])
			}
		})
	}
}

func TestValidateDigits(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"single digit", "5", false},
		{"multiple digits", "12345", false},
		{"with letters", "123abc", true},
		{"negative number", "-1", true},
		{"decimal", "1.5", true},
		{"empty string", "", true},
		{"spaces", "1 2", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDigits(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDigits(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateInteger(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"positive", "42", false},
		{"negative", "-42", false},
		{"zero", "0", false},
		{"float", "3.14", true},
		{"letters", "abc", true},
		{"empty string", "", true},
		{"overflow", "99999999999999999999", true},
		{"leading zero", "007", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInteger(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateInteger(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateFloat(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"integer", "42", false},
		{"positive float", "3.14", false},
		{"negative float", "-2.71", false},
		{"zero", "0.0", false},
		{"scientific notation", "1e10", false},
		{"letters", "abc", true},
		{"empty string", "", true},
		{"mixed", "12.3abc", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFloat(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateFloat(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid email", "user@example.com", false},
		{"with display name", "User <user@example.com>", false},
		{"subdomain", "user@sub.example.com", false},
		{"missing at", "userexample.com", true},
		{"missing domain", "user@", true},
		{"empty string", "", true},
		{"just text", "notanemail", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"http url", "http://example.com", false},
		{"https url", "https://example.com", false},
		{"with path", "https://example.com/path", false},
		{"with query", "https://example.com?q=1", false},
		{"no scheme", "example.com", true},
		{"no host", "http://", true},
		{"empty string", "", true},
		{"just text", "notaurl", true},
		{"ftp scheme", "ftp://files.example.com", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateURL(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateURL(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateUnknownRule(t *testing.T) {
	err := validate("anything", ValidationRule("unknown"))
	if err == nil {
		t.Fatal("validate with unknown rule should return an error")
	}
	if err.Error() != "unknown validation rule" {
		t.Errorf("got error %q, want %q", err.Error(), "unknown validation rule")
	}
}
