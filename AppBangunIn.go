package main

import "fmt"

const NMAX = 100

type Supplier struct {
	ID       int
	NamaPT   string
	Kontak   string
	Rating   float64
	Provinsi string
	Kota     string

	RiwayatTanggal  [NMAX]string
	RiwayatMaterial [NMAX]string
	RiwayatJumlah   [NMAX]int
	JumlahRiwayat   int
}

type Suppliers [NMAX]Supplier

func main() {
	var daftarSupplier Suppliers
	var JumlahSupplier int = 0
	var pilihanMenu int
	var keluar bool = false

	for !keluar {
		fmt.Println("\n======================================================")
		fmt.Println("       APLIKASI MANAJEMEN SUPPLIER BANGUNIN     ")
		fmt.Println("========================================================")
		fmt.Println("1. Kelola Data Supplier (Tambah/Hapus/Ubah)  ")
		fmt.Println("2. Tampilkan Semua Daftar Supplier")
		fmt.Println("3. Cari Data Supplier di Daftar Supplier")
		fmt.Println("4. Riwayat Pelayanan Supplier")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih Opsi (1-5) : ")
		fmt.Scan(&pilihanMenu)
		fmt.Println("--------------------------------------")

		switch pilihanMenu {
		case 1:
			KelolaSupplier(&daftarSupplier, &JumlahSupplier)
		case 2:
			PrintSuppliers(&daftarSupplier, JumlahSupplier)
		case 3:
			CariDataSupplier(&daftarSupplier, JumlahSupplier)
		case 4:
			RiwayatPelayananSupplier(&daftarSupplier, JumlahSupplier)
		case 5:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
			keluar = true
		default:
			fmt.Println("Opsi tidak valid. Silakan pilih antara 1-3.")
		}
	}
}

func KelolaSupplier(suppliers *Suppliers, JumlahSupplier *int) {
	var indexHapus, pilihan int
	var idCari int
	var i int
	var ketemu bool
	var keluar bool = false

	for !keluar {
		fmt.Println("\n=== Kelola Supplier ===")
		fmt.Println("1. Tambah Supplier")
		fmt.Println("2. Hapus Supplier")
		fmt.Println("3. Ubah Supplier")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Pilih Opsi (1-4) : ")
		fmt.Scan(&pilihan)
		fmt.Println("------------------------------")

		switch pilihan {
		case 1:
			if *JumlahSupplier >= NMAX {
				fmt.Println("Jumlah supplier sudah mencapai batas maksimum.")

			} else {
				fmt.Println("Tambah Supplier")

				fmt.Print("Masukkan ID Supplier: ")
				fmt.Scan(&(*suppliers)[*JumlahSupplier].ID)

				fmt.Print("Masukkan Nama PT Supplier (Tanpa Spasi): ")
				fmt.Scan(&(*suppliers)[*JumlahSupplier].NamaPT)

				fmt.Print("Masukkan Provinsi Supplier (Tanpa Spasi): ")
				fmt.Scan(&(*suppliers)[*JumlahSupplier].Provinsi)

				fmt.Print("Masukkan Kota Supplier (Tanpa Spasi): ")
				fmt.Scan(&(*suppliers)[*JumlahSupplier].Kota)

				fmt.Print("Masukkan Kontak Supplier (Nomor Telepon) (Tanpa Spasi): ")
				fmt.Scan(&(*suppliers)[*JumlahSupplier].Kontak)

				fmt.Print("Masukkkan Rating Supplier (0.0 - 5.0): ")
				fmt.Scan(&(*suppliers)[*JumlahSupplier].Rating)

				*JumlahSupplier++
				fmt.Println("Supplier berhasil ditambahkan.")
			}
		case 2:
			if *JumlahSupplier == 0 {
				fmt.Println("Tidak ada supplier yang dapat dihapus.")
			} else {
				fmt.Print("Masukkan ID Supplier yang ingin dihapus: ")
				fmt.Scan(&idCari)

				ketemu = false
				i = 0
				for i < *JumlahSupplier && !ketemu {
					if (*suppliers)[i].ID == idCari {
						ketemu = true
						indexHapus = i
					} else {
						i++
					}
				}
				if ketemu {
					i = indexHapus
					for i < *JumlahSupplier-1 {
						(*suppliers)[i] = (*suppliers)[i+1]
						i++
					}
					*JumlahSupplier--
					fmt.Println("Supplier berhasil dihapus.")
				} else {
					fmt.Println("Supplier dengan ID", idCari, "tidak ditemukan.")
				}
			}
		case 3:
			if *JumlahSupplier == 0 {
				fmt.Println("Tidak ada supplier yang dapat diubah.")
			} else {
				fmt.Print("Masukkan ID Supplier yang ingin diubah: ")
				fmt.Scan(&idCari)

				ketemu = false
				i = 0
				for i < *JumlahSupplier && !ketemu {
					if (*suppliers)[i].ID == idCari {
						ketemu = true
					} else {
						i++
					}
				}
				if ketemu {
					fmt.Print("Masukkan ID Supplier Baru (Angka): ")
					fmt.Scan(&(*suppliers)[i].ID)

					fmt.Print("Masukkan Nama PT Supplier (Tanpa Spasi): ")
					fmt.Scan(&(*suppliers)[i].NamaPT)

					fmt.Print("Masukkan Provinsi Supplier (Tanpa Spasi): ")
					fmt.Scan(&(*suppliers)[i].Provinsi)

					fmt.Print("Masukkan Kota Supplier (Tanpa Spasi): ")
					fmt.Scan(&(*suppliers)[i].Kota)

					fmt.Print("Masukkan Kontak Supplier (Nomor Telepon) (Tanpa Spasi): ")
					fmt.Scan(&(*suppliers)[i].Kontak)

					fmt.Print("Masukkkan Rating Supplier (0.0 - 5.0): ")
					fmt.Scan(&(*suppliers)[i].Rating)

					fmt.Println("Supplier berhasil diubah.")
				} else {
					fmt.Println("Supplier dengan ID", idCari, "tidak ditemukan.")
				}
			}
		case 4:
			keluar = true
		default:
			fmt.Println("Opsi tidak valid. Silakan pilih antara 1-4.")
		}
	}
}

