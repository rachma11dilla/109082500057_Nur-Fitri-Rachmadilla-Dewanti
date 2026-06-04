# <h1 align="center">Laporan Praktikum Modul 12 - ... </h1>
<p align="center">[Nur Fitri Rachmadilla Dewanti] - [109082500057]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
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

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Println(i, ":", suara[i])
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul12/output/output-soal1.png)
[Program ini digunakan untuk menghitung hasil suara pada pemilihan ketua RT. Program akan membaca semua data suara yang di input, lalu memeriksa apakah suara tersebut valid atau tidak. Setelah itu, program akan menghitung jumlah suara masuk, jumlah suara sah, dan menampilkan berapa suara yang diperoleh masing masing calon.]



## Unguided 

### 2. [Soal]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul12/output/output-soal2.png)
[Program ini digunakan untuk menentukan ketua RT dan wakil ketua RT berdasarkan hasil pemungutan suara. Setelah seluruh suara dihitung, program akan mencari calon yang memperoleh suara paling banyak sebagai ketua RT dan calon dengan suara terbanyak berikutnya sebagai wakil ketua RT. Jika ada jumlah suara yang sama, calon dengan nomor yang lebih kecil akan dipriotaskan.]


## Unguided  

### 3. [Soal]
#### soal3.go

```go
package main
import "fmt"

const NMAX = 1000000
var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kr := 0
	kn := n - 1

	for kr <= kn {
		med := (kr + kn) / 2

		if data[med] == k {
			return med
		} else if data[med] < k {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}
	return -1
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	isiArray(n)

	hasil := posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul12/output/output-soal3.png)
[Program ini digunakan untuk mencari suatu bilangan pada kumpulan data yang sudah terurut dengan cara binary search, yaitu dengan membandingkan bilangan yang dicari dengan nilai ditengah data. Cara ini lebih cepat dari pada memeriksa data satu per satu. Jika bilangan sudah ditemukan, program akan menampilkan output posisinya. Jika tidak ditemukan, program akan menampilkan output "TIDAK ADA".]