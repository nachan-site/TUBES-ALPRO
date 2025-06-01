package main

import "fmt"

type Pakaian struct {
	nama     string
	kategori string
	warna    string
	musim    string
	jenis    string
	lainnya  string
}

type outfit struct {
	hari           string
	tanggal        int
	bulan          int
	tahun          int
	atasan         string
	bawahan        string
	sepatu         string
	tambahan       string
	warna_atasan   string
	warna_bawahan  string
	warna_sepatu   string
	warna_tambahan string
}

// konstanta untuk kapasitas maksimum pakaian di lemari
const NMAX int = 1000

type dafPakaian [NMAX]Pakaian
type dafOutfit [7]outfit

var pilih int
var nPakaian, nOutfit int
var daftar dafPakaian
var dOutfit dafOutfit

//var dCari dafCari
var kata string

func main() {
	menu()
}

//procedure menu utama dari semua fitur
func menu() {
	// IS: pilih tidak terdefinisi
	// FS: nilai pilih diperbarui berdasarkan input pengguna, sehingga procedure dipanggil berdasarkan pilihan pengguna

	var nama, warna, namaBaru string
	fmt.Println("\n----------------------------------------")
	fmt.Println("              OOTD Planner              ")
	fmt.Println("----------------------------------------")
	fmt.Println("1. Tambah Pakaian")
	fmt.Println("2. Lihat Isi Lemari")
	fmt.Println("3. Buat Outfit")
	fmt.Println("4. Lihat Outfit")
	fmt.Println("5. Edit atau Hapus Pakaian")
	fmt.Println("6. Urutkan Pakaian berdasarkan Tingkat Formalitas")
	fmt.Println("0. Keluar")
	fmt.Println("----------------------------------------")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		tambahkanPakaian(&daftar, &nPakaian)
	} else if pilih == 2 {
		tampilkanPakaian(daftar, nPakaian)
	} else if pilih == 3 {
		buatOutfit(daftar, nPakaian)
	} else if pilih == 4 {
		fmt.Println("1. Lihat Outfit secara keseluruhan \n2. Cari jadwal Outfit berdasarkan tanggal \n3. Lihat jadwal Outfit terdekat \n4. Lihat jadwal Outfit terjauh \n0. Kembali")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			lihatOutfit(&dOutfit, nOutfit)
		} else if pilih == 2 {
			tanggalOutfit(&dOutfit, nOutfit)
		} else if pilih == 3 {
			jadwalTerdekat(&dOutfit, nOutfit)
		} else if pilih == 4 {
			jadwalTerjauh(&dOutfit, nOutfit)
		} else {
			menu()
		}
	} else if pilih == 5 {
		fmt.Println("1. Edit Pakaian \n2. Hapus Pakaian \n0. Kembali")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			renamePakaian(&daftar, nPakaian, &nama, &warna, &namaBaru)
		} else if pilih == 2 {
			hapusPakaian(&daftar, &nPakaian, &nama, &warna)
		} else {
			menu()
		}
	} else if pilih == 6 {
		urutkanFormalitas(&daftar, nPakaian)
	} else if pilih == 0 {
		fmt.Println("Okey! Kalau mau ootd-an, kesini lagi aja ya!\n")
	}

}

