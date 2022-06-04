package main

import "todo-api/internal"

func main() {
	(&internal.Service{}).Run()
}
