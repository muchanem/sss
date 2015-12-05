package main

import (
	e "development/modules/email"
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-1/utils/input"
	s "github.com/whitman-colm/go-1/utils/other"
	"os"
	"strings"
)

func main() {
	fmt.Println(c.CL + c.B3 + "SkilStak Support System by whitman-colm on git")
	fmt.Println(c.O + "This system runs on the following module versions")
	fmt.Println(c.B1 + "S3 Main             |" + c.V + "V β-S-3.1.0")
	fmt.Println(c.B1 + "Support By Tickets  |" + e.Version)
	s.Spacer(1)
	fmt.Println(c.B2 + "How can we help you today?")
	fmt.Println(c.R+"{A}", c.G, "Get programming help.")
	//fmt.Println(c.R+"{B}", c.G, "Get miscellaneous help.")
	fmt.Println(c.R+"{B}", c.G, "I'm fine, thanks for asking.")
	function, _ := i.Prompt(c.B + "> " + c.M)
	function = strings.ToLower(function)
	done := false
	for done == false {
		switch function {
		case "a":
			e.Startup()
		//case "b":
		//Setup slack helper lib. STOP SLACKING!
		case "b":
			os.Exit(-1)
		default:
			fmt.Println("Sorry, try entering one of the letters in red")
		}
	}
}
