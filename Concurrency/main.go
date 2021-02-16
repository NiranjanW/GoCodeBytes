package main


import (
	_ "./Concurrency"
	"errors"
	"fmt"
	"log"
	"time"
)

type Greeter struct {
	Format string
}
func greeter(msg string) string {
	return ("%s Hello + $msg")
}

func NewGreeter (format string)(*Greeter ,error) {
	if len(format) == 0 {
		return nil ,errors.New("format required")
	}
	g := &Greeter{
		 format,
	}
	return g ,nil
}
//Greet mtd with a receiver ie : Fun that operates on the struc no this.
func (g *Greeter) Greet(name string) string {
	return fmt.Sprintf(g.Format , name)
}


func noTimes( inp string)  {
	dict :=make( map[string]int )

	for _,r := range inp{
		c := string(r)
		_ , exists :=  dict[c]
		if exists {
			dict[c] +=1
		}else {
			dict[c] = 1
		}
	}

	for k ,v := range dict {
		fmt.Printf("Key[%s] value[%d]" ,  k ,v)
	}
	//var result int = 0
	//for i:=0 ; i <= len(inp) ; i++ {
	//	for j:=1 ; j<= len(inp) -i ; j++ {
	//		if inp[i] == inp[j] {
	//			result ++
	//		}
	//	}
	//}
	//fmt.Printf(" no of time %d" , result)
}
func main(){

	greeter , err := NewGreeter("Hello There %s")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeter.Greet("Niran"))
	fmt.Println(greeter.Greet("Nayan"))
	resCh := make(chan bool)

	go func(ch chan<- bool) {
		time.Sleep(time.Second)
		ch <- true
	}(resCh)

	res := <-resCh
	fmt.Print(res)

	var s = "aaabbc"
	noTimes(s)
	for i , rune :=range "Hello"{
		fmt.Printf("%d  : %c\n" , i ,rune)
	}


}
