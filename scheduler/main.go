package main

import (
	"fmt"

	"github.com/anwam/devil-butler/lib"
)

func main() {
	numberString := "123.45"
	value := getFloat64(numberString)
	fmt.Printf("%.2f", value)
}

func getFloat64(s string) float64 {
	return lib.StringToNumber[float64](s)
}
