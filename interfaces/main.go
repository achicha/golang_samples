/*learn interfaces by implementing language Bots*/
package main

import "fmt"

// custom types
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

// main
func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

// functions
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//func (eb englishBot) getGreeting() string {
	// some custom logic
	// !!! if we do not use `eb` variable in the body of our function, we can just remove it from the head.
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	// some another custom logic
	return "Hola!"
}
