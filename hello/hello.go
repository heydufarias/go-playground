package main

import (
	"fmt"
)

const (
	spanish    = "Spanish"
	french     = "French"
	portuguese = "Portuguese"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchhHelloPrefix = "Bonjour, "
	portuguesePrefix   = "Olá, "
)

// public functions start with a capital letter
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

	// if language == spanish {
	// 	return spanishHelloPrefix + name
	// }

	// if language == french {
	// 	return frenchhHelloPrefix + name
	// }

	// return englishHelloPrefix + name
}

// private functions start with a lowercase letter
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchhHelloPrefix
	case portuguese:
		prefix = portuguesePrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Eduardo", ""))
}
