package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){

	var w sync.WaitGroup

	fmt.Println("Starting channels example")
	fmt.Println("spawing all of you pals")
	c:= make(chan string, 1)
	for i:= 0; i < 10; i++{
		w.Add(1)
		go func(i int){
			fmt.Println("channel", i, " just waiting");
			<-c
			fmt.Println("hi from channel", i, "now terminating");
			w.Done()
		}(i)
	}

	fmt.Println("\nIm going to sleep now");
	fmt.Println();

	time.Sleep(time.Second * 5)
	fmt.Println("\nHey Guys, here from main function ,how are you doing!");
	fmt.Println();

	for j:=0; j<10; j++{
		c <- "Hi to every rutine, please respond"
	}
	w.Wait()

	fmt.Println("\nNice to have news from you guys");


}