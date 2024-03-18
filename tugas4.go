package main

import "fmt"

func main() {
	mahasiswa := map[string]string{
		"aldi": "150 cm",
		"budi": "170 cm",
	}

	tampil_mahasiswa(mahasiswa)
}

func tampil_mahasiswa(mahasiswa map[string]string) {
	for key, value := range mahasiswa {
		fmt.Println(key, ":", value)
	}
}