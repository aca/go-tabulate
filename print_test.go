package tabulate_test

import (
	"os"

	"github.com/aca/go-tabulate"
	"github.com/olekukonko/tablewriter"
)

func ExampleStructSlice() {
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
	// +------+-----+
	// | NAME | AGE |
	// +------+-----+
	// | john |   3 |
	// | kim  |   3 |
	// +------+-----+
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
	// +------+-----+
	// | NAME | AGE |
	// +------+-----+
	// | john |   3 |
	// | kim  |   3 |
	// +------+-----+
}

func ExampleFprint() {
	type Human struct {
		Username string `tabulate:"name"`
		Age      int    `tabulate:"age"`
		Secret   string `tabulate:"-"`
	}

	d := []Human{
		{"john", 3, "password1"},
		{"kim", 3, "password2"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetFooter([]string{"TOTAL", "2"})

	tabulate.Fprint(table, d)
	// Output:
  // +-------+-----+
  // | NAME  | AGE |
  // +-------+-----+
  // | john  |   3 |
  // | kim   |   3 |
  // +-------+-----+
  // | TOTAL |  2  |
  // +-------+-----+
}
