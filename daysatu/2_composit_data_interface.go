package daysatu

import "fmt"

// Definisi interface
type Kalkulator interface {
	Tambah() float64
	Kurang() float64
	Kali() float64
	Bagi() float64
}

// Implementasi interface untuk struct Masukkan
type Masukkan struct {
	prm1 float64
	prm2 float64
}

func (r Masukkan) Tambah() float64 {
	return r.prm1 + r.prm2
}

func (r Masukkan) Kurang() float64 {
	return r.prm1 - r.prm2
}

func (c Masukkan) Kali() float64 {
	return c.prm1 * c.prm2
}

func (c Masukkan) Bagi() float64 {
	return c.prm1 / c.prm2
}

func InterfaceType() {
	// Penggunaan interface
	var kalkulator Kalkulator
	input := Masukkan{
		prm1: 10,
		prm2: 2,
	}
	kalkulator = input
	kalkulator.Bagi()
	fmt.Printf("%v+%v:%v\n", input.prm1, input.prm2, kalkulator.Tambah())
	fmt.Printf("%v-%v:%v\n", input.prm1, input.prm2, kalkulator.Kurang())
	fmt.Printf("%vx%v:%v\n", input.prm1, input.prm2, kalkulator.Kali())
	fmt.Printf("%v/%v:%v\n", input.prm1, input.prm2, kalkulator.Bagi())
}
