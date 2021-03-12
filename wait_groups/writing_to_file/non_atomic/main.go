package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func main(){
	var w sync.WaitGroup
	fmt.Println("Creating a file here")
	file, err := os.OpenFile("to_modify.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err!= nil{
		panic(err)
	}

	for i:= 0; i<10; i++{
		w.Add(1)
		go func(i int){
			_, err = file.Write([]byte(fmt.Sprint("worker ", i, " writing", "\n")))
			w.Done()
		}(i)
	}
	w.Wait()
	

	fmt.Println("Now reading file")
	data, err := ioutil.ReadFile("to_modify.txt")

	if err != nil{
		panic(err)
	}

	_, err = fmt.Print(string(data))

	if err != nil{
		panic(err)
	}

	

}