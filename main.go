package main

import "fmt"

func main() {
	const (
		i = 7
		j
		k
	)

	fmt.Println(i, j, k)

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]
	news := append(s, 55, 66)

	fmt.Println(len(news), cap(news))

	d := []string{"1", "2", "3", "4", "5"}
	d = d[:0]

	fmt.Println(d, len(d), cap(d))
}

func hh() {

	fmt.Println("hhhh")
}
