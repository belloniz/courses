package main

import (
	"fmt"
	"math"
)

type errorString struct {
	s string

}

type areaError struct {
	err string
	radius float64
}

func New(text string) error  {
	return &errorString{text}
}

func (e *errorString) Error() string  {
	return e.s
}

func (e *areaError) Error() string  {
	return fmt.Sprintf("Radius %0.2f: %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error)  {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func main()  {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle1 %0.2f", area)
}