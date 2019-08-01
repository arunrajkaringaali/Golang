package main

import "fmt"

func main(){
	n, _ := fmt.Println("Introducing Usage of Functions!")
	fmt.Println(n)
	myfunc1()
	myfunc2()
	func_with_args("Arunraj")
	sp := func_with_args_ret("Koya")
	fmt.Println(sp)
	
	str,boo := func_with_args_multi_ret("Kriste","geenoy")
	fmt.Println(str)
	fmt.Println(boo)
	
	slice := []int{1,2,3,4,5}
	fmt.Println("Sum of 1,2,3,5 is ", variadic_func(1,2,3,4,5))
	fmt.Println("Sum of 1,2,3,5 is ", variadic_func()) //Zero number of args also works fine without errors
	fmt.Println("Sum of 1,2,3,5 is ", variadic_func(slice ...)) //unfurling_slice
	fmt.Println("Sum of 1,2,3,5 is ", variadic_func2("Arunraj")) 
	fmt.Println("Sum of 1,2,3,5 is ", variadic_func2("Arunraj",1,2,3,4,5)) 
	
	
	fmt.Println("Sum of 1,2,3,5 is ", unfurling_slice(slice ...))
	
	defer_ing() //defer
	anonemous_func()
	func_expression()//Func expression  - like function pointer in C++ 
	ret_func_example()
	incrementor_example()//Nice one!!
	recursion_example()
	
	
	excercise1()
	fmt.Println("Exiting from main!!")
}

func variadic_excercise(x ...int){
	for _,v := range x{
		fmt.Println(v)
	}
}
func excercise1(){
	variadic_excercise(1)
	variadic_excercise(1,2,3)
}

func myfunc1(){
	fmt.Println("In myfunc1")
}

func myfunc2(){
	fmt.Println("In myfunc2")
}
func func_with_args(s string){
	fmt.Println(s)
}
func func_with_args_ret(s string) string {
	return fmt.Sprint("Hello ",s)
}
func func_with_args_multi_ret(fn string, ln string) (string, bool) {
	return fmt.Sprint(fn, " " , ln, " Says ", `"hello!"`), true
}
func variadic_func(x ...int) int {
	sum := 0
	for _,v := range x {
		sum += v
	}
	
	return sum
}
func variadic_func2(s string, x ...int) int {
	sum := 0
	fmt.Println(s)
	for _,v := range x {
		sum += v
	}
	
	return sum
}
func unfurling_slice(x ...int) int {
	sum := 0
	for _,v := range x {
		sum += v
	}
	
	return sum
}

func foo(){
	fmt.Println("foo")
}

func bar(){
	fmt.Println("bar")
}
func defer_ing(){
	foo()
	bar()
	
	defer foo()
	bar()
	foo()
	bar()
	fmt.Println("Going to exit")
}

func anonemous_func() {
	
	func() {
		fmt.Println("Inside my Anonemous function")
	}()
}
func func_expression(){
	f := func(){
		fmt.Println("Inside my function expression")
	}
	
	f()
}

func return_func() func() string {
	return func() string {
		return "Hello"
	}
}

type convert func() string

func foo1 () string{
	return "World"
}

func ret_func_example(){
	fmt.Println(return_func()())
	fmt.Println(passing_func(foo1))
}


func passing_func( fn func() string) string{
	return fn()
}

func incrementor_example(){
	inc := incrementor()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
func factorial(x int) int {
	if(x==0){
		return 1
	}
	return x * factorial(x-1)
}
func recursion_example(){
	fmt.Printf("Factorial of %v is : %v\n",5,factorial(5))
}

