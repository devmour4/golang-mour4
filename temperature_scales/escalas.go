//Escalas Termométricas//

package main

import "fmt"
// Boiling point Kelvin
var Kelvin = 373.0

func main() {
	tempK := Kelvin

//conversion formula
	Celsius := tempK - 273.0

//result
	fmt.Printf("A temperatura de ebulição da água em °K = %g, e em °C é = %g", tempK,Celsius)
}