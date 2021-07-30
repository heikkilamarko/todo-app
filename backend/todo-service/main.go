package main

import "todo-service/service"

func main() {
	(&service.Service{}).Run()
}
