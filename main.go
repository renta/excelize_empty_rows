package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("test_20.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.Rows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var i int
	for rows.Next() {
		i++
		row, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(fmt.Sprintf("Row № %d: ", i))
		for _, colCell := range row {
			fmt.Print(colCell, " ")
		}
		fmt.Println()
	}
	if err = rows.Close(); err != nil {
		fmt.Println(err)
	}
}