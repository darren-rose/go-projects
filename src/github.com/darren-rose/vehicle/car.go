package vehicle

import "fmt"

type Car struct {
	Make string
	Model string
	Started bool
	Gear int
}

func (car *Car) Start() {
	car.Started = true
}

func (car *Car) ShiftUp() {
	car.Gear = car.Gear + 1
}

func (car *Car) ToString() string {
	started := "N"
	if car.Started == true { started = "Y" }
	return fmt.Sprintf("Make:%s Model:%s Started:%s Gear:%d", car.Make, car.Model, started, car.Gear)
}