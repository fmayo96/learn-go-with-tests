package main

import "fmt"

func Hello(name ...string) string {
	if len(name) == 0 {
		return "Hello World!"
	}
	return "Hello " + name[0] + "!"
}

func main() {
	fmt.Println(Hello("World"))
}
