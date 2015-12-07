package main

import (
	l "development/modules/UITest"
	e "development/modules/email"
	i "development/modules/other"
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
	"strings"
)

func main() {
	l.Output()
	fmt.Println()
	done := false
	for done == false {
		function, _ := i.Input("~", "sss")
		function = strings.ToLower(function)
		switch function {
		case "a":
			e.Startup()
			done = true
		//case "b":
		//Setup slack helper lib. STOP SLACKING!
		case "b":
			os.Exit(-1)
		default:
			fmt.Println(c.B00 + "Sorry, try entering one of the letters in red")
		}
	}
}
