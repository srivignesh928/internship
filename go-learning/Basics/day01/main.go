package main

import "fmt"

// functions
func add(a int, b int) int {
	return a + b
}

//variable declaration nd condition nd loop
func main() {
	name := "srivignesh"
	age := 20
	rollno := 131

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Roll No:", rollno)

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	for i := 1; i < 3; i++ {
		fmt.Println("Count:", i)
	}

	sum := add(12330954, 4839818)
	fmt.Println("Sum:", sum)
}