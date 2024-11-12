package main

import (
	"fmt"
)

type ToDo struct {
	list []string
}

func (t *ToDo) tambah(todo string) {
	t.list = append(t.list, todo)
}

func (t *ToDo) lihat() {
	if len(t.list) == 0 {
		fmt.Println("Tidak ada todo list yang tersedia.")
	} else {
		fmt.Println("Daftar To-Do List:")
		for i, todo := range t.list {
			fmt.Printf("%d. %s\n", i+1, todo)
		}
	}
}

func (t *ToDo) hapus(index int) {
	if index >= 0 && index < len(t.list) {
		t.list = append(t.list[:index], t.list[index+1:]...)
		fmt.Println("Berhasil menghapus todo list")
	} else {
		fmt.Println("Indeks tidak valid!")
	}
}

func (t *ToDo) update(index int, newToDo string) {
	if index >= 0 && index < len(t.list) {
		t.list[index] = newToDo
		fmt.Println("Berhasil mengupdate todo list")
	} else {
		fmt.Println("Indeks tidak valid!")
	}
}

func main() {
	var todoApp ToDo
	var pilihan int
	var input string

	for {
		fmt.Println("\nPROGRAM ToDo - List")
		fmt.Println("--------------------")
		fmt.Println("1. Menambah toDo List")
		fmt.Println("2. Melihat toDo List")
		fmt.Println("3. Menghapus toDo List")
		fmt.Println("4. Mengupdate toDo List")
		fmt.Println("5. Keluar applikasi")
		fmt.Print("Masukan Angka yang dipilih: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("\nMasukkan To-Do: ")
			fmt.Scanln(&input)
			todoApp.tambah(input)

		case 2:
			todoApp.lihat()

		case 3:
			todoApp.lihat()
			if len(todoApp.list) > 0 {
				fmt.Print("\nMasukkan nomor To-Do yang ingin dihapus: ")
				var index int
				fmt.Scanln(&index)
				todoApp.hapus(index - 1)
			}

		case 4:
			todoApp.lihat()
			if len(todoApp.list) > 0 {
				fmt.Print("\nMasukkan nomor To-Do yang ingin diupdate: ")
				var index int
				fmt.Scanln(&index)
				fmt.Print("Masukkan To-Do baru: ")
				fmt.Scanln(&input)
				todoApp.update(index-1, input)
			}

		case 5:
			fmt.Println("Terima kasih! Keluar dari aplikasi.")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		}
	}
}