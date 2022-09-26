package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	// get input Argument
	var argsRaw = os.Args
	var num = argsRaw[1]

	// convert string menjadi int
	number, _ := strconv.Atoi(num)
	// fmt.Printf("-> %#v\n", number)

	// data siswa
	var dataSiswa = []Person{
		{Nama: "Ikhwan", Alamat: "Bandung", Pekerjaan: "Mahasiswa", Alasan: "Ingin tahu lebih banyak"},
		{Nama: "Shalih", Alamat: "Bandung", Pekerjaan: "Mahasiswa", Alasan: "Supaya jadi programmer handal"},
		{Nama: "Dimas", Alamat: "Bandung", Pekerjaan: "Mahasiswa", Alasan: "Mau jadi backend engineer"},
		{Nama: "Ahmad", Alamat: "Bandung", Pekerjaan: "Mahasiswa", Alasan: "Modal buat cari kerja"},
		{Nama: "Firman", Alamat: "Bandung", Pekerjaan: "Mahasiswa", Alasan: "Memperdalam ilmu coding"},
		{Nama: "Rizal", Alamat: "Bandung", Pekerjaan: "Mahasiswa", Alasan: "Dapetin skill baru"},
		{Nama: "Farid", Alamat: "Bandung", Pekerjaan: "Pengangguran", Alasan: "Nyari kesibukan"},
		{Nama: "Fatiya", Alamat: "Bandung", Pekerjaan: "Mahasiswa", Alasan: "Diajak teman"},
		{Nama: "Ikhsan", Alamat: "Bandung", Pekerjaan: "Mahasiswa", Alasan: "Disuruh orang tua"},
		{Nama: "Farah", Alamat: "Bandung", Pekerjaan: "Mahasiswa", Alasan: "Dapetin sertifikat"},
	}

	// check data jika tidak ada
	if number > len(dataSiswa) {
		showError()
	} else {
		// jika data siswa ada panggil fungsi showdata
		showData(dataSiswa[number-1])
	}
}

func showData(mahasiswa Person) {
	fmt.Println("Nama      : \t", mahasiswa.Nama)
	fmt.Println("Alamat    : \t", mahasiswa.Alamat)
	fmt.Println("Pekerjaan : \t", mahasiswa.Pekerjaan)
	fmt.Println("Alasan    : \t", mahasiswa.Alasan)
}

func showError() {
	fmt.Println("No Absen Tidak Ditemukan")
}