func PrintSuppliers(suppliers *Suppliers, JumlahSupplier int) {
	var i int
	var pilihanSort int

	if JumlahSupplier == 0 {
		fmt.Println("\n==========================================================================================")
		fmt.Println("          KOSONG: Belum ada data supplier yang terdaftar.     ")
		fmt.Println("============================================================================================ ")
	} else {
		fmt.Println("\n==========================================================================================")
		fmt.Println("                                   DAFTAR MITRA SUPPLIER                                  ")
		fmt.Println("============================================================================================")
		fmt.Println("\n==========================================================================================")
		fmt.Printf("%-5s | %-20s | %-15s | %-15s | %-15s | %-6s\n", "ID", "Nama PT", "Provinsi", "Kota", "Kontak", "Rating")
		fmt.Println("==========================================================================================")
		for i = 0; i < JumlahSupplier; i++ {
			fmt.Printf("%-5d | %-20s | %-15s | %-15s | %-15s | %-6.1f\n",
				(*suppliers)[i].ID,
				(*suppliers)[i].NamaPT,
				(*suppliers)[i].Provinsi,
				(*suppliers)[i].Kota,
				(*suppliers)[i].Kontak,
				(*suppliers)[i].Rating)
		}
		fmt.Println("==========================================================================================")
		fmt.Printf("Total Supplier: %d\n", JumlahSupplier)
		fmt.Println("==========================================================================================")

		fmt.Println("Menu Sorting:")
		fmt.Println("1. Urutkan berdasarkan Rating (Ascending/Membesar)")
		fmt.Println("2. Urutkan berdasarkan Rating (Descending/Mengecil)")
		fmt.Println("3. Tampilkan Statistik Wilayah dan Rata-Rata Rating")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Pilih Opsi (1-4) : ")

		fmt.Scan(&pilihanSort)

		switch pilihanSort {
		case 1:
			RatingAscendingSelection(suppliers, JumlahSupplier)
			fmt.Println("\n=====Daftar Supplier setelah diurutkan berdasarkan Rating Terendah - Tertinggi(Ascending): ====")
			CetakIsiTabel(suppliers, JumlahSupplier)
		case 2:
			RatingDescendingInsertion(suppliers, JumlahSupplier)
			fmt.Println("\n=====Daftar Supplier setelah diurutkan berdasarkan Rating Tertinggi - Terendah (Descending): ====")
			CetakIsiTabel(suppliers, JumlahSupplier)
		case 3:
			StatistikWilayah(suppliers, JumlahSupplier)
		case 4:
			return
		default:
			fmt.Println("Opsi tidak valid.")
		}
	}
}

