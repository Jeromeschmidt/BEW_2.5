package main

import "fmt"

func main() {
    var sum_of_weights int
    var temp_weight int

    sum_of_weights = 0

    for i := 0; i < 5; i++ {
        fmt.Println("Enter the weight for person: ")
        fmt.Scanln(&temp_weight)
        sum_of_weights = sum_of_weights + temp_weight
    }

    fmt.Println("The Average weight of the five people entered is: ")
    fmt.Println(sum_of_weights/5)
}
