package main

import (
	"fmt"
	"os"
)

func main() {
	//if len(os.Args) > 1 {
	//	fmt.Println("Hello World", os.Args[1])
	//}
	//fmt.Println(os.Args)
	fmt.Println(os.Executable())
	os.Exit(-1)
}
