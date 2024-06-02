package main

import "fmt"

const NMAX = 1000

type Barang struct {
	idBarang, jumlahBarang, hargaBarang int
	kategori, namaBarang, tanggalMasuk string
}

type DataAdmin struct {
	namaLengkap, email, password string
	statusLogin bool
}

type DataBarang struct {
	list       [NMAX]Barang
	count      int
	transaksi  DataTransaksi
}

type Transaksi struct {
	idTransaksi, idBarang, jumlahBarang int
	tanggalKeluar                    string
}

type DataTransaksi struct {
	list  [NMAX]Transaksi
	count int
}

func main() {
	var barang2 DataBarang
	mainMenu(&barang2)
}

func mainMenu(barang2 *DataBarang) {
	var pilih int
	var dataEmail, dataPassword string
	var A DataAdmin
	stop := false
	for !stop {
		if !A.statusLogin {
			fmt.Println("=================================")
			fmt.Println("|   Aplikasi Inventori Barang   |")
			fmt.Println("=================================")
			fmt.Println("|             MENU              |")
			fmt.Println("=================================")
			fmt.Println("| 1. Sign Up                    |")
			fmt.Println("| 2. Login                      |")
			fmt.Println("| 3. Exit                       |")
			fmt.Println("=================================")
			fmt.Print("Pilih Menu (1/2/3) : ")
			fmt.Scanln(&pilih)
			switch pilih {
			case 1:
				signUp(&A)
			case 2:
				login(&A, &dataEmail, &dataPassword)
			case 3:
				fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
				stop = true
			default:
				fmt.Println("Pilihan tidak valid!")
			}
		} else {
			menu(barang2, &A)
		}
	}
}

func signUp(A *DataAdmin) {
	fmt.Print("Masukkan nama lengkap: ")
	fmt.Scanln(&A.namaLengkap)
	fmt.Print("Masukkan Email: ")
	fmt.Scanln(&A.email)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&A.password)
	fmt.Println("Akun Berhasil Dibuat")
}

func login(A *DataAdmin, dataEmail, dataPassword *string) {
	fmt.Print("Masukkan Email: ")
	fmt.Scanln(dataEmail)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(dataPassword)
	if cekDaftar(A, *dataEmail, *dataPassword) {
		fmt.Print("Login Berhasil, Welcome ")
		fmt.Printf("%s", A.namaLengkap)
		fmt.Println()
		A.statusLogin = true
	} else {
		fmt.Println("Login Gagal")
	}
}

func cekDaftar(A *DataAdmin, dataEmail, dataPassword string) bool {
	if A.email == dataEmail {
		if A.password == dataPassword {
			return true
		} else {
			fmt.Println("Password Salah")
			return false
		}
	}
	fmt.Println("Email Tidak Terdaftar")
	return false
}

