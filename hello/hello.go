package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchhHelloPrefix = "Bonjour, "
const portuguesePrefix = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchhHelloPrefix
	case portuguese:
		prefix = portuguesePrefix
	}
	return prefix + name

	// if language == spanish {
	// 	return spanishHelloPrefix + name
	// }

	// if language == french {
	// 	return frenchhHelloPrefix + name
	// }

	// return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Eduardo", ""))
}
