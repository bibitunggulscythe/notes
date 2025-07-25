package feature

import (
  "fmt"
  "notes/files"
)

func HapusCatatan(Notes map[string]files.Note) {
	judul := Inputan("Masukkan judul yang ingin dihapus: ")
	if _, ok := Notes[judul]; ok {
		delete(Notes, judul)
		fmt.Println("🗑️  Catatan dihapus.")
		files.SimpanJSON(Notes)
	} else {
		fmt.Println("❌ Judul tidak ditemukan.")
	}
}