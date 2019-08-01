package main

import ("fmt"
		"runtime"
		"sync"
)

var wg sync.WaitGroup
var mtx sync.Mutex

var counter int

func main() {
	
	fmt.Println("Helloworld!!")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	
	Init_()
	
	for i := 0;i<100; i++ {
		go foo()
		go foo()
		go bar()
		go bar()
	}
	
	fmt.Println(runtime.NumGoroutine())
	
	wg.Wait()
	fmt.Println("Exit from main()!")
	
}
func Init_() {
	wg.Add(400)
}
func foo(){
	for i:= 0; i<50; i++ {
		mtx.Lock()
		counter ++
		fmt.Println("<<<<foo>>>> Count is ", counter)
		runtime.Gosched()
		mtx.Unlock()
	}
	wg.Done()
}

func bar(){
	for i:= 0; i<50; i++ {
		mtx.Lock()
		counter ++
		fmt.Println("+++++++++++bar++++ Count is ", counter)
		runtime.Gosched()
		mtx.Unlock()
	}
	wg.Done()
}

