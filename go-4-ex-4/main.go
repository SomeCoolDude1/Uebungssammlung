package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(C float64) float64 {
	F := C*(9.0/5.0) + 32
	return F
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(F float64) float64 {
	C := (F - 32) * (5.0 / 9.0)
	return C
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	test1 := convertCelsiusToFahrenheit(3)  //37.4
	test2 := convertCelsiusToFahrenheit(10) //50
	test3 := convertCelsiusToFahrenheit(20) //68
	// TODO: call the function convertFahrenheitToCelsius
	fmt.Println(convertFahrenheitToCelsius(test1))
	fmt.Println(convertFahrenheitToCelsius(test2))
	fmt.Println(convertFahrenheitToCelsius(test3))
}
