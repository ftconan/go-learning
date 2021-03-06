// Author: magician
// File:   demo1.go
// Date:   2020/6/17
package main

import (
"log"
"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// ...lots of work...
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}
