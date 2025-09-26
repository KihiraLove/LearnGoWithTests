package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const sentenceEnd = "!"

func Hello(name, language string) string {
	prefix := englishHelloPrefix
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}
	if name == "" {
		name = "World"
	}
	return prefix + name + sentenceEnd
}

func main() {
	fmt.Println(Hello("Doma", "English"))
}
