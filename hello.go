package main

import "fmt"

const englishGreetingPrefix = "Hello, "
const spanishGreetingPrefix = "Hola, "
const frenchGreetingPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func greeting(s string) string {
	if s == "" {
		return "World"
	}
	return s
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchGreetingPrefix
	case spanish:
		prefix = spanishGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}
	return
}

func Hello(name string, language string) string {
	return greetingPrefix(language) + greeting(name)
}

func main() {
	fmt.Println(Hello("Chris", "French"))
}
