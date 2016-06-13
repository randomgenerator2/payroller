package main

import (
	"fmt"
	"os"
	"strings"
)

var nameIndex, idIndex, dateIndex, startIndex, endIndex int

func main() {
	data := ReadFile(os.Args[1])
	header := strings.Split(data[0], ",")
	HeaderIndexes(header)
	employees, shifts := ReadData(data, header)
	dateshifts := CreateAllDateshifts(employees, shifts)

	for _, p := range employees {
		for _, ds := range dateshifts[p] {
			ds.CalculatePayRoll()
		}
	}
	wages := make([]float64, len(employees))
	for i, p := range employees {
		for _, ds := range dateshifts[p] {
			wages[i] += ds.payroll
		}
	}
	for i, p := range employees {
		fmt.Printf("%s wage is $%.2f\n", p.name, round(wages[i], 2))
	}

}
