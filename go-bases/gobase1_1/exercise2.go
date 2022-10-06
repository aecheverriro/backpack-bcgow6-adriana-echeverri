package main

import "fmt"

func main() {
	var temperature float32
	var moisture float32
	var pressure int

	temperature = 17
	moisture = 63
	pressure = 760
	fmt.Printf("Temperature: %v°C\nMoisture: %v%%\nPressure: %vmmHg\n", temperature, moisture, pressure)
}
