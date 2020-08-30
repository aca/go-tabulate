package tabulate_test

import (
	"github.com/aca/go-tabulate"
)

func ExampleStructSlice() {
	type Human struct {
		Username string `tabulate:"name"`
		Age      int    // table header will be set to field name by default
		Secret   string `tabulate:"-"` // "-" to ignore certain field
	}

	d := []Human{
		{"john", 3, "password1"},
		{"kim", 3, "password2"},
	}

	tabulate.Print(d)

	// Output:
	// | NAME | AGE |
	// |------|-----|
	// | john |   3 |
	// | kim  |   3 |
}

func ExamplePointerSlice() {
	type Human struct {
		Username string `tabulate:"name"`
		Age      int    `tabulate:"age"`
		Secret   string `tabulate:"-"`
	}

	d := []Human{
		{"john", 3, "password1"},
		{"kim", 3, "password2"},
	}

	tabulate.Print(d)

	// Output:
	// | NAME | AGE |
	// |------|-----|
	// | john |   3 |
	// | kim  |   3 |
}
