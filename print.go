package tabulate

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

func AppendBulk(table *tablewriter.Table, in interface{}) error {
	switch reflect.TypeOf(in).Kind() {
	case reflect.Slice, reflect.Array:
		et := reflect.TypeOf(in).Elem()

		// if element is pointer try to dereference
		switch et.Kind() {
		case reflect.Ptr:
			et = et.Elem()
			if et.Kind() != reflect.Struct {
				return errors.New("tabulate: array/slice element must be struct or pointer to struct")
			}
		case reflect.Struct:
		default:
			return errors.New("tabulate: not a array/slice")
		}

		// set header
		header := make([]string, 0)
		for i := 0; i < et.NumField(); i++ {
			f := et.Field(i)
			tag := f.Tag.Get("tabulate")
			if tag == "-" {
				continue
			} else if tag == "" {
				header = append(header, f.Name)
			} else {
				header = append(header, tag)
			}
		}

		table.SetHeader(header)

		// set data
		s := reflect.ValueOf(in)
		for i := 0; i < s.Len(); i++ {
			ev := reflect.Indirect(s.Index(i))
			data := []string{}
			for i := 0; i < ev.NumField(); i++ {
				f := ev.Type().Field(i)
				if f.Tag.Get("tabulate") != "-" {
					data = append(data, fmt.Sprint(ev.Field(i).Interface()))
				}
			}
			table.Append(data)
		}
	default:
		return errors.New("tabulate: not a slice or array")
	}
	return nil
}

func Print(in interface{}) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	err := AppendBulk(table, in)
	if err != nil {
		return err
	}
	table.Render()
	return nil
}
