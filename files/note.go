package files

import (
  "fmt"
  "encoding/json"
  "os"
)

type Note struct {
	Judul   string
	Catatan string
}

// === FILE HANDLING ===

func SimpanJSON(Notes map[string]Note) {
	file, err := os.Create("catatan.json")
	if err != nil {
		fmt.Println("❌ Gagal menyimpan file:", err)
		return 
	} 
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(Notes)
	if err != nil {
		fmt.Println("❌ Gagal menyimpan data JSON:", err)
	}
}

func MuatDariJSON(Notes *map[string]Note) {
	file, err := os.Open("catatan.json")
	if err != nil {
	  fmt.Println("⚠️ Gagal membuka catatan.json")
		*Notes = make(map[string]Note)
		return 
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(Notes)
	if err != nil {
		fmt.Println("⚠️  File rusak atau kosong, mulai baru.")
		*Notes = make(map[string]Note)
	}
}