func CetakIsiTabel(suppliers *Suppliers, JumlahSupplier int) {
	var i int
	fmt.Println("\n==========================================================================================")
	fmt.Printf("%-5s | %-20s | %-15s | %-15s | %-15s | %-6s\n", "ID", "Nama PT", "Provinsi", "Kota", "Kontak", "Rating")
	fmt.Println("==========================================================================================")
	for i = 0; i < JumlahSupplier; i++ {
		fmt.Printf("%-5d | %-20s | %-15s | %-15s | %-15s | %-6.1f\n",
			(*suppliers)[i].ID,
			(*suppliers)[i].NamaPT,
			(*suppliers)[i].Provinsi,
			(*suppliers)[i].Kota,
			(*suppliers)[i].Kontak,
			(*suppliers)[i].Rating)
	}
	fmt.Println("==========================================================================================")
	fmt.Printf("Total Supplier: %d\n", JumlahSupplier)
	fmt.Println("==========================================================================================")
}

// b. Sistem dapat mencatat riwayat pelayanan dan detail kontak dari setiap mitra supplier
func RiwayatPelayananSupplier(suppliers *Suppliers, JumlahSupplier int) {
	var pilihan, target, idxSupplier, i int
	var ketemu bool
	var keluar bool = false

	for !keluar {
		fmt.Println("\n=== Menu Riwayat Pelayanan Supplier ===")
		fmt.Println("1. Tambah Riwayat Pelayanan Baru")
		fmt.Println("2. Tampilkan Riwayat Pelayanan Supplier")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilih Opsi (1-3) : ")
		fmt.Scan(&pilihan)
		fmt.Println("------------------------------")

		switch pilihan {
		case 1:
			if JumlahSupplier == 0 {
				fmt.Println("Tidak ada supplier yang terdaftar. Silakan tambah supplier terlebih dahulu.")
			} else {
				fmt.Print("Masukkan ID Supplier: ")
				fmt.Scan(&target)

				// Mencari supplier berdasarkan ID
				ketemu = false
				i = 0
				for i < JumlahSupplier && !ketemu {
					if (*suppliers)[i].ID == target {
						ketemu = true
						idxSupplier = i
					} else {
						i++
					}
				}

				if ketemu {
					if (*suppliers)[idxSupplier].JumlahRiwayat >= NMAX {
						fmt.Println("Gagal: Riwayat pelayanan untuk supplier ini sudah penuh (Maksimal 50).")
					} else {
						fmt.Println("\n--- Input Data Riwayat Pelayanan ---")
						// Memasukkan data ke array sejajar di indeks JumlahRiwayat saat ini
						fmt.Print("Masukkan Tanggal Pelayanan (DD-MM-YYYY): ")
						fmt.Scan(&(*suppliers)[idxSupplier].RiwayatTanggal[(*suppliers)[idxSupplier].JumlahRiwayat])

						fmt.Print("Masukkan Nama Material yang Disuplai (Tanpa Spasi): ")
						fmt.Scan(&(*suppliers)[idxSupplier].RiwayatMaterial[(*suppliers)[idxSupplier].JumlahRiwayat])

						fmt.Print("Masukkan Jumlah Material (Angka): ")
						fmt.Scan(&(*suppliers)[idxSupplier].RiwayatJumlah[(*suppliers)[idxSupplier].JumlahRiwayat])
						// Menaikkan kapasitas hitung riwayat supplier yang bersangkutan
						(*suppliers)[idxSupplier].JumlahRiwayat++
						fmt.Println("Riwayat pelayanan berhasil dicatat.")
					}
				} else {
					fmt.Println("Supplier dengan ID", target, "tidak ditemukan.")
				}
			}

		case 2:
			if JumlahSupplier == 0 {
				fmt.Println("Tidak ada supplier yang terdaftar.")
			} else {
				fmt.Print("Masukkan ID Supplier untuk melihat riwayat: ")
				fmt.Scan(&target)

				// Mencari supplier
				ketemu = false
				i = 0
				for i < JumlahSupplier && !ketemu {
					if (*suppliers)[i].ID == target {
						ketemu = true
						idxSupplier = i
					} else {
						i++
					}
				}

				if ketemu {
					fmt.Println("\n==========================================================================================")
					fmt.Printf("        RIWAYAT PELAYANAN MITRA: %-31s ID: %-5d | Kontak: %-15s \n", (*suppliers)[idxSupplier].NamaPT, (*suppliers)[idxSupplier].ID, (*suppliers)[idxSupplier].Kontak)
					fmt.Println("==========================================================================================")

					if (*suppliers)[idxSupplier].JumlahRiwayat == 0 {
						fmt.Println("              Belum ada riwayat pelayanan yang tercatat untuk supplier ini.              ")
					} else {
						fmt.Printf("%-5s | %-15s | %-30s | %-15s\n", "No", "Tanggal", "Material", "Jumlah")
						fmt.Println("------------------------------------------------------------------------------------------")
						for i = 0; i < (*suppliers)[idxSupplier].JumlahRiwayat; i++ {
							fmt.Printf("%-5d | %-15s | %-30s | %-15d\n",
								i+1,
								(*suppliers)[idxSupplier].RiwayatTanggal[i],
								(*suppliers)[idxSupplier].RiwayatMaterial[i],
								(*suppliers)[idxSupplier].RiwayatJumlah[i])
						}
					}
					fmt.Println("==========================================================================================")
				} else {
					fmt.Println("Supplier dengan ID", target, "tidak ditemukan.")
				}
			}

		case 3:
			keluar = true
		default:
			fmt.Println("Opsi tidak valid. Silakan pilih antara 1-3.")
		}
	}
}