// procedure untuk menambah pakaian ke lemari
func tambahkanPakaian(A *dafPakaian, n *int) {
	// IS: array daftar pakaian (A) dan jumlah array pakaian (n) tidak terdifinisi
	// FS: array daftar pakaian (A) terisi dan jumlah array pakaian (n) diperbarui sesuai dengan berapa banyak total inputan pengguna

	var noKat, noMus, noJen int
	var nama, kategori, warna, musim, jenis, lainnya string

	for *n < NMAX && pilih == 1 {
		fmt.Print("\nMasukkan nama pakaian: ")
		fmt.Scan(&nama)

		fmt.Print("Pilih nomor kategori (1.Atasan | 2.Bawahan | 3.Sepatu | 4.Lainnya): ")
		fmt.Scan(&noKat)
		if noKat == 1 {
			kategori = "atasan"
		} else if noKat == 2 {
			kategori = "bawahan"
		} else if noKat == 3 {
			kategori = "sepatu"
		} else {
			kategori = "lainnya"
			fmt.Print("Masukkan nama kategori: ")
			fmt.Scan(&lainnya)
		}

		fmt.Print("Warna: ")
		fmt.Scan(&warna)

		fmt.Print("Cocok dipakai di musim apa? (1.Panas | 2.Dingin | 3.Semua): ")
		fmt.Scan(&noMus)
		if noMus == 1 {
			musim = "panas"
		} else if noMus == 2 {
			musim = "dingin"
		} else {
			musim = "semua"
		}

		fmt.Print("Pilih tingkat formalitas? (1.Formal | 2.Semi-Formal | 3.Non-Formal): ")
		fmt.Scan(&noJen)
		if noJen == 1 {
			jenis = "formal"
		} else if noJen == 2 {
			jenis = "semi-formal"
		} else {
			jenis = "tidak formal (kasual)"
		}
		A[*n].nama = nama
		A[*n].kategori = kategori
		A[*n].warna = warna
		A[*n].musim = musim
		A[*n].jenis = jenis
		A[*n].lainnya = lainnya
		*n++

		fmt.Printf("Item '%s %s' berhasil ditambahkan ke lemari. \n", nama, warna)
		fmt.Println()
		fmt.Println("Mau tambah item lagi, ga?")
		fmt.Println("1. Tambah lagi dong!")
		fmt.Println("2. Udah dulu deh")
		fmt.Println("3. Lihat isi lemari")
		fmt.Print("Pilih nomornya ya: ")
		fmt.Scan(&pilih)
		if pilih == 2 {
			menu()
		} else if pilih == 3 {
			tampilkanPakaian(daftar, nPakaian)
		} else {
			for pilih != 1 && pilih != 2 && pilih != 3 {
				fmt.Println("Ketik nomornya sesuai yang ada di pilihan aja ya ^^")
				fmt.Scan(&pilih)
				if pilih == 2 {
					menu()
				} else if pilih == 3 {
					tampilkanPakaian(daftar, nPakaian)
				}
			}
		}
	}
}

// procedure untuk melihat item apa saja yang ada di lemari setelah diinput oleh pengguna
func tampilkanPakaian(A dafPakaian, n int) {
	// IS: array daftar pakaian (A) dan jumlah array pakaian (n) terdifinisi
	// FS: menampilkan array daftar pakaian (A)
	// Jika pengguna belum memasukkan pakaian, maka akan ada pilihan untuk memanggil procedure tambahkanPakaian

	fmt.Println("\n----------------------------------------")
	fmt.Println(" Ini dia daftar pakaian di lemari kamu!")
	fmt.Println("----------------------------------------")

	if n == 0 {
		fmt.Println("Yahh lemari kamu masih kosong nih :(")
		fmt.Println("\nIsi lemari kamu dulu, yuk! ^^")
		fmt.Println("Ketik 1 untuk isi lemari!")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahkanPakaian(&daftar, &nPakaian)
		} else {
			for pilih != 1 {
				fmt.Println("ketik 1 bang, bukan yang lain :( ")
				fmt.Scan(&pilih)
				if pilih == 1 {
					tambahkanPakaian(&daftar, &nPakaian)
				}
			}
		}
	} else {
		fmt.Println("KATEGORI ATASAN")
		for i := 0; i < n; i++ {
			if A[i].kategori == "atasan" {
				fmt.Printf("- %s %s (%s, cocok dipakai di cuaca: %s)\n", A[i].nama, A[i].warna, A[i].jenis, A[i].musim)
			}
		}
		fmt.Println("\nKATEGORI BAWAHAN")
		for i := 0; i < n; i++ {
			if A[i].kategori == "bawahan" {
				fmt.Printf("- %s %s (%s, cocok dipakai di cuaca: %s)\n", A[i].nama, A[i].warna, A[i].jenis, A[i].musim)
			}
		}
		fmt.Println("\nKATEGORI SEPATU")
		for i := 0; i < n; i++ {
			if A[i].kategori == "sepatu" {
				fmt.Printf("- %s %s (%s, cocok dipakai di cuaca: %s)\n", A[i].nama, A[i].warna, A[i].jenis, A[i].musim)
			}
		}
		fmt.Println("\nKATEGORI LAINNYA")
		for i := 0; i < n; i++ {
			if A[i].kategori == "lainnya" {
				fmt.Printf("- %s %s (%s, cocok dipakai di cuaca: %s)\n", A[i].nama, A[i].warna, A[i].jenis, A[i].musim)
			}
		}
		fmt.Println("----------------------------------------\n")
		fmt.Println("Ketik 1 untuk kembali ke menu utama!")
		fmt.Scan(&pilih)
		if pilih == 1 {
			menu()
		}
	}

}

