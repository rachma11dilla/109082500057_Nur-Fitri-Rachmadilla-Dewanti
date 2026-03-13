package main
import "fmt"

func main() {

	var berat, kg, sisa int
	var biaya_kg, biaya_sisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)

	kg = berat / 1000
	sisa = berat % 1000

	biaya_kg = kg * 10000

	if kg > 10 {
		biaya_sisa = 0
	} else {
		if sisa >= 500 {
			biaya_sisa = sisa * 5
		} else {
			biaya_sisa = sisa * 15
		}
	}

	total = biaya_kg + biaya_sisa

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biaya_kg, "+ Rp.", biaya_sisa)
	fmt.Println("Total biaya: Rp.", total)

}