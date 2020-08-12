package main

import "github.com/wuxiaobu/gotoolkit/utility"

func main() {
	record := []string{"id", "name"}
	_ := utility.CsvWriteLine("./a", record)

}
