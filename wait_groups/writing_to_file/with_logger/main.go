package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func main(){
	var w sync.WaitGroup
	err:= os.Remove("to_modify.txt")
	if err != nil{
		fmt.Println("There is not a file to write in this folder")
	}else{
		fmt.Println("Removed an old file in order to start with a new one")
	}
	fmt.Println("Creating a file here")
	file, err := os.OpenFile("to_modify.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logger := log.New(file, "[from logger] : ", 0)

	if err!= nil{
		panic(err)
	}

	for i:= 0; i<10; i++{
		w.Add(1)
		go func(i int){
			fmt.Println("writing worker number", i)
			logger.Output(2, fmt.Sprint("worker ", i, " writing", "\n"))
			w.Done()
		}(i)
	}
	w.Wait()
	

	fmt.Println("\nNow reading file")
	data, err := ioutil.ReadFile("to_modify.txt")

	if err != nil{
		panic(err)
	}

	_, err = fmt.Print(string(data))

	if err != nil{
		panic(err)
	}

}