// function untuk mencari index pakaian menggunakan sequential search
func cariIndexPakaian(A dafPakaian, n int, nama, warna string) int {
	// IS: array daftar pakaian (A), jumlah array pakaian (n), nama serta warna pakaian terdifinisi
	// Proses: mencari data di array yang sesuai dengan nama dan warna
	// FS: mengembalikan index dari data yang dicari
	idx := -1
	for i := 0; i < n; i++ {
		if A[i].nama == nama && A[i].warna == warna {
			idx = i
		}
	}
	return idx
}

// procedure untuk menghapus item pakaian
func hapusPakaian(A *dafPakaian, n *int, nama, warna *string) {
	// IS: array daftar pakaian (A), jumlah array pakaian (n), nama serta warna pakaian terdifinisi
	// Proses: menggunakan function cariIndexPakaian untuk mencari index pakaian yang ingin dihapus, jika ada maka item akan dihapus
	// FS: array daftar pakaian (A) diperbarui, dan jumlah array pakaian (n) berkurang 1
	var i, j int

	fmt.Println("\n----------------------------------------")
	for k := 0; k < *n; k++ {
		if A[k].kategori == "atasan" {
			fmt.Printf("- %s %s (%s, cocok dipakai di cuaca: %s)\n", A[k].nama, A[k].warna, A[k].jenis, A[k].musim)
		} else if A[k].kategori == "bawahan" {
			fmt.Printf("- %s %s (%s, cocok dipakai di cuaca: %s)\n", A[k].nama, A[k].warna, A[k].jenis, A[k].musim)
		} else if A[k].kategori == "sepatu" {
			fmt.Printf("- %s %s (%s, cocok dipakai di cuaca: %s)\n", A[k].nama, A[k].warna, A[k].jenis, A[k].musim)
		} else if A[k].kategori == "lainnya" {
			fmt.Printf("- %s %s (%s, cocok dipakai di cuaca: %s)\n", A[k].nama, A[k].warna, A[k].jenis, A[k].musim)
		}
	}

	fmt.Println("----------------------------------------")
	fmt.Print("\nMasukkan nama + warna pakaian yang mau kamu hapus: ")
	fmt.Scan(&*nama, &*warna)
	i = cariIndexPakaian(*A, *n, *nama, *warna)
	if i != -1 {
		fmt.Println("Apakah kamu yakin ingin menghapus pakaian ini?")
		fmt.Printf("%s %s (%s, cocok dipakai pada cuaca: %s)\n", A[i].nama, A[i].warna, A[i].jenis, A[i].musim)
		fmt.Println("1. Ya \n2. Tidak")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			for j = i; j < *n; j++ {
				A[i] = A[i+1]
			}
			*n--
			fmt.Println("Item telah dihapus!")
		}

	} else {
		fmt.Println("Item tidak ditemukan.")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Masih ada pakaian yang mau di-hapus lagi, ga?")
	fmt.Println("1. Masih ada!")
	fmt.Println("2. Udah, itu aja")
	fmt.Print("Pilih nomornya ya: ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		hapusPakaian(A, n, nama, warna)
	} else {
		menu()
	}
}

