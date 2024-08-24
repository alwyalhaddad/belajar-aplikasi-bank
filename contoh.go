package main

import (
	"errors"
	"fmt"
	"os"
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
	//sukses
	return user, nil
}

// func untuk tarik saldo
func (u *User) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Penarikan Tidak Boleh NOL")
	}
	if amount > u.Saldo {
		return errors.New("Saldo Anda Kurang")
	}
	// sukses withdrawl
	u.Saldo -= amount
	return nil
}

func (u *User) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Pengisian Minimal Adalah Rp.10.000")
	}
	// sukses deposit
	u.Saldo += amount
	return nil
}

func main() {
	user, err := NewUser("Alwy", 1231231230, 750000)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var choice int

	for {
		// menampilan Menu
		fmt.Println("Silahkan Pilih 1 Angka:")
		fmt.Println("1.Cek Rekening")
		fmt.Println("2.Tarik Tunai")
		fmt.Println("3.Deposit")
		fmt.Println("4.Keluar Aplikasi")
		fmt.Print("")

		// menu scanner
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("input salah", err)
			break
		}

		switch choice {
		// informasi data pengguna
		case 1:
			fmt.Println("Nama Pengguna: ", user.Name)
			fmt.Println("Nomor Rekening:", user.NoRek)
			fmt.Println("Sisa Saldo: ", user.Saldo)
		// penarikan tunai
		case 2:
			var withdraw_amountStr string
			fmt.Print("Masukan Jumlah Penarikan: Rp.")
			fmt.Scanln(&withdraw_amountStr)
			fmt.Println("Transaksi Berhasil.")
		// deposit
		case 3:
			var deposit_amountStr string
			fmt.Print("Mau Topup Berapa? Rp.")
			fmt.Scanln(&deposit_amountStr)
			fmt.Println("Transaksi Berhasil.")
		// keluar dari aplikasi
		case 4:
			fmt.Println("Keluar Dari Aplikasi...")
			os.Exit(0)
		default:
			fmt.Println("input salah")
		}
	}
}
