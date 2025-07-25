package feature

import (
  "fmt"
  "notes/files"
)

func TambahCatatan(Notes map[string]files.Note) {
	judul := Inputan("Masukkan judul: ")
	if _, ok := Notes[judul]; ok {
		fmt.Println("⚠️  Judul sudah ada. Gunakan judul lain.")
		return
	}
	isi := Inputan("Tulis catatan: ")

	Notes[judul] = files.Note{Judul: judul, Catatan: isi}
	fmt.Println("✅ Catatan ditambahkan.")
	files.SimpanJSON(Notes)
}