// c. Pengguna dapat mencari data supplier berdasarkan nama perusahaan atau lokasi menggunakan Sequential dan Binary Search
func CariDataSupplier(suppliers *Suppliers, JumlahSupplier int) {
	var i, j, pilihan int
	var target string
	var ketemu bool
	var keluar bool = false
	var tempSupplier Supplier
	var left, right, mid int

	for !keluar {
		fmt.Println("\n=== Mencari Data Supplier ===")
		fmt.Println("1. Pencarian Berurutan (Sequential) - Berdasarkan Rating Terbesar")
		fmt.Println("2. Pencarian Spesifik di Tengah (Binary) - Berdasarkan Nama Perusahaan atau Lokasi")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilih Opsi (1-3) : ")
		fmt.Scan(&pilihan)
		fmt.Println("------------------------------")

		switch pilihan {
		case 1:
			if JumlahSupplier == 0 {
				fmt.Println("Tidak ada supplier yang dapat dicari. Mohon kembali ke menu utama untuk menambahkan Data Supplier.")
			} else {
				for i = 0; i < JumlahSupplier-1; i++ {
					for j = 0; j < JumlahSupplier-i-1; j++ {
						if (*suppliers)[j].Rating < (*suppliers)[j+1].Rating {
							tempSupplier = (*suppliers)[j]
							(*suppliers)[j] = (*suppliers)[j+1]
							(*suppliers)[j+1] = tempSupplier
						}
					}
				}

				fmt.Print("Masukkan Nama Perusahaan atau Lokasi (Provinsi/Kota) yang dicari: ")
				fmt.Scan(&target)
				fmt.Println("\n==========================================================================================")
				fmt.Println("              HASIL PENCARIAN (BERDASARKAN RATING TERTINGGI (SEQUENTIAL))                 ")
				fmt.Println("==========================================================================================")
				fmt.Printf("%-5s | %-30s | %-15s | %-15s | %-15s | %-6s\n", "ID", "Nama PT", "Provinsi", "Kota", "Kontak", "Rating")
				fmt.Println("==========================================================================================")

				ketemu = false
				// Mencari secara Sequential
				for i = 0; i < JumlahSupplier; i++ {
					if (*suppliers)[i].NamaPT == target || (*suppliers)[i].Provinsi == target || (*suppliers)[i].Kota == target {
						fmt.Printf("%-5d | %-30s | %-15s | %-15s | %-15s | %-6.1f\n",
							(*suppliers)[i].ID, (*suppliers)[i].NamaPT, (*suppliers)[i].Provinsi,
							(*suppliers)[i].Kota, (*suppliers)[i].Kontak, (*suppliers)[i].Rating)
						ketemu = true
					}
				}
				fmt.Println("==========================================================================================")

				if !ketemu {
					fmt.Printf("Supplier dengan kata kunci '%s' tidak ditemukan.\n", target)
				}
			}

		case 2:
			if JumlahSupplier == 0 {
				fmt.Println("Tidak ada supplier yang dapat dicari. Mohon kembali ke menu utama untuk menambahkan Data Supplier.")
			} else {
				for i = 0; i < JumlahSupplier-1; i++ {
					for j = 0; j < JumlahSupplier-i-1; j++ {
						if (*suppliers)[j].NamaPT > (*suppliers)[j+1].NamaPT {
							tempSupplier = (*suppliers)[j]
							(*suppliers)[j] = (*suppliers)[j+1]
							(*suppliers)[j+1] = tempSupplier
						}
					}
				}
				fmt.Print("Masukkan Nama Perusahaan atau Lokasi (Provinsi/Kota) yang spesifik dicari: ")
				fmt.Scan(&target)

				// Mencari Secara Binary
				left = 0
				right = JumlahSupplier - 1
				ketemu = false
				for left <= right && !ketemu {
					mid = (left + right) / 2

					if (*suppliers)[mid].NamaPT == target || (*suppliers)[mid].Provinsi == target || (*suppliers)[mid].Kota == target {
						ketemu = true
					} else if (*suppliers)[mid].NamaPT < target || (*suppliers)[mid].Provinsi < target || (*suppliers)[mid].Kota < target {
						left = mid + 1
					} else {
						right = mid - 1
					}
				}

				fmt.Println("\n==========================================================================================")
				fmt.Println("                                HASIL PENCARIAN (BINARY SEARCH)                           ")
				fmt.Println("==========================================================================================")
				fmt.Printf("%-5s | %-30s | %-15s | %-15s | %-15s | %-6s\n", "ID", "Nama PT", "Provinsi", "Kota", "Kontak", "Rating")
				fmt.Println("==========================================================================================")

				if ketemu {
					fmt.Printf("%-5d | %-30s | %-15s | %-15s | %-15s | %-6.1f\n",
						(*suppliers)[mid].ID, (*suppliers)[mid].NamaPT, (*suppliers)[mid].Provinsi,
						(*suppliers)[mid].Kota, (*suppliers)[mid].Kontak, (*suppliers)[mid].Rating)
				} else {
					fmt.Printf("Supplier dengan kata kunci '%s' tidak ditemukan.\n", target)
				}
				fmt.Println("==========================================================================================")
			}

		case 3:
			keluar = true

		default:
			fmt.Println("Opsi tidak valid. Silakan pilih antara 1-3.")
		}
	}
}

