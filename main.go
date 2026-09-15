package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Definisikan Struct Task
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// Slice buat nampung task di memori
var tasks []Task
var nextID = 1

func addTask(title string) {
	newTask := Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}
	tasks = append(tasks, newTask)
	nextID++
	fmt.Println("✓ Tugas berhasil ditambahkan!")
}

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("Belum ada tugas tersimpan.")
		return
	}
	fmt.Println("\n--- DAFTAR TUGAS ---")
	for _, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
	}

	fmt.Println("--------------------")
}

func completeTask(id int) {
	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			found = true
			fmt.Printf("✓ Tugas $%d ditandai selesai!\n", id)
			break
		}
	}
	if !found {
		fmt.Printf("! Tugas dengan ID %d tidak ditemukan.\n", id)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== TASK CLI GO ===")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Lihat Semua Tugas")
		fmt.Println("3. Selesaikan Tugas")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih Menu (1-4): ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Masukkan Judul Tugas: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input != "" {
				addTask(input)
			} else {
				fmt.Println("Judul tugas tidak boleh kosong.")
			}
		case 2:
			showTasks()
		case 3:
			fmt.Print("Masukkan ID tugas yang selesai: ")
			var id int
			fmt.Scanln(&id)
			completeTask(id)
		case 4:
			fmt.Println("Terima kasih, program selesai!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

// // 1. Definisikan Struct Task
// type Task struct {
// 	ID        int
// 	Title     string
// 	Completed bool
// }

// // 2. Kumpulan tugas disimpan di slice
// var tasks []Task
// var nextID = 1

// // Fungsi untuk menambah tugas baru
// func addTask(title string) {
// 	newTask := Task{
// 		ID:        nextID,
// 		Title:     title,
// 		Completed: false,
// 	}
// 	tasks = append(tasks, newTask)
// 	nextID++
// 	fmt.Println("✓ Tugas berhasil ditambahkan!")
// }

// // Fungsi untuk menampilkan semua tugas
// func showTasks() {
// 	if len(tasks) == 0 {
// 		fmt.Println("Belum ada tugas tersimpan.")
// 		return
// 	}

// 	fmt.Println("\n--- DAFTAR TUGAS ---")
// 	for _, task := range tasks {
// 		status := "[ ]"
// 		if task.Completed {
// 			status = "[x]"
// 		}
// 		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
// 	}
// 	fmt.Println("--------------------")
// }

// // Fungsi untuk menandai tugas selesai (pakai pointer)
// func completeTask(id int) {
// 	found := false
// 	for i := range tasks {
// 		if tasks[i].ID == id {
// 			tasks[i].Completed = true
// 			found = true
// 			fmt.Printf("✓ Tugas #%d ditandai selesai!\n", id)
// 			break
// 		}
// 	}
// 	if !found {
// 		fmt.Printf("! Tugas dengan ID %d tidak ditemukan.\n", id)
// 	}
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Println("\n=== TASK CLI GO ===")
// 		fmt.Println("1. Tambah Tugas")
// 		fmt.Println("2. Lihat Semua Tugas")
// 		fmt.Println("3. Selesaikan Tugas")
// 		fmt.Println("4. Keluar")
// 		fmt.Print("Pilih menu (1-4): ")

// 		var choice int
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case 1:
// 			fmt.Print("Masukkan judul tugas: ")
// 			input, _ := reader.ReadString('\n')
// 			input = strings.TrimSpace(input)
// 			if input != "" {
// 				addTask(input)
// 			} else {
// 				fmt.Println("Judul tugas tidak boleh kosong.")
// 			}
// 		case 2:
// 			showTasks()
// 		case 3:
// 			fmt.Print("Masukkan ID tugas yang selesai: ")
// 			var id int
// 			fmt.Scanln(&id)
// 			completeTask(id)
// 		case 4:
// 			fmt.Println("Terima kasih, program selesai!")
// 			return
// 		default:
// 			fmt.Println("Pilihan tidak valid, coba lagi.")
// 		}
// 	}
// }
