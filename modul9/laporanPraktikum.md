# <h1 align="center">Laporan Praktikum Modul 9 - ... </h1>
<p align="center">[Nur Fitri Rachmadilla Dewanti] - [109082500057]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

type titik struct {
	x, y int
}

type lingkaran struct {
	c titik
	r int
}

func didalam(c lingkaran, p titik) bool {
	dx := p.x - c.c.x
	dy := p.y - c.c.y

	return dx*dx+dy*dy <= c.r*c.r
}

func main() {
	var c1, c2 lingkaran
	var p titik

	fmt.Scan(&c1.c.x, &c1.c.y, &c1.r)
	fmt.Scan(&c2.c.x, &c2.c.y, &c2.r)
	fmt.Scan(&p.x, &p.y)

	in1 := didalam(c1, p)
	in2 := didalam(c2, p)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul9/output/output-soal1.png)
[Program ini dibuat untuk mengecek apakah sebuah titik berada di dalam satu atau dua lingkaran. (x, y) untuk menyimpan posisi titik pasangan koordinat dan lingkaran yang memiliki (pusat & jari-jari). Lalu ada fungsi untuk menghitung jarak titik ke pusat lingkaran. Jika jaraknya tidak lebih besar dari jari-jari berarti titiknya ada di dalam lingkaran. Inputan yang berupa dua lingkaran dan satu titik, kemudian program akan mengecek posisi titik tersebut terhadap masing-masing lingkaran, selanjutnya menampilkan apakah titik itu ada di kedua lingkaran, hanya di salah satu, atau di luar keduanya.] 



## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main
import "fmt"

const NMAX = 100

func main() {
	var A [NMAX]int
	var n int

	fmt.Print("jumlah elemen: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	fmt.Print("semua: ")
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	sum := 0
	for i := 0; i < n; i++ {
		sum += A[i]
	}
	mean := float64(sum) / float64(n)
	fmt.Println("rata-rata:", mean)

	var jumlah float64
	for i := 0; i < n; i++ {
		selisih := float64(A[i]) - mean
		jumlah += selisih * selisih
	}
	variansi := jumlah / float64(n)

	fmt.Println("variansi:", variansi)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul9/output/output-soal2.png)
[Program ini digunakan untuk mengolah data angka dalam array, dimulai dari input "jumlah elemen (n)", lalu menginput nilai-nilai nya ke dalam array, selanjutnya menampilkan semua data yang dimasukkan dan menghitung rata-rata dengan menjumlahkan semua nilai lalu dibagi jumlah data. Step terakhir menghitung variansi dengan cara melihat seberapa jauh tiap nilai dari rata-rata lalu di kuadratkan dan di rata-ratakan, sehingga penyebaran datanya diketahui.]



## Unguided  

### 3. [Soal]
#### soal3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul9/output/output-soal3.png)
[Program ini digunakan untuk mencatat hasil beberapa pertandingan antara dua klub, input awal berupa nama kedua klub, lalu inputan berupa skor tiap pertandingan secara berulang sampai ada skor negatif yang berarti inputan akan berhenti, setiap hasil pertandingan dicek yang menang atau seri, lalu akan di simpan ke dalam array (nama klub pemenang atau draw). Setelah input selesai, program akan menampilkan daftar hasil tiap pertandingan yang sudah di mainkan lalu menampilkan output "pertandingan selesai".]


## Unguided  

### 4. [Soal]
#### soal4.go

```go
package main
import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	for {
		fmt.Scanf("%c", &ch)

		if ch == '.' || *n >= NMAX {
			break
		}
		(*t)[*n] = ch
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := (*t)[i]
		(*t)[i] = (*t)[n-1-i]
		(*t)[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks : ")
	isiArray(&tab, &m)

	fmt.Print("Palindrom ? ")
	if palindrom(tab, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul9/output/output-soal4(1).png)
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul9/output/output-soal4(2).png)
[Program ini digunakan untuk mengolah teks karakter satu per satu, di mulai dari membaca input huruf hingga tanda titik sebagai penanda input selesai, lalu di simpan ke dalam array. Setelah itu program akan mengecek apakah teks tersebut merupakan palindrom (di baca sama dari depan dan belakang), kemudian membalik urutan karakter dalam array, dan akan menampilkan output hasil teks yang sudah di balik.]