// procedure untuk mengedit item pada pakaian
func renamePakaian(A *dafPakaian, n int, nama *string, warna *string, itembaru *string) {
	// IS: array daftar pakaian (A), jumlah array pakaian (n) pakaian terdifinisi
	// Proses: menggunakan function cariIndexPakaian untuk mencari index pakaian yang ingin di-rename
	// FS: array daftar pakaian (A) diperbarui, dan data dari index tersebut akan diperbarui

	var i, noKat, noMus, noJen int

	fmt.Println("\n----------------------------------------")
	for k := 0; k < n; k++ {
		if A[k].kategori == "atasan" {
			fmt.Printf("- %s %s (%s, cocok dipakai pada cuaca: %s)\n", A[k].nama, A[k].warna, A[k].jenis, A[k].musim)
		} else if A[k].kategori == "bawahan" {
			fmt.Printf("- %s %s (%s, cocok dipakai pada cuaca: %s)\n", A[k].nama, A[k].warna, A[k].jenis, A[k].musim)
		} else if A[k].kategori == "sepatu" {
			fmt.Printf("- %s %s (%s, cocok dipakai pada cuaca: %s)\n", A[k].nama, A[k].warna, A[k].jenis, A[k].musim)
		} else if A[k].kategori == "lainnya" {
			fmt.Printf("- %s %s (%s, cocok dipakai pada cuaca: %s)\n", A[k].nama, A[k].warna, A[k].jenis, A[k].musim)
		}
	}

	fmt.Println("----------------------------------------")
	fmt.Print("Mausukkan nama dan warna pakaian yang ingin di-edit: ")
	fmt.Scan(nama, warna)
	i = cariIndexPakaian(*A, n, *nama, *warna)
	if i != -1 {
		fmt.Println("\nApakah kamu yakin ingin mengedit pakaian ini?")
		fmt.Printf("%s %s (%s, cocok dipakai pada cuaca: %s)\n", A[i].nama, A[i].warna, A[i].jenis, A[i].musim)
		fmt.Println("1. Ya \n2. Tidak")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			fmt.Println("\nMau edit apanya, nih?")
			fmt.Println("1. Nama \n2. Kategori \n3. Warna \n4. Musim \n5. Jenis")
			fmt.Print("Pilih: ")
			fmt.Scan(&pilih)

			if pilih == 1 {
				fmt.Print("\nMasukkan nama pakaian yang baru: ")
				fmt.Scan(itembaru)
				A[i].nama = *itembaru
			} else if pilih == 2 {
				fmt.Print("Pilih nomor kategori yang baru (1.Atasan | 2.Bawahan | 3.Sepatu | 4.Lainnya): ")
				fmt.Scan(&noKat)
				if noKat == 1 {
					A[i].kategori = "atasan"
				} else if noKat == 2 {
					A[i].kategori = "bawahan"
				} else if noKat == 3 {
					A[i].kategori = "sepatu"
				} else {
					A[i].kategori = "lainnya"
					fmt.Print("Masukkan nama kategori: ")
					fmt.Scan(&A[i].lainnya)
				}
			} else if pilih == 3 {
				fmt.Print("Warna: ")
				fmt.Scan(itembaru)
				A[i].warna = *itembaru
			} else if pilih == 4 {
				fmt.Print("Cocok dipakai di musim apa? (1.Panas | 2.Dingin | 3.Semua): ")
				fmt.Scan(&noMus)
				if noMus == 1 {
					A[i].musim = "panas"
				} else if noMus == 2 {
					A[i].musim = "dingin"
				} else {
					A[i].musim = "semua"
				}
			} else if pilih == 5 {
				fmt.Print("Pilih tingkat formalitas? (1.Formal | 2.Semi-Formal | 3.Non-Formal): ")
				fmt.Scan(&noJen)
				if noJen == 1 {
					A[i].jenis = "formal"
				} else if noJen == 2 {
					A[i].jenis = "semi-Formal"
				} else {
					A[i].jenis = "non-Formal"
				}
			}

			fmt.Printf("Pakaian berhasil di-edit menjadi %s %s! \n(%s, yang cocok dipakai pada cuaca: %s)\n", A[i].nama, A[i].warna, A[i].jenis, A[i].musim)
		}
	} else {
		fmt.Println("Pakaian tidak ditemukan")
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Masih ada pakaian yang mau di-edit lagi, ga?")
	fmt.Println("1. Masih ada!")
	fmt.Println("2. Udah, itu aja")
	fmt.Print("Pilih nomornya ya: ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		renamePakaian(A, n, nama, warna, itembaru)
	} else {
		menu()
	}
}

