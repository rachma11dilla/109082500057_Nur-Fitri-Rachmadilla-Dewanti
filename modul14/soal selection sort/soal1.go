package main
import "fmt"

type arrInt [1000000]int

func selectionSort(T *arrInt, n int) {
	var i, j, idxMin, temp int

	i = 1
	for i <= n-1 {
		idxMin = i - 1
		j = i
		for j < n {
			if T[idxMin] > T[j] {
				idxMin = j
			}
			j++
		}
		temp = T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = temp
		i++
	}
}

func main() {
	var n, m, i, j int
	var data arrInt
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&m)

		for j = 0; j < m; j++ {
			fmt.Scan(&data[j])
		}
		selectionSort(&data, m)

		for j = 0; j < m; j++ {
			fmt.Print(data[j], " ")
		}
		fmt.Println()
	}
}