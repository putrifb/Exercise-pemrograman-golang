package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	values := strings.Fields(calculate)
	if len(values) == 0 {
		return 0.0
	}
	result, _ := strconv.ParseFloat(values[0], 32)
	var operator string
	for i := 1; i < len(values); i++ {
		if i%2 == 0 {
			continue
		}
		operator = values[i]
		num, _ := strconv.ParseFloat(values[i+1], 32)
		switch operator {
		case "+":
			result += num
		case "-":
			result -= num
		case "*":
			result *= num
		case "/":
			result /= num
		}
	}
	return float32(result)
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
