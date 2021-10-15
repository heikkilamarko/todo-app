package workflow

import "fmt"

func GreetActivity() error {
	fmt.Printf("Greetings from greet activity :)")
	return nil
}
