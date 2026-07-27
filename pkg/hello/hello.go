package hello

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if(name == "") {
		return englishHelloPrefix + "World"
	}
	return englishHelloPrefix + name 
}