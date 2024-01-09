package helloWorld

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	greetingPrefixInEnglish = "Hello, "
	greetingPrefixInSpanish = "Hola, "
	greetingPrefixInFrench  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (greetingPrefix string) {
	switch language {
	case spanish:
		greetingPrefix = greetingPrefixInSpanish
	case french:
		greetingPrefix = greetingPrefixInFrench
	default:
		greetingPrefix = greetingPrefixInEnglish
	}

	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
