package nil

import (
	id "development/modules/TimeID"
	e "development/modules/email"
	"fmt"
	c "github.com/skilstak/go/colors"
)

const version string = c.V + "V S-1.0.0"

func Output() {
	fmt.Print(c.Clear)
	fmt.Println(c.B3 + "╔═════════════════════════════════════════════════════╗")
	fmt.Println(c.B3 + "║ " + c.Y + "SkilStak Support System, By: " + c.G + "whitman-colm & tsnlc04 " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.O + "This is run by the following modules:               " + c.B3 + "║")
	fmt.Println(c.B3 + "║                                                     ║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS Main  |  " + c.V + "V S-1.0.0                              " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS Email |  " + e.Version + "                              " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS UUID  |  " + id.Version + "                              " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS UI    |  " + version + "                              " + c.B3 + "║")
	fmt.Println(c.B3 + "╚═════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println(c.M + "How may we help you today?")
	fmt.Println(c.R + "[A]" + c.B + " Get programming help!")
	fmt.Println(c.R + "[B]" + c.B + " Check the FAQs " + c.B01 + "[Coming soon!]")
	fmt.Println(c.R + "[C]" + c.B + " Just stoppin' by.")
	fmt.Pintln(c.Y + "When you know what you need just type the corrosponding letter in the command line")
}
