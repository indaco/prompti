package choose_test

import (
	"fmt"

	"github.com/indaco/prompti/choose"
)

func ExampleGetItemsKeys() {
	items := []choose.Item{
		{Name: "Go", Desc: "A compiled language"},
		{Name: "Rust", Desc: "A systems language"},
		{Name: "Python", Desc: "A scripting language"},
	}
	keys := choose.GetItemsKeys(items)
	fmt.Println(keys)
	// Output:
	// [Go Rust Python]
}

func ExampleToItems() {
	items := choose.ToItems([]string{"apple", "banana", "cherry"})
	for _, item := range items {
		fmt.Printf("Name=%s Desc=%s\n", item.Name, item.Desc)
	}
	// Output:
	// Name=apple Desc=apple
	// Name=banana Desc=banana
	// Name=cherry Desc=cherry
}

func ExampleItem_FilterValue() {
	item := choose.Item{Name: "Go", Desc: "A compiled language"}
	fmt.Println(item.FilterValue())
	// Output:
	// Go
}
