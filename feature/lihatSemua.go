package feature

import (
  "fmt"
  "notes/files"
)

func LihatSemua(Notes map[string]files.Note) {
	if len(Notes) == 0 {
		fmt.Println("📭 Tidak ada catatan.")
		return
	}
	i := 1
	for key := range Notes {
		fmt.Printf("%d. %s\n", i, key)
		i++
	}
}