# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>
<p align="center">[Nur Fitri Rachmadilla Dewanti] - [109082500057]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul10/output/output-soal1.png)
[Program ini dibuat untuk mencari berat kelinci dari yang paling terkecil hingga terbesar dari data kelinci yang di input. Input pertama berupa jumlah kelinci, input kedua berupa berat kelinci yang di simpan ke dalam array. Cara kerja nya dengan membandingkan setiap elemen array satu per satu dari data pertama hingga data terakhir. Setelah seluruh data selesai di cek, hasil output berupa berat kelinci yang terkecil dan terbesar.] 



## Unguided 

### 2. [Soal]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul10/output/output-soal2.png)
[Program ini digunakan untuk menghitung total berat ikan pada tiap wadah dan mencari rata-rata berat dari seluruh wadah yang digunakan. Input pertama berupa jumlah ikan yang akan di jual. Input kedua berupa berat ikan dan berat masing-masing ikan di simpan ke dalam array. Cara kerja nya yaitu, ikan akan di kelompokkan ke dalam beberapa wadah sesuai urutan input dan kapasitas wadah. Setiap berat ikan pada wadah yang sama akan di jumlahkan, hasilnya yaitu total berat untuk masing-masing wadah. Output berupa total berat tiap wadah dan rata-rata berat seluruh wadah dengan menjumlahkan seluruh total wadah lalu dibagi dengan jumlah wadah yang terbentuk.]


## Unguided  

### 3. [Soal]
#### soal3.go

```go
package main
import "fmt"

const NMAX int = 100
type arrBalita [NMAX]float64

func inputData(A *arrBalita, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&A[i])
	}
}

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total = total + arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	inputData(&data, n)
	hitungMinMax(data, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(data, n))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul10/output/output-soal3.png)
[Program ini digunakan untuk mencari berat balita dari terkecil ke terbesar dan rata-rata. Input pertama berupa jumlah balita, input kedua hingga seterusnya berupa berat balita dan disimpan ke dalam array. Cara kerjanya yaitu, membandingkan setiap elemen array satu per satu. Selanjutnya program akan menghitung total seluruh berat balita dengan menjumlahkannya, lalu dibagi dengan jumlah balita untuk menghasilkan rata-rata. Output nya berupa berat balita minimum, berat balita maksimum, dan rata-ratanya.]
