package main

import (
	id "development/modules/UUID"
	e "development/modules/email"
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-1/utils/input"
)

func Output() {
	fmt.Println(c.B3 + "╔═════════════════════════════════════════════════════╗")
	fmt.Println(c.B3 + "║ " + c.Y + "SkilStak Support System, By: " + c.G + "@whitman-colm          " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.O + "This is run by the following modules:               " + c.B3 + "║")
	fmt.Println(c.B3 + "║                                                     ║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS Main  |  " + c.X + "V β-S-3.1.0                      " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS Email |  " + c.X + e.Version + "                      " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS UUID  |  " + c.X + id.Version + "                      " + c.B3 + "║")
	fmt.Println(c.B3 + "╚═════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println(c.M + "How can I help you today?")
	fmt.Println(c.R + "[A]" + c.B + " I want programming help!")
	fmt.Println(c.R + "[B]" + c.B + " Just stoppin' by.")
}

func main() {
	fmt.Print(c.Clear)
	Output()
	i.Prompt("> ")
}
