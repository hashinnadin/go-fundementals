// package main

// import "fmt"

// func main() {

// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			fmt.Println("Recovered :", r)

// 			switch v := r.(type) {
// 			case string:
// 				fmt.Println("Type is string ", v)
// 			case int:
// 				fmt.Println("Type is int", v)
// 			default:
// 				fmt.Println("Inavalid error", v)
// 			}
// 		}
// 	}()
// 	panic(404)

// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {

// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered :", r)

// 			if err, ok := r.(error); ok {
// 				fmt.Println("It was an error:", err.Error())
// 			}
// 		}
// 	}()
// 	panic(errors.New("database connections failed "))
// }

// package main

// import "fmt"

// type Rectangle struct {
// 	width, hieght int
// }

// type Shape interface{
// 	Area() int
// }

// func (r Shape)  {

// }
// func (r Rectangle) Area() int {
// 	return r.hieght * r.width
// }

// func (r Rectangle) Perim() int {
// 	return r.hieght + r.width*2
// }
// func newRect(width int, hieght int) *Rectangle {
// 	return &Rectangle{
// 		width:  width,
// 		hieght: hieght,
// 	}
// }

// func main() {

// 	r := Rectangle{width: 30, hieght: 20}

// 	fmt.Println("Area :", r.Area())
// 	fmt.Println("Perim :", r.Perim())
// 	r2 := newRect(20, 30)
// 	fmt.Println(r2.Perim())
// }

// package main

// import "fmt"

// type Rectangle struct {
// 	width, hieght int
// }

// type Circle struct {
// 	Radius int
// }
// type Square struct {
// 	Value int
// }

// type Shape interface {
// 	area() int
// }

// type Worker interface {
// 	Work()
// }

// type Teacher struct {
// 	Value string
// }
// type Principal struct {
// 	Value string
// }

// func (T Teacher) Work() {
// 	fmt.Println("He Was Not Working " + T.Value)
// }

// func (P Principal) Work() {
// 	fmt.Println("He Was Not Working " + P.Value)
// }
// func (s Square) area() int {
// 	return s.Value * s.Value
// }
// func (r Rectangle) area() int {
// 	return r.hieght * r.width
// }
// func (c Circle) area() int {
// 	return 3 * c.Radius * c.Radius
// }

// func printOur(s Shape) {
// 	fmt.Println(s.area())
// }
// func printWorker(W Worker) {
// 	W.Work()
// }

// func main() {
// 	r1 := Rectangle{hieght: 20, width: 15}
// 	printOur(r1)
// 	c1 := Circle{Radius: 20}
// 	printOur(c1)
// 	s1 := Square{Value: 10}
// 	printOur(s1)
// 	p1 := Teacher{Value: "Silpa"}
// 	p2 := Principal{Value: "velayudan"}
// 	printWorker(p1)
// 	printWorker(p2)
// }