// mengurutkan pakaian berdasarkan tingkat formalitas menggunakan Selection Sort
func urutkanFormalitas(A *dafPakaian, n int) {
	// IS: array daftar pakaian (A) dan jumlah array pakaian (n) terdifinisi
	// FS: array daftar pakaian (A) diurutkan berdasarkan tingkat formalitas (formal ke non formal)

	fmt.Println("\n--------------------------------------------------")
	fmt.Println("Pakaian diurutkan berdasarkan tingkat formalitasnya")
	fmt.Println("---------------------------------------------------")

	var i, idx, pass int
	var temp Pakaian

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].jenis < A[idx].jenis {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}

	for k := 0; k < n; k++ {
		fmt.Printf("- %s %s (%s, cocok dipakai di cuaca: %s)\n", A[k].nama, A[k].warna, A[k].jenis, A[k].musim)
	}

	fmt.Println("---------------------------------------------------\n")
	fmt.Println("Ketik 1 untuk kembali ke menu utama!")
	fmt.Scan(&pilih)
	if pilih == 1 {
		menu()
	}
}

// procedure untuk membuat Outfit
func buatOutfit(A dafPakaian, n int) {
	// IS: array daftar pakaian (A) dan jumlah array pakaian (n) terdifinisi
	// Proses: memanggil procedure pilihKategori untuk memilih kombinasi item untuk satu outfit
	// FS: -

	fmt.Println("\n--------------------------------------------")
	fmt.Println("Ayo buat Outfit yang sesuai dengan style kamu!")
	fmt.Println("--------------------------------------------")
	pilihKategori(daftar, nPakaian, &dOutfit, &nOutfit)

	fmt.Println()
	fmt.Println("Mau buat outfit lagi, ga?")
	fmt.Println("1. Buat lagi!")
	fmt.Println("2. Udah dulu")
	fmt.Println("3. Lihat daftar outfit")
	fmt.Print("Pilih nomornya ya: ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		buatOutfit(A, n)
	} else if pilih == 2 {
		menu()
	} else if pilih == 3 {
		lihatOutfit(&dOutfit, nOutfit)
	}
}

