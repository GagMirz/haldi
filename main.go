package main

import (
	"fmt"
	"haldi/internal/manager"
)

func main() {
	err := manager.InitManager("./config.json")
	if (err != nil) {
		fmt.Printf("%s", err)
	}
}
