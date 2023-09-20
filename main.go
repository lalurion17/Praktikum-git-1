package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int = 100

	fmt.Println("Bilangan prima antara 2 dan", n, "adalah:")
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}

	fmt.Println("\nbilangan kelipatan 7")

	var kelipatan7 int = 100
	var i int = 1

	for ; i <= kelipatan7; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("\nmenhitung Luas Trapesium")

	var a, b, h float64
	a = 10
	b = 15
	h = 7

	luas := 0.5 * (a + b) * h
	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}
