package main

import ("fmt"
		"net/http")

const MAX TipeInt = 100

type TipeInt int

type Tugas struct {
	judul    string
	kategori string
	deadline string 
	status   string 
}

type DaftarTugas [MAX]Tugas

var listTugas DaftarTugas
var jumlahTugas int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World dari Go Backend!")
}

func main() {
	var pilihan string

	menuAksi := map[string]func(){
		"1": tambahTugas,
		"2": ubahTugas,
		"3": hapusTugas,
		"4": menuPencarian,
		"5": menuPengurutan,
	}

	for {
		cetakMenu()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		if pilihan == "6" {
			fmt.Println("\nTerima kasih telah menggunakan aplikasi ini!")
			return
		}

		if aksi, ditemukan := menuAksi[pilihan]; ditemukan {
			aksi() 
		} else {
			fmt.Println("\nPilihan tidak valid! Silakan masukkan angka dari 1 sampai 6.\n")
		}
	}
}

func cetakMenu() {
	fmt.Println("\n=== APLIKASI MANAJEMEN TUGAS HARIAN ===")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Ubah Tugas")
	fmt.Println("3. Hapus Tugas")
	fmt.Println("4. Cari Tugas")
	fmt.Println("5. Urutkan & Tampilkan Tugas")
	fmt.Println("6. Keluar")
}

func tampilkanSemua() {
	if jumlahTugas == 0 {
		fmt.Println("Daftar tugas masih kosong.")
		return
	}
	fmt.Printf("\n%-3s | %-20s | %-12s | %-12s | %-10s\n", "No", "Judul Tugas", "Kategori", "Deadline", "Status")
	fmt.Println("----------------------------------------------------------------------")
	var i int = 0
	for i < jumlahTugas {
		fmt.Printf("%-3d | %-20s | %-12s | %-12s | %-10s\n", i+1, listTugas[i].judul, listTugas[i].kategori, listTugas[i].deadline, listTugas[i].status)
		i++
	}
}


func tambahTugas() {
	if jumlahTugas >= int(MAX) {
		fmt.Println("Gagal menambahkan tugas, memori penuh!")
		return
	}
	fmt.Println("\n--- TAMBAH TUGAS ---")
	fmt.Print("Judul Tugas: ")
	fmt.Scanln(&listTugas[jumlahTugas].judul)
	fmt.Print("Kategori   : ")
	fmt.Scanln(&listTugas[jumlahTugas].kategori)
	fmt.Print("Deadline (DD-MM-YYYY): ")
	fmt.Scanln(&listTugas[jumlahTugas].deadline)
	listTugas[jumlahTugas].status = "belum" 

	jumlahTugas++
	fmt.Println("Tugas berhasil ditambahkan!")
}

func ubahTugas() {
	fmt.Println("\n--- UBAH TUGAS ---")
	tampilkanSemua()
	if jumlahTugas == 0 {
		return
	}

	var judulCari string
	fmt.Print("Masukkan Judul Tugas yang ingin diubah: ")
	fmt.Scanln(&judulCari)

	var indeks int = cariIndeksSesuaiJudul(judulCari)

	if indeks != -1 {
		var pilihan int
		fmt.Println("\nData ditemukan. Apa yang ingin diubah?")
		fmt.Println("1. Judul")
		fmt.Println("2. Kategori")
		fmt.Println("3. Deadline")
		fmt.Println("4. Status (selesai/belum)")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			fmt.Print("Masukkan Judul Baru: ")
			fmt.Scanln(&listTugas[indeks].judul)
		} else if pilihan == 2 {
			fmt.Print("Masukkan Kategori Baru: ")
			fmt.Scanln(&listTugas[indeks].kategori)
		} else if pilihan == 3 {
			fmt.Print("Masukkan Deadline Baru (DD-MM-YYYY): ")
			fmt.Scanln(&listTugas[indeks].deadline)
		} else if pilihan == 4 {
			var stat int
			fmt.Println("Pilih status: 1. selesai | 2. belum")
			fmt.Print("Pilihan: ")
			fmt.Scanln(&stat)
			if stat == 1 {
				listTugas[indeks].status = "selesai"
			} else {
				listTugas[indeks].status = "belum"
			}
		}
		fmt.Println("Data tugas berhasil diperbarui!")
	} else {
		fmt.Println("Tugas dengan judul tersebut tidak ditemukan.")
	}
}

func hapusTugas() {
	fmt.Println("\n--- HAPUS TUGAS ---")
	tampilkanSemua()
	if jumlahTugas == 0 {
		return
	}

	var judulCari string
	fmt.Print("Masukkan Judul Tugas yang ingin dihapus: ")
	fmt.Scanln(&judulCari)

	var indeks int = cariIndeksSesuaiJudul(judulCari)

	if indeks != -1 {
		var i int = indeks
		for i < jumlahTugas-1 {
			listTugas[i] = listTugas[i+1]
			i++
		}
		jumlahTugas--
		fmt.Println("Tugas berhasil dihapus!")
	} else {
		fmt.Println("Tugas dengan judul tersebut tidak ditemukan.")
	}
}

func cariIndeksSesuaiJudul(judul string) int {
	var i int = 0
	var ditemukan int = -1
	for i < jumlahTugas && ditemukan == -1 {
		if listTugas[i].judul == judul {
			ditemukan = i
		}
		i++
	}
	return ditemukan
}

