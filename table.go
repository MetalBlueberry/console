package console

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func Table(dataSliceOrObject any, columns ...[]string) {
	panic("Not Implemented Yet")
	DefaultConsole.Table(dataSliceOrObject, columns...)
}

var DefaultTableWriter = func() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	return table
}

func (c *Console) Table(dataSliceOrObject any, columns ...[]string) {
	panic("Not Implemented Yet")

	// tw := c.TableWriter()
	// tw.SetHeader()
	// c.TableWriter().AppendBulk()
	// tw.Render()
}

func (c *Console) TableData(dataSliceOrObject any) [][]string {
	panic("Not Implemented Yet")

	var data [][]string
	switch reflect.TypeOf(dataSliceOrObject).Kind() {
	case reflect.Slice:
		// Slice of ...

		s := reflect.ValueOf(dataSliceOrObject)

		data = make([][]string, s.Len())

		for i := 0; i < s.Len(); i++ {
			item := s.Index(i)
			switch item.Kind() {
			case reflect.Slice:
				// Slice of Slices
			default:
				// Slice of objects
				if data[0] == nil {
					data[0] = make([]string, 0)
					data[1] = make([]string, 0)
				}
				data[0] = append(data[0], strconv.Itoa(i))
				data[1] = append(data[1], fmt.Sprintf("%#v", item))
			}
		}
	default:
		// Object
		// iterate over object fields :harold:
	}
	return data
}
