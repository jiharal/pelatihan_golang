package daysatu

import "fmt"

// Definisi struct
type Person struct {
	Name    string
	Age     int
	Address string
}

func StructType() {
	// Inisialisasi struct
	var person1 Person
	person1.Name = "John Doe"
	person1.Age = 30
	person1.Address = "123 Main St"

	// Inisialisasi struct secara langsung
	// person1 := Person{Name: "John Doe", Age: 30, Address: "123 Main St"}
	fmt.Println(person1) // Output: {John Doe 30 123 Main St}
}
