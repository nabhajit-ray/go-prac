package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchhHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	} else if language == french {
		return frenchhHelloPrefix + name
	}
	return englishHelloPrefix + name

}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
