package confirm

import "testing"

func TestSetOkButtonLabel(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{"default", Config{}, "Yes"},
		{"custom", Config{OkButtonLabel: "Agree"}, "Agree"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setOkButtonLabel(&tt.cfg); got != tt.want {
				t.Errorf("setOkButtonLabel() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSetCancelButtonLabel(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{"default", Config{}, "No"},
		{"custom", Config{CancelButtonLabel: "Reject"}, "Reject"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setCancelButtonLabel(&tt.cfg); got != tt.want {
				t.Errorf("setCancelButtonLabel() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSetCursor(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{"default", Config{}, ">"},
		{"custom", Config{Cursor: "*"}, "*"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setCursor(&tt.cfg); got != tt.want {
				t.Errorf("setCursor() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSetDivider(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{"default", Config{}, "/"},
		{"custom", Config{Divider: "|"}, "|"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setDivider(&tt.cfg); got != tt.want {
				t.Errorf("setDivider() = %q, want %q", got, tt.want)
			}
		})
	}
}
