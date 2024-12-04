package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	switch {
	case x%10 == 0:
		fmt.Print("Kategori: Bilangan Kelipatan 10\nHasil pembagian antara", x, " / 10 = ", x/10)
	default:
		switch {
		case x%5 == 0 && x != 5:
			fmt.Print("Kategori: Bilangan Kelipatan 5 \nHasil kuadrat dari ", x, "^2 = ", x*x)
		default:
			switch {
			case x%2 == 0:
				fmt.Print("Kategori: Bilangan Genap \nHasil perkalian dengan bilangan berikutnya ", x, " * ", x+1, " = ", x*(x+1))
			case x%2 != 0:
				fmt.Print("Kategori: Bilangan Ganjil \nHasil penjumlahan dengan bilangan berikutnya ", x, " + ", x+1, " = ", x+(x+1))
			default:
				fmt.Print("invalid")
			}
		}
	}
}
