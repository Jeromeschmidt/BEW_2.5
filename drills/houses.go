package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    )

type House struct {
    numberOfRooms int
    city string
    address string
    price int
}

func printHouses(houses []House) {

    // Iterate over all houses in area and print a formatted string for each
    for i := 0; i < len(houses); i++ {
        fmt.Printf( "%-40s %-20s %2d Rooms   $%d \n", houses[i].address, houses[i].city, houses[i].numberOfRooms, houses[i].price)
    }
}

func getHouses() []House {

    reader := bufio.NewReader(os.Stdin)

    var check int
    check = 0

    var house House
    var numberOfRooms int
    var city string
    var address string
    var price int

    var houses []House

    // Loop and gather information on each house until exit condition is met
    for check != 1 {

        fmt.Println("Enter Number of Rooms: ")
        fmt.Scanln(&numberOfRooms)
        house.numberOfRooms = numberOfRooms

        fmt.Println("Enter the city: ")
        city, _ = reader.ReadString('\n')
        city = strings.TrimSuffix(city, "\n")
        house.city = city

        fmt.Println("Enter the address: ")
        address, _ = reader.ReadString('\n')
        address = strings.TrimSuffix(address, "\n")
        house.address = address

        fmt.Println("Enter the price: ")
        fmt.Scanln(&price)
        house.price = price

        houses = append(houses, house)

        fmt.Println("Press 1 to exit or 0 to continue: ")
        fmt.Scanln(&check)
    }

    return houses
}


func main() {
    printHouses(getHouses())
}
