package main

import (
	"aplikasi_bank/menu_package"
	"fmt"
)

func main() {
	users := menu_package.InisialisasiPengguna()
	user1 := users[1231231230] // use user1 for example

	var choice int

	for {
		// [[DISPLAY MENU]]
		fmt.Println("Silahkan Pilih 1 Angka:")
		fmt.Println("1. Cek Rekening")
		fmt.Println("2. Tarik Tunai")
		fmt.Println("3. Deposit")
		fmt.Println("4. Transfer")
		fmt.Println("5. Keluar Aplikasi")
		fmt.Print("")

		// scanner menu
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("input salah", err)
			break
		}

		switch choice {
		case 1:
			menu_package.CekRekening(user1)
		case 2:
			menu_package.TarikTunai(user1)
		case 3:
			menu_package.DepositSaldo(user1)
		case 4:
			menu_package.TransferSaldo(user1, users)
		case 5:
			fmt.Println("Keluar Dari Aplikasi...")
			return
		default:
			fmt.Println("input salah")
		}
	}
}
