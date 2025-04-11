package main

import (
	"fmt"

	elocalculator "example/elocalculator/calculator"
)

func main() {
	writeToConsole := func(line string) { fmt.Println(line) }
	elocalculator.EloCalculator(writeToConsole)
}
