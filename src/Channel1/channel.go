package main

import ( "fmt" )


func main() {
	
	basic()
	//deadlock()//fatal error: all goroutines are asleep - deadlock!
	sendRecieve()
	rangeChannel()
	selectusecase()
}

func basic(){
	var c chan int
	fmt.Println(c)
	fmt.Printf("Type of C : %T\n",c)
}

func deadlock(){
	var c chan int
	c <- 26
	fmt.Println(c)
	fmt.Printf("Type of C : %T\n",c)
}

func sendRecieve(){
	//var c chan int
	c := make(chan int, 0)
	go sendFoo(c)
	recieveFoo(c)
}

func sendFoo(c chan <- int){ //Send Channel
	c <- 12
}

func recieveFoo(c <- chan int){// Recieve Channel
	fmt.Println(<- c)
}

func rangeChannel(){
	c := make(chan int)
	go sendFoo1(c)
	recieveFoo1(c)
}
func sendFoo1(c chan <- int){ //Send Channel
	for i:= 0; i<5; i++ {
		c <- i
	}
	close(c)
}

func recieveFoo1(c <- chan int){// Recieve Channel
	for v:= range c{
		fmt.Println(v)
	}
}

func selectusecase(){
	c := make(chan int,0)
	q := make(chan bool)
	
	go send_data(c,q)
	
	v:= recieve_data(c,q)
	
	fmt.Println(v)
}

func send_data(c chan<- int, q chan<- bool){
	for i:=0; i<5; i++{
		c <- i
	}
	close(c)
	q <- true
}

func recieve_data(c <-chan int, q <-chan bool) []int {
	data := make([]int,0)
	
	for {
		select {
			case v := <- c:
				data = append(data, v)
				fmt.Println(v)
			case v := <- q:
				if v == true {
					return data
				} 
		}
	}
}
