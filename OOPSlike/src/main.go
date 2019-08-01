package main

import (
	"fmt"
)

func main() {
	methods() //Methods - like Class methods in C++
	why_interfaces()
	interface_example()

}

///////// Human STructure ///////////
type doctor struct {
		fn string
		ln string
	}
	
func (d doctor) speak(){
	fmt.Println(d.fn, d.ln, "is speaking!")
}

///////// Person STructure ///////////
type engineer struct {
		fn string
		ln string
	}
	
func (e engineer) speak(){
	fmt.Println(e.fn, e.ln, "is speaking!")
}

func methods(){  
	d1 := doctor {
		fn : "Arunraj",
		ln : "Karingali",
	}
	d1.speak()
	
	e1 := engineer {
		fn : "Kora",
		ln : "Kuruppu",
	}
	e1.speak()
}


func foo(i doctor) {
	fmt.Println("Inside foo : ", i)
}



func why_interfaces(){
	d1 := doctor {
		fn : "Arunraj",
		ln : "Karingali",
	}
	d1.speak()
	
	e1 := engineer {
		fn : "Kora",
		ln : "Kuruppu",
	}
	e1.speak()
	foo(d1)
	//foo(e1) //This gives error! because function have arg with type doctor. 
			  //thats the reason we need interfaces 
			  //to pass any type to a single interface 
}

type human interface {
	speak() //Tells if you have this methode you are my type
}

func bar(h human){
	fmt.Println("\nInside foo : ", h)
	h.speak()
	fmt.Printf("Type of h is %T\n\n",h)
}

func interface_example(){
	d1 := doctor {
		fn : "Arunraj",
		ln : "Karingali",
	}
	d1.speak()
	
	e1 := engineer {
		fn : "Kora",
		ln : "Kuruppu",
	}
	e1.speak()
	bar(d1)
	bar(e1) //This gives error! because 
}