func menu(barang2 *DataBarang, A *DataAdmin) {
	fmt.Println("=============================================")
	fmt.Println("|          Aplikasi Inventori Barang        |")
	fmt.Println("=============================================")
	fmt.Println("|     By : Fazari Razka Davira              |")
	fmt.Println("|          Muhamad Rizky Kurniawan          |")
	fmt.Println("=============================================")
	fmt.Println("|                    MENU                   |")
	fmt.Println("=============================================")
	fmt.Println("| 1. Tambah Barang                          |")
	fmt.Println("| 2. Ubah Barang                            |")
	fmt.Println("| 3. Hapus Barang                           |")
	fmt.Println("| 4. Tambah Transaksi                       |")
	fmt.Println("| 5. Lihat Transaksi                        |")
	fmt.Println("| 6. Cari Barang Dari Kata Kunci            |")
	fmt.Println("| 7. Cari Barang Dari Kategori              |")
	fmt.Println("| 8. Urut Barang Berdasarkan Banyak Barang  |")
	fmt.Println("| 9. Urut Barang Berdasarkan Kategori       |")
	fmt.Println("| 10. Lihat Barang                          |")
	fmt.Println("| 11. Logout                                |")
	fmt.Println("=============================================")

	var pilih int
	var enter string
	fmt.Print("Pilih Menu (1/2/3/.../11) : ")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		tambahBarang(barang2)
		fmt.Println("!! Enter Untuk Kembali Ke Menu !!")
		fmt.Scanln(&enter)
	case 2:
		ubahBarang(barang2)
		fmt.Println("!! Enter Untuk Kembali Ke Menu !!")
		fmt.Scanln(&enter)
	case 3:
		hapusBarang(barang2)
		fmt.Println("!! Enter Untuk Kembali Ke Menu !!")
		fmt.Scanln(&enter)
	case 4:
		tambahTransaksi(barang2)
		fmt.Println("!! Enter Untuk Kembali Ke Menu !!")
		fmt.Scanln(&enter)
	case 5:
		lihatTransaksi(barang2)
		fmt.Println("!! Enter Untuk Kembali Ke Menu !!")
		fmt.Scanln(&enter)
	case 6:
		cariBarangByKataKunci(barang2)
		fmt.Println("!! Enter Untuk Kembali Ke Menu !!")
		fmt.Scanln(&enter)
	case 7:
		cariBarangByKategori(barang2)
		fmt.Println("!! Enter Untuk Kembali Ke Menu !!")
		fmt.Scanln(&enter)
	case 8:
		urutBarangByJumlah(barang2)
		fmt.Println("!! Enter Untuk Kembali Ke Menu !!")
		fmt.Scanln(&enter)
	case 9:
		urutBarangByKategori(barang2)
		fmt.Println("!! Enter Untuk Kembali Ke Menu !!")
		fmt.Scanln(&enter)
	case 10:
		lihatBarang(barang2)
		fmt.Println("!! Enter Untuk Kembali Ke Menu !!")
		fmt.Scanln(&enter)
	case 11:
		A.statusLogin = false
		fmt.Print("Logout Berhasil, Terimakasih ")
		fmt.Printf("%s", A.namaLengkap)
		fmt.Println()
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func tambahBarang(A *DataBarang) {
    if A.count >= NMAX {
        fmt.Println("Inventori Penuh")
    } else {
        var newBarang Barang
        newBarang.idBarang = A.count + 1
        fmt.Print("Masukkan Nama Barang: ")
        fmt.Scanln(&newBarang.namaBarang)
        fmt.Print("Masukkan Kategori Barang: ")
        fmt.Scanln(&newBarang.kategori)
        fmt.Print("Masukkan Jumlah Barang: ")
        fmt.Scanln(&newBarang.jumlahBarang)
        fmt.Print("Masukkan Harga Barang: ")
        fmt.Scanln(&newBarang.hargaBarang)
        fmt.Print("Masukkan Tanggal Masuk (YYYY-MM-DD): ")
        fmt.Scanln(&newBarang.tanggalMasuk)

        A.list[A.count] = newBarang
        A.count++

		fmt.Println()
        fmt.Println("Barang berhasil ditambahkan!")
		fmt.Println()
    }
}


func ubahBarang(A *DataBarang) {
    var id int
    var enter string

    berhenti := false

    if !berhenti {
        lihatBarang(A)
        if A.count == 0 {
            berhenti = true
        }
    }

    if !berhenti {
        fmt.Print("Masukkan ID Barang yang akan diubah: ")
        fmt.Scan(&id)
        index := cariBarangBinarySearch(A, id)
        if index == -1 {
			fmt.Println()
            fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
			fmt.Println()
			fmt.Scanln()
            berhenti = true
        }

        if !berhenti {
            fmt.Print("Masukkan Nama Barang baru: ")
            fmt.Scanln(&enter) // Baca enter sebagai placeholder
            fmt.Scanln(&A.list[index].namaBarang)

            fmt.Print("Masukkan Kategori Barang baru: ")
            fmt.Scanln(&A.list[index].kategori)

            fmt.Print("Masukkan Jumlah Barang baru: ")
            fmt.Scanln(&A.list[index].jumlahBarang)

            fmt.Print("Masukkan Harga Barang baru: ")
            fmt.Scanln(&A.list[index].hargaBarang)

            fmt.Print("Masukkan Tanggal Masuk baru (YYYY-MM-DD): ")
            fmt.Scanln(&A.list[index].tanggalMasuk)

			fmt.Println()
            fmt.Println("Barang berhasil diubah!")
			fmt.Println()
        }
    }
}

func hapusBarang(A *DataBarang) {
	berhenti := false
	if !berhenti {
        lihatBarang(A)
        if A.count == 0 {
            berhenti = true
        }
    }
	if !berhenti {
		var id int
		fmt.Print("Masukkan ID Barang yang akan dihapus: ")
		fmt.Scan(&id)
		index := cariBarangBinarySearch(A, id)
		if index == -1 {
			fmt.Println()
            fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
			fmt.Println()
			fmt.Scanln()
			berhenti = true
		} else {
			for i := index; i < A.count-1; i++ {
				A.list[i] = A.list[i+1]
			}
			A.count--
			fmt.Println()
			fmt.Println("Barang berhasil dihapus!")
			fmt.Println()
		}
	}
}


func tambahTransaksi(A *DataBarang) {
    berhenti := false
    if A.transaksi.count >= NMAX {
        fmt.Println("Data transaksi penuh.")
        berhenti = true
    }
    if !berhenti {
        lihatBarang(A)
        if A.count == 0 {
            berhenti = true
        }
    }
    if !berhenti {
        var newTransaksi Transaksi
        newTransaksi.idTransaksi = A.transaksi.count + 1
        fmt.Print("Masukkan ID Barang Yang Ingin Ditransaksi: ")
        fmt.Scan(&newTransaksi.idBarang)
        index := cariBarangBinarySearch(A, newTransaksi.idBarang)
        if index == -1 {
            fmt.Println()
            fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
			fmt.Println()
			fmt.Scanln()
			berhenti = true
        }
		if !berhenti {
			fmt.Print("Masukkan Jumlah Barang Yang Ingin Ditransaksi: ")
			fmt.Scan(&newTransaksi.jumlahBarang)
			fmt.Print("Masukkan Tanggal Keluar Transaksi (YYYY-MM-DD): ")
			fmt.Scan(&newTransaksi.tanggalKeluar)
			if !berhenti {
				// Periksa apakah jumlah barang yang akan ditransaksikan tersedia
				if A.list[index].jumlahBarang < newTransaksi.jumlahBarang {
					fmt.Println()
					fmt.Println("Jumlah barang tidak mencukupi untuk transaksi.")
					fmt.Println()
					berhenti = true
				}
			}
		}

        // Periksa apakah ID barang ada di inventori
		

        if !berhenti {
            // Kurangi jumlah barang di inventori
            A.list[index].jumlahBarang -= newTransaksi.jumlahBarang

            // Tambahkan transaksi ke data transaksi
            A.transaksi.list[A.transaksi.count] = newTransaksi
            A.transaksi.count++

			fmt.Println()
            fmt.Println("Transaksi berhasil ditambahkan!")
			fmt.Println()
			fmt.Scanln()
        }
    }
}


func lihatTransaksi(A *DataBarang) {
    if A.transaksi.count == 0 {
		fmt.Println()
        fmt.Println("Tidak ada transaksi.")
		fmt.Println()
    } else {
        for i := 0; i < A.transaksi.count; i++ {
            fmt.Println("======================================")
            fmt.Printf("ID Transaksi             : %d\n", A.transaksi.list[i].idTransaksi)
            fmt.Printf("ID Barang                : %d\n", A.transaksi.list[i].idBarang)
            fmt.Printf("Jumlah                   : %d\n", A.transaksi.list[i].jumlahBarang)
            fmt.Printf("Tanggal Keluar Transaksi : %s\n", A.transaksi.list[i].tanggalKeluar)
        }
        fmt.Println("======================================")
        fmt.Println()
    }
}

func cariBarangBinarySearch(A *DataBarang, id int) int {
	var found, tengah, kiri, kanan int 
	found = -1
	kiri = 0
	kanan = A.count - 1
	for kiri <= kanan && found == -1 {
		tengah = (kiri + kanan) / 2
		if id < A.list[tengah].idBarang {
			kanan = tengah - 1
		}else if id > A.list[tengah].idBarang {
			kiri = tengah + 1
		}else {
			found = tengah
		}
	}
	return found
}

func lihatBarang(A *DataBarang) {
	berhenti := false
	if !berhenti {
		if A.count == 0 {
			fmt.Println()
			fmt.Println("Tidak ada barang di inventori.")
			fmt.Println()
			berhenti = true 
		}
	}

	if !berhenti {
		for i := 0; i < A.count; i++ {
			fmt.Println("======================================")
			fmt.Printf("ID            : %d\n", A.list[i].idBarang)
			fmt.Printf("Nama          : %s\n", A.list[i].namaBarang)
			fmt.Printf("Kategori      : %s\n", A.list[i].kategori)
			fmt.Printf("Jumlah        : %d\n", A.list[i].jumlahBarang)
			fmt.Printf("Harga         : %d\n", A.list[i].hargaBarang)
			fmt.Printf("Tanggal Masuk : %s\n", A.list[i].tanggalMasuk)
		}
		fmt.Println("======================================")
		fmt.Println()
	}
}



func cariBarangByKataKunci(A *DataBarang) {
	var kataKunci string
	found := false
	if !found {
        if A.count == 0 {
			found = true
			lihatBarang(A)
        }
    }
	if !found {
		fmt.Print("Masukkan Nama Barang Yang Ingin Dicari: ")
		fmt.Scanln(&kataKunci)
	}
	for i := 0; i < A.count; i++ {
		if A.list[i].namaBarang == kataKunci {
			fmt.Println("======================================")
			fmt.Printf("ID            : %d\n", A.list[i].idBarang)
			fmt.Printf("Nama          : %s\n", A.list[i].namaBarang)
			fmt.Printf("Kategori      : %s\n", A.list[i].kategori)
			fmt.Printf("Jumlah        : %d\n", A.list[i].jumlahBarang)
			fmt.Printf("Harga         : %d\n", A.list[i].hargaBarang)
			fmt.Printf("Tanggal Masuk : %s\n", A.list[i].tanggalMasuk)
			found = true
			fmt.Println("======================================")
			fmt.Println()
		}
	}
	if !found {
		fmt.Println()
		fmt.Println("Barang dengan kata kunci tersebut tidak ditemukan.")
		fmt.Println()
	}
}

func cariBarangByKategori(A *DataBarang) {
	var kategori string
	found := false
	if !found {
        if A.count == 0 {
            found = true
			lihatBarang(A)
        }
    }
	if !found {
		fmt.Print("Masukkan Kategori : ")
		fmt.Scanln(&kategori)
	}
	for i := 0; i < A.count; i++ {
		if A.list[i].kategori == kategori {
			fmt.Println("======================================")
			fmt.Printf("ID            : %d\n", A.list[i].idBarang)
			fmt.Printf("Nama          : %s\n", A.list[i].namaBarang)
			fmt.Printf("Kategori      : %s\n", A.list[i].kategori)
			fmt.Printf("Jumlah        : %d\n", A.list[i].jumlahBarang)
			fmt.Printf("Harga         : %d\n", A.list[i].hargaBarang)
			fmt.Printf("Tanggal Masuk : %s\n", A.list[i].tanggalMasuk)
			found = true
			fmt.Println("======================================")
			fmt.Println()
		}
	}
	if !found {
		fmt.Println()
		fmt.Println("Barang dengan kategori tersebut tidak ditemukan.")
		fmt.Println()
	}
}

func urutBarangByJumlah(A *DataBarang) {
	berhenti := false
	if !berhenti {
		if A.count == 0 {
			berhenti = true
			lihatBarang(A)
        }
	}
	if !berhenti {
		// Melakukan insertion sort berdasarkan jumlah barang
		for i := 1; i < A.count; i++ {
			key := A.list[i]
			j := i - 1
			// Pindahkan elemen yang lebih besar dari key ke satu posisi di depan posisi saat ini
			for j >= 0 && A.list[j].jumlahBarang > key.jumlahBarang {
				A.list[j+1] = A.list[j]
				j = j - 1
			}
			A.list[j+1] = key
		}
		lihatBarang(A)
		fmt.Println("Barang berhasil diurutkan berdasarkan jumlah.")
		fmt.Println()
	}
}


func urutBarangByKategori(A *DataBarang) {
	berhenti := false
	if !berhenti {
		if A.count == 0 {
			berhenti = true
			lihatBarang(A)
		}
	}
	if !berhenti {
		for i := 0; i < A.count-1; i++ {
			max := i
			for j := i + 1; j < A.count; j++ {
				if A.list[j].kategori < A.list[max].kategori {
					max = j
				}
			}
			temp := A.list[max]
			A.list[max] = A.list[i]
			A.list[i] = temp
		}
		lihatBarang(A)
		fmt.Println("Barang berhasil diurutkan berdasarkan kategori.")
		fmt.Println()
		berhenti = true
	}
}
