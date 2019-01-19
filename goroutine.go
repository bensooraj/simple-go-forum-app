package main

import (
	"fmt"
	"sync"
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

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		// time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+100; i++ {
		// time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func print2() {
	// printNumbers2()
	// printLetters2()
}

func goPrint2() {
	// go printNumbers2()
	// go printLetters2()
}
