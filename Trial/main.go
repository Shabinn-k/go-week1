// package main

// import "fmt"

// type Student struct {
// 	Name  string
// 	Score []int
// }

// func (s Student) Get() float64 {
// 	if len(s.Score) == 0 {
// 		return 0
// 	}
// 	sum := 0
// 	for _, v := range s.Score {
// 		sum += v
// 	}
// 	return float64(len(s.Score)) / float64(sum)
// }
// func main() {
// 	marks := Student{
// 		Name:  "Shabin",
// 		Score: []int{10, 20, 30, 40, 50},
// 	}
// 	fmt.Println("Average",marks.Get())
// }

// package main
// import (
// 	"fmt"
// 	"strings"
// )
// func main() {
// 	text := "go is fun and go is fast"
// 	words:=strings.Fields(text)
// 	m:=make(map[string]int)
// 	for _,v:=range words{
// 		m[v]++
// 	}
// 	fmt.Println(m)
// }

// package main
// import "fmt"
// func topper(m map[string]int){
// 	name:=""
// 	mark:=0
// 	for i,v:=range m{
// 		if v>mark{
// 			mark=v
// 			name=i
// 		}
// 	}
// 	fmt.Println(name,mark)
// }
// func main() {
// 	m:=map[string]int{
// 		"shabin":20,
// 		"rahul":10,
// 		"messi":100,
// 		"modi":9,
// 	}
// 	delete(m,"modi")
// 	fmt.Println(m)
// 	v,ok:=m["messi"]
// 	if ok{fmt.Println(v)}else{fmt.Println("not found")}
// 	topper(m)
// }

// package main
// import "fmt"
// func main() {
// 	arr:=[]int{1,2,2,3,4,4,5}
// 	og:=[]int{}
// 	m:=make(map[int]int)
// 	for _,v:=range arr{
// 		m[v]++
// 		if m[v]!=2{
// 			og=append(og, v)
// 		}
// 	}
// 	fmt.Println(og)
// }

// package main
// import "fmt"
// func main() {
// 	arr:=[]int{5,4,3,2,1}
// 	f:=0
// 	l:=len(arr)-1
// 	for f<l{
//    arr[f],arr[l]=arr[l],arr[f]
//    f++
//    l--
// 	}
// 	arr = append(arr[:2],append([]int{6},arr[2:]...)...)
// 	fmt.Println(arr)
// }

package main

import "fmt"
func trial(a,b *int){
	*a=100
	*b=200	
}
func main() {
	x:=10
	y:=20
	trial(&x,&y) 
	fmt.Println(x,y)
}