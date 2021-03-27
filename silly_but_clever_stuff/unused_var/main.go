package main

import (
	"fmt"
)

func main(){
	// unusedString will throw an error, but the exported one from ./unused_but_chill will not
	// Try commenting line bellow and you will see
	var unusedString string = "use me"
	fmt.Println("[From Main]: there is a unused variable here")
}