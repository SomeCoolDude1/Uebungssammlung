package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDay struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  int16
}

type Profile struct {
	Name             FullName
	Birth            BirthDay
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var name FullName = FullName{
		"Ioannis", "Ligkas",
	}
	var birth BirthDay = BirthDay{
		2, 12, 2007,
	}
	var me = Profile{
		Name:             name,
		Birth:            birth,
		NumberOfSiblings: 1,        // TODO: adjust
		ZodiacSign:       '\u2650', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
