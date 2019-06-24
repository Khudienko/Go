package main

import "fmt"

func main() {
fib(5,10)
}

func fib (from int,till int){
	x,y := 0,1
	for i:=0; i<till; i++ {
		x,y= y,x+y
		if x>=from {
			fmt.Println(x)
		}
	}
}
