package main

import (
	"fmt"
	"urlsh/internal/config"
)

func main() {
	fmt.Println("start")
	cfg := config.Load()
	fmt.Println(cfg)
}
