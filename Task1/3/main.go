package main

func main() {
	mas := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	sum := 0
	for _, elem := range mas {
		go Sum(elem*elem, ch)
		a := <-ch
		sum += a
	}
	println(sum)
}

func Sum(elem int, ch chan int) {
	ch <- elem
}
