package main

import "fmt"

func main() {
	var suhu float32
	fmt.Println("Masukkan Suhu yang Akan Dikonversi (dalam Celcius)")
	fmt.Scanln(&suhu)
	fmt.Println()

	var pilih int
	fmt.Print("Konversi ke\n1. Fahrenheit\n2. Kelvin\n3. Reamur\n")
	fmt.Print("Pilihan Konversi: ")
	fmt.Scanln(&pilih)

	switch pilih {
		case 1:
			fmt.Println("Hasil Konversi = ", suhu/5*9+32, "derajat Fahrenheit")
			break
		case 2:
			fmt.Println("Hasil Konversi = ", suhu+273, "Kelvin")
			break
		case 3:
			fmt.Println("Hasil Konversi = ", suhu/5*4, "derajat Reamur")
			break
		default:
			fmt.Println("Pilihan Anda Tidak Valid, Ulangi Kembali");
		}
	}