package daysatu

import "fmt"

func ArrayType() {
	// Deklarasi array dengan panjang 3 yang berisi integer
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	// Inisialisasi array secara langsung
	// var numbers = [3]int{1, 2, 3}
	fmt.Println(numbers) // Output: [1 2 3]
}

func SliceType() {
	// Deklarasi slice
	var fruits []string

	// Inisialisasi slice dengan elemen langsung
	// fruits := []string{"Apple", "Banana", "Orange"}

	// Menambahkan elemen ke slice
	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Banana")
	fruits = append(fruits, "Orange")

	fmt.Println(fruits) // Output: [Apple Banana Orange]
}

func MapType() {
	// Deklarasi dan inisialisasi map
	person := map[string]string{
		"name":  "John Doe",
		"email": "john.doe@example.com",
	}

	// Menambahkan elemen ke map
	person["phone"] = "123-456-789"

	fmt.Println(person) // Output: map[name:John Doe email:john.doe@example.com phone:123-456-789]
}