func menuPencarian() {
	var opsi int
	fmt.Println("\n--- MENU PENCARIAN TUGAS ---")
	fmt.Println("1. Cari Berdasarkan Kategori")
	fmt.Println("2. Cari Berdasarkan Tanggal Deadline")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&opsi)

	if opsi == 1 {
		var kat string
		fmt.Print("Masukkan Kategori yang dicari: ")
		fmt.Scanln(&kat)
		cariBerdasarkanKategoriSequential(kat)
	} else if opsi == 2 {
		var dl string
		fmt.Print("Masukkan Deadline yang dicari (DD-MM-YYYY): ")
		fmt.Scanln(&dl)
		cariBerdasarkanDeadlineBinary(dl)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func cariBerdasarkanKategoriSequential(kategori string) {
	var i int = 0
	var ketemu bool = false
	fmt.Printf("\nHasil pencarian kategori '%s':\n", kategori)
	fmt.Printf("%-3s | %-20s | %-12s | %-12s | %-10s\n", "No", "Judul Tugas", "Kategori", "Deadline", "Status")
	fmt.Println("----------------------------------------------------------------------")

	for i < jumlahTugas {
		if listTugas[i].kategori == kategori {
			fmt.Printf("%-3d | %-20s | %-12s | %-12s | %-10s\n", i+1, listTugas[i].judul, listTugas[i].kategori, listTugas[i].deadline, listTugas[i].status)
			ketemu = true
		}
		i++
	}
	if !ketemu {
		fmt.Println("Tidak ada tugas dalam kategori tersebut.")
	}
}

func cariBerdasarkanDeadlineBinary(deadline string) {
	urutkanDeadlineInsertion(true)

	var left int = 0
	var right int = jumlahTugas - 1
	var indeksKetemu int = -1

	for left <= right && indeksKetemu == -1 {
		var mid int = (left + right) / 2
		if listTugas[mid].deadline == deadline {
			indeksKetemu = mid
		} else if listTugas[mid].deadline < deadline {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if indeksKetemu != -1 {
		fmt.Println("\nTugas ditemukan:")
		fmt.Printf("%-20s | %-12s | %-12s | %-10s\n", "Judul Tugas", "Kategori", "Deadline", "Status")
		fmt.Println("----------------------------------------------------------------------")

		j := indeksKetemu
		for j > 0 && listTugas[j-1].deadline == deadline {
			j--
		}

		for j < jumlahTugas && listTugas[j].deadline == deadline {
			fmt.Printf("%-20s | %-12s | %-12s | %-10s\n", listTugas[j].judul, listTugas[j].kategori, listTugas[j].deadline, listTugas[j].status)
			j++
		}
	} else {
		fmt.Println("Tugas dengan deadline tersebut tidak ditemukan.")
	}
}

func menuPengurutan() {
	if jumlahTugas == 0 {
		fmt.Println("\nTidak dapat mengurutkan dan menampilkan") 
		fmt.Println("daftar tugas masih kosong.")
		return
	}
	
	var kriteria, urutan int
	fmt.Println("\n--- MENU PENGURUTAN TUGAS ---")
	fmt.Println("Urutkan Berdasarkan:")
	fmt.Println("1. Deadline")
	fmt.Println("2. Status")
	fmt.Print("Pilihan Kriteria: ")
	fmt.Scanln(&kriteria)

	fmt.Println("Metode Urutan:")
	fmt.Println("1. Terdekat")
	fmt.Println("2. Terlama")
	fmt.Print("Pilihan Urutan: ")
	fmt.Scanln(&urutan)

	varisAscending := (urutan == 1)

	if kriteria == 1 {
		urutkanDeadlineInsertion(varisAscending)
		fmt.Println("\nBerhasil diurutkan berdasarkan Deadline!")
	} else if kriteria == 2 {
		urutkanStatusSelection(varisAscending)
		fmt.Println("\nBerhasil diurutkan berdasarkan Status!")
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	tampilkanSemua()
}

func urutkanDeadlineInsertion(isAscending bool) {
	var i int = 1
	for i < jumlahTugas {
		key := listTugas[i]
		var j int = i - 1

		if isAscending {
			for j >= 0 && listTugas[j].deadline > key.deadline {
				listTugas[j+1] = listTugas[j]
				j--
			}
		} else {
			for j >= 0 && listTugas[j].deadline < key.deadline {
				listTugas[j+1] = listTugas[j]
				j--
			}
		}
		listTugas[j+1] = key
		i++
	}
}

func urutkanStatusSelection(isAscending bool) {
	var i int = 0
	for i < jumlahTugas-1 {
		idxTarget := i
		var j int = i + 1
		for j < jumlahTugas {
			if isAscending {
				if listTugas[j].status < listTugas[idxTarget].status {
					idxTarget = j
				}
			} else {
				if listTugas[j].status > listTugas[idxTarget].status {
					idxTarget = j
				}
			}
			j++
		}
		// Tukar tempat (Swap)
		temp := listTugas[i]
		listTugas[i] = listTugas[idxTarget]
		listTugas[idxTarget] = temp
		i++
	}
}
