package main

import "fmt"

// Create a new custom type called "distance" based on the built-in type "float32"
type distance float32

func main() {
	var miles distance // Declare variable "miles" of type "distance"
	miles = 20
	km := miles.ToKm()

	fmt.Println(km)
}

// Create a method "ToKm()" on custom type "distance"; meaning all instances of type "distance"
// will have access to this method
func (d distance) ToKm() distance {
	return d * 1.6
}
