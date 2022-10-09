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
```cmd
go build namaFile.go
```

Untuk melakukan run code (testing) tanpa melakukan build gunakan perintah dibawah ini
```cmd
go run namaFile.go
```

## Hello World
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
- Angka
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
