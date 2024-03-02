package concurrency

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func Channels() {
	s := []int{7, 2, 8, 9, 4, 0}

	c := make(chan int)

	a := s[:len(s)/2]
	b := s[len(s)/2:]

	fmt.Println(a)
	fmt.Println(b)

	go sum(a, c)
	go sum(b, c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
