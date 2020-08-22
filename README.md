# go-tabulate

[![PkgGoDev](https://pkg.go.dev/badge/aca/go-tabulate)](https://pkg.go.dev/aca/go-tabulate)

Simple library to tabulate your slice/array data with [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter).

- `Print(in interface{})` simply prints slice to stdout
  ```
  type Human struct {
    Username string `tabulate:"name"`
    Age      int   
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
  ```

- `Fprint(table *tablewriter.Table, in interface{})` to use customized tablewriter
  ```
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
  ```
