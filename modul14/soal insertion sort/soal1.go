package main
import "fmt"

type arrInt [1000]int

func insertionSort(T *arrInt, n int) {
	var i, j, temp int

	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j--
		}
		T[j] = temp
		i++
	}
}

func main() {
	var data arrInt
	var x, n, i int
	var tetap bool
	var selisih int
	n = 0

	for {
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		data[n] = x
		n++
	}
	insertionSort(&data, n)

	for i = 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	if n <= 2 {
		fmt.Println("Data berjarak tetap")
		return
	}

	selisih = data[1] - data[0]
	tetap = true

	for i = 2; i < n; i++ {
		if data[i]-data[i-1] != selisih {
			tetap = false
		}
	}

	if tetap {
		fmt.Println("Data berjarak", selisih)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}