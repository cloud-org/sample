// author: ashing
// time: 2020/4/16 10:55 上午
// mail: axingfly@gmail.com
// Less is more.

package main

import "fmt"

const (
	helloPrefix        = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = helloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
