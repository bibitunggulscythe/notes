package feature

import(
  "fmt"
  "notes/files"
)

func GantiCatatan(Notes map[string]files.Note) {
  judul := Inputan("Masukkan judul yang ingin di ganti: ")
   if _, ok := Notes[judul]; !ok {
    fmt.Println("⚠️ Judul tidak ditemukan")
    return
  }
  
  delete(Notes, judul)
  judul = Inputan("Masukan judul baru: ")
  if _, ok := Notes[judul]; ok {
    fmt.Println("⚠️ Judul sudah ada gunakan judul lain")
    return
  }
  isi := Inputan("Masukkan isi baru: ")
  Notes[judul] = files.Note{Judul: judul, Catatan: isi}
  files.SimpanJSON(Notes)
}