// procedure untuk memilih pakaian berdasarkan kategorinya
func pilihKategori(A dafPakaian, np int, B *dafOutfit, no *int) {
	// IS: array daftar pakaian (A) dan jumlah array pakaian (n) terdifinisi, serta kategori dari masing-masing array terdefinisi
	// Proses: program akan menampilkan array daftar pakaian dan pengguna akan menginput nama pakaian berdasarkan kategori
	// FS: menampilkan kombinasi pakaian yang telah diisi pengguna sebelumnya, yaitu atasan, bawahan, sepatu, dan tambahan lainnya.
	var atasan, bawahan, sepatu, tambahan, wAtasan, wBawahan, wSepatu, wTambahan, hari string
	var tanggal, bulan, tahun int
	for *no < NMAX {
		fmt.Println("Mau dipake kapan Outfit-nya?")

		fmt.Print("Masukkan hari: ")
		fmt.Scan(&hari)
		fmt.Print("Masukkan tanggal, bulan, tahun: ")
		fmt.Scan(&tanggal, &bulan, &tahun)

		B[*no].hari = hari
		B[*no].tanggal = tanggal
		B[*no].bulan = bulan
		B[*no].tahun = tahun

		fmt.Println("\n!!PERHATIAN!!")
		fmt.Println("Ketik sesuai format Nama + Warna, ya!")
		fmt.Println("Kalau mau kosongin, ketik - - aja")

		fmt.Println("\nPilih atasan yang tersedia!")
		for i := 0; i < np; i++ {
			if A[i].kategori == "atasan" {
				fmt.Printf("- %s %s \n", A[i].nama, A[i].warna)
			}

		}

		fmt.Print("Nama atasan: ")
		fmt.Scan(&atasan)
		fmt.Scan(&wAtasan)
		B[*no].atasan = atasan
		B[*no].warna_atasan = wAtasan

		fmt.Println("\nPilih bawahan yang tersedia!")
		for i := 0; i < np; i++ {
			if A[i].kategori == "bawahan" {
				fmt.Printf("- %s %s \n", A[i].nama, A[i].warna)
			}

		}
		fmt.Print("Nama bawahan: ")
		fmt.Scan(&bawahan)
		fmt.Scan(&wBawahan)
		B[*no].bawahan = bawahan
		B[*no].warna_bawahan = wBawahan

		fmt.Println("\nPilih sepatu yang tersedia!")
		for i := 0; i < np; i++ {
			if A[i].kategori == "sepatu" {
				fmt.Printf("- %s %s \n", A[i].nama, A[i].warna)
			}

		}
		fmt.Print("Nama sepatu: ")
		fmt.Scan(&sepatu)
		fmt.Scan(&wSepatu)
		B[*no].sepatu = sepatu
		B[*no].warna_sepatu = wSepatu

		fmt.Println("\nPilih tambahan lainnya yang tersedia!")
		for i := 0; i < np; i++ {
			if A[i].kategori == "lainnya" {
				fmt.Printf("- %s %s (%s) \n", A[i].nama, A[i].warna, A[i].lainnya)
			}

		}
		fmt.Print("Nama tambahan: ")
		fmt.Scan(&tambahan)
		fmt.Scan(&wTambahan)
		B[*no].tambahan = tambahan
		B[*no].warna_tambahan = wTambahan
		*no++

		fmt.Println("\nYeay! Oufit kamu udah jadi nih! >⩊<")
		fmt.Println("Atasan:", atasan, wAtasan)
		fmt.Println("Bawahan:", bawahan, wBawahan)
		fmt.Println("Sepatu:", sepatu, wSepatu)
		fmt.Println("Tambahan:", tambahan, wTambahan)
		fmt.Printf("Dipakai pada hari: %s, tanggal: %d-%d-%d\n", hari, tanggal, bulan, tahun)
		fmt.Println("--------------------------------------------")
		return
	}

}

// procedure untuk melihat outfit yang terlah dibuat
func lihatOutfit(B *dafOutfit, no int) {
	// IS: array daftar outfit (B) dan jumlah array outfit (no) terdifinisi
	// FS: menampilkan daftar outfit berdasarkan inputan terakhir user
	fmt.Println("\n----------------------------------------")
	fmt.Println("     Daftar Outfit yang kamu buat!")
	fmt.Println("----------------------------------------")

	urutkanTanggal(B, no)
	for i := 0; i < no; i++ {
		fmt.Printf("%s, %v-%v-%v\n", B[i].hari, B[i].tanggal, B[i].bulan, B[i].tahun)
		fmt.Println("Atasan:", B[i].atasan, B[i].warna_atasan)
		fmt.Println("Bawahan:", B[i].bawahan, B[i].warna_bawahan)
		fmt.Println("Sepatu:", B[i].sepatu, B[i].warna_sepatu)
		fmt.Println("Tambahan:", B[i].tambahan, B[i].warna_tambahan)
		fmt.Println("----------------------------------------")
	}

	fmt.Println("Ketik 1 untuk kembali ke menu utama!")
	fmt.Scan(&pilih)
	if pilih == 1 {
		menu()
	}
}

