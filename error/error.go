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
// //	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printEven(oddChan, evenChan chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 2; i <= 10; i += 2 {
// 		<-evenChan
// 		fmt.Println("Even Number:", i)
// 		oddChan <- true
// 	}
// }

// func printOdd(oddChan, evenChan chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 9; i += 2 {
// 		<-oddChan
// 		fmt.Println("Odd:", i)
// 		evenChan <- true
// 	}
// }

// func main() {
// 	evenChan := make(chan bool)
// 	oddChan := make(chan bool)

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go printEven(oddChan, evenChan, &wg)
// 	go printOdd(oddChan, evenChan, &wg)

// 	oddChan <- true // start with odd

// 	wg.Wait()
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func oddNum(evenChan, oddChan chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 19; i += 2 {
// 		<-oddChan
// 		fmt.Println("Odd Number: ", i)
// 		evenChan <- true
// 	}
// }

// func evenNum(oddChan, evenChan chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 2; i <= 20; i += 2 {
// 		<-evenChan
// 		fmt.Println("Even Number: ", i)
// 		if i != 20 {
// 			oddChan <- true
// 		}

// 	}
// }

// func main() {
// 	evenChan := make(chan bool)
// 	oddChan := make(chan bool)
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go oddNum(evenChan, oddChan, &wg)
// 	go evenNum(oddChan, evenChan, &wg)
// 	oddChan <- true
// 	wg.Wait()
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func Producer(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 10; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

// func Consumer(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := range ch {
// 		fmt.Println("Proceccing :", i)
// 	}
// }

// func main() {
// 	ch := make(chan int)
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go Producer(ch, &wg)
// 	go Consumer(ch, &wg)
// 	wg.Wait()
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func Worker(id int, jobs chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for job := range jobs {
// 		fmt.Printf("Worker %d proceccing %d\n", id, job)
// 	}
// }

// func main() {
// 	jobs := make(chan int)
// 	var wg sync.WaitGroup

//		for i := 1; i <= 3; i++ {
//			wg.Add(1)
//			go Worker(i, jobs, &wg)
//		}
//		go func() {
//			for i := 1; i <= 15; i++ {
//				jobs <- i
//			}
//			close(jobs)
//		}()
//		wg.Wait()
// //	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	result := make(chan string)
// 	cancel := make(chan bool)

// 	go func() {
// 		select {
// 		case <-time.After(3 * time.Second):
// 			result <- "Job Done"
// 		case <-cancel:
// 			fmt.Println("Worker cancelled")
// 			return
// 		}
// 	}()

// 	select {
// 	case res := <-result:
// 		fmt.Println(res)
// 	case <-time.After(2 * time.Second):
// 		fmt.Println("Timeout!")
// 		cancel <- true
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func CheckAge(age int) error {
// 	if age < 18 {
// 		return fmt.Errorf("Age must %d be 18 above", age)
// 	} else {
// 		fmt.Println("Selected")
// 	}
// 	return nil
// }

// func main() {
// 	err := CheckAge(19)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func CheckAge(age int) error {
// 	if age < 18 {
// 		return errors.New("Age must be 18 above")
// 	} else {
// 		fmt.Println("Selected")
// 	}
// 	return nil
// }

// func main() {
// 	err := CheckAge(17)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// package main

// import "fmt"

// type Person struct {
// 	Nmae string
// 	Age  int
// }
// func Greet(p Person) {
// 	fmt.Println("Name :", p.Nmae)
// 	fmt.Println("Age :", p.Age)
// }
// func newPerson(Name string, Age int) *Person {
// 	return &Person{
// 		Nmae: Name,
// 		Age:  Age,
// 	}
// }
// func main() {
// 	Greet(*newPerson("HAshin", 21))
// }

// package main

// import "fmt"

// func main() {

// 	var num1, num2 float64
// 	var operator string

// 	fmt.Println("-----Calculator-----")
// 	fmt.Print("Enter First Number")
// 	fmt.Scan(&num1)

// 	fmt.Print("Enter Operator")
// 	fmt.Scan(&operator)

// 	fmt.Print("Enter Second Number")
// 	fmt.Scan(&num2)

// 	switch operator {
// 	case "+":
// 		fmt.Println("Result :", num1+num2)
// 	case "-":
// 		fmt.Println("Result :", num1-num2)
// 	default:
// 		fmt.Println("Inavalid")
// 	}

// // }

// package main

// import "fmt"

// func test() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered:", r)
// 		}
// 	}()

// 	panic("Error!")
// }

// func main() {
// 	fmt.Println("Start")

// 	test()
// 	arr := []int{1, 2, 3, 4, 5, 6}
// 	abc := arr[1:4]
// 	// arr = append(arr[:ind], arr[ind+1:]...)
// 	fmt.Println(abc)
// }

// package main

// import "fmt"

// func Greet() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovred :", r)
// 		}
// 	}()
// 	panic("Program crashed")
// }

// func main() {
// 	fmt.Println("Hello")
// 	Greet()
// 	fmt.Println("End")
// }

// package main

// import "fmt"

// func modifyarry(arr *[3]int) {
// 	arr[2] = 7
// }
// func main() {
// 	a := [3]int{1, 2, 3}
// 	modifyarry(&a)
// 	fmt.Println(a)
// }

// package main

// import "fmt"

// func main() {
// 	ch1 := make(chan int, 3)
// 	ch1 <- 10
// 	ch1 <- 20
// 	b := <-ch1
// 	ch1 <- 30
// 	ch1 <- 40
// 	fmt.Println(b)
// 	fmt.Println(<-ch1, <-ch1, <-ch1)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	ch := make(chan int, 2)
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go func() {
// 		ch <- 31
// 		ch <- 23
// 		ch <- 34
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println(<-ch)
// 		fmt.Println(<-ch)
// 		fmt.Println(<-ch)
// 	}()
// 	wg.Wait()
// // }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printOdd(evenChan, oddChan chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 20; i += 2 {
// 		<-oddChan
// 		fmt.Println("Odd Number :", i)
// 		evenChan <- true
// 	}
// }

// func printEven(evenChan, oddChan chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 2; i <= 20; i += 2 {
// 		<-evenChan
// 		fmt.Println("Even number :", i)
// 		if i != 20 {
// 			oddChan <- true
// 		}
// 	}
// }

// func main() {
// 	evenChan := make(chan bool)
// 	oddChan := make(chan bool)
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go printEven(evenChan, oddChan, &wg)
// 	go printOdd(evenChan, oddChan, &wg)
// 	oddChan <- true
// 	wg.Wait()
// }

// package main

// import "fmt"

// type rule interface {
// 	Area() float64
// 	Perim() float64
// }

// type Rectangle struct {
// 	hieght, width float64
// }

// func (r *Rectangle) Area() float64 {
// 	return r.hieght * r.width
// }

// func (r Rectangle) Perim() float64 {
// 	return r.hieght*2 + r.width*2
// }

// func main() {
// 	a1 := Rectangle{width: 20, hieght: 30}
// 	fmt.Println(a1.Area())
// 	fmt.Println(a1.Perim())
// }

// package main

// import "fmt"

// func printArray(n ...int) {
// 	fmt.Println("Num :", n)
// }

// func main() {
// 	printArray(1, 2, 3, 4, 5)
// }

// package main

// import "fmt"

// type Student struct {
// 	Name string
// 	age  int
// }

// type Worker interface {
// 	work()
// }

// func (s *Student) work(newName string) {
// 	s.Name = newName
// 	fmt.Println(newName, s.age)
// }

// func rules(w *Worker)  {

// }

// func main() {
// 	s1 := Student{Name: "Hashin", age: 21}
// 	s1.work("Shanu")
// }

// package main

// import "fmt"

// func main() {
// 	m1 := map[int]string{
// 		1: "Hashin",
// 		2: "nabeel",
// 		3: "Manhar",
// 	}

// 	value, ok := m1[3]
// 	if ok {
// 		fmt.Println(value)
// 	}

// }

// package main

// import "fmt"

// func removeDuplicates(arr []int) []int {
// 	m1 := make(map[int]bool)
// 	result := []int{}

// 	for _, v := range arr {
// 		if !m1[v] {
// 			m1[v] = true
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }

// func main() {
// 	arr := []int{1, 2, 3, 4, 1, 2, 6, 3, 7}
// 	fmt.Println(arr)
// 	fmt.Println(removeDuplicates(arr))
// }

// package main

// import "fmt"

// type Student struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	m1 := make(map[int]Student)

// 	m1[1] = Student{
// 		Name: "Hashin",
// 		Age:  21,
// 	}

// 	fmt.Println(m1[1].Name)
// }

// package main

// import "fmt"

// func printNum(n ...int) int {
// 	fmt.Println(n)
// 	return 8
// }

// func main() {
// 	n1 := printNum(1, 2, 3, 4, 5)
// 	fmt.Println(n1)

// }
// package main

// import "fmt"

// func changeNum(num *int) {
// 	*num = 20
// 	fmt.Println(*num)
// }

// func main() {
// 	value := 1
// 	changeNum(&value)
// 	// fmt.Println(value)
// }

// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7}
// 	arr = append(arr[:4], arr[4+2:]...)
// 	fmt.Println(arr)
// 	arr = append(arr[:0], arr[0+2:]...)
// 	fmt.Println(arr)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func oddNum(oddChan, evenChan chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 20; i += 2 {
// 		<-oddChan
// 		fmt.Println("oddnum :", i)
// 		evenChan <- true
// 	}

// }

// func evenNum(evenChan, oddChan chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 2; i <= 20; i += 2 {
// 		<-evenChan
// 		fmt.Println("evenNum :", i)

// 		if i != 20 {
// 			oddChan <- true
// 		}
// 	}
// }

// func main() {
// 	oddChan := make(chan bool)
// 	evenChan := make(chan bool)
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go oddNum(oddChan, evenChan, &wg)
// 	go evenNum(evenChan, oddChan, &wg)
// 	oddChan <- true
// 	wg.Wait()
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var counter int
// var wg sync.WaitGroup
// var mu sync.Mutex

// func increment() {
// 	defer wg.Done()

// 	for i := 0; i <= 20; i++ {
// 		mu.Lock()
// 		counter++
// 		mu.Unlock()
// 	}
// 	fmt.Println(counter)
// }

// func main() {
// 	wg.Add(2)
// 	go increment()
// 	go increment()
// 	wg.Wait()
// 	fmt.Println("Final counter :", counter)

// }
// }
// package main

// import (
// 	"fmt"
// )

// func elon(msg string) error {
// 	return fmt.Errorf("Error: %s", msg)
// }

// func main() {
// 	err := elon("Something went wrong")

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

//lln package handle go rutine ,, fan in fan out gorutine package, workerfool,