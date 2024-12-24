package main

import (
	"blog_agreegator/internal/config"
	"fmt"
)

func main() {
	file := config.Read()
	fmt.Println(file)
	file.SetUser("thana")
	fmt.Println(file)
}
