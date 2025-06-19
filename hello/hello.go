package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Aisa"
	currentTime := time.Now().Format("2006-01-02")
	fmt.Printf("Hello, %s!\nToday is %s\n", name, currentTime)
}
