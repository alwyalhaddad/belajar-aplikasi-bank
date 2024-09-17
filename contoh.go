package main

import (
	"errors"
	"fmt"
	"strconv"
)

type User struct {
	Name  string
	NoRek int
	Saldo float64
}

// membuat data untuk pengguna baru
func NewUser(name string, norek int, saldo float64) (*User, error) {
	if name == "" {
		return nil, errors.New("Nama Tidak Boleh Kosong")
	}
	if norek <= 0 {
		return nil, errors.New("Nomor Reking Harus Berisikan 10 Digit")
	}
	if saldo < 0 {
		return nil, errors.New("Saldo Kurang")
	}

	user := &User{
		Name:  name,
		NoRek: norek,
		Saldo: saldo,
	}
	// sukses
	return user, nil
}

// func untuk withdraw (tarik saldo)
func (u *User) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Penarikan Tidak Boleh NOL")
	}
	if amount > u.Saldo {
		return errors.New("Saldo Anda Kurang")
	}
	// sukses withdraw
	u.Saldo -= amount
	return nil
}

// func untuk deposit (masukan saldo)
func (u *User) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit Tidak Boleh Nol")
	}
	// sukses deposit
	u.Saldo += amount
	return nil
}

// func untuk transfer
func (u *User) Transfer(users map[int]*User, norekTarget int, amount float64) error {
	target, exists := users[norekTarget]
	if !exists {
		return errors.New("Pengguna Tidak di Temukan")
	}
	if amount <= 0 {
		return errors.New("Jumlah Transfer Tidak Boleh NOL")
	}
	if amount > u.Saldo {
		return errors.New("Saldo Anda Kurang")
	}
	// sukses
	u.Saldo -= amount
	target.Saldo += amount
	return nil
}

// Fungsi untuk menampilkan informasi rekening
func CekRekening(user *User) {
	fmt.Println("Nama Pengguna: ", user.Name)
	fmt.Println("Nomor Rekening:", user.NoRek)
	fmt.Println("Sisa Saldo: Rp.", user.Saldo)
}

// Fungsi untuk penarikan saldo
func TarikTunai(user *User) {
	fmt.Print("Masukan Jumlah Penarikan: Rp.")
	var withdraw_amountStr string
	fmt.Scanln(&withdraw_amountStr)
	withdrawamount, err := strconv.ParseFloat(withdraw_amountStr, 64)

	if err != nil {
		fmt.Println("input tidak valid", err)
		return
	}

	err = user.Withdraw(withdrawamount)
	if err != nil {
		fmt.Println("Gagal Melakukan Penarikan", err)
	} else {
		fmt.Println("Transaksi Berhasil, sisa saldo anda Rp.", user.Saldo)
	}
}

// Fungsi untuk deposit saldo
func DepositSaldo(user *User) {
	fmt.Print("Mau Topup Berapa ? Rp.")
	var deposit_amountStr string
	fmt.Scanln(&deposit_amountStr)
	depositamount, err := strconv.ParseFloat(deposit_amountStr, 64)
	if err != nil {
		fmt.Println("input tidak valid")
		return
	}

	err = user.Deposit(depositamount)
	if err != nil {
		fmt.Println("Transaksi Gagal", err)
	} else {
		fmt.Println("Transaksi Berhasil, Saldo Sekarang Rp.", user.Saldo)
	}
}

// Fungsi untuk transfer saldo
func TransferSaldo(user *User, users map[int]*User) {
	fmt.Println("Masukan Nomor Rekening Tujuan:")
	var norekTarget int
	fmt.Scanln(&norekTarget)

	fmt.Println("Masukan Jumlah Transfer: Rp.")
	var transfer_amountStr string
	fmt.Scanln(&transfer_amountStr)
	transferamount, err := strconv.ParseFloat(transfer_amountStr, 64)
	if err != nil {
		fmt.Println("input tidak valid", err)
		return
	}

	err = user.Transfer(users, norekTarget, transferamount)
	if err != nil {
		fmt.Println("Gagal Melakukan Transfer", err)
	} else {
		fmt.Println("Transfer Berhasil, Sisa Saldo Anda: Rp.", user.Saldo)
		fmt.Println("Saldo Penerima: Rp.", users[norekTarget].Saldo)
	}
}

// Fungsi untuk inisialisasi pengguna
func InisialisasiPengguna() map[int]*User {
	users := make(map[int]*User)

	// menambahkan pengguna
	user1, _ := NewUser("Alwy Alhaddad", 1231231230, 500000)
	user2, _ := NewUser("Anas Sufi", 3213213210, 300000)

	users[user1.NoRek] = user1
	users[user2.NoRek] = user2

	return users
}

// Fungsi main untuk menampilkan menu dan memanggil fungsi lain
func main() {
	// Inisialisasi pengguna
	users := InisialisasiPengguna()
	user1 := users[1231231230] // Menggunakan user pertama (contoh)

	var choice int

	for {
		// menampilkan Menu
		fmt.Println("Silahkan Pilih 1 Angka:")
		fmt.Println("1. Cek Rekening")
		fmt.Println("2. Tarik Tunai")
		fmt.Println("3. Deposit")
		fmt.Println("4. Transfer")
		fmt.Println("5. Keluar Aplikasi")
		fmt.Print("")

		// menu scanner
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("input salah", err)
			break
		}

		// Pemanggilan fungsi berdasarkan pilihan user
		switch choice {
		case 1:
			CekRekening(user1)
		case 2:
			TarikTunai(user1)
		case 3:
			DepositSaldo(user1)
		case 4:
			TransferSaldo(user1, users)
		case 5:
			fmt.Println("Keluar Dari Aplikasi...")
			return
		default:
			fmt.Println("input salah")
		}
	}
}
