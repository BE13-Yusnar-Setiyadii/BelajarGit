package main

import "fmt"

func main() {
	var T float64 = 20
	var r float64 = 4
	var Pi float64 = 3.14
	Lp := (2 * Pi * r * r) + (2 * Pi * r * T)
	fmt.Println(Lp)
}
