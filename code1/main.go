package main

import (
	"fmt"
	"os"
	"strconv"
)

type Peson struct {
	Id                              int
	Nama, Alamat, Pekerjaan, Alasan string
}

func main() {
	persons := []Peson{
		{Id: 1, Nama: "John", Alamat: "Depok", Pekerjaan: "Programmer", Alasan: "Switch carrer"},
		{Id: 2, Nama: "Doe", Alamat: "Jakarta", Pekerjaan: "Atlet", Alasan: "Suka aja"},
		{Id: 3, Nama: "Jane", Alamat: "Bandung", Pekerjaan: "Accounting", Alasan: "Up skill"},
		{Id: 4, Nama: "Smith", Alamat: "Bali", Pekerjaan: "Driver", Alasan: "Biar keren"},
	}

	args := os.Args[1]
	id, _ := strconv.ParseInt(args, 10, 64)

	for _, val := range persons {
		if val.Id == int(id) {
			fmt.Printf("Data teman dengan absen %d\n", id)
			fmt.Printf("Nama : %s\n", val.Nama)
			fmt.Printf("Alamat : %s\n", val.Alamat)
			fmt.Printf("Pekerjaan : %s\n", val.Pekerjaan)
			fmt.Printf("Alasan : %s\n", val.Alasan)
			return
		}
	}

	fmt.Println("Id tidak ditemukan!")
}
