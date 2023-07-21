package main

import "fmt"

func cekKondisi() {
	var nilai uint8 = 200

	if nilai >= 90 {
		fmt.Println("A")
		fmt.Printf("tipe data %T nilai %d\n", nilai, nilai)
	} else if nilai >= 80 {
		fmt.Println("B")
		fmt.Printf("tipe data %T nilai %d\n", nilai, nilai)
	} else if nilai >= 70 {
		fmt.Println("C")
		fmt.Printf("tipe data %T nilai %d\n", nilai, nilai)
	} else if nilai >= 60 {
		fmt.Println("D")
		fmt.Printf("tipe data %T nilai %d\n", nilai, nilai)
	} else {
		fmt.Println("GAGAL")
		fmt.Printf("tipe data %T nilai %d\n", nilai, nilai)
	}
}

func temporary() {
	const nilai uint16 = 1000
	if percent := nilai / 100; percent >= 100 {
		fmt.Printf("%d%s, bagus!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%d%s, bagus!\n", percent, "%")
	} else if percent >= 50 {
		fmt.Printf("%d%s, bagus!\n", percent, "%")
	} else {
		fmt.Printf("%d%s, bagus!\n", percent, "%")
	}
}

func switchCase() {
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

func switchGayaIfElse() {
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

func keyFallthrough() {
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

func main() {
	fmt.Println("--------------------")
	cekKondisi()
	fmt.Println("--------------------")
	temporary()
	fmt.Println("--------------------")
	switchCase()
	fmt.Println("--------------------")
	switchGayaIfElse()
	fmt.Println("--------------------")
	keyFallthrough()
	fmt.Println("--------------------")
}
