package main

import (
	"fmt"

	"devgym-exc.com/lru/internal"
)

func main() {
	lru := internal.NewLRU(3)
	lru.Set("Mark", 12344)
	lru.Set("Mark2", 12354)
	lru.Set("Mark3", 1235544)
	lru.Set("Mark4", 123554466)
	fmt.Println(lru)
	fmt.Println(lru.Keys)
}
