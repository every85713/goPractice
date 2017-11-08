package main

import(
	"fmt"
	"controller"
	//"net/http"
)

func main() {
	fmt.Println("Main run")
	defer fmt.Println("Main end")
	controller.Routing()
}