package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	n, error := fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	n2, error2 := fmt.Fprintln(os.Stderr, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go>eyes.txt es steht in die standard ausgabe weil für mich der Augenzahl frei zum sehen sein soll
	// go run ex3/main.go2>dice.log es steht in die standard error weil ein log öffters für admins ist
	fmt.Print(n, n2, error, error2)
}
