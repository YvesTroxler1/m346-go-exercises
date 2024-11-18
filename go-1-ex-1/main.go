package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!

	var firstName string = "Yves";
	var lastName string = "Troxler";

	var dayOfBirth uint = 3;
	var monthOfBirth uint = 8;
	var yearOfBirth uint16 = 2008;

	var numberOfSiblings uint = 2;
	
	var heightInMeters float32 = 1.78;

	var zodiacSign rune = '\u264C';

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
