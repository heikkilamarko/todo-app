package main

import "todo-api/internal/service"

func main() {
	(&service.Service{}).Run()
}
