package main

func main() {
	v := 5
	p := &v
	println(*p)

	changePointer1(p)
	println(*p)
}

func changePointer1(p *int) {
	v := 3
	p = &v
}
