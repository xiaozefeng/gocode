package main

import "github.com/xiaozefeng/gocode/pattern/builder/car"

func main() {
	builder := car.NewBuilder().Color(car.BlueColor)

	familyCar := builder.Wheels(car.SportWheels).TopSpeed(50 * car.MPH).Build()
	familyCar.Drive()

	sportCar := builder.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
	sportCar.Drive()
}
