package main

import "fmt"

const mail = `
Pemain yang terhormat,

Maaf TOP UP Anda tidak dapat diselesaikan saat ini.

Jika Anda menggunakan Google Pay, pastikan saldo Google Pay mencukupi dan saldo metode pembayaran terikatnya seperti Dana Gopay mencukupi; jika Anda menggunakan pembayaran pulsa telepon seperti Telkomsel, XL, Indosat, Tri, dll, pastikan Anda Saldo kredit kartu 1 mencukupi;

Jika Anda memiliki pertanyaan, silakan hubungi CS kami yang berdedikasi melalui metode berikut : https://linktr.ee/Csroyaldomino

Saya berharap Anda beruntung dalam permainan!
`
func main() {
	fmt.Printf("%v\n", mail)
}