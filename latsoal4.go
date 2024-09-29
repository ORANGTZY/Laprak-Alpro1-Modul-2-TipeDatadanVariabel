package main

import "fmt"

func main() {
	var fahrenheit, celsius float64

	// Meminta input suhu dalam Fahrenheit dari pengguna
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// Mengonversi suhu dari Fahrenheit ke Celcius
	celsius = (fahrenheit - 32) * 5 / 9

	// Menampilkan hasil konversi suhu dalam Celcius
	fmt.Printf("Suhu dalam Celcius: %.2f\n", celsius)
}
