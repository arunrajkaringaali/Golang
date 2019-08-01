package  main

import "fmt"

var myGlobVar = "ArunRaj"
var mystrGlobVar string
var myintGlobVar int
//mynewVar := 2 //syntax error: non-declaration statement outside function body

type arun int
var myNewVar arun //Making my own data type variable

func main(){
	fmt.Println("I am in main()")
	
	//basics()
	//exercise1()
	//exercise2()
	//exercise3()
	//exercise4()
	//exercise5()
	//exercise6()
	//arrays()
	//map_()
	//arrays_excercise()
	//string_excercise()
	//struct_excercise()
	vehicle_example()
}

func basics(){
	
	myLocVar := 2	//Works fine only if its used atleast once
	var myLocVar1 = 2 //Works fine only if its used atleast once
	//mynewLocVar = 4 //undefined: mynewLocVar
	
	fmt.Println(myGlobVar)
	fmt.Printf("Type of myGlobVar: %T\n",myGlobVar)	
	fmt.Printf("Type of mystrGlobVar: %T\n",mystrGlobVar)	
	fmt.Printf("Type of myintGlobVar: %T\n",myintGlobVar)
	
	//myLocVar = "Hello" //cannot use "Hello" (type string) as type int in assignment
	
	fmt.Println(myLocVar)
	fmt.Printf("Type of myLocVar: %T\n",myLocVar)
	fmt.Printf("Type of myLocVar1: %T\n",myLocVar1)
	
	fmt.Printf("Type of myNewVar: %V\n",myNewVar)
	fmt.Println(myNewVar)
	fmt.Printf("Type of myNewVar: %T\n",myNewVar)
	
	//myNewVar = myLocVar //cannot use myLocVar (type int) as type arun in assignment
	myNewVar = arun(myLocVar) //Go TypeCasting This is possible !!
	
}
func exercise1(){
	fmt.Println("\nexercise1")
	x := 42
	y := "James Bond"
	z := true
	
	fmt.Printf("%V, %V, %V\n",x,y,z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
func exercise2(){
	fmt.Println("\nexercise2")
	var x int
	var y string
	var z bool
	
	fmt.Printf("%V, %V, %V\n",x,y,z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
func exercise3(){
	fmt.Println("\nexercise3")
	var x int
	var y string
	var z bool
	
	strr := fmt.Sprintln(x,y,z)
	strr = fmt.Sprintf("%v %v %v\n",x,y,z)
	fmt.Println(strr)
}
func exercise4(){
	fmt.Println("\nexercise4")
	type arun int
	type raj string
	
	
	var OwnDataType arun = 35
	fmt.Println(OwnDataType)
	fmt.Printf("Data type of OwnDataType is : %T\n",OwnDataType)
	
	var OwnDataType2 arun = OwnDataType
	fmt.Println(OwnDataType2)
	fmt.Printf("Data type of OwnDataType2 is : %T\n",OwnDataType2)
	
	var OwnDataType3 raj = "ArunRaj"
	fmt.Println(OwnDataType3)
	
	//OwnDataType = arun(OwnDataType3)
	//fmt.Println(OwnDataType3) //cannot convert OwnDataType3 (type raj) to type arun
	
	var newwInt int
	//newwInt = int(OwnDataType3)// Not possible from a conversion from a data type which is not derived from same primitive type
	newwInt = int(OwnDataType)
	fmt.Println(newwInt)
	
}

func exercise5(){
	fmt.Println("\nexercise5")
	s := "Hello world"
	
	//C Style
	for i := 0; i< len(s); i++{
		fmt.Println(s[i])
	}
	//Python Style
	for i, v := range(s){
		fmt.Printf("%d   have %#x\n",i,v)
		//	fmt.Println(i,v)
	}
	
}

func exercise6(){
	fmt.Println("\nExcercise6")
	for i :=0X41; i< (0X41+26) ; i++ {
		fmt.Printf("%#U\n",i)
		if i>=0x41 {
			for j:=0; j<2; j++ {
				fmt.Printf("%#U\n",i)
			}
		}
		fmt.Println("\n")
	} 
	
	i:= 0
	
	for i<10 {
		fmt.Printf("%v ",i)
		i++
	}
	
	fmt.Println("")
	i = 0
	for {
		fmt.Printf("%v ",i)
		if i>=10 {
			break
		}
		i++
	}
	fmt.Println("")
}

func arrays(){
	fmt.Println("\nArrays")
	
	var arr1 [5]int
	fmt.Println(arr1)
	fmt.Printf("Type of arr1 is %T\n", arr1)
	
	
	//Composing Data
	arr2 := []int {100,200,300,400,500}
	fmt.Println(arr2)
	fmt.Printf("Type of arr2 is %T\n", arr2)
	for i,v := range(arr2){
		fmt.Printf("%v\t%v\n",i,v)
	}
	fmt.Printf("Slice it : arr2[1:3] is %v\n",arr2[1:3])
	//fmt.Printf("Slice it : arr[-1] is %v\n",arr2[-1])
	fmt.Printf("Slice it : arr[2:] is %v\n",arr2[2:])
	fmt.Printf("Slice it : arr[:4] is %v\n",arr2[:4])
	fmt.Printf("Length of arr2 is %v\n",len(arr2))
	
	arr2 = append(arr2, 600)
	fmt.Println(arr2)
	
	//arr3 := [5]int{150,250,350,450,550}
	//arr2 = append(arr2, arr3...)
	//fmt.Println(arr2)
	
	//make
	x := make([]int, 10, 10) 
	fmt.Println(x)
	fmt.Printf("Length of x is : %v\n",len(x))
	fmt.Printf("Capacity of x is : %v\n",cap(x))
	
	x = append(x, 11)
	fmt.Println(x)
	fmt.Printf("Length of x is : %v\n",len(x))
	fmt.Printf("Capacity of x is : %v\n",cap(x))
	
	x = append(x, 12)
	fmt.Println(x)
	fmt.Printf("Length of x is : %v\n",len(x))
	fmt.Printf("Capacity of x is : %v\n",cap(x))
	
	x = append(x, 13)
    fmt.Println(x)
    fmt.Printf("Length of x is : %v\n",len(x))
    fmt.Printf("Capacity of x is : %v\n",cap(x))
    
	//2D Arrays
	
	names := []string{"Arun", "Soumya"}
	fruits := []string{"Apple", "Orange"}
	
	fruitmap := [][]string{names, fruits}
	fmt.Println(fruitmap)
	fmt.Printf("Length of x is : %v\n",len(fruitmap))
    fmt.Printf("Capacity of x is : %v\n",cap(fruitmap))
    
}

func map_(){
	
	fmt.Println("\nmap_")
	mymap := map[string]int{
		"Arun": 7760,
	}
	fmt.Println(mymap)
	mymap["Arun"] = 240
	mymap["Soumya"] = 500
	fmt.Println(mymap)
	
	value, iskeyexist := mymap["Arun"]
	fmt.Printf("iskeyexist: %v, value: %v\n",iskeyexist,value)
	
	value, iskeyexist = mymap["Ram"]
	fmt.Printf("iskeyexist: %v, value: %v\n",iskeyexist,value)
	
	for key, value := range mymap {
		fmt.Printf("Key: %v, Value: %v\n",key,value)
	}
	//delete(mymap["Arun"])
	
	
}
func arrays_excercise(){
	x := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(x)
	for i,v := range x{
		fmt.Println(i,v)
	}
	y := x[:5]
	x = append(x,y...)
	fmt.Println(x)
	fmt.Println(y)
	
	x = append(x[:3],y[2:]...)
	fmt.Println(x)
}	
func string_excercise() {
	s := make([]string,0)
	
	fmt.Println(s)
	s = append(s, "Bangalore", "Kozhikode")
	fmt.Println(s)
	for _,v := range s{
		fmt.Println(v)
	}
	
	m := map[string][]string{
		"Maharashtra" 	: []string{"Satara","Sangli","Kohlapur","Jath"},
		"Delhi" 		: []string{"Maharani Bagh"},
		"Bangalore" 	: []string{"Marathahalli"},
		}
	fmt.Println(m)
	fmt.Println("\n")
	for key, value := range m{
		fmt.Printf("%v\n",key)
		for _,place := range value{
			fmt.Printf("\t %v\n",place)
		} 
	}
}

func struct_excercise() {
	type person struct {
		first string
		last string
	}
	type employee  struct {
		person
		emp_id uint32
	}
	
	person1 := person {
		first 	: "Arunraj",
		last	: "Karingali",
	} 
	fmt.Println(person1)
	fmt.Println(person1.first,person1.last)
	
	employee1 := employee {
		person : person1,
		emp_id : 222,
	}
	fmt.Println(employee1)
	
	//Anonemous Struct Example 
	ps := struct {
		name	string 
		length uint32
	}{
		name 	: "circle",
		length 	: 320,
	}
	fmt.Println(ps)
}

func vehicle_example(){
	type vehicle struct {
		doors 	uint32
		color	string
	}
	
	type truck struct {
		vehicle
		four_wheel bool
	}
	type sedan struct {
		vehicle
		laxury bool
	}
	
	newTruck := truck {
		vehicle : vehicle {
			doors : 6,
			color : "Black",
		},
		four_wheel : false,
	}
	
	fmt.Println(newTruck)
	fmt.Println(newTruck.vehicle.doors)
	fmt.Println(newTruck.vehicle.color)
	fmt.Println(newTruck.four_wheel)
	
	
	newSedan := sedan {
		vehicle : vehicle {
			doors : 4,
			color : "Red",
		},
		laxury : true,
	}
	fmt.Println(newSedan)
}
