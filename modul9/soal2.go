package main
import "fmt"

const NMAX = 100

func main() {
	var A [NMAX]int
	var n int

	fmt.Print("jumlah elemen: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	fmt.Print("semua: ")
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	sum := 0
	for i := 0; i < n; i++ {
		sum += A[i]
	}
	mean := float64(sum) / float64(n)
	fmt.Println("rata-rata:", mean)

	var jumlah float64
	for i := 0; i < n; i++ {
		selisih := float64(A[i]) - mean
		jumlah += selisih * selisih
	}
	variansi := jumlah / float64(n)

	fmt.Println("variansi:", variansi)
}