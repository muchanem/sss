package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-1/utils/input"
	s "github.com/whitman-colm/go-1/utils/other"
	e "github.com/whitman-colm/sss/src/development/modules/email"
	"os"
	"strings"
)

func main() {
	fmt.Println(c.B3+"SkilStak Script Support", c.V, "V S-1.0.0")
	s.Spacer(2)
	fmt.Println(c.B2 + "How can we help you today?")
	fmt.Println(c.R+"{A}", c.G, "Submit a ticket")
	fmt.Println(c.R+"{B}", c.G, "I'm fine, thanks for asking")
	function, _ := i.Prompt(c.B + "> " + c.M)
	function = strings.ToLower(function)
	switch function {
	case "a":
		e.Startup()
	case "b":
		os.Exit(-1)
	default:
		fmt.Println("Sorry, try entering one of the letters in red")
	}
}
