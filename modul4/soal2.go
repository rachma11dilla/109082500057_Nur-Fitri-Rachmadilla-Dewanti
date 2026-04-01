package main
import "fmt"

func hitungskor(soal []int, jumlah *int, skor *int) {
	*jumlah = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		if soal[i] < 301 {
			*jumlah++
			*skor += soal[i]
		}
	}
}

func main() {
	var nama, pemenang string
	var soal [8]int
	var jumlah, skor int
	var maxSoal, minSkor int

	maxSoal = -1
	minSkor = 999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}
		for i := 0; i < 8; i++ {
			fmt.Scan(&soal[i])
		}

		hitungskor(soal[:], &jumlah, &skor)

		if jumlah > maxSoal || (jumlah == maxSoal && skor < minSkor) {
			maxSoal = jumlah
			minSkor = skor
			pemenang = nama
		}
	}
	fmt.Println(pemenang, maxSoal, minSkor)
}