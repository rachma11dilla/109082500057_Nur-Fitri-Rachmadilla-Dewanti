package main
import "fmt"

type titik struct {
	x, y int
}

type lingkaran struct {
	c titik
	r int
}

func didalam(c lingkaran, p titik) bool {
	dx := p.x - c.c.x
	dy := p.y - c.c.y

	return dx*dx+dy*dy <= c.r*c.r
}

func main() {
	var c1, c2 lingkaran
	var p titik

	fmt.Scan(&c1.c.x, &c1.c.y, &c1.r)
	fmt.Scan(&c2.c.x, &c2.c.y, &c2.r)
	fmt.Scan(&p.x, &p.y)

	in1 := didalam(c1, p)
	in2 := didalam(c2, p)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}