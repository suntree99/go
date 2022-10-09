# GO Lang

## Installation
- Download dan install GO compiler di https://go.dev/
- Cek status instalasi di cmd : `go version`

Pengaturan GOPATH dan GOROOT
- Buka Environment Variables
- Pada User Variables, buat GOPATH mengacu pada alamat workspace
- Pada System Variables, buat GOROOT mengacu pada alamat instalasi
- Cek go environment di cmd : `go env`
- Pastikan GO111MODULE=off : `go env -w GO111MODULE=off`

## Basic Command
Untuk melakukan build atau compile code gunakan perintah dibawah ini
```go
go build namaFile.go
```

Untuk melakukan run code (testing) tanpa melakukan build gunakan perintah dibawah ini
```go
go run namaFile.go
```

## Hello World (Basic Structure)
```go
package main

import "fmt"

func main() {

	fmt.Println("Hello World")

}
```

## Data Type
- String
```go
	fmt.Println("Budi Darmawan") // Budi Darmawan
	fmt.Println(len("Budi Darmawan")) // 13
	fmt.Println("Budi Darmawan"[0]) // 66 (angka ASCII dari B)
```
- Number
```go
	fmt.Println("Integer =", 1) // Integer =  1
	fmt.Println("Float =", 3.5) // Float =  3.5
```
- Boolean
```go
	fmt.Println("Benar =", true) // Benar = true
	fmt.Println("Salah =", false) // Salah = false
```

## Variable
```go
	var name string
	name = "Budi Darmawan"
	fmt.Println(name) // Budi Darmawan
	
	var age = 30
	fmt.Println(age) // 30

	country := "Indonesia"
	fmt.Println(country) // Indonesia

	var (
		firstName = "Budi"
		lastName = "Darmawan"
	)
	fmt.Println(firstName) // Budi
	fmt.Println(lastName) // Darmawan
```

## Constant
```go
	const value = 1000

	const (
		firstName string = "Budi"
		lastName	 = "Darmawan"
	)

	fmt.Println(firstName) // Budi
	fmt.Println(lastName) // Darmawan
	fmt.Println(value) // 1000
```

## Conversion
```go
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 =  int8(nilai32)

	fmt.Println(nilai32) // 129
	fmt.Println(nilai64) //129
	fmt.Println(nilai8) // -127 (Overflow range int8 -> -128 s.d 127)

	var name = "Budi"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name) // Budi
	fmt.Println(eString) // B
```

## Type Declaration
```go
	type NoKTP string
	type Married bool

	var noKtpBudi NoKTP = "3603090710890001"
	var marriedStatus Married = true
	fmt.Println(noKtpBudi) // 3603090710890001
	fmt.Println(marriedStatus) // true
```

## Math Operation
```go
	var result = 10 + 10
	fmt.Println(result) // 20

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c) // 100

	var i = 10
	i += 10 // i = i + 10 
	fmt.Println(i) // 20

	i++
	fmt.Println(i) // 21

	var status = true
	var notStatus = !status
	fmt.Println(status) // true
	fmt.Println(notStatus) // false

	var negatif = -100
	var positif = 100
	fmt.Println(negatif) // -100
	fmt.Println(positif) //100
```

## Comparation Operator
```go
	var name1 = "Budi"
	var name2 = "budi"

	var result bool = name1 == name2
	fmt.Println(result) // false

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2) // false
	fmt.Println(value1 < value2) // true
	fmt.Println(value1 == value2) // false
	fmt.Println(value1 != value2) // true
```

## Boolean Operator
```go
	var ujian = 80
	var absensi = 75

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian) // true
	fmt.Println(lulusAbsensi) // false

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus) // false

	fmt.Println(ujian >= 80 && absensi >= 80) // false
```

## Array
```go
	var names [3]string

	names[0] = "Budi"
	names[1] = "Darmawan"
	names[2] = "Suntree"

	fmt.Println(names[0]) // Budi
	fmt.Println(names[1]) // Darmawan
	fmt.Println(names[2]) // Suntree

	var value = [3]int {
		90,
		95,
		80,
	}

	fmt.Println(value) // [90 95 80]

	var lagi [10]string
	
	fmt.Println(len(names)) // 3
	fmt.Println(len(value)) // 3
	fmt.Println(len(lagi)) // 10
```

## Slice
```go
	var month = [...]string {
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	month[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "Diganti"
	fmt.Println(month)

	var slice2 = month[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Tambah")
	fmt.Println(slice3)
	
	fmt.Println(slice2)
	fmt.Println(month)

	newSlice := make ([]string, 2, 5)

	newSlice[0] = "Budi"
	newSlice[1] = "Darmawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	
	copySlice := make ([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int {1, 2, 3, 4, 5}
	iniSlice := []int {1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
```

## Judul
```go
	Isi
```

##
##

# Heading 1 / Judul Utama (gunakan #)

## Heading 2 / Sub Judul (gunakan ##)

Text biasa (ditulis biasa tanpa format apapun)

[Hyperlink](https://www.google.com) (nama hyperlink dibungkus kurung siku, urlnya dibungkus tanda kurung biasa)

```bash
git add .
git commit -m "baris code menggunakan backtick 3x di awal(sertakan bahasanya) dan akhir code"
git push
```

Untuk `menyoroti` bungkus text dengan backtick 1x

Update README.md
