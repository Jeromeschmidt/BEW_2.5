package main

import "fmt"

func main() {
    var birth_year int
    var age int

    fmt.Println("Enter Your Birth Year: ")
    fmt.Scanln(&birth_year)

    fmt.Println("Enter Your Age: ")
    fmt.Scanln(&age)

    fmt.Print("The Year is: ")

    year := birth_year+age

    fmt.Println(year)
}
