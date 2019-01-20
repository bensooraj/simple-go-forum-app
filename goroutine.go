package main

import (
	"fmt"
	"time"
)

func printNumbers1() {
	for i := 0; i < 100; i++ {
		// fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 100; i++ {
		// fmt.Printf("%c ", i)
	}
}

func print1() {
	printNumbers1()
	printLetters1()
}

func goPrint1() {
	go printNumbers1()
	go printLetters1()
}

func printNumbers2(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	w <- true
}

func printLetters2(w chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	fmt.Println()
	w <- true
}

func print2() {
	// printNumbers2()
	// printLetters2()
}

func goPrint2() {
	// go printNumbers2()
	// go printLetters2()
}

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Threw  >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("Caught  <<", num)
	}
}

func callerA(c chan string) {
	c <- "Hello, World!"
	close(c)
}

func callerB(c chan string) {
	c <- "Hola, Mundo!"
	close(c)
}
