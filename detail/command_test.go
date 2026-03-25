package detail

import "testing"

func TestSetCollapsedIndicator(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{"default", Config{}, "\u25b6"},
		{"custom", Config{CollapsedIndicator: ">"}, ">"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setCollapsedIndicator(&tt.cfg); got != tt.want {
				t.Errorf("setCollapsedIndicator() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSetExpandedIndicator(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{"default", Config{}, "\u25bc"},
		{"custom", Config{ExpandedIndicator: "v"}, "v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setExpandedIndicator(&tt.cfg); got != tt.want {
				t.Errorf("setExpandedIndicator() = %q, want %q", got, tt.want)
			}
		})
	}
}
