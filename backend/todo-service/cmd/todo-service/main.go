package main

import "todo-service/internal"

func main() {
	(&internal.Service{}).Run()
}
