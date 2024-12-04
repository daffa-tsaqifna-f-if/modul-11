package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	switch {
	case x >= 0 && x <= 14:
		switch {
		case x >= 6.5 && x <= 8.6:
			fmt.Print("Air layak minum")
		case x < 6.5 || x > 8.6:
			fmt.Print("Air tidak layak minum ")
		}
	default:
		fmt.Print("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	}
}
