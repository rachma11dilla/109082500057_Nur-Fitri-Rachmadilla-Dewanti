# <h1 align="center">Laporan Praktikum Modul 5 - ... </h1>
<p align="center">[Nur Fitri Rachmadilla Dewanti] - [109082500057]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

func main () {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}

func fibonacci (n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul5/output/output-soal1.png)
[Program ini dibuat untuk menampilkan deret fibonacci sampai suku ke-n. Di func fibonacci (n) ada 2 base-case, yaitu n == 0 akan mengembalikan ke 0 dan n == 1 akan mengembalikan ke 1. Selain itu, ada juga rekursif, yaitu fibonacci(n-1) dan fibonacci(n-2). Program melakukan perulangan dari 0 sampai n untuk menghasilkan output fibonacci. Oleh karena itu, rekursif ini digunakan untuk menghitung nilai tiap suku.] 



## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main
import "fmt"

func main () {
	var n int
	fmt.Scan(&n)
	bintang(n,1)
}

func bintang (n int, i int) {
	if i > n {
		return
	}
	for j := 1; j <= i; j++ {
		fmt.Print("*")
	}
	fmt.Println()
	bintang(n, i+1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul5/output/output-soal2.png)
[Program ini digunakan untuk menampilkan segitiga ke bawah dengan menggunakan pola * (bintang) sesuai angka yg di input. func bintang (n, i) ini adalah rekursif dengan parameter i dan akan menjalankan func bintang (n, i+1) untuk melanjutkan ke baris berikutnya. Ini yang membuat pola bintang terbentuk dari 1 sampai n secara bertahap.]



## Unguided  

### 3. [Soal]
#### soal3.go

```go
package main
import "fmt"

func main () {
	var n int
	fmt.Scan(&n)
	faktor (n, 1)
}

func faktor (n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktor(n, i+1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul5/output/output-soal3.png)
[Program ini digunakan untuk mencari dan menghasilkan faktor dari suatu bilangan (n). func faktor (n, i) digunakan untuk mengecek setiap bilangan dari 1 sampai n. Base-case ini terjadi di i > n, dan proses rekursif akan berhenti. Program akan mengecek apakah n%1 == 0. Jika benar maka i adalah faktor dari n dan akan di tampilkan di output. Setelah nya, fungsi akan melanjutkan i+1 sampai semua kemungkinan di cek.]



## Unguided  

### 4. [Soal]
#### soal4.go

```go
package main
import "fmt"

func main () {
	var n int
	fmt.Scan(&n)
	baris(n)
}

func baris (n int) {
	if n == 1 {
		fmt.Print(1, " ")
	} else {
		fmt.Print(n, " ")
		baris(n-1)
		fmt.Print(n, " ")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul5/output/output-soal4.png)
[Program ini digunakan untuk menampilkan pola bilangan dari n ke 1 (mundur), lalu kembali lagi ke n (maju). func baris (n) ini ada base-case nya yaitu n == 1, yang berfungsi program hanya menghasilkan output angka 1. Selain itu, program juga akan menampilkan output  nilai n, dan memanggil func baris(n-1). Setelah proses rekursif selesai, program akan kembali menampilkan output n. Ini yang membuat output menjadi mundur lalu maju kembali (turun ke naik).]



## Unguided  

### 5. [Soal]
#### soal5.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul5/output/output-soal5.png)
[Program ini digunakan untuk menampilkan bilangan ganjil dari 1 sampai n. func ganjil (n,i) akan berjalan dari i = 1 sampai n. Base-case terjadi di i > n. Program akan mengecek apakah i%2 != 0. Jika benar, maka i adalah bilangan ganjil dan akan di tampilkan di output. Kemudian fungsi akan kembali berjalan dengan i+1 untuk melanjutkan ke angka berikutnya.]



## Unguided  

### 6. [Soal]
#### soal6.go

```go
package main
import "fmt"

func main () {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(pangkat(x, y))
}

func pangkat (x int, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * pangkat(x, y-1)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul5/output/output-soal6.png)
[Program ini digunakan untuk menghitung hasil pangkat suatu bilangan menggunakan rekursif. func pangkat (x, y) memiliki base-case di y == 0, yang hasilnya adalah 1 (karena semua bilangan yang di pangkatkan 0 hasilnya adalah 1). Fungsi akan berjalan di x * pangkat(x, y-1). Maksudnya adalah nilai x di kali terus sampai nilai y nya = 0.]