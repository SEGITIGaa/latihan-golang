package main

import "fmt"

func sambutan() {
	fmt.Println("ini adalah latihan function golang!")
}

func celToKel(cel int) float64 {
	return float64(cel) + 273.15
}

func celToFahr(cel int) int {
	return (cel * 9 / 5) + 32
}

func main() {
	sambutan()
	fmt.Println("================================")
	// konversi suhu go
	var celcius int
	fmt.Println("masukan suhu celsius: ")
	fmt.Scanf("%d", &celcius)

	kelvin := celToKel(celcius)
	fahrenheit := celToFahr(celcius)

	fmt.Println("================================")

	// %d itu buat angka 
	// %s itu buat string 
	// %.2f itu buat angka desimal 
	fmt.Printf("%d celcius sama dengan %.2f k ", celcius, kelvin)
	fmt.Printf("dan juga sama dengan %d f ", fahrenheit)
}
