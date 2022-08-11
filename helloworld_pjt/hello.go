package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHolaPrefix = "Hola, "
const frenchBonjourPrefix = "Bonjour, "

func hello(name string , language string) string{
	if (name == ""){
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string{
	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = frenchBonjourPrefix
	case "Spanish":
		prefix = spanishHolaPrefix
	}

	return prefix
}

func main(){
	fmt.Println(hello("Chris", "English"))
}