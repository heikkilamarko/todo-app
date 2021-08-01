package main

import "todo-service/internal/service"

func main() {
	(&service.Service{}).Run()
}
