package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main(){

	var w sync.WaitGroup

	fmt.Println("[ From MAIN ]: Starting channels example")
	fmt.Println("[ From MAIN ]: spawing gorutines")
	c:= make(chan string, 10)
	for i:= 0; i < 10; i++{
		w.Add(1)
		go func(i int){
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
			c <- fmt.Sprintf("[ From channel %v ]: hi there!", i);
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
			fmt.Println(fmt.Sprintf("[ From channel %v ]: now terminating", i))
			w.Done()
		}(i)
	}

	fmt.Println("\n[ From MAIN ]: Hey Guys, here from main function ,how are you doing!");
	fmt.Println();

	msgs:= 0;
	for msg:= range c{
		fmt.Println(msg)
		msgs++
		if msgs == 10{
			fmt.Println("\n[ From MAIN ]: Closing channel")
			fmt.Println()
			close(c)
		}
	}
	fmt.Println("\n[ From MAIN ]: I will wait until all of you are done")
	fmt.Println()


	w.Wait()




	fmt.Println("\n[ From MAIN ]: Nice to have news from you guys");


}