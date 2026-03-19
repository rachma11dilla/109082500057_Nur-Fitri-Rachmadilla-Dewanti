# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Nur Fitri Rachmadilla Dewanti] - [109082500057]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a,c))
	fmt.Println(permutation(b, d), combination(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul3/output/output-soal1.png)
[Program ini digunakan untuk menghitung permutasi dan kombinasi dari dua pasang bilangan, yaitu (a, c) dan (b, d). Perhitungan dilakukan menggunakan fungsi faktorial, lalu digunakan dalam fungsi permutasi dan kombinasi. Input berupa empat bilangan, kemudian hasilnya ditampilkan dalam dua baris sesuai pasangan masing-masing. Fungsi digunakan agar prgram menjadi lebih rapih]



## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main
import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)
	hasil1 := f(g(h(a)))
	hasil2 := g(h(f(b)))
	hasil3 := h(f(g(c)))

	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul3/output/output-soal2.png)
[Program ini dibuat untuk menghitung 3 fungsi komposisi, yaitu f(x), g(x), dan h(x). Nilai dari fungsi akan di masukkan ke fungsi lain sesuai dengan urutan yang di minta atau di input. Input berupa tiga bilangan, kemudian program akan menghitung hasil komposisi fungsi dan akan menampilkannya dalam tiga baris. Fungsi yang digunakan membuat program lebih terlihat rapih dan mudah dipahami]



## Unguided 

### 3. [Soal]
#### soal3.go

```go
package main
import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(x, y, cx, cy) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	in1 := didalam(cx1, cy1, r1, x, y)
	in2 := didalam(cx2, cy2, r2, x, y)

	if in1 && in2 {
		fmt.Println("Titik didalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik diluar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul3/output/output-soal3.png)
[Program ini digunakan untuk menentukan posisi suatu titik terhadap dua lingkaran. Cara kerja program yaitu, menghitung jarak titik ke pusat masing-masing lingkaran, lalu akan membandingkannya dengan jari-jari lingkaran. Jika jarak lebih kecil sama dengan jari-jari, maka titik berada di dalam lingkaran. Hasil output akan menampilkan apakah titik berada di dalam lingkaran 1 atau 2, atau di dalam lingkaran 1 dan 2, atau di luar lingkaran1 dan 2]