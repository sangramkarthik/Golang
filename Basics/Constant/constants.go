package main

import "fmt"

const name = "sangram"

func main() {

	const pi = 3.14178

	const (
		age = 21

		interest = "Boxing"
	)

	fmt.Printf("Hi!,I'm %v. I'm %v.\n", name, age)
	fmt.Printf("I like %v\n", interest)
	fmt.Printf("The value of Pi is %v \n", pi)
}
