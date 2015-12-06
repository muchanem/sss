package main

import (
	l "development/modules/UITest"
	e "development/modules/email"
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-1/utils/input"
	"os"
	"strings"
)

func main() {
	l.Output()
	fmt.Println()
	function, _ := i.Prompt(c.B + "sss@skilstak:~$ " + c.M)
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
