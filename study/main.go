//Reverse a String
// package main
// import "fmt"
// func main() {
// 	str:="hello world"
// 	r:=[]rune(str)
// 	f:=0
// 	l:=len(str)-1
// 	for f<l{
// 		r[f],r[l]=r[l],r[f]
// 		f++
// 		l--
// 	}
// 	fmt.Println(string(r))
// }

//Count Vowels
// package main
// import "fmt"
// func main() {
// 	str:="hello world"
// 	v:=0
// 	c:=0
// 	for i:=0;i<len(str);i++{
// 		ch:=str[i]
// 		if ch=='a'||ch=='A'||ch=='e'||ch=='E'||ch=='i'||ch=='I'||ch=='o'||ch=='O'||ch=='u'||ch=='U'{
// 			v++
// 		}else{
// 			c++
// 		}
// 	}
// 	fmt.Println("vowels -",v)
// 	fmt.Println("Consonants -",c)
// }

//find larget number
package main
import "fmt"
func main() {
	arr:= []int{10, 20, 4, 45, 99}
	max:=0
	min:=0
	for i:=0;i<len(arr);i++{
		if arr[i]>max{
			max=arr[i]
		}else{
			min=arr[i]
		}
	}
	fmt.Println(max,"-",min)
}

//Check Palindrome
// package main
// import "fmt"
// func main() {
// 	p:=true
// 	str:="malayalam"
// 	r:=[]rune(str)
// 	f:=0
// 	l:=len(str)-1
// 	for f<l{
// 		if r[f]!=r[l]{
// 			p=false
// 			break
// 		}
// 		f++
// 		 l--
// 	}
// 	fmt.Println(p)
// }

//Remove Duplicates from Slice
// package main
// import "fmt"
// func main() {
// 	arr:=[]int{1,2,3,12,3,1,1,3,53,52,2,5,53,100}
// 	m:=make(map[int]int)
// 	plane:=[]int{}
// 	for _,val:=range arr{
// 		m[val]++
// 		if m[val]==1{
// 			plane=append(plane, val)
// 		}
// 	}
// 	og:=[]int{}
// 	dupe:=[]int{}
// 	for k,c:=range m{
// 		if c>1{
// 			dupe = append(dupe, k)
// 		}else{
// 			og = append(og, k)
// 		}
// 	}
// 	fmt.Println("Array ",arr)
// 	fmt.Println("removed spam",plane)
// 	fmt.Println("duplicates",dupe)
// 	fmt.Println("not duplicates",og)
// }

// repeating words
// package main
// import (
// 	"fmt"
// 	"strings"
// )
// func main() {
// 	str := `then I asked him with my eyes to ask again yes and then he
// 	        asked me would I yes to say yes my mountain flower and first
// 	        I put my arms around him yes and drew him down to me so he
// 	        could feel my breasts all perfume yes and his heart was
// 	        going like mad and yes I said yes I will Yes`
// 	words:=strings.Fields(str)
// 	m:=make(map[string]int)
// 	for _,v:=range words{
// 		m[v]++
// 	}
// 	fmt.Println(m)
// }

//average using struct
// package main
// import "fmt"
// type Student struct{
// 	Name string
// 	Score []int
// }
// func (s Student)Grade()float64{
// 	if len(s.Score)==0{
// 		return 0
// 	}
// 	sum:=0
// 	for _,val:=range s.Score{
// 		sum+=val
// 	}
// 	return float64(sum)/float64(len(s.Score))
// }
// func main() {
// 	st:=Student{
// 		Name: "Shabin",
// 		Score: []int{10,20,30,40,50},
// 	}
// 	fmt.Println(st.Name)
// 	fmt.Println(st.Grade())
// }

//pointer
// package main
// import "fmt"
// func trial(n *int) {
// 	*n= *n+*n
// }
// func update(x **int) {
//     n := 100
// 	*x=&n
// }
// func main() {
// 	num:=20
// 	ptr:=&num
// 	update(&ptr)
// 	fmt.Println(*ptr)
// }

// package main
// import (
// 	"fmt"
// 	"net/http"
// 	"io"
// )	
// func main() {
// 	resp,err:=http.Get("https://jsonplaceholder.typicode.com/posts/1")
// 	if err!=nil{
// 		fmt.Println("Error",err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	body,err:=io.ReadAll(resp.Body)
// 	if err!=nil{
// 		fmt.Println("err")
// 		return
// 	}
// 	fmt.Println(string(body))
// }
 
// package main
// import (
// 	"fmt"
// 	"sync"
// )
// var count int
// var mu sync.Mutex
// func trial(wg *sync.WaitGroup){
// 	defer wg.Done()
// 	for i:=0;i<1000;i++{
// 		mu.Lock()
// 		count++
// 		mu.Unlock()
// 	}
// }
// func main() {
// 	var wg sync.WaitGroup
// 	for i:=0;i<5;i++{
// 		wg.Add(1)
// 		go trial(&wg)
// 	}
// 	wg.Wait()
// 	fmt.Println(count)
// }	