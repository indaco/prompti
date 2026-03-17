package choose

import "testing"

func TestSetListHeight(t *testing.T) {
	tests := []struct {
		name       string
		cfg        Config
		numOfItems int
		want       int
	}{
		{
			name:       "zero height no extras",
			cfg:        Config{},
			numOfItems: 5,
			want:       15,
		},
		{
			name:       "zero height with ShowHelp",
			cfg:        Config{ShowHelp: true},
			numOfItems: 5,
			want:       20,
		},
		{
			name:       "zero height with ShowStatusBar",
			cfg:        Config{ShowStatusBar: true},
			numOfItems: 3,
			want:       12,
		},
		{
			name:       "zero height with EnableFiltering",
			cfg:        Config{EnableFiltering: true},
			numOfItems: 4,
			want:       16,
		},
		{
			name:       "zero height with all extras",
			cfg:        Config{ShowHelp: true, ShowStatusBar: true, EnableFiltering: true},
			numOfItems: 2,
			want:       8,
		},
		{
			name:       "custom height ignores numOfItems",
			cfg:        Config{ListHeight: 10},
			numOfItems: 5,
			want:       10,
		},
		{
			name:       "custom height with extras still returns custom",
			cfg:        Config{ListHeight: 7, ShowHelp: true},
			numOfItems: 3,
			want:       7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := setListHeight(&tt.cfg, tt.numOfItems)
			if got != tt.want {
				t.Errorf("setListHeight() = %d, want %d", got, tt.want)
			}
		})
	}
}
