package user_package

import (
	"errors"
)

type User struct {
	Name  string
	NoRek int
	Saldo float64
}

// membuat data untuk pengguna baru
func NewUser(name string, norek int, saldo float64) (*User, error) {
	if name == "" {
		return nil, errors.New("nama tidak boleh kosong")
	}
	if norek <= 0 {
		return nil, errors.New("nomor rekening harus berisi 10 digit")
	}
	if saldo < 0 {
		return nil, errors.New("saldo kurang")
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
		return errors.New("penarikan tidak boleh nol")
	}
	if amount > u.Saldo {
		return errors.New("saldo anda kurang")
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
		return errors.New("pengguna tidak di temukan")
	}
	if amount <= 0 {
		return errors.New("jumlah transfer tidak boleh nol")
	}
	if amount > u.Saldo {
		return errors.New("saldo anda kurang")
	}
	// sukses
	u.Saldo -= amount
	target.Saldo += amount
	return nil
}
