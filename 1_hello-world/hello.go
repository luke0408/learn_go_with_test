package main

import "fmt"

const engHelloPrefix = "Hello, "
const korHelloPrefix = "안녕하세요, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Korean" {
		return korHelloPrefix + name
	}

	return engHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
