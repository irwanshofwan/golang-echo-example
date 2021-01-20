package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		fmt.Println("this is example", rand.Intn(100))
		time.Sleep(10 * time.Second)
	}
}
