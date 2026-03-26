package main

import "fmt"

func main() {

	// db.select().filter().where().limit().offset().exec()
	r := NewClac(10.5).Add(12.3).Add(10.4).Sub(5.4).Mul(8).Div(4).Get()
	fmt.Printf("Result:%.2f\n", r)
}

type ICalc interface {
	Add(d float64) ICalc
	Sub(d float64) ICalc
	Mul(d float64) ICalc
	Div(d float64) ICalc
	Get() float64
}

type Calc struct {
	data float64
}

func NewClac(d float64) *Calc {
	return &Calc{d}
}

func (c *Calc) Add(d float64) ICalc {
	c.data += d
	return c
}

func (c *Calc) Sub(d float64) ICalc {
	c.data -= d
	return c
}
func (c *Calc) Mul(d float64) ICalc {
	c.data *= d
	return c
}

func (c *Calc) Div(d float64) ICalc {
	c.data /= d
	return c
}

func (c *Calc) Get() float64 {
	return c.data
}
