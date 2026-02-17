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

// package main

// import "fmt"

// type Empolyee interface {
// 	CalCulateSalary() float32
// 	GetName() string
// }

// type FullTimeEmpolyee struct {
// 	Name          string
// 	MonthlySalary float32
// }

// type PartTimeEmployee struct {
// 	Name         string
// 	HourlyRate   float32
// 	HouresWorked float32
// }

// func (F *FullTimeEmpolyee) CalCulateSalary() float32 {
// 	return F.MonthlySalary * 12
// }
// func (N FullTimeEmpolyee) GetName() string {
// 	return N.Name
// }
// func (P *PartTimeEmployee) CalCulateSalary() float32 {
// 	return P.HouresWorked * P.HourlyRate
// }
// func (N PartTimeEmployee) GetName() string {
// 	return N.Name
// }

// func PrintSalry(P Empolyee) {
// 	fmt.Println("Name " + P.GetName())
// 	fmt.Println("Salary :", P.CalCulateSalary())
// 	fmt.Println("-------------------")
// }

// func main() {
// 	p1 := FullTimeEmpolyee{Name: "Sajad", MonthlySalary: 800}
// 	p2 := PartTimeEmployee{Name: "Hashin", HouresWorked: 6, HourlyRate: 550}

// 	PrintSalry(&p1)
// 	PrintSalry(&p2)

// }

// package main

// import "fmt"

// type Student struct {
// 	ID   int
// 	Name string
// 	Mark int
// }

// func UpdateStudent(s *Student, newName string, newMark int) {
// 	s.Name = newName
// 	s.Mark = newMark
// }

// func DeletedStudent(student []Student, id int) []Student {
// 	for i, s := range student {
// 		if s.ID == id {
// 			return append(student[:i], student[i+1:]...)
// 		}
// 	}
// 	return student
// }

// func main() {
// 	student := []Student{
// 		{ID: 1, Name: "Hashin", Mark: 100},
// 		{ID: 2, Name: "salman", Mark: 96},
// 		{ID: 3, Name: "fakru", Mark: 100},
// 	}
// 	UpdateStudent(&student[0], "Shanu", 98)
// 	fmt.Println(student)
// 	student = DeletedStudent(student, 3)
// 	fmt.Println(student)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printNumber(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i)
// 	}
// }

// func evenNumber(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i <= 20; i += 2 {
// 		if i%2 == 0 {
// 			fmt.Println("Even Num :", i)
// 		}
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go evenNumber(&wg)
// 	go printNumber(&wg)
// 	wg.Wait()
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var count int
// var mu sync.Mutex

// func increment(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 100; i++ {
// 		mu.Lock()
// 		count++
// 		mu.Unlock()
// 	}

// }
// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(4)
// 	go increment(&wg)
// 	go increment(&wg)
// 	go increment(&wg)
// 	go increment(&wg)

// 	wg.Wait()
// 	fmt.Println("Conter :", count)

// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printEven(oddChan, evenChan chan bool, wg *sync.WaitGroup) {

// 	defer wg.Done()
// 	for i := 2; i <= 20; i += 2 {
// 		<-evenChan
// 		fmt.Println("Even :", i)
// 		oddChan <- true
// 	}
// }
// func Oddnum(evenChan, oddChan chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 20; i += 2 {
// 		<-oddChan
// 		fmt.Println("Even :", i)
// 		evenChan <- true
// 	}

// }

// func main() {
// 	oddChan := make(chan bool)
// 	evenChan := make(chan bool)

// 	var wg sync.WaitGroup

//		wg.Add(2)
//		go printEven(oddChan, evenChan, &wg)
//		go Oddnum(oddChan, evenChan, &wg)
//		evenChan <- true
//		wg.Wait()
//	}
package main

import (
	"fmt"
	"sync"
)

func printEven(oddChan, evenChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-evenChan
		fmt.Println("Even Number:", i)
		oddChan <- true
	}
}

func printOdd(oddChan, evenChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 9; i += 2 {
		<-oddChan
		fmt.Println("Odd:", i)
		evenChan <- true
	}
}

func main() {
	evenChan := make(chan bool)
	oddChan := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go printEven(oddChan, evenChan, &wg)
	go printOdd(oddChan, evenChan, &wg)

	oddChan <- true // start with odd

	wg.Wait()
}