// mengurutkan outfit berdasarkan tanggal pemakaian secara ascending menggunakan Insertion Sort
func urutkanTanggal(B *dafOutfit, n int) {
	// IS: array daftar outfit (B) dan jumlah array outfit (no) terdifinisi
	// FS: array daftar outfit (B) diurutkan berdasarkan tanggal secara ascending
	var i, pass int
	var temp outfit

	pass = 1
	for pass < n {
		i = pass
		temp = B[pass]
		for (i > 0 && temp.tahun < B[i-1].tahun) || (i > 0 && temp.tahun == B[i-1].tahun && temp.bulan < B[i-1].bulan) || (i > 0 && temp.bulan == B[i-1].bulan && temp.tanggal < B[i-1].tanggal) {
			B[i] = B[i-1]
			i = i - 1
		}
		B[i] = temp
		pass = pass + 1
	}
}

// mencari tanggal menggunakan Binary Search
func cariTanggal(B *dafOutfit, no, tgl, bln, thn int) int {
	// IS: array daftar outfit (B), jumlah array outfit (no), tanggal (tgl), bulan (bln), tahun (thn) dipakainya outfit terdifinisi
	// Proses: mencari data di array yang sesuai dengan  tanggal (tgl), bulan (bln), tahun (thn)
	// FS: mengembalikan index dari data yang dicari

	urutkanTanggal(B, no)
	var left, right, mid int
	var idx int

	left = 0
	right = no - 1
	idx = -1

	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if (thn < B[mid].tahun) || (thn == B[mid].tahun && bln < B[mid].bulan) || (thn == B[mid].tahun && bln == B[mid].bulan && tgl < B[mid].tanggal) {
			right = mid - 1
		} else if (thn > B[mid].tahun) || (thn == B[mid].tahun && bln > B[mid].bulan) || (thn == B[mid].tahun && bln == B[mid].bulan && tgl > B[mid].tanggal) {
			left = mid + 1
		} else {
			idx = mid
		}
	}

	return idx

}

// menampilkan hasil pencarian berdasarkan tanggal dipakai outfitnya
func tanggalOutfit(B *dafOutfit, no int) {
	// IS: array daftar outfit (B) dan jumlah array outfit (no) terdifinisi
	// Proses: memanggil function cariTanggal untuk mencari index tanggal yang diinginkan
	// FS: menampilkan outfit yang dibuat pada tanggal yang dicari

	var tanggal, bulan, tahun, i int

	fmt.Println("----------------------------------------")
	fmt.Print("Masukkan tanggal, bulan, dan tahun yang ingin dicari: ")
	fmt.Scan(&tanggal, &bulan, &tahun)
	i = cariTanggal(B, no, tanggal, bulan, tahun)
	if i == -1 {
		fmt.Println("Kamu belum buat Outfit untuk tanggal ini.")
	} else {
		fmt.Println("\n----------------------------------------")
		fmt.Println("      Daftar Outfit yang kamu buat!")
		fmt.Println("----------------------------------------")
		fmt.Printf("Hari: %s, Tanggal: %v-%v-%v\n", B[i].hari, B[i].tanggal, B[i].bulan, B[i].tahun)
		fmt.Println("Atasan:", B[i].atasan, B[i].warna_atasan)
		fmt.Println("Bawahan:", B[i].bawahan, B[i].warna_bawahan)
		fmt.Println("Sepatu:", B[i].sepatu, B[i].warna_sepatu)
		fmt.Println("Tambahan:", B[i].tambahan, B[i].warna_tambahan)

	}

	fmt.Println("----------------------------------------")
	fmt.Println("Ketik 1 untuk kembali ke menu utama!")
	fmt.Scan(&pilih)
	if pilih == 1 {
		menu()
	}
}

