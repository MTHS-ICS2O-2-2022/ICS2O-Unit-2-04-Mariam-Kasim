// Created by: Mariam Kasim
// Created on: March 2023
//
// This program finds the area of a triangle

package main

import "fmt"

func main() {
	// This function finds the area of a triangle
	var base int
	var height int
	var area int

	// input
	fmt.Println("This program finds the area of a triangle.")
	fmt.Println()
	fmt.Print("Enter the base (in): ")
	fmt.Scanln(&base)
	fmt.Print("Enter the height (in): ")
	fmt.Scanln(&height)

	// process
	area = (base * height) / 2

	// output
	fmt.Println()
	fmt.Println("The area is:", area, "in².")
	fmt.Println("\nDone.")
}
