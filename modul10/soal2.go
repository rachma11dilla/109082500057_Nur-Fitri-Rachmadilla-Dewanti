package main
import "fmt"

const NMAX int = 1000
type arrFloat [NMAX]float64

func main() {
	var ikan arrFloat
	var wadah arrFloat
	var x, y int
	var jumlahWadah int
	var rata float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}
	var idx int = 0

	for idx < x {
		var total float64 = 0

		for j := 0; j < y && idx < x; j++ {
			total = total + ikan[idx]
			idx++
		}
		wadah[jumlahWadah] = total
		rata = rata + total
		jumlahWadah++
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f ", wadah[i])
	}
	fmt.Println()

	rata = rata / float64(jumlahWadah)

	fmt.Printf("Rata-rata: %.2f\n", rata)
}