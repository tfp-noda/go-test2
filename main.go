package main

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println("main start")
	now := time.Now()

	// open template excel
	file, _ := excelize.OpenFile("./template.xlsx")

	// remove row
	for i := 0; i < 10000; i++ {
		file.RemoveRow("Sheet1", 1)
	}

	// output
	file.SaveAs("output.xlsx")

	fmt.Println("processing time: ", time.Since(now).Milliseconds(), " ms")
	fmt.Println("main end")
}
