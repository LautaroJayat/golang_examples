package main

import "fmt"

func genValues(n ...int) <-chan int{
	outChannel := make (chan int)
	go func(){
		for _, nums:= range n{
			outChannel<- nums
		}
		close(outChannel)
	}()
	return outChannel

}

func printValue(inputChannel <-chan int) <-chan int {
	outChannel := make(chan int)
	go func(){
		for toPrint:= range inputChannel{
			fmt.Println("[From Print Stage]: recived value ", toPrint)
			fmt.Println("")

			outChannel<-toPrint
		}
		close(outChannel)

	}()
	return outChannel
}


func multiplyBySix(inputChannel <-chan int) <-chan int {
	outChannel := make(chan int)
	go func(){
		for toMultiply:= range inputChannel{
			fmt.Println("[From Multiply Stage]: recived value ", toMultiply)
			fmt.Println("[From Multiply Stage]: multiplying by six and sending to next stage")
			fmt.Println("")
			outChannel<- toMultiply * 6
		}
		close(outChannel)
	}()
	return outChannel
}


func main(){
	fmt.Println("[From Main]: Starting...")
	fmt.Println("[From Main]: Creating Value Generator")
	fmt.Println("")
	valueGenOut := genValues(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("[From Main]: Piping to Printing Stage")
	fmt.Println("")
	printingOut := printValue(valueGenOut)
	fmt.Println("[From Main]: Piping to Multiply By Six Stage")
	fmt.Println("")

	multiplyOut := multiplyBySix(printingOut)

	for n:= range multiplyOut {
		fmt.Println("[From Main]: Recived from multiplyBySix Stage ",n )
		fmt.Println("")

	}

}