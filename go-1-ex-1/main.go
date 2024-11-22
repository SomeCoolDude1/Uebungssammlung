package main

import "fmt"

var firstName, lastName, dayOfBirth, monthOfBirth, yearOfBirth, numberOfSiblings, heightInMeters, zodiacSign = "Ioannis", "Ligkas", 2, 12, 2007, 1, 1.89, '\u2650'

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
