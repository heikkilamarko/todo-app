package main

import "todo-api/service"

func main() {
	(&service.Service{}).Run()
}
