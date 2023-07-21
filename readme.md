# pengkondisian dan perulangan

## pengkondisian

digunakan untuk mengkontrol alur program. Yang digunakan sebagai acuan oleh seleksi kondisi disini adalah nilai bertipe `bool`, bisa berasal dari variabel, ataupun operasi perbandingan.

GO memiliki 2 macam keyword yang digunakan untuk melakukan seleksi kondisi, yaitu `if else` dan `switch`.

1. seleksi kondisi dengan menggunakan `if`, `else if`, dan `else`.

penerapannya didalam GO sama seperti pada kebanyakan bahasa pemrograman, yang membedakan hanyalah tanda kurung () dan di go tidak perlu di tulis.

2. variabel temporary pada `if`, `else`

yaitu variabel yang hanya bisa digunakan pada block seleksi kondisi dimana ia di tempatkan, penggunaan variabel ini membawa beberapa manfaat antara lain :

- scope atau cangkupan variabel jelas, hanya bisa digunakan pada seleksi kondisi itu saja.
- ketika nilai variabel tersebut didapat dari sebuah komputasi, perhitungan tidak perlu dilakukan didalam block masing masing.

deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang menggunakan tanda `:=` dan penggunaan keyword var disitu tidak diperbolehkan karena akan menyebabkan error.

3. seleksi kondisi dengan menggunakan keyword `switch` `case`

merupakan seleksi kondisi yang fokus pada satu variabel, lalu kemdian di cek nilainya. Contoh sederhananya seperti penentuan apakah variabel x nilainya adalah 1,2,3, atau yang lain. pada `switch` statement di GO, ketika sebuah case terpenuhi maka tidak akan di lanjutkan ke pengecekan selanjutnya, meskipun tidak ada keyword `break` disitu, konsep ini berbanding terbalik dengan `switch` sattement pada bahasa pemrograman pada umumnya, dimana jika sebuah case terpenuhi dan tidak ada keyword `break` maka akan lanjut ke pengecekan selanjutnya.

sebuah `case` pada `switch` statement pada go dapat menampung banyak kondisi, cara penerapannya yaitu dengan menuliskan nilai nilai yang akan di lakukan pengecekan setelah keyword `case` dan dipisahkan dengan tanda `,`.

```go
var nilai byte = 60

switch nilai {
case 100:
	fmt.Printf("Good Job, %d\n", nilai)
case 80:
	fmt.Printf("Perfect, %d\n", nilai)
case 70, 60, 50:
	fmt.Printf("Awesome, %d\n", nilai)
default:
	fmt.Printf("Not Bad, %d\n", nilai)
}
```

selain itu tanda kurung kurawal (`{}`) bisa diterapkan pada keyword `case` dan `default` pada `switch` statement, tanda ini opsional bisa dipakai bisa tidak.

penggunaan tanda kurung kurawal bagus digunakan ketika terjadi pada blok kondiis didalamnya terdapat banyak statement, dan kode akan terlihat menjadi lebih rapi.

```go
func main() {
	var nilai byte = 30

	switch nilai {
	case 100:
		fmt.Printf("Good Job, %d\n", nilai)
	case 80:
		fmt.Printf("Perfect, %d\n", nilai)
	case 70, 60, 50:
		fmt.Printf("Awesome, %d\n", nilai)
	default:
		{
			fmt.Printf("Not Bad, %d\n", nilai)
			fmt.Println("Maaf anda gagal")
		}
	}
}
```

menariknya di GO, `switch` statement bisa digunakan dengan menggunakan gaya `if else`. Nilai yang akan dibandingkan tidak ditulis setelah keyword `switch`, melainkan akan ditulis langsung dalam bentuk perbandingan dalam keyword `case`.

```go
func main() {
	const nilai uint8 = 50

	// penggunaan switch case dengan menggunakan gaya if else
	switch {
	case nilai >= 100:
		fmt.Printf("Perfect %d\n", nilai)
	case nilai >= 80:
		fmt.Printf("Good %d\n", nilai)
	case (nilai <= 80) && (nilai >= 40):
		{
			fmt.Printf("Awesome %d\n", nilai)
			fmt.Println("Tidak terlalu buruk")
		}
	default:
		fmt.Printf("Not Bad %d\n", nilai)
	}
}
```

kebanyakan dalam bahasa pemrograman pada `switch` statement jika ingin sebuah kondisi tetap dijalankan meskipun bernilai false, kita bisa menghilangkan keyword `break` pada `switch` statement tersebut. Berbeda dengan GO, ketika sebuah kondisi terpenuhi (bernilai true) maka secara otomatis akan langsung keluar dari `switch` statement tersebut, jika kita ingin membuat `switch` statement tersebut tidak langsung keluar melainkan juga akan mengeksekusi kondisi setelahnya meskipun kondisinya bernilai falsy, kita bisa menggunakan keyword `fallthrough` dalam `switch` statement nya.

keyword `fallthrough` akan memaksa proses pengecekan untuk diteruskan ke `case` berikutnya **tanpa menghiraukan nilai kondisinya**, jadi satu case selanjutnya akan selalu dianggak true (meskipun aslinya adalah false)

```go
func main() {
	const point byte = 6

	switch point {
	case 10, 9:
		fmt.Println("perfect ", point)
	case 8, 7, 6, 5:
		fmt.Println("Awesome", point)
		fallthrough
		// penggunaan keyword ini, case berikutnya akan dihiraukan kebenarannya (menghiraukan nilai kondisinya) dan akan dijalankan
	case 4, 3, 2:
		fmt.Println("Bad", point)
	default:
		fmt.Println("?")
	}
}
```