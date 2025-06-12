package main

import (
	"fmt"

	"github.com/evilmagics/go-redfox"
)

func main() {
	err := redfox.New("SERVER_ERROR", "internal server error")

	fmt.Print(err.C().Error())
}
