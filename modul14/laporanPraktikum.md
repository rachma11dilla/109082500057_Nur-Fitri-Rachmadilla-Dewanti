# <h1 align="center">Laporan Praktikum Modul 14 - Sorting </h1>
<p align="center">[Nur Fitri Rachmadilla Dewanti] - [109082500057]</p>

## Unguided 

### 1. [Soal-SelectionSort]
#### soal selection sort/soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul14/output/selection_sort_soal1.png)
[Program ini digunakan untuk mengurutkan nomor rumah kerabat Hercules disetiap daerah dari yang terkecil hingga terbesar. Input berupa jumlah daerah, kemudian memasukkan jumlah rumah kerabat dan nomor rumah pada masing-masing daerah. Setelah semua data diterima, program akan mengurutkan nomor rumah menggunakan algoritma Selection Sort dengan cara mencari nomor rumah terkecil dan menempatkannya pada posisi yang sesuai secara berulang hingga seluruh data urut. Output berupa daftar nomor rumah yang tersusun secara menaik pada setiap daerah sehingga lebih mudah dibaca dan digunakan untuk menentukan jumlah urutan kunjungan ke rumah para kerabat.]



## Unguided 

### 2. [Soal-SelectionSort]
#### soal selection sort/soal2.go

```go
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
			if data[j]%2 != 0 {
				fmt.Print(data[j], " ")
			}
		}
		for j = m - 1; j >= 0; j-- {
			if data[j]%2 == 0 {
				fmt.Print(data[j], " ")
			}
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul14/output/selection_sort_soal2(1).png)
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul14/output/selection_sort_soal2(2).png)
[Program ini dibuat untuk membantu Hercules mengurutkan nomor rumah kerabatnya disetiap daerah dgn menggunakan Selection Sort. Setelah seluruh nomor rumah diurutkan secara menaik, program akan menampilkan nomor rumah ganjil terlebih dahulu secara terurut membesar, kemudian diikuti nomor rumah genap secara terurut mengecil. Output pada modul beda dengan deskripsi soal, seperti munculnya angka yang tidak terdapat pada data masukan dan hilangnya beberapa angka yang seharusnya ditampilkan. Oleh karena itu, program ini dibuat berdasarkan aturan yang sesuai dengan soal.]


## Unguided  

### 3. [Soal-InsertionSort]
#### soal insertion sort/soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul14/output/insertion_sort_soal1(1).png)
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul14/output/insertion_sort_soal1(2).png)
[Program ini digunakan untuk membaca sekumpulan bilangan bulat non-negatif yang diinput hingga ditemukan bilangan negatif sebagai tanda akhir input. Seluruh bilangan yang diinput akan di urutkan dari yang terkecil hingga terbesar menggunakan Insertion Sort. Setelah data terurut, program akan mengecek apakah selisih antara setiap dua bilangan yang berurutan memiliki nilai yang sama. Jika semua selisih berinilai sama, program akan menghasilkan output data memiliki jarak tetap beserta nilai jaraknya. Jika terdapat perbedaan selisih pada salah satu pasangan bilangan, output akan menampilkan data berjarak tidak tetap. Program ini tidak hanya mengurutkan data, tapi juga menganalisis jarak antar bilangan data.]

## Unguided  

### 4. [Soal-InsertionSort]
#### soal insertion sort/soal2.go

```go
package main
import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var i, idxMax int

	idxMax = 0

	for i = 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}

	fmt.Println("Buku Terfavorit:")
	fmt.Println("Judul    :", pustaka[idxMax].judul)
	fmt.Println("Penulis  :", pustaka[idxMax].penulis)
	fmt.Println("Penerbit :", pustaka[idxMax].penerbit)
	fmt.Println("Tahun    :", pustaka[idxMax].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var i, j int
	var temp Buku

	i = 1
	for i <= n-1 {

		j = i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
		i++
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i, batas int

	if n < 5 {
		batas = n
	} else {
		batas = 5
	}

	fmt.Println("5 Buku Dengan Rating Tertinggi:")

	for i = 0; i < batas; i++ {
		fmt.Println(i+1, ".", pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	var kiri, kanan, tengah int
	var ketemu bool

	kiri = 0
	kanan = n - 1
	ketemu = false

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2
		if pustaka[tengah].rating == r {

			fmt.Println("Data Buku Ditemukan:")
			fmt.Println("Judul     :", pustaka[tengah].judul)
			fmt.Println("Penulis   :", pustaka[tengah].penulis)
			fmt.Println("Penerbit  :", pustaka[tengah].penerbit)
			fmt.Println("Tahun     :", pustaka[tengah].tahun)
			fmt.Println("Eksemplar :", pustaka[tengah].eksemplar)
			fmt.Println("Rating    :", pustaka[tengah].rating)

			ketemu = true

		} else if r > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int
	fmt.Scan(&n)

	DaftarkanBuku(&pustaka, n)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	fmt.Scan(&ratingCari)
	CariBuku(pustaka, n, ratingCari)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul14/output/insertion_sort_soal2(1).png)
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul14/output/insertion_sort_soal2(2).png)
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul14/output/insertion_sort_soal2(3).png)
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul14/output/insertion_sort_soal2(4).png)
[Program ini digunakan untuk mengelola data buku pada sebuah perpustakaan. Pengguna dapat memasukkan sejumlah data buku yang terdiri dari ID buku, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. Setelah data diinput, program akan mencari dan menampilkan buku terfavorit berdasarkan rating tertinggi. Setelah itu, data buku akan diurutkan menurun (descending) berdasarkan rating menggunakan Insertion Sort, kemudian program menampilkan 5 judul buku dgn rating tertinggi. Program juga dapat digunakan mencari buku berdasarkan rating dgn menggunakan Binary Search. Jika data yang dicari ada maka akan menampilkan output informasi lengkap buku tsb, jika tidak ada program akan menampilkan output bahwa tidak ada buku dengan rating yang dicari.]