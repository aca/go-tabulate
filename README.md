# go-tabulate

[![PkgGoDev](https://pkg.go.dev/badge/aca/go-tabulate)](https://pkg.go.dev/aca/go-tabulate)

Simple library to tabulate your slice/array data with [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter).

Use `tabulate` struct tag to customize table header.


- `Print(in interface{})` simply prints slice to stdout in markdown format
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
  ```

  ```
  | NAME | AGE |
  |------|-----|
  | john |   3 |
  | kim  |   3 |
  ```


- `AppendBulk(table *tablewriter.Table, in interface{})` instead of tablewriter's `AppendBulk`
  ```
  table := tablewriter.NewWriter(os.Stdout)
  tabulate.AppendBulk(table, d)
  table.Render()
  ```
  ```
  +-------+-----+
  | NAME  | AGE |
  +-------+-----+
  | john  |   3 |
  | kim   |   3 |
  +-------+-----+
  ```
