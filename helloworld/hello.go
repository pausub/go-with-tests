package helloworld

import "fmt"

const (
	helloEnglishPrefix = "Hello, "
	helloSpanishPrefix = "Hola, "
	helloFrenchPrefix  = "Bonjour, "

	languageSpanish = "Spanish"
	languageFrench  = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case languageSpanish:
		prefix = helloSpanishPrefix
	case languageFrench:
		prefix = helloFrenchPrefix
	default:
		prefix = helloEnglishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Paulius", ""))
}
