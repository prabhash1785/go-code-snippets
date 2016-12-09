package main

import (
	"fmt"
)

func Sqrt(x, epsilon float64) float64 {
	z := float64(1)
	
	for z < x {
		fmt.Println("Intermediate sqrt value = ", z)
                
		// calculate next value of z using Newton's formula
		prev := z
		a := z * z
		b := a - x
		c := b / (2 * z)
		
		z = z - c 
		
		diff := z - prev
		
		// return as soon results are less than or equal to epsilon
		if diff >= 0 && diff <= epsilon {
			return z
		}
	}
	
	return z
}

func main() {
	fmt.Println("Square Root = ", Sqrt(2, 0.0001))
}

