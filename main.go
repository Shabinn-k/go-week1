// package main

// import "fmt"

// func main() {
// 	// arr := []int{1, 2, 2, 3, 4, 4, 5, 6, 7, 7}
// 	// seen := make(map[int]int)
// 	// duplicate := []int{}
// 	// for _, value := range arr {
// 	// 	seen[value]++
// 	// 	if seen[value] == 2 {
// 	// 		duplicate = append(duplicate, value)
// 	// 	}
// 	// }
// 	// fmt.Println(arr)
// 	// fmt.Println(duplicate)
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)
// 	ch1 <-"shabin"
// 	select {
// 	case msg1 := <-ch1:
// 		fmt.Println("Received from ch1:", msg1)

// 	case msg2 := <-ch2:
// 		fmt.Println("Received from ch2:", msg2)

// 	default:
// 		fmt.Println("No channel ready")
// 	}
// }

// package main
// import (
// 	"fmt"
// 	"sync"
// )
// var mu sync.Mutex

// var count int

// func add(wg *sync.WaitGroup){
// 	defer wg.Done()
// 	for i:=0;i<1000;i++{
// 		mu.Lock()
// 		count++
// 		mu.Unlock()
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go add(&wg)
// 	}

// 	wg.Wait()
// 		fmt.Println(count)
// }

// package main

// import (
// 	"fmt"
// )

// func sendData(ch chan string) {
// 	ch <- "Hello from Goroutine!"
// }

// func main() {
// 	ch := make(chan string )

// 	go sendData(ch)

// 	message := <-ch

// 	fmt.Println(message)
// }

// package main
// import "fmt"
// type Student struct {
// 	Name string
// 	Age  int
// 	Mark int
// }
// var s1=Student{
// 	Name: "Shabin",
// 	Age:  18,
// 	Mark: 95,
// }
// func (s Student) greet() {

// 	fmt.Println("Hello,", s.Name,s.Age)
// }

// func main() {
// 	s1.greet()
// }
// package main
// import "fmt"

// func greet() {
// 	fmt.Println("Hello")
// }

// func main() {
// 	go greet()
// 	fmt.Println("Main function")
// }

// package main

// import "fmt"

// type Shape interface{
// 	Area() int
// }

// type A struct{
// 	Width int
// 	Height int
// }

// func (a A)Area()int{
// 	return a.Width*a.Height
// }
// func hs(s Shape){
// 	fmt.Println(s.Area())
// }
// func main() {
// 	d:=A{Width: 10,Height: 100}
// 	hs(d)
// }

// package main

// import (
// "errors"
// "fmt"
// )
// func hy(a int)error{
//  if a%2==0{
// 	return  errors.New("it is even")
//  }
//  return nil
// }
// func risk(){
// 	defer func(){
// 		if r:=recover();r!=nil{
// 			fmt.Println("recovered from",r)
// 		}
// 	}()
// 	panic("panic occured")
// }
// func main() {
// err:=hy(10)
// if err!=nil{
// 	fmt.Println(err)
// }
// risk()
// }

// package main

// import "fmt"

// type Shape interface {
// 	Area() int
// }

// type Sq struct {
// 	Width  int
// 	Height int
// }

// func (s Sq) Area() int {
// 	return s.Width * s.Height
// }

// func S(s Shape) {
// 	fmt.Println("Area:", s.Area())
// }

// func main() {
// 	square := Sq{
// 		Width:  5,
// 		Height: 5,
// 	}

// 	S(square)
// }
// func incr(ch chan int){
// for i:=0;i<10;i++{
// 	if i%2==0{
// 		ch<- i
// 	}
// }
// close(ch)
// }

// package main

// import( "fmt"
// "time")
// var count int=0
// func print(){

//  for i:=0;i<1000;i++{
// 	count++
// }
// }

// func main() {

// for i:=0;i<5;i++{

// 	go print()
// }
// time.Sleep(time.Second)
// fmt.Println(count)
// }
// package main

// import "fmt"

// func update(num *int) {
// 	*num = 10
// }

// func main() {
// 	x := 5
// 	fmt.Println("Before:", x)

// 	update(&x)

// fmt.Println("After:", x)
// 	arr:=[]int{1,23,4,5,67,1,1,1,3,4,44,4,4,5,5}
// 	m:=make(map[int]int)
// 	duplicate:=[]int{}
// 	for _,v:=range arr{
// 		m[v]++
// 		if m[v]==2{
// 			duplicate=append(duplicate,v)
// 		}
// 	}
// 	fmt.Println(duplicate)
// }

// package main
// import "fmt"
// func main() {
// 		s:=[]int{1,2,4,5,6}
// 		i:=2
// 		v:=3
// 		s=append(s[:i],append([]int{v},s[i:]...)...)
// 		s=append(s[:4],s[6:]...)
// 		fmt.Println(s)
// }

// package main
// import "fmt"
// func main() {
// 	var a int
// 	var b int
// 	var op string
// 	fmt.Println("enter 2 number")
// 	fmt.Scan(&a,&b)
// 	fmt.Println("enter a sign")
// 	fmt.Scan(&op)
// 	switch op{
// 	case "+":fmt.Println(a+b)
// 	case "*":fmt.Println(a*b)
// 	default:fmt.Println("over")
// 	}
// }

// package main
// import "fmt"
// type Student struct{
// 	Name string
// }
// func (s Student)Trial() {
// 	fmt.Println(s.Name)
// }
// func main(){
// 	n:=Student{
// 		Name: "shabin",
// 	}
// 	n.Trial()
// }

// package main

// import (
// 	"fmt"
// 	 "golang/Practice" 
// )
// func main() { 
// 	m:=map [int]string{
// 		1:"shabin",
// 	}
// 	v,ok:=m[2]
// 	if ok{fmt.Println(v)}else{fmt.Println("not found")}
// 	num:=shabin.Add(10,20)
// 	fmt.Println(num)

// }

package main
import f "fmt"
func main() {
	x:=10
	p:=&x
	*p=40
	f.Println(x)
}