// mencari indeks jadwal terdekat outfit akan dipakai menggunakan Nilai Ekstrim Minimum
func Terdekat(B dafOutfit, no int) int {
	// IS: array daftar outfit (B) dan jumlah array outfit (no) terdifinisi
	// FS: mengembalikan index dari tanggal yang paling dekat/cepat

	var min int = 0
	for i := 0; i < no; i++ {
		if (B[i].tahun < B[min].tahun) || (B[i].tahun == B[min].tahun && B[i].bulan < B[min].bulan) || (B[i].tahun == B[min].tahun && B[i].bulan == B[min].bulan && B[i].bulan < B[min].tanggal) {
			min = i
		}
	}
	return min
}

// mencari indeks jadwal terjauh outfit akan dipakai menggunakan Nilai Ekstrim Maksimum
func Terjauh(B dafOutfit, no int) int {
	// IS: array daftar outfit (B) dan jumlah array outfit (no) terdifinisi
	// FS: mengembalikan index dari tanggal yang paling jauh/lama

	var max int = 0
	for i := 0; i < no; i++ {
		if (B[i].tahun > B[max].tahun) || (B[i].tahun == B[max].tahun && B[i].bulan > B[max].bulan) || (B[i].tahun == B[max].tahun && B[i].bulan == B[max].bulan && B[i].tanggal > B[max].tanggal) {
			max = i
		}
	}
	return max
}

// menampilkan daftar outfit terdekat outfit akan dipakai
func jadwalTerdekat(B *dafOutfit, no int) {
	// IS: array daftar outfit (B) dan jumlah array outfit (no) terdifinisi
	// FS: menampilkan daftar outfit berdasarkan tanggal terdekat/tercepat yang akan dipakai

	var i int

	fmt.Println("\n----------------------------------------")
	fmt.Println("        Daftar Outfit terdekat")
	fmt.Println("----------------------------------------")

	urutkanTanggal(B, no)
	i = Terdekat(*B, no)

	fmt.Printf("Hari: %s, Tanggal: %v-%v-%v\n", B[i].hari, B[i].tanggal, B[i].bulan, B[i].tahun)
	fmt.Println("Atasan:", B[i].atasan, B[i].warna_atasan)
	fmt.Println("Bawahan:", B[i].bawahan, B[i].warna_bawahan)
	fmt.Println("Sepatu:", B[i].sepatu, B[i].warna_sepatu)
	fmt.Println("Tambahan:", B[i].tambahan, B[i].warna_tambahan)

	fmt.Println("----------------------------------------")
	fmt.Println("Ketik 1 untuk kembali ke menu utama!")
	fmt.Scan(&pilih)
	if pilih == 1 {
		menu()
	}
}

// menampilkan daftar outfit terjauh outfit akan dipakai
func jadwalTerjauh(B *dafOutfit, no int) {
	// IS: array daftar outfit (B) dan jumlah array outfit (no) terdifinisi
	// FS: menampilkan daftar outfit berdasarkan tanggal terjauh/terlama yang akan dipakai

	var i int

	fmt.Println("\n----------------------------------------")
	fmt.Println("        Daftar Outfit terjauh")
	fmt.Println("----------------------------------------")

	urutkanTanggal(B, no)
	i = Terjauh(*B, no)

	fmt.Printf("Hari: %s, Tanggal: %v-%v-%v\n", B[i].hari, B[i].tanggal, B[i].bulan, B[i].tahun)
	fmt.Println("Atasan:", B[i].atasan, B[i].warna_atasan)
	fmt.Println("Bawahan:", B[i].bawahan, B[i].warna_bawahan)
	fmt.Println("Sepatu:", B[i].sepatu, B[i].warna_sepatu)
	fmt.Println("Tambahan:", B[i].tambahan, B[i].warna_tambahan)

	fmt.Println("----------------------------------------")
	fmt.Println("Ketik 1 untuk kembali ke menu utama!")
	fmt.Scan(&pilih)
	if pilih == 1 {
		menu()
	}
}
