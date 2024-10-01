package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Description string
	Completed   bool
}

var tasks []Task

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Tambah tugas")
		fmt.Println("2. Lihat semua tugas")
		fmt.Println("3. Hapus tugas")
		fmt.Println("4. Tandai tugas selesai")
		fmt.Println("5. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Masukkan tugas baru: ")
			scanner.Scan()
			task := scanner.Text()
			tasks = append(tasks, Task{Description: task, Completed: false})
			fmt.Println("Tugas berhasil ditambahkan!")
		case "2":
			fmt.Println("Daftar Tugas:")
			for i, task := range tasks {
				status := "Belum selesai"
				if task.Completed {
					status = "Selesai"
				}
				fmt.Printf("%d. %s [%s]\n", i+1, task.Description, status)
			}
		case "3":
			fmt.Print("Masukkan nomor tugas yang ingin dihapus: ")
			scanner.Scan()
			numStr := scanner.Text()
			num, err := strconv.Atoi(numStr)
			if err != nil || num < 1 || num > len(tasks) {
				fmt.Println("Nomor tugas tidak valid.")
			} else {
				tasks = append(tasks[:num-1], tasks[num:]...)
				fmt.Println("Tugas berhasil dihapus!")
			}
		case "4":
			fmt.Print("Masukkan nomor tugas yang ingin ditandai selesai: ")
			scanner.Scan()
			numStr := scanner.Text()
			num, err := strconv.Atoi(numStr)
			if err != nil || num < 1 || num > len(tasks) {
				fmt.Println("Nomor tugas tidak valid.")
			} else {
				tasks[num-1].Completed = true
				fmt.Println("Tugas berhasil ditandai selesai!")
			}
		case "5":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}