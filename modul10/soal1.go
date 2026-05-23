package main
import "fmt"

const NMAX int = 1000
type arrFloat [NMAX]float64

func inputData(A *arrFloat, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}

func cariMin(A arrFloat, n int) float64 {
	var min float64 = A[0]
	for i := 1; i < n; i++ {
		if A[i] < min {
			min = A[i]
		}
	}
	return min
}

func cariMax(A arrFloat, n int) float64 {
	var max float64 = A[0]
	for i := 1; i < n; i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	return max
}

func main() {
	var A arrFloat
	var n int
	var min, max float64
	inputData(&A, &n)
	min = cariMin(A, n)
	max = cariMax(A, n)

	fmt.Printf("Berat kelinci terkecil: %.2f\n", min)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", max)
}