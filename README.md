# Reverse String

Fungsi `ReverseString` menerima sebuah string sebagai input dan mengembalikan string tersebut dalam urutan terbalik. Fungsi ini menggunakan perulangan untuk mengakses setiap karakter dari string asli dan membangun string baru dengan urutan karakter yang terbalik.

## Contoh Penggunaan

Berikut adalah contoh penggunaan fungsi `ReverseString`:

```go
// Contoh 1: Membalikkan string "indra"
indraReverse := ReverseString("indra")
fmt.Println(indraReverse) // Output: ardni

// Contoh 2: Membalikkan string "golang"
golangReverse := ReverseString("golang")
fmt.Println(golangReverse) // Output: gnalog

// Contoh 3: Membalikkan string "hello world"
helloReverse := ReverseString("hello world")
fmt.Println(helloReverse) // Output: dlrow olleh
```

## Cara Menjalankan Program

Untuk menjalankan program ini, anda dapat menggunakan perintah `go run main.go` di terminal.

```
 git clone https://github.com/funukonta/reverse-string.git
 cd reverse-string
 go run main.go
```