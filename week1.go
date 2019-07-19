package main

import (
	"fmt"
	"os"
)

func evalFormula(first, second int, arimethic string) {
	switch arimethic {
	case "+":
		fmt.Printf("%d + %d = %d\n", first, second, first + second)
	case "-":
		fmt.Printf("%d - %d = %d\n", first, second, first - second)
	case "*":
		fmt.Printf("%d * %d = %d\n", first, second, first * second)
	case "/":
		fmt.Printf("%d / %d = %d\n", first, second, first / second)
	return
	}
}

func checkDivideByZero(arimethic string, secondNum int) bool {
	return secondNum == 0 && arimethic == "/"
}

func checkSyntaxArimethic(arimethic string) bool {
	return arimethic != "+" && arimethic != "-" && arimethic != "*" && arimethic != "/"
}

func main() {
	var first, second int
	var arimethic string
	fmt.Print("Enter formula: ")
	_, err := fmt.Scanf("%d %s %d", &first, &arimethic, &second)
	if err != nil || checkSyntaxArimethic(arimethic) {
		fmt.Println("Error syntax")
		os.Exit(0)
	}
	if checkDivideByZero(arimethic, second) {
		fmt.Println("Divide by zero")
		os.Exit(0)
	}
	evalFormula(first, second, arimethic)
}
