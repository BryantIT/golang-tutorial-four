package main

import "fmt"

func main() {
	fmt.Printf("%T ", 10)
	fmt.Printf("%v ", 10)
	fmt.Printf("%t ", true)

	fmt.Printf("%o ", 1025)

	fmt.Printf("%e ", 3.15487465748876545757)
	fmt.Printf("%f ", 3.15487465748876545757)
	fmt.Printf("%.f ", 3.15487465748876545757)
	fmt.Printf("%.2f ", 3.15487465748876545757)

	fmt.Printf("%s ", "Bryant")
	fmt.Printf("%q ", "Bryant")

}
