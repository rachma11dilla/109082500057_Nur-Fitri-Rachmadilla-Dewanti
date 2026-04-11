package main
import "fmt"

func main () {
	var n int
	fmt.Scan(&n)
	ganjil(n, 1)
}

func ganjil (n int, i int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	ganjil(n, i+1)
}