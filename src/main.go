package main

import (
	"fmt"

	basic "github.com/Ufigat/basic/src/service"
)

func main() {
	Ivan := basic.NewHuman("Ivan")
	fmt.Println("Hello, " + Ivan.GetHumanName())
}
