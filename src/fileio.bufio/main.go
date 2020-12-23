package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func errorCheck(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	file, err := os.Open("sample.csv")
	errorCheck(err)

	defer file.Close()

	rr := csv.NewReader(bufio.NewReader(file))

	row, err := rr.Read()
	errorCheck(err)

	fmt.Println(row)
	fmt.Println("=============================")

	rows, err := rr.ReadAll()
	errorCheck(err)

	fmt.Println(rows)
	fmt.Println("=============================")

	fmt.Println(rows[5][1], rows[5][2], rows[5][1:3])
	fmt.Println("=============================")

	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s\t", rows[i][j])
		}
		fmt.Println()
	}
	fmt.Println("=============================")
}
