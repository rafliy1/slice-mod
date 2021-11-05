package main

import "fmt"

func main() {
	namaTeman := [...]string{
		"jae",
		"rafli",
		"nida",
		"shafa",
		"rizal",
		"pimen",
		"nadhar",
		"josi",
		"bintang",
		"alpat",
	}

	fmt.Println("namaTeman", namaTeman)

	slicenamaTeman := namaTeman[1:6]

	fmt.Println(slicenamaTeman)
}
