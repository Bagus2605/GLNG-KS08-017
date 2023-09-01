package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Struct untuk merepresentasikan data mahasiswa
type Student struct {
	ID          string `json:"id"`
	StudentCode string `json:"student_code"`
	Name        string `json:"student_name"`
	Address     string `json:"student_address"`
	Occupation  string `json:"student_occupation"`
	Reason      string `json:"joining_reason"`
}

// Struct untuk membaca file JSON
type Students struct {
	Participants []Student `json:"participants"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Masukan Parameter Kode Mahasiswa Setelah .go")
		return
	}

	// Membaca file JSON
	jsonFile, err := os.Open("participants.json")
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer jsonFile.Close()

	// Membaca isi file JSON
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Mendekode JSON ke dalam struct
	var students Students
	err = json.Unmarshal(byteValue, &students)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Mengambil KODE mahasiswa dari argumen baris perintah
	codeMahasiswa := os.Args[1]

	// Mencari dan menampilkan data mahasiswa berdasarkan ID
	searchByCODE(students.Participants, codeMahasiswa)
}

func searchByCODE(students []Student, code string) {
	for _, student := range students {
		if student.StudentCode == code {
			fmt.Printf("\n")
			fmt.Printf("ID		: %s\n", student.ID)
			fmt.Printf("Kode Mahasiswa	: %s\n", student.StudentCode)
			fmt.Printf("Nama		: %s\n", student.Name)
			fmt.Printf("Alamat		: %s\n", student.Address)
			fmt.Printf("Pekerjaan	: %s\n", student.Occupation)
			fmt.Printf("Alasan		: %s\n", student.Reason)
			fmt.Printf("\n")
			return
		}
	}
	fmt.Println("Mahasiswa dengan Kode", code, "tidak ditemukan.")
}
