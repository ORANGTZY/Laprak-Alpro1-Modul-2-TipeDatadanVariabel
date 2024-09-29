package main

import "fmt"

func main() {
	var r, luas float64
	const pi = 3.14

	// Meminta input jari-jari lingkaran dari pengguna
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)

	// Menghitung luas lingkaran menggunakan rumus π * r^2
	luas = pi * r * r

	// Menampilkan hasil perhitungan luas lingkaran
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", r, luas)
}
