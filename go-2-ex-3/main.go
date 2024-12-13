package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "Datenmodel Implementieren",
		117: "Informatik- und Netzwerkinfrastrucktur",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 346)
	// TODO: add one
	modules[320] = "Objektorientiertes Programmieren"
	// TODO: replace one
	delete(modules, 104)
	modules[107] = "Blockchain Technologie umsetzen"

	fmt.Println(modules)
}
