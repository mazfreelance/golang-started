package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var name string = "John Doe"
	var age int = 32
  
	fmt.Println("My name is", name)
	fmt.Println("I am", age, "years old")

	
	for i := 0; i < 5; i++ {
		fmt.Println("Counting:", i)
	  }

	greet("Jaapar")
}

func greet(name string) {
	fmt.Println("Hello, " + name + "!")
  }