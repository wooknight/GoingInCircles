package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer func() {
		fmt.Println("Inside main defer")

		if r := recover(); r != nil && r == "Aaankh lag gaya" {
			printStack()
			fmt.Println("I love kitties ")
		}
	}()
	fmt.Println("Inside main and getting ready to call bho")
	bho()
}

func bho() {
	defer func() {
		fmt.Println("Inside bho defer\nI love Aaankh my dear")
	}()

	fmt.Println("Getting ready to do a panic")
	panic("Aaankh lag gaya")
}

func printStack() {
	var buf [8192]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
