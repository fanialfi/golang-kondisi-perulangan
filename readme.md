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

## Perulangan

merupakan proses mengulang-ulang eksekusi block code tanpa henti selama kondisi yang dijadikan sebagai acuan masih terpenuhi, biasanya disiapkan variabel untuk diiterasi atau variabel penanda kapan perulangan akan berhenti.

di GO keyword perulangan hanya `for` saja, namun kemampuannya merupakan gabungan dari `for`, `forEach`, dan `while` dalam kebanyakan bahasa pemrograman.

1. Perulangan menggunakan keyword `for`

Ada beberapa cara standar menggunakan keyword `for`, cara pertama dengan memasukkan variabel counter perulangan beserta kondisi setelah keyword.

```go
for i := 0; i < 100; i++ {
		fmt.Println("angka", i)
}
```

perulangan di atas akan berjalan ketika variabel `i` bernilai dibawah 100, dengan ketentuan setiap perulangan variabel `i` akan di increment 1 (i++ artinya akan di tambah satu (i = i + 1)), karena i pada awalnya bernilai 0, maka perulangan akan berjalan sebanyak 100 kali.

2. Penggunaan `for` dengan argument hanya kondisi

cara kedua adalah dengan menuliskan kondisi setelah keyword `for`, konsepnya mirip seperti `while` pada bahasa pemrograman lain.

```go
var i = 0

for i <= 50 {
	fmt.Println("angka", i)
	i++
}
```

3. Penggunaan `for` tanpa argument

dengan begini akan dihasilkan perulangan tanpa henti, Pemberhentian peruangan dilakukan dengan menggunakan keyword `break`.

```go
var i = 0

for {
	if i >= 10 {
		break
	}
	i++
	fmt.Println("nilai", i)
}
```

4. Penggunaan keyword `break` dan `continue`

keyword `break` digunakan untuk menghentikan sebuah perulangan, sedangkan keyword `continue` digunakan untuk memajukan ke perulangan selanjutnya.

```go
for i := 1; i < 50; i++ {

	if i%2 == 1 {
		continue
	}

	if i >= 20 {
		break
	}

	fmt.Println("Angka Genap :", i)
}
```

kode di atas akan mencetak semua angka genap ke layar, dan ketika variabel `i` di modulo (%) 2 masih sisa 1 (bilangan ganjil), maka akan di lanjut ke perulangan selanjutnya dengan menggunakan keyword `continue` tanpa mengeksekusi kode setelahnya, dan ketika variabel `i` berisi 20 maka perulangan akan berhenti yang di hentikan dengan menggunakan keyword `break`.

5. Perulangan bersarang

merupakan perulangan, yang berada di dalam perulangan, yang mungkin berada didalam perulangan lagi dan seterusnya.

```go
for i := 1; i < 10; i++ {
	for j := i; j < 10; j++ {
		fmt.Print(j, " ")
	}
	fmt.Println()
}
```

6. Pemanfaatan label dalam perulangan

di perulangan bersarang, `break` dan `continue` akan berlaku pada block perulangan dimana ia digunakan saja. Penggunaan label digunakan untuk memberikan identifikasi pada sebuah perulangan, sehingga dapat mengontrol aliran eksekusi program dengan baik.

Label digunakan untuk mengubah aliran eksekusi program dengan mengarahkan perulangan atau statement lainnya untuk melompat ke posisi tertentu dalam kode.

Penggunaan label pada perulangan bersarang berguna ketika kita ingin melakukan perulangan berdasarkan kondisi tertentu yang berbeda dari perulangan yang diinginkan, atau ketika ingin keluar dari beberapa tingkat perulangan secara bersamaan.