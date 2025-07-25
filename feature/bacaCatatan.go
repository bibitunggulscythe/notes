package feature

import (
  "fmt"
  "notes/files"
)

func BacaCatatan(Notes map[string]files.Note) {
	judul := Inputan("Masukkan judul catatan: ")
	if note, ok := Notes[judul]; ok {
		fmt.Println("\n---")
		fmt.Printf("Judul: %s\nIsi: %s\n", note.Judul, note.Catatan)
		fmt.Println("---")
	} else {
		fmt.Println("❌ Judul tidak ditemukan.")
	}
}