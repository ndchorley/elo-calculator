package main

import (
	"fmt"

	elocalculator "example/elocalculator/calculator"
)

func WriteToConsole(line string) {
	fmt.Println(line)
}

func main() {
	elocalculator.EloCalculator(WriteToConsole)
}
