package main

import "fmt"

func main() {
	var jumlah, bilangan int

	jumlah = 0

	fmt.Scan(&bilangan)
	for bilangan = 1; bilangan <= 100; bilangan++ {

		jumlah = jumlah + bilangan
	}

	fmt.Println(jumlah)
}
