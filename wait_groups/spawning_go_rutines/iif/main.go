package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var w sync.WaitGroup
	fmt.Println("Starting example")

	for i:=0; i<10; i++{
		w.Add(1)
		go worker(i, &w)
	}
	fmt.Println("waiting for all of this guys")
	w.Wait()
	fmt.Println("Done in main function")

}

func worker(i int, w *sync.WaitGroup){
	fmt.Println("starting worker", i)
			time.Sleep(time.Second)
			fmt.Println("done here with worker", i, "!")
			w.Done()
}