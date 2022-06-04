package main

import "todo-worker/internal"

func main() {
	(&internal.Service{}).Run()
}
