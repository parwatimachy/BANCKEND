package main

import "fmt"

func main() {
	for {
		fmt.Println("PROGRAM KALKULATOR SEDERHANA DIGOLANG")
		fmt.Println("1.OPERASI PERKALIAN")
		fmt.Println("2.OPERASI PEMBAGIAN")
		fmt.Println("3.OPERASI PENAMBAHAN")
		fmt.Println("4.OPERASI PENGURANGAN")
		fmt.Println("5.OPERASI MODULUS")
		fmt.Println("6.EXIT")
		fmt.Print("Pilih salah satu angka diatas :")

		var pilihan int
		fmt.Scan(&pilihan)

		if pilihan == 6 {
			fmt.Println("PROGRAM DIHENTIKAN")
			break
		}
		if pilihan < 1 || pilihan > 6 {
			fmt.Println("PILIHAN TIDAK VALID.SILAHKAN COBA LAGI")
			continue
		}

		var nilai1, nilai2 int
		fmt.Print("Masukkan angka pertama :")
		fmt.Scan(&nilai1)
		fmt.Print("Masukkan angka kedua :")
		fmt.Scan(&nilai2)

		switch pilihan {
		case 1:
			fmt.Println("HASIL PERKALIAN", nilai1*nilai2)
		case 2:
			if nilai2 == 0 {
				fmt.Println("PEMBAGIAN TIDAK VALID")
			} else {
				fmt.Println("HASIL PEMBAGIAN", nilai1/nilai2)
			}
		case 3:
			fmt.Println("HASIL PENAMBAHAN", nilai1+nilai2)
		case 4:
			fmt.Println("HASIL PENGURANGAN", nilai1/-nilai2)
		case 5:
			if nilai2 == 0 {
				fmt.Println("MODULUS TIDAK VALID")
			} else {
				fmt.Println("HASIL MODULUS", nilai1%nilai2)
			}
		}
	}
}