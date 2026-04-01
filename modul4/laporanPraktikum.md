# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>
<p align="center">[Nur Fitri Rachmadilla Dewanti] - [109082500057]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i<=n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var fn, fnr int
	faktorial(n, &fn)
	faktorial(n-r, &fnr)
	*hasil = fn / fnr
}

func kombinasi(n, r int, hasil  *int) {
	var fn, fr, fnr int
	faktorial(n, &fn)
	faktorial(r, &fr)
	faktorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int
	fmt.Scan(&a, &b, &c, &d)

	permutasi(a, c, &p1)
	kombinasi(a, c, &c1)

	permutasi(b, d, &p2)
	kombinasi(b, d, &c2)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul4/output/output-soal1.png)
[Program ini dibuat untuk menghitung nilai permutasi dan kombinasi dari dua pasang bilangan, yaitu (a, c) dan (b, d). Program yang di input menggunakan fungsi faktorial untuk menghitung nilai faktorial dari suatu bilangan, caranya mengalikan angka dari 1 sampai n secara berulang. Hasil faktorial ini digunakan dalam fungsi permutasi dengan rumus P(n, r) = n! / (n - r)! dan fungsi kombinasi dengan rumus C(n, r) = n! / (r! × (n - r)!). Setelah semua perhitungan selesai, program menampilkan output yaitu permutasi dan kombinasi untuk kedua pasangan bilangan dalam dua baris output. Program ini menunjukkan penggunaan fungsi, perulangan, serta konsep faktorial.]



## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main
import "fmt"

func hitungskor(soal []int, jumlah *int, skor *int) {
	*jumlah = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		if soal[i] < 301 {
			*jumlah++
			*skor += soal[i]
		}
	}
}

func main() {
	var nama, pemenang string
	var soal [8]int
	var jumlah, skor int
	var maxSoal, minSkor int

	maxSoal = -1
	minSkor = 999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}
		for i := 0; i < 8; i++ {
			fmt.Scan(&soal[i])
		}

		hitungskor(soal[:], &jumlah, &skor)

		if jumlah > maxSoal || (jumlah == maxSoal && skor < minSkor) {
			maxSoal = jumlah
			minSkor = skor
			pemenang = nama
		}
	}
	fmt.Println(pemenang, maxSoal, minSkor)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul4/output/output-soal2.png)
[Program ini digunakan untuk menentukan pemenang dalam kompetisi pemrograman dengan melihat jumlah soal yang berhasil diselesaikan dan total waktu pengerjaan. Setiap peserta memiliki nama serta delapan data waktu untuk setiap soal. Jika waktu pengerjaan kurang dari 301 menit, maka soal dianggap berhasil dan waktunya akan dijumlahkan, sedangkan jika 301 menit atau lebih maka dianggap tidak selesai. Program menggunakan prosedur hitungSkor untuk menghitung jumlah soal yang berhasil diselesaikan dan total waktunya. Di bagian utama, data peserta dibaca satu per satu hingga ditemukan input “Selesai”, kemudian setiap peserta dibandingkan untuk menentukan pemenang, yaitu yang menyelesaikan soal paling banyak, dan jika jumlahnya sama maka dipilih yang memiliki total waktu paling kecil. Pada akhirnya, program menampilkan nama pemenang beserta jumlah soal yang diselesaikan dan total waktu yang diperoleh.]


## Unguided 

### 3. [Soal]
#### soal3.go

```go
package main
import "fmt"

func cetakDeret(n int) {
	for {
		fmt.Print(n, " ")

		if n == 1 {
			break
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	cetakDeret(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul4/output/output-soal3.png)
[Program ini digunakan untuk mencetak deret bilangan yang dimulai dari sebuah bilangan awal n. Jika bilangan tersebut genap, maka nilai berikutnya adalah n dibagi 2, jika ganjil maka nilai berikutnya adalah 3n + 1. Proses ini dilakukan terus-menerus hingga mencapai nilai 1. Program menggunakan fungsi cetakDeret untuk mencetak angka-angka deret satu per satu dalam satu baris. Input berupa sebuah bilangan, bilangan yang di input akan diproses untuk menghasilkan deret hingga berakhir di angka 1.