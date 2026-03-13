# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Nur Fitri Rachmadilla Dewanti] - [109082500057]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"
func main() {
	var (
		satu, dua, tiga string
		temp string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul1/output/output_soal1.png)
[Program ini digunakan untuk menerima 3 input string, lalu menampilkan urutan string yang di inputkan. Setelah itu, program akan menukar posisi string dari ketiga variabel tersebut. Variabe temp di gunakan untuk proses penukaran nya. Urutan nilai akan berubah dari (satu, dua, tiga) menjadi (dua, tiga, satu). Singkatnya, program melakukan pergeseran posisi pada 3 buah string yang di input kan]



## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main
import "fmt"
func main() {

	var w1, w2, w3, w4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("percobaan ", i, ": ")
		fmt.Scanln(&w1, &w2, &w3, &w4)

		if w1 != "merah" || w2 != "kuning" || w3 != "hijau" || w4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Println("berhasil: ", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul1/output/output_soal2.png)
[Program ini dibuat untuk mengecek apakah hasil percobaan kimia yang dilakukan sebanyak 5 kali percobaan memiliki warna yang benar atau tidak. Variabel w1, w2, w3, w4 digunakan untuk menyimpan warna cairan pada tabung reaksi. Jika hasil input tidak sesuai dengan ketentuan yang di minta, maka variabel "berhasil" akan berubah menjadi "false" karena menggunakan kondisi if dengan operator OR (atau)]



## Unguided 

### 3. [Soal]
#### soal3.go

```go
package main
import "fmt"

func main() {

	var berat, kg, sisa int
	var biaya_kg, biaya_sisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)

	kg = berat / 1000
	sisa = berat % 1000

	biaya_kg = kg * 10000

	if kg > 10 {
		biaya_sisa = 0
	} else {
		if sisa >= 500 {
			biaya_sisa = sisa * 5
		} else {
			biaya_sisa = sisa * 15
		}
	}

	total = biaya_kg + biaya_sisa

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biaya_kg, "+ Rp.", biaya_sisa)
	fmt.Println("Total biaya: Rp.", total)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachma11dilla/109082500057_Nur-Fitri-Rachmadilla-Dewanti/blob/main/modul1/output/output_soal3.png)
[Program ini dibuat untuk menghitung biaya pengiriman parsel berdasarkan beratnya dalam gram. Input berupa berat parsel (gram). Kemudian program akan menghitung jumlah kilogram dengan membagi dengan 1000, dan sisa gram menggunakan modulus dengan biaya pengiriman setiap kilogramnya Rp10.000. Setelah itu program akan menghitung biaya dari sisa gram. Jika total berat lebih dari 10kg, maka sisa gram di gratiskan. Jika tidak, maka sisa gram akan dikenakan biaya Rp5 per gram jika >= 500 gram, atau Rp15 jika < 500 gram]