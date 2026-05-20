package main

import "fmt"

func main() {
    num1 := 10
    num2 := 5

    fmt.Println("Addition:", num1+num2)
    fmt.Println("Subtraction:", num1-num2)
    fmt.Println("Multiplication:", num1*num2)

    if num2 != 0 {
        fmt.Println("Division:", num1/num2)
    } else {
        fmt.Println("Cannot divide by zero")
    }
}