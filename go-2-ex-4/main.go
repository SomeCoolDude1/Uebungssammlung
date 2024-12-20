package main

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Students []Student
	}
	// TODO: declare a map of modules being attended by multiple classes
	modules := make(map[uint][]Class, 3)
	// TODO: output everything using fmt.Println()
	var Student1 = Student{
		FirstName: "Rat",
		LastName:  "Rattus",
	}
	var Student2 = Student{
		FirstName: "Test",
		LastName:  "Tester",
	}
	var Student3 = Student{
		FirstName: "Potato",
		LastName:  "Starchius",
	}
	var Student4 = Student{
		FirstName: "Max",
		LastName:  "Muster",
	}
	var Student5 = Student{
		FirstName: "Maxi",
		LastName:  "Musterfrau",
	}
	var Student6 = Student{
		FirstName: "Probe",
		LastName:  "Probius",
	}
	var Student7 = Student{
		FirstName: "Ligkas",
		LastName:  "Ioannis",
	}
	var Student8 = Student{
		FirstName: "Timon",
		LastName:  "Soom",
	}
	var Student9 = Student{
		FirstName: "Last",
		LastName:  "Guy",
	}

	var BMSE26a = Class{
		Students: []Student{Student1, Student2, Student3},
	}
	var BMSE26b = Class{
		Students: []Student{Student4, Student5, Student6},
	}
	var BMSE26d = Class{
		Students: []Student{Student7, Student8, Student9},
	}

}
