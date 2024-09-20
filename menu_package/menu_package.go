package menu_package

import (
	"aplikasi_bank/user_package"
	"fmt"
	"strconv"
)

// Fungsi untuk menampilkan informasi rekening
func CekRekening(user *user_package.User) {
	fmt.Println("Nama Pengguna: ", user.Name)
	fmt.Println("Nomor Rekening:", user.NoRek)
	fmt.Println("Sisa Saldo: Rp.", user.Saldo)
}

// Fungsi untuk penarikan saldo
func TarikTunai(user *user_package.User) {
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
func DepositSaldo(user *user_package.User) {
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
func TransferSaldo(user *user_package.User, users map[int]*user_package.User) {
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
func InisialisasiPengguna() map[int]*user_package.User {
	users := make(map[int]*user_package.User)

	// menambahkan pengguna
	user1, _ := user_package.NewUser("Alwy Alhaddad", 1231231230, 500000)
	user2, _ := user_package.NewUser("Anas Sufi", 3213213210, 300000)

	users[user1.NoRek] = user1
	users[user2.NoRek] = user2

	return users
}
