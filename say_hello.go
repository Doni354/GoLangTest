package go_say_hello
import "fmt"
func SayHello() {
	var a string
	fmt.Println("Hello World!")
	fmt.Println("What's your name?")
	fmt.Scan(&a)
	fmt.Println("My Name is ", a)
	fmt.Println("Nice to see you", a)
}