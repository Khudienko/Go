package main

import (
	"fmt"
	"math"
)

func main() {
printNumbers(100)
}

func squareRoot(n int) int {
	s := math.Sqrt(float64(n))
	return int (s)
}

func printNumbers (n int){
	temp:=squareRoot(n)
	for i:=1;i<=temp;i++{
		if i<temp {
			fmt.Print(i,", ")
		}else {
			fmt.Print(i,".")
		}
	}
}
