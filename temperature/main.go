package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Constants for conversion factors
const cToFMultiplier float64 = 9.0 / 5.0
const fToCMultiplier float64 = 5.0 / 9.0
const offset float64 = 32.0

// Temperature struct represents a temperature value and its unit
type Temperature struct {
	Value float64
	Unit  string // C for Celsius, F for Fahrenheit
}

// Function to convert Celsius to Fahrenheit
func celsiusToFahrenheit(temp *Temperature) *Temperature {
	convertedValue := temp.Value*cToFMultiplier + offset
	return &Temperature{Value: convertedValue, Unit: "F"}
}

// Function to convert Fahrenheit to Celsius
func fahrenheitToCelsius(temp *Temperature) *Temperature {
	convertedValue := (temp.Value - offset) * fToCMultiplier
	return &Temperature{Value: convertedValue, Unit: "C"}
}

func main() {
	// Check if the correct number of arguments are provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: tempconverter [temperature] [unit]")
		fmt.Println("Example: tempconverter 100 C")
		os.Exit(1)
	}

	// Read the temperature value and unit from command-line arguments
	tempStr := os.Args[1]
	unit := strings.ToUpper(os.Args[2])

	// Convert the temperature string to a float64
	tempValue, err := strconv.ParseFloat(tempStr, 64)
	if err != nil {
		fmt.Println("Invalid temperature value. Please enter a numeric value.")
		os.Exit(1)
	}

	// Create a pointer to a Temperature struct
	temp := &Temperature{Value: tempValue, Unit: unit}

	// Pointer to hold the converted temperature
	var convertedTemp *Temperature

	// Perform the conversion based on the unit
	switch temp.Unit {
	case "C":
		convertedTemp = celsiusToFahrenheit(temp)
	case "F":
		convertedTemp = fahrenheitToCelsius(temp)
	default:
		fmt.Println("Invalid unit. Please provide 'C' for Celsius or 'F' for Fahrenheit.")
		os.Exit(1)
	}

	// Output the converted temperature
	fmt.Printf("%.0f %s\n", convertedTemp.Value, convertedTemp.Unit)
}
