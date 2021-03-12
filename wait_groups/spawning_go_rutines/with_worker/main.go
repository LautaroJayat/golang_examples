package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var w sync.WaitGroup
	fmt.Println("Starting example")
	fmt.Println()
	for i:=0; i<10; i++{
		w.Add(1)
		fmt.Println("Workers added:", i + 1)
		go func(i int){
			fmt.Println("starting worker", i)
			time.Sleep(time.Second)
			fmt.Println("done here with worker", i, "!")
			w.Done()
		}(i)
	}
	fmt.Println()
	fmt.Println("waiting for all of this guys")
	fmt.Println()

	w.Wait()
	fmt.Println()
	fmt.Println("Done in main function")
	

}