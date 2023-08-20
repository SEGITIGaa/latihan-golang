package main

import "fmt"
// import "rsc.io/quote"
func main() {
	var usia int
	var asal_sekolah string
	fmt.Println("program input usia dan nama dari go")
	// INPUT USIA
	fmt.Println("usia kamu berapa: ")
	fmt.Scanf("%d", &usia)

	fmt.Scanln()

	//  INPUT SEKOLAH
	if usia != 0 {
		fmt.Println("kamu sekolah di mana : ")
		fmt.Scanf("%s", &asal_sekolah)
	}

	if usia != 0 && asal_sekolah != "" {
		fmt.Printf("usia saya : %d\n", usia)
		fmt.Printf("sekolah saya  saya : %s\n", asal_sekolah)
		// fmt.Printf(quote.Go())
	}
}
