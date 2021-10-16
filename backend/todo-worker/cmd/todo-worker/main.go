package main

import "todo-worker/internal/service"

func main() {
	(&service.Service{}).Run()
}
