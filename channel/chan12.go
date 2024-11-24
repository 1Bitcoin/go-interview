package main

func main() {
	var a chan int

	//<- a
	a <- 1223
}
