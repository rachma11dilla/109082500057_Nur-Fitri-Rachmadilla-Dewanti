package main
import "fmt"

func main() {
	var suara [21]int
	var x int
	var totalMasuk, totalSah int

	fmt.Scan(&x)

	for x != 0 {
		totalMasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			totalSah++
		}
		fmt.Scan(&x)
	}

	ketua := 1

	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	wakil := -1

	for i := 1; i <= 20; i++ {
		if i != ketua {
			if wakil == -1 || suara[i] > suara[wakil] {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}