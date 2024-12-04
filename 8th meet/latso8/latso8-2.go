package main

import "fmt"

func main() {
	var x int
	var y string
	fmt.Scan(&y)
	switch {
	case y == "motor" || y == "Motor" || y == "MOTOR":
		fmt.Scan(&x)
		x = x * 2000
		fmt.Print("Rp. ", x)
	case y == "mobil" || y == "Mobil" || y == "MOBIL":
		fmt.Scan(&x)
		x = x * 5000
		fmt.Print("Rp. ", x)
	case y == "truk" || y == "Truk" || y == "TRUK":
		fmt.Scan(&x)
		x = x * 8000
		fmt.Print("Rp. ", x)
	default:
		fmt.Print("invalid input")
	}
}
