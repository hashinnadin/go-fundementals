// package main

// import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

// func newPerson(name string, age int) *person {

// 	return &person{
// 		name: name,
// 		age:  age,
// 	}ee
// }

// func main() {

// 	p1 := newPerson("hashin", 21)
// 	fmt.Println(p1)

// }

// package main

// import "fmt"

// func main() {
// 	valu := []int{}

// 	valu = append(valu, 2)
// 	valu[0] = 1
// 	fmt.Println(valu)
// }

// package main

// import "fmt"

// func co(num *int) {
// 	*num = 0
// }
// func main() {
// 	val := 1

// 	fmt.Println(val)
// 	co(&val)
// 	fmt.Println(val)
// }

// package main

// import "fmt"

// type rect struct {
// 	width, hiegt int
// }

// func (r rect) area() int {
// 	return r.width * r.hiegt
// }

// func newRect(width int, hiegt int) rect {
// 	if width <= 0 {
// 		width = 1
// 	}
// 	if hiegt <= 0 {
// 		hiegt = 1
// 	}
// 	return rect{
// 		hiegt: hiegt,
// 		width: width,
// 	}
// }

// func (p rect) preim() int {
// 	return 2*p.hiegt + 2*p.width
// }

// func main() {
// 	r1 := rect{width: 20, hiegt: 46}
// 	fmt.Println(r1.preim())

// 	r2 := newRect(0, 15)
// 	fmt.Println(r2.area())
// }

// package main

// import "fmt"

// func Add(x *int) {
// 	*x++
// }
// func main() {

// 	x := 10
// 	var mark *int
// 	mark = &x
// 	// Add(mark)
// 	fmt.Println(*mark)
// }

package main

import "fmt"

func main() {

	m := map[int]string{
		1: "Hashin",
		2: "Shanu",
		3: "Rishu",
	}

	value, ok := m[4]

	if ok {
		fmt.Println("Found :", value)
	} else {
		fmt.Println("Not found")
	}

}
