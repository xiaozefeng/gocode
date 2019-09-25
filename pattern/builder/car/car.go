package car

import "fmt"

type CarBuilder struct {
	Car Car
}

type Car struct {
	Color  Color
	Speed  Speed
	Wheels Wheels
}

func (c Car) Drive() error {
	fmt.Printf("color:%s,speed:%.2f,wheels:%s is running\n", c.Color, c.Speed, c.Wheels)
	return nil
}

func (c Car) Stop() error {
	fmt.Printf("color:%s,speed:%.2f,wheels:%s is running\n", c.Color, c.Speed, c.Wheels)
	return nil
}

func (c *CarBuilder) Color(color Color) Builder {
	c.Car.Color = color
	return c
}

func (c *CarBuilder) Wheels(wheels Wheels) Builder {
	c.Car.Wheels = wheels
	return c
}

func (c *CarBuilder) TopSpeed(speed Speed) Builder {
	c.Car.Speed = speed
	return c
}

func (c *CarBuilder) Build() Interface {
	return c.Car
}

func NewBuilder() Builder{
	return &CarBuilder{}
}
