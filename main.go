package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/kvrasuli/hexlet-go/greeting"
)

func main() {
	fmt.Println(greeting.Hello())

	color.Red("We have red")
	color.Magenta("And many others ..")
}
