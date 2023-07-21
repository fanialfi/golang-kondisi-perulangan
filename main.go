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
func perulanganFor() {
	for i := 0; i < 100; i++ {
		fmt.Println("angka", i)
	}
}

func forHanyaKondisi() {
	var i = 0
	for i <= 50 {
		fmt.Println("angka", i)
		i++
	}
}

func tanpaArgument() {
	var i = 0

	for {
		if i >= 10 {
			break
		}
		i++
		fmt.Println("nilai", i)
	}
}

func breakContinue() {
	for i := 1; i < 50; i++ {
		if i%2 == 1 {
			continue
		}

		if i >= 20 {
			break
		}

		fmt.Println("Angka Genap :", i)
	}
}

func bersarang() {
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func label() {
outerLoop: // penamaan label
	for i := 0; i < 5; i++ {
	innerLoop:
		for j := 0; j < 5; j++ {
			if j == 3 {
				fmt.Printf("Keluar dari inner loop\n")
				break innerLoop
			}
			if i == 3 {
				fmt.Printf("Keluar dari inner dan outer loop\n")
				break outerLoop
			}
			fmt.Print("Matriks [", i, "][", j, "]", "\n")
		}
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
	perulanganFor()
	fmt.Println("--------------------")
	forHanyaKondisi()
	fmt.Println("--------------------")
	tanpaArgument()
	fmt.Println("--------------------")
	breakContinue()
	fmt.Println("--------------------")
	bersarang()
	fmt.Println("--------------------")
	label()
	fmt.Println("--------------------")
}
