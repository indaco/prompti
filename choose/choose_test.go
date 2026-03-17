package choose

import "testing"

func TestGetItemsKeys(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
		want  []string
	}{
		{"nil input", nil, []string{}},
		{"empty slice", []Item{}, []string{}},
		{"single item", []Item{{Name: "a", Desc: "desc-a"}}, []string{"a"}},
		{"multiple items", []Item{
			{Name: "x", Desc: "dx"},
			{Name: "y", Desc: "dy"},
			{Name: "z", Desc: "dz"},
		}, []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetItemsKeys(tt.items)
			if len(got) != len(tt.want) {
				t.Fatalf("len = %d, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("index %d = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestToItems(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []Item
	}{
		{"nil input", nil, []Item{}},
		{"empty slice", []string{}, []Item{}},
		{"single string", []string{"foo"}, []Item{{Name: "foo", Desc: "foo"}}},
		{"multiple strings", []string{"a", "b"}, []Item{
			{Name: "a", Desc: "a"},
			{Name: "b", Desc: "b"},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToItems(tt.input)
			if len(got) != len(tt.want) {
				t.Fatalf("len = %d, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if got[i].Name != tt.want[i].Name || got[i].Desc != tt.want[i].Desc {
					t.Errorf("index %d = %+v, want %+v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestItem_FilterValue(t *testing.T) {
	item := Item{Name: "test-name", Desc: "test-desc"}
	if got := item.FilterValue(); got != "test-name" {
		t.Errorf("FilterValue() = %q, want %q", got, "test-name")
	}
}

func TestItem_String(t *testing.T) {
	item := Item{Name: "test-name", Desc: "test-desc"}
	if got := item.String(); got != "test-name" {
		t.Errorf("String() = %q, want %q", got, "test-name")
	}
}
