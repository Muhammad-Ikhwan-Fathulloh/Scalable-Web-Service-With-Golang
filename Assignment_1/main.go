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
		{Nama: "Ikhwan", Alamat: "Jl telaga bogor", Pekerjaan: "Mahasiswa", Alasan: "Ingin tahu lebih banyak"},
		{Nama: "Shalih", Alamat: "Jl perkutut bintaro", Pekerjaan: "Mahasiswa", Alasan: "Supaya jadi programmer handal"},
		{Nama: "Dimas", Alamat: "Jl Bulak indah jakbar", Pekerjaan: "Mahasiswa", Alasan: "Mau jadi backend engineer"},
		{Nama: "Ahmad", Alamat: "Jl Paal V Pondok indah", Pekerjaan: "Mahasiswa", Alasan: "Modal buat cari kerja"},
		{Nama: "Firman", Alamat: "Jl Kasang Ciledug", Pekerjaan: "PNS", Alasan: "Memperdalam ilmu coding"},
		{Nama: "Rizal", Alamat: "Jl Kumpeh Banten", Pekerjaan: "Mahasiswa", Alasan: "Dapetin skill baru"},
		{Nama: "Farid", Alamat: "Jl Senopati Jakarta", Pekerjaan: "Pengangguran", Alasan: "Nyari kesibukan"},
		{Nama: "Fatiya", Alamat: "Jl Kenangan Bandung barat", Pekerjaan: "Mahasiswa", Alasan: "Diajak teman"},
		{Nama: "Ikhsan", Alamat: "Jl Sudirman", Pekerjaan: "Mahasiswa", Alasan: "Disuruh orang tua"},
		{Nama: "Farah", Alamat: "Jl Kebangsaan", Pekerjaan: "Freelancer", Alasan: "Dapetin sertifikat"},
	}

	// check data jika tidak ada
	if number > len(dataSiswa) {
		showError()
	} else {
		// jika data siswa ada panggil fungsi showdata
		showData(dataSiswa[number-1])
	}
}

func showData(siswa Person) {
	fmt.Println("Nama      : \t", siswa.Nama)
	fmt.Println("Alamat    : \t", siswa.Alamat)
	fmt.Println("Pekerjaan : \t", siswa.Pekerjaan)
	fmt.Println("Alasan    : \t", siswa.Alasan)
}

func showError() {
	fmt.Println("No Absen Tidak Ditemukan")
}
