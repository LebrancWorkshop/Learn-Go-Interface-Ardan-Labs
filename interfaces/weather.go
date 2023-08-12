package interfaces

import "fmt"

type Celcius float64

func (c Celcius) Unit() string {
	return "°C"
}

func (c Celcius) Amount() float64 {
	return float64(c)
}

func Current() (Celcius, error) {
	c := Celcius(23.4)
	u := c.Unit()

	fmt.Printf("C: %v, U: %v", c, u)

	return c, nil
}


