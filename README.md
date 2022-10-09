# GO Lang

## Instalasi
- Download dan install GO compiler di https://go.dev/
- Cek status instalasi di cmd : `go version`

Pengaturan GOPATH dan GOROOT
- Buka Environment Variables
- Pada User Variables, buat GOPATH mengacu pada alamat workspace
- Pada System Variables, buat GOROOT mengacu pada alamat instalasi
- Cek go environment di cmd : `go env`
- Pastikan GO111MODULE=off : `go env -w GO111MODULE=off`

## Perintah Dasar
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

## Tipe Data
- String
```go
	fmt.Println("Budi Darmawan") // Budi Darmawan
	fmt.Println(len("Budi Darmawan")) // 13
	fmt.Println("Budi Darmawan"[0]) // 66 (angka ASCII dari B)
```
- Angka
```go
	fmt.Println("Satu = ", 1) // Satu =  1
	fmt.Println("Tiga Koma Lima = ", 3.5) // Tiga Koma Lima =  3.5
```
- Boolean
```go
	fmt.Println("Benar = ", true) // Benar = true
	fmt.Println("Salah = ", false) // Salah = false
```

## Variable
- String
```go
	var name string
	name = "Budi"
	fmt.Println(name) // Budi
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
