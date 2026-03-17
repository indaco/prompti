package input

import "testing"

func TestSetDefaultPlaceholderMsg(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{
			name: "no initial value",
			cfg:  Config{Placeholder: "Enter name"},
			want: "Enter name",
		},
		{
			name: "with initial value",
			cfg:  Config{Placeholder: "Enter name", Initial: "Alice"},
			want: "Enter name (Default: Alice)",
		},
		{
			name: "empty placeholder no initial",
			cfg:  Config{},
			want: "",
		},
		{
			name: "empty placeholder with initial",
			cfg:  Config{Initial: "Bob"},
			want: " (Default: Bob)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setDefaultPlaceholderMsg(&tt.cfg); got != tt.want {
				t.Errorf("setDefaultPlaceholderMsg() = %q, want %q", got, tt.want)
			}
		})
	}
}
