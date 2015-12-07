package main

import (
	l "development/modules/UITest"
	e "development/modules/email"
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-1/utils/input"
	"os"
	"os/user"
	"strings"
)

func main() {
	usr, _ := user.Current()
	l.Output()
	fmt.Println()
	done := false
	for done == false {
		function, _ := i.Prompt(c.B1 + usr.Username + c.B01 + "@" + c.B00 + "sss:" + c.Y + "~" + c.C + "$ " + c.X)
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
			fmt.Println("Sorry, try entering one of the letters in red")
		}
	}
}
