package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Name     string
	Rekening int
}

func (u User) InfoUser(name string, rek int) {
	fmt.Println("Nama : ", name, "Rekening", rek)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	// Sukses

	switch char {
	case 'A':
		fmt.Println("A key Pressed")
		break
	case 'a':
		fmt.Println("a Key Pressed")
		break
	}

	// var anas User
	// anas.Name = "Anas"
	// anas.Rekening = 123123123

	// fmt.Println("Nama : ", anas.Name, "\nNo Rekening : ", anas.Rekening)

}
