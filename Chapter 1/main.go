package main

import (
	"os"
	"fmt"
)

func main() {
	var _, getMulitpleValues = multipleReturnValues()
	fmt.Println(getMulitpleValues)
}

// Simple function that shows os Args from commandline and FMT println
func osArgs() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	fmt.Println("It's over", os.Args[1])
}

func workingWithNumbers() {

	// Statically typed language
	var power int = 9000
	power++

	// Were infering the type here
	higherPower := 9000
	higherPower++

}

func multipleReturnValues() (string, int) {
	return "Hello", 1
}
