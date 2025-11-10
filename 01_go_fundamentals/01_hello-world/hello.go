package hello

import "fmt"

const (
	korean  = "Korean"
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	koreanHelloPrefix  = "안녕하세요, "
	spainHelloPrefix   = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

/**
 * language에 따라 다른 인사말을 반환한다.
 * @param language
 * @return prefix
 */
func greetingPrefix(language string) (prefix string) {
	switch language {
	case korean:
		prefix = koreanHelloPrefix
	case spanish:
		prefix = spainHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
