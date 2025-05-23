package main

import "fmt"

type PekerjaanSampingan struct {
	Nama         string
	Bidang       string
	TipeUpah     string
	BesarUpah    int
	Deskripsi    string
	Keterampilan string
	Fleksibel    bool
	Alamat       string
}

const maxData = 40
const maxFavorit = 10

var pekerjaan [maxData]PekerjaanSampingan
var favorit [maxFavorit]PekerjaanSampingan
var totalFavorit int = 0

func main() {
	initData()

	var pilihan int
	for {
		fmt.Println("\n=== SHIP (Side Hustle Income Planner) ===")
		fmt.Println("1. Cari pekerjaan sampingan")
		fmt.Println("2. Hitung potensi penghasilan")
		fmt.Println("3. Simpan pekerjaan favorit")
 		fmt.Println("4. Perbandingan Potensi Pekerjaan")
		fmt.Println("5. Mengurutkan Pekerjaan dari Nama")
        fmt.Println("6. Mengurutkan Pekerjaan dari Upah")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			menuCariPekerjaan()
		} else if pilihan == 2 {
			menuHitungPenghasilan()
		} else if pilihan == 3 {
			menuFavorit()
		} else if pilihan == 4 {
    	    menuPerbandingan()
		} else if pilihan == 5 {
    	    mengurutkanNama()
		} else if pilihan == 6 {
    	    menuCariUpah()
        } else if pilihan == 7 {
	        fmt.Println("Terima kasih telah menggunakan SHIP!")
	        break
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func initData() {
	pekerjaan[0] = PekerjaanSampingan{"Online Web Developer", "Teknologi", "Per proyek", 500000, "Membangun website untuk klien", "HTML, CSS, JS", true, "https://freelancer.com"}
	pekerjaan[1] = PekerjaanSampingan{"Desainer Grafis Freelance", "Desain", "Per proyek", 300000, "Desain logo dan media sosial", "CorelDraw, Photoshop", true, "https://fiverr.com"}
	pekerjaan[2] = PekerjaanSampingan{"Penulis Artikel", "Konten", "Per artikel", 100000, "Menulis artikel SEO", "Menulis, Riset", true, "https://sribulancer.com"}
	pekerjaan[3] = PekerjaanSampingan{"Guru Les Privat", "Pendidikan", "Per jam", 75000, "Mengajar siswa secara langsung", "Mengajar, Kesabaran", false, "Jl. Melati No. 12, Bandung"}
	pekerjaan[4] = PekerjaanSampingan{"Admin Sosial Media", "Pemasaran", "Per bulan", 1500000, "Kelola akun Instagram/FB", "Copywriting, Canva", true, "https://remoteok.io"}
	pekerjaan[5] = PekerjaanSampingan{"Freelance Data Entry", "Administrasi", "Per jam", 35000, "Input data dari dokumen ke sistem", "Excel, Teliti", true, "https://projects.upwork.com"}
	pekerjaan[6] = PekerjaanSampingan{"Petugas Inventaris Gudang", "Logistik", "Per shift", 120000, "Mengecek dan mendata barang masuk/keluar", "Ketelitian, Excel", false, "Jl. Industri No. 12, Jakarta"}
	pekerjaan[7] = PekerjaanSampingan{"Video Editor YouTube", "Media", "Per video", 300000, "Mengedit video YouTube", "Premiere Pro, Storytelling", true, "https://freelancer.com"}
	pekerjaan[8] = PekerjaanSampingan{"Barista Paruh Waktu", "Kuliner", "Per jam", 40000, "Membuat kopi dan minuman", "Meracik, Melayani", false, "Jl. Cendana No. 7, Bandung"}
	pekerjaan[9] = PekerjaanSampingan{"Dropshipper Online", "E-Commerce", "Per produk", 15000, "Menjual barang secara online", "Marketing, Negosiasi", true, "https://shopee.co.id"}
	pekerjaan[10] = PekerjaanSampingan{"Programmer Python Junior", "Teknologi", "Per proyek", 700000, "Membuat aplikasi sederhana", "Python, Debugging", true, "https://remoteok.io"}
	pekerjaan[11] = PekerjaanSampingan{"Ilustrator Buku Anak", "Desain", "Per ilustrasi", 250000, "Menggambar untuk buku anak", "Ilustrasi, Warna", true, "https://sribulancer.com"}
	pekerjaan[12] = PekerjaanSampingan{"Content Creator TikTok", "Konten", "Per video", 100000, "Membuat konten pendek kreatif", "Video Editing, Ide Kreatif", true, "https://tiktok.com"}
	pekerjaan[13] = PekerjaanSampingan{"Tutor Bahasa Jepang", "Pendidikan", "Per jam", 100000, "Mengajar dasar-dasar Bahasa Jepang", "Bahasa Jepang, Mengajar", false, "Jl. Sakura No. 88, Bekasi"}
	pekerjaan[14] = PekerjaanSampingan{"Jasa Desain Undangan Digital", "Pemasaran", "Per desain", 100000, "Membuat undangan pernikahan digital", "Canva, Kreativitas", true, "https://canva.com"}
	pekerjaan[15] = PekerjaanSampingan{"Jasa Pengetikan Skripsi", "Administrasi", "Per halaman", 3000, "Mengetik dokumen dan skripsi mahasiswa", "Ms Word, Rapi", true, "https://freelancer.com"}
	pekerjaan[16] = PekerjaanSampingan{"Kurir Makanan Lokal", "Logistik", "Per order", 15000, "Mengantar makanan lokal ke pelanggan", "Motor, Aplikasi delivery", true, "https://grab.com"}
	pekerjaan[17] = PekerjaanSampingan{"Penyiar Podcast Indie", "Media", "Per episode", 100000, "Membuat dan menyebarkan podcast", "Public Speaking, Editing", true, "https://spotify.com"}
	pekerjaan[18] = PekerjaanSampingan{"Penjual Kue Online", "Kuliner", "Per box", 80000, "Membuat dan menjual kue homemade", "Memasak, Promosi", true, "https://instagram.com"}
	pekerjaan[19] = PekerjaanSampingan{"Admin Marketplace", "E-Commerce", "Per bulan", 1000000, "Kelola toko online di marketplace", "Shopee, Tokopedia", true, "https://tokopedia.com"}
	pekerjaan[20] = PekerjaanSampingan{"Frontend Developer", "Teknologi", "Per proyek", 800000, "Membangun antarmuka web responsif", "ReactJS, CSS", true, "https://remoteok.io"}
	pekerjaan[21] = PekerjaanSampingan{"Desainer UI/UX", "Desain", "Per proyek", 600000, "Desain pengalaman pengguna aplikasi", "Figma, Wireframing", true, "https://figma.com"}
	pekerjaan[22] = PekerjaanSampingan{"Penulis Caption Sosmed", "Konten", "Per caption", 25000, "Membuat teks menarik untuk media sosial", "Copywriting, Kreativitas", true, "https://instagram.com"}
	pekerjaan[23] = PekerjaanSampingan{"Guru Privat Matematika", "Pendidikan", "Per jam", 80000, "Mengajar matematika tingkat SMA", "Matematika, Mengajar", false, "Jl. Aljabar No. 22, Bandung"}
	pekerjaan[24] = PekerjaanSampingan{"Digital Marketer", "Pemasaran", "Per kampanye", 1200000, "Mengelola kampanye digital", "SEO, Google Ads", true, "https://linkedin.com"}
	pekerjaan[25] = PekerjaanSampingan{"Data Entry Asisten", "Administrasi", "Per jam", 30000, "Input data survey ke spreadsheet", "Excel, Ketelitian", true, "https://upwork.com"}
	pekerjaan[26] = PekerjaanSampingan{"Staf Gudang Harian", "Logistik", "Per hari", 100000, "Membantu pengemasan dan sortir barang", "Tenaga fisik, Rapi", false, "Jl. Gudang No. 77, Surabaya"}
	pekerjaan[27] = PekerjaanSampingan{"Editor Video Reels", "Media", "Per video", 200000, "Mengedit konten pendek Instagram", "CapCut, Storytelling", true, "https://instagram.com"}
	pekerjaan[28] = PekerjaanSampingan{"Penjual Makanan Ringan", "Kuliner", "Per pack", 10000, "Menjual snack buatan sendiri", "Masak, Kemasan", true, "https://shopee.co.id"}
	pekerjaan[29] = PekerjaanSampingan{"Spesialis Produk Online", "E-Commerce", "Per bulan", 1300000, "Mengelola katalog dan promosi toko", "Shopee, Optimasi", true, "https://tokopedia.com"}
	pekerjaan[30] = PekerjaanSampingan{"Mobile App Developer", "Teknologi", "Per proyek", 1000000, "Membuat aplikasi Android/iOS sederhana", "Flutter, Dart", true, "https://freelancer.com"}
	pekerjaan[31] = PekerjaanSampingan{"Animator 2D", "Desain", "Per animasi", 400000, "Membuat animasi karakter 2D", "Adobe Animate, Krita", true, "https://sribulancer.com"}
	pekerjaan[32] = PekerjaanSampingan{"Penulis Script YouTube", "Konten", "Per naskah", 75000, "Membuat script untuk channel edukasi", "Riset, Menulis", true, "https://youtube.com"}
	pekerjaan[33] = PekerjaanSampingan{"Guru Les Bahasa Inggris", "Pendidikan", "Per jam", 90000, "Mengajar speaking dan grammar dasar", "Bahasa Inggris, Teaching", false, "Jl. Bahasa No. 3, Malang"}
	pekerjaan[34] = PekerjaanSampingan{"Spesialis Iklan Facebook", "Pemasaran", "Per kampanye", 1100000, "Menjalankan iklan Facebook Ads", "Targeting, Analitik", true, "https://facebook.com/ads"}
	pekerjaan[35] = PekerjaanSampingan{"Asisten Administrasi Sekolah", "Administrasi", "Per bulan", 900000, "Membantu pengarsipan dan data siswa", "Ms Office, Rapi", false, "Jl. Sekolah No. 12, Depok"}
	pekerjaan[36] = PekerjaanSampingan{"Driver Logistik Harian", "Logistik", "Per trip", 70000, "Mengantar barang dari gudang ke toko", "SIM A, Navigasi", false, "Jl. Kirim No. 89, Bekasi"}
	pekerjaan[37] = PekerjaanSampingan{"Kreator Podcast Edukasi", "Media", "Per episode", 120000, "Podcast materi edukasi pelajar", "Voice clarity, Topik menarik", true, "https://spotify.com"}
	pekerjaan[38] = PekerjaanSampingan{"Penjual Nasi Kotak Online", "Kuliner", "Per kotak", 20000, "Menerima pesanan catering rumahan", "Memasak, Manajemen pesanan", true, "https://whatsapp.com"}
	pekerjaan[39] = PekerjaanSampingan{"Admin Toko Online Fashion", "E-Commerce", "Per bulan", 950000, "Mengelola pesanan dan chat pembeli", "Marketplace, Chat cepat", true, "https://shopee.co.id"}
}


func menuCariPekerjaan() {
	var bidang string
	fmt.Println("\nPilih bidang:")
	fmt.Println("1. Teknologi")
	fmt.Println("2. Desain")
	fmt.Println("3. Konten")
	fmt.Println("4. Pendidikan")
	fmt.Println("5. Pemasaran")
	fmt.Println("6. Administrasi")
	fmt.Println("7. Logistik")
	fmt.Println("8. Media")
	fmt.Println("9. Kuliner")
	fmt.Println("10. E-Commerce")
	fmt.Print("Masukkan pilihan (1-10): ")
	var pilih int
	fmt.Scan(&pilih)

	if pilih == 1 {
		bidang = "Teknologi"
	} else if pilih == 2 {
		bidang = "Desain"
	} else if pilih == 3 {
		bidang = "Konten"
	} else if pilih == 4 {
		bidang = "Pendidikan"
	} else if pilih == 5 {
		bidang = "Pemasaran"
	} else if pilih == 6 {
		bidang = "Administrasi"
	} else if pilih == 7 {
		bidang = "Logistik"
	} else if pilih == 8 {
		bidang = "Media"
	} else if pilih == 9 {
		bidang = "Kuliner"
	} else if pilih == 10 {
		bidang = "E-Commerce"
	} else {
		fmt.Println("Input tidak valid.")
		return
	}

	fmt.Println("\nDaftar Pekerjaan di bidang", bidang, ":")
	var ditemukan bool = false
	for i := 0; i < maxData; i++ {
		if pekerjaan[i].Bidang == bidang {
			fmt.Printf("\n%d. %s\n", i+1, pekerjaan[i].Nama)
			fmt.Printf("   Tipe Upah     : %s\n", pekerjaan[i].TipeUpah)
			fmt.Printf("   Besar Upah    : Rp%d\n", pekerjaan[i].BesarUpah)
			fmt.Printf("   Deskripsi     : %s\n", pekerjaan[i].Deskripsi)
			fmt.Printf("   Keterampilan  : %s\n", pekerjaan[i].Keterampilan)
			fmt.Printf("   Fleksibel     : %v\n", pekerjaan[i].Fleksibel)
			fmt.Printf("   Alamat/Link   : %s\n", pekerjaan[i].Alamat)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ada pekerjaan di bidang ini.")
	}
}

func menuHitungPenghasilan() {
	fmt.Println("\n=== Hitung Potensi Penghasilan ===")
	fmt.Println("Pilih bidang:")
	fmt.Println("1. Teknologi")
	fmt.Println("2. Desain")
	fmt.Println("3. Konten")
	fmt.Println("4. Pendidikan")
	fmt.Println("5. Pemasaran")
	fmt.Println("6. Administrasi")
	fmt.Println("7. Logistik")
	fmt.Println("8. Media")
	fmt.Println("9. Kuliner")
	fmt.Println("10. E-Commerce")
	fmt.Print("Masukkan pilihan (1-10): ")

	var pilih int
	fmt.Scan(&pilih)

	var bidang string
	if pilih == 1 {
		bidang = "Teknologi"
	} else if pilih == 2 {
		bidang = "Desain"
	} else if pilih == 3 {
		bidang = "Konten"
	} else if pilih == 4 {
		bidang = "Pendidikan"
	} else if pilih == 5 {
		bidang = "Pemasaran"
	} else if pilih == 6 {
		bidang = "Administrasi"
	} else if pilih == 7 {
		bidang = "Logistik"
	} else if pilih == 8 {
		bidang = "Media"
	} else if pilih == 9 {
		bidang = "Kuliner"
	} else if pilih == 10 {
		bidang = "E-Commerce"
	} else {
		fmt.Println("Input tidak valid.")
		return
	}

	count := 0
	for i := 0; i < maxData; i++ {
		if pekerjaan[i].Bidang == bidang {
			count++
			fmt.Printf("%d. %s (%s - Rp%d)\n", i+1, pekerjaan[i].Nama, pekerjaan[i].TipeUpah, pekerjaan[i].BesarUpah)
		}
	}

	if count == 0 {
		fmt.Println("Tidak ada pekerjaan di bidang ini.")
		return
	}

	fmt.Print("\nPilih pekerjaan (nomor sesuai daftar): ")
	var idx int
	fmt.Scan(&idx)

	if idx < 1 || idx > maxData || pekerjaan[idx-1].Bidang != bidang {
		fmt.Println("Input tidak valid.")
		return
	}

	var jumlah int
	fmt.Print("Berapa kali dikerjakan per bulan? ")
	fmt.Scan(&jumlah)

	if jumlah < 0 {
		fmt.Println("Jumlah tidak valid.")
		return
	}

	total := pekerjaan[idx-1].BesarUpah * jumlah
	fmt.Printf("Estimasi penghasilan bulanan: Rp%d\n", total)
}


func menuFavorit() {
	fmt.Println("\n=== Simpan Pekerjaan Favorit ===")
	for i := 0; i < maxData; i++ {
		fmt.Printf("%d. %s (%s - Rp%d)\n", i+1, pekerjaan[i].Nama, pekerjaan[i].TipeUpah, pekerjaan[i].BesarUpah)
	}

	fmt.Print("Pilih pekerjaan untuk disimpan (1-6, 0 untuk lihat favorit): ")
	var idx int
	fmt.Scan(&idx)

	if idx == 0 {
		fmt.Println("\nDaftar Favorit:")
		if totalFavorit == 0 {
			fmt.Println("Belum ada favorit.")
			return
		}
		for i := 0; i < totalFavorit; i++ {
			fmt.Printf("%d. %s (%s - Rp%d)\n", i+1, favorit[i].Nama, favorit[i].TipeUpah, favorit[i].BesarUpah)
		}
		return
	} else if idx < 1 || idx > maxData {
		fmt.Println("Input tidak valid.")
		return
	}

	if totalFavorit >= maxFavorit {
		fmt.Println("Favorit sudah penuh.")
		return
	}

	favorit[totalFavorit] = pekerjaan[idx-1]
	totalFavorit++
	fmt.Println("Pekerjaan ditambahkan ke favorit.")
}

func menuPerbandingan() {
	fmt.Println("\n=== Perbandingan Potensi Pekerjaan ===")

	// Tampilkan semua pekerjaan
	for i := 0; i < maxData; i++ {
		fmt.Printf("%d. %s (%s - Rp%d)\n", i+1, pekerjaan[i].Nama, pekerjaan[i].Bidang, pekerjaan[i].BesarUpah)
	}

	// Pilih dua pekerjaan
	var idx1, idx2 int
	fmt.Print("\nPilih pekerjaan pertama (nomor): ")
	fmt.Scan(&idx1)
	fmt.Print("Pilih pekerjaan kedua (nomor): ")
	fmt.Scan(&idx2)

	if idx1 < 1 || idx1 > maxData || idx2 < 1 || idx2 > maxData || idx1 == idx2 {
		fmt.Println("Pilihan tidak valid atau sama.")
		return
	}

	// Input jumlah per bulan untuk masing-masing
	var jml1, jml2 int
	fmt.Printf("Berapa kali '%s' dikerjakan per bulan? ", pekerjaan[idx1-1].Nama)
	fmt.Scan(&jml1)
	fmt.Printf("Berapa kali '%s' dikerjakan per bulan? ", pekerjaan[idx2-1].Nama)
	fmt.Scan(&jml2)

	total1 := pekerjaan[idx1-1].BesarUpah * jml1
	total2 := pekerjaan[idx2-1].BesarUpah * jml2

	// Hasil perbandingan
	fmt.Println("\n=== Hasil Perbandingan ===")
	fmt.Printf("%s: Rp%d/bulan\n", pekerjaan[idx1-1].Nama, total1)
	fmt.Printf("%s: Rp%d/bulan\n", pekerjaan[idx2-1].Nama, total2)

	if total1 > total2 {
		fmt.Printf("%s menghasilkan lebih banyak.\n", pekerjaan[idx1-1].Nama)
	} else if total2 > total1 {
		fmt.Printf("%s menghasilkan lebih banyak.\n", pekerjaan[idx2-1].Nama)
	} else {
		fmt.Println("Keduanya menghasilkan jumlah yang sama.")
	}
}

func mengurutkanNama() {
    for i := 0; i < maxData-1; i++ {
        minIdx := i
        for j := i + 1; j < maxData; j++ {
            if pekerjaan[j].Nama < pekerjaan[minIdx].Nama {
                minIdx = j
            }
        }
        if minIdx != i {
            pekerjaan[i], pekerjaan[minIdx] = pekerjaan[minIdx], pekerjaan[i]
        }
    }

    fmt.Println("\n=== Pekerjaan setelah diurutkan berdasarkan Nama ===")
    for i := 0; i < maxData; i++ {
        fmt.Printf("%d. %s (%s - Rp%d)\n", i+1, pekerjaan[i].Nama, pekerjaan[i].Bidang, pekerjaan[i].BesarUpah)
    }
}

func menuCariUpah() {
	// Sorting berdasarkan besar upah ascending dengan selection sort
	for i := 0; i < maxData-1; i++ {
		minIdx := i
		for j := i + 1; j < maxData; j++ {
			if pekerjaan[j].BesarUpah < pekerjaan[minIdx].BesarUpah {
				minIdx = j
			}
		}
		if minIdx != i {
			pekerjaan[i], pekerjaan[minIdx] = pekerjaan[minIdx], pekerjaan[i]
		}
	}

	fmt.Println("\nPekerjaan berhasil diurutkan berdasarkan upah (terendah ke tertinggi):")
	for i := 0; i < maxData; i++ {
		fmt.Printf("%d. %s (%s - Rp%d)\n", i+1, pekerjaan[i].Nama, pekerjaan[i].Bidang, pekerjaan[i].BesarUpah)
	}
}



