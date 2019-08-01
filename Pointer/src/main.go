package main

import ("fmt")

func main() {
	fmt.Println("Inside main()")
	address_example()
	struct_pointer()
	fmt.Println("Exit main()")
}

type human struct {
	name string
	age uint32
}

func foo(h *human){
	fmt.Println(h.name) 	//Direct . operator also works fine unlike C++
	fmt.Println((*h).age)	//This method also works fine like native C++
	  
}

func struct_pointer(){
	
	p1 := human{
		name 	: "Arunraj",
		age 	: 30,
	}
	foo(&p1)
}

func address_example(){
	x := 10
	fmt.Printf("Value of x is %v and its type is: %T\n",x,x)
	fmt.Printf("Address of x is %v and its type is %T\n",&x,&x)
	  
	px := &x
	fmt.Printf("Value of px is %v and its type is %T\n",px,px)
	fmt.Printf("Dereference of px is %v and its type is %T\n",*px,*px)
	
	*px = 20
	fmt.Printf("Value of x is %v and its type is: %T\n",x,x)
	fmt.Printf("Dereference of px is %v and its type is %T\n",*px,*px)
	
	

}
