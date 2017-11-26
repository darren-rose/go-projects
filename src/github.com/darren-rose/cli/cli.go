package main

import (
	"fmt"
	"os"
	"github.com/darren-rose/vehicle"
)

func main() {
	if len(os.Args) != 3 {
		println("usage: %s make model", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("%v argument supplied was: %v\n", os.Args[0], os.Args[1], os.Args[2])

	car := &vehicle.Car{Make: os.Args[1], Model: os.Args[2]}

	println(car.ToString())

	car.Start()
	car.ShiftUp()

	println(car.ToString())

}