// d. Pengguna dapat mengurutkan data mitra berdasarkan rating layanan tertinggi atau terendah menggunakan Selection dan Insertion Sort.
func RatingAscendingSelection(suppliers *Suppliers, JumlahSupplier int) {
	var i, j int
	var idxMin int
	var tempData Supplier

	for i = 0; i < JumlahSupplier-1; i++ {
		idxMin = i
		for j = i + 1; j < JumlahSupplier; j++ {
			if (*suppliers)[j].Rating < (*suppliers)[idxMin].Rating {
				idxMin = j
			}
		}
		if idxMin != i {
			tempData = (*suppliers)[i]
			(*suppliers)[i] = (*suppliers)[idxMin]
			(*suppliers)[idxMin] = tempData
		}
	}
}

func RatingDescendingInsertion(suppliers *Suppliers, JumlahSupplier int) {
	var i int
	var temp Supplier
	var pass int

	pass = 1
	for pass < JumlahSupplier {
		i = pass
		temp = (*suppliers)[pass]

		for i > 0 && (*suppliers)[i-1].Rating < temp.Rating {
			(*suppliers)[i] = (*suppliers)[i-1]
			i--
		}
		(*suppliers)[i] = temp
		pass++
	}
}

// e. Sistem dapat menampilkan statistik jumlah supplier untuk setiap wilayah dan rata-rata skor kepuasan mitra
func StatistikWilayah(suppliers *Suppliers, JumlahSupplier int) {
	var i, j int
	var proses bool
	var totalSupplierWilayah int
	var totalRatingWilayah float64
	var rataRataRating float64

	if JumlahSupplier == 0 {
		fmt.Println("\n==========================================================================================")
		fmt.Println("   KOSONG: Belum ada data supplier yang terdaftar.     ")
		fmt.Println("============================================================================================ ")
	}

	fmt.Println("\n==========================================================================================")
	fmt.Println("                             STATISTIK SUPPLIER BERDASARKAN WILAYAH                            ")
	fmt.Println("==========================================================================================")

	for i = 0; i < JumlahSupplier; i++ {
		proses = false
		j = 0

		for j < i && !proses {
			if (*suppliers)[i].Provinsi == (*suppliers)[j].Provinsi && (*suppliers)[i].Kota == (*suppliers)[j].Kota {
				proses = true
			}
			j++
		}

		if !proses {
			totalSupplierWilayah = 0
			totalRatingWilayah = 0.0

			for j = 0; j < JumlahSupplier; j++ {
				if (*suppliers)[j].Provinsi == (*suppliers)[i].Provinsi && (*suppliers)[j].Kota == (*suppliers)[i].Kota {
					totalSupplierWilayah++
					totalRatingWilayah += (*suppliers)[j].Rating
				}
			}

			rataRataRating = totalRatingWilayah / float64(totalSupplierWilayah)

			fmt.Printf("%-15s | %-15s | %-20d | %-20.1f\n",
				(*suppliers)[i].Provinsi,
				(*suppliers)[i].Kota,
				totalSupplierWilayah,
				rataRataRating)
		}
	}
	fmt.Println("==========================================================================================")
}
