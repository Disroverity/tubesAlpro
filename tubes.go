package main

import "fmt"

const maxData int = 100

type Aktivitas struct {
	nama        string
	tanggal     string
	durasi      float64
	penghasilan int
	efisiensi   float64
}

var dataAktivitas [maxData]Aktivitas
var banyakAktivitas int = 0

func main() {
	var pilihan int
	for {
		fmt.Println("==== SIDE HUSTLE INCOME PLANNER (SHIP) ====")
		fmt.Println("1. Tambah aktivitas baru")
		fmt.Println("2. Lihat semua aktivitas")
		fmt.Println("3. Cari aktivitas")
		fmt.Println("4. Urutkan aktivitas")
		fmt.Println("5. Lihat aktivitas paling efisien")
		fmt.Println("6. Analisis ketahanan finansial Side Hustle")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahAktivitas(&dataAktivitas, &banyakAktivitas)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahAktivitas(A *[maxData]Aktivitas, banyakAktivitas *int) {
	var nama, tanggal string
	var durasi, efisiensi float64
	var penghasilan int

	if *banyakAktivitas > maxData {
		fmt.Println("Data sudah penuh. Tidak bisa menambah aktivitas baru.")
	} else {
		fmt.Print("Nama Aktivitas           ")
		fmt.Scan(&nama)
		fmt.Print("Tanggal (DD-MM-YYYY)     ")
		fmt.Scan(&tanggal)
		fmt.Print("Durasi Kerja (jam)       ")
		fmt.Scan(&durasi)
		fmt.Print("Uang Diterima (Rp)       ")
		fmt.Scan(&penghasilan)

		efisiensi = float64(penghasilan) / durasi
		A[*banyakAktivitas] = Aktivitas{
			nama:        nama,
			tanggal:     tanggal,
			durasi:      durasi,
			penghasilan: penghasilan,
			efisiensi:   efisiensi,
		}

		*banyakAktivitas++
		fmt.Println("Aktivitas berhasil ditambahkan!")
	}
}

func lihatAktivitas()
