package main

import (
	"fmt"
	"notes/files"
	"notes/feature"
)

func main() {
  Notes := make(map[string]files.Note)
	files.MuatDariJSON(&Notes)

	for {
		fmt.Println("\n=== APLIKASI CATATAN CLI ===")
		fmt.Println("1. Tambah catatan")
		fmt.Println("2. Lihat semua catatan")
		fmt.Println("3. Baca isi catatan")
		fmt.Println("4. Ganti catatan")
		fmt.Println("5. Hapus catatan")
		fmt.Println("6. Keluar")

		pilihan := feature.Inputan("Pilih menu: ")

		switch pilihan {
		case "1":
			feature.TambahCatatan(Notes)
		case "2":
			feature.LihatSemua(Notes)
		case "3":
			feature.BacaCatatan(Notes)
		case "4":
		  feature.GantiCatatan(Notes)
		case "5":
			feature.HapusCatatan(Notes)
		case "6":
			fmt.Println("👋 Sampai jumpa!")
			return
		default:
			fmt.Println("⚠️  Menu tidak valid.")
		}
	}
}