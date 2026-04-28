package main
import "fmt"

const NMAX = 100

type arrString [NMAX]string

func main() {
	var klubA, klubB string
	var hasil arrString
	var n int = 0

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	var skorA, skorB int
	i := 1

	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[n] = klubA
		} else if skorB > skorA {
			hasil[n] = klubB
		} else {
			hasil[n] = "Draw"
		}
		n++
		i++
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}