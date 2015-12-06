package nil

import (
	id "development/modules/UUID"
	e "development/modules/email"
	"fmt"
	c "github.com/skilstak/go/colors"
)

func Output() {
	fmt.Println(c.B3 + "╔═════════════════════════════════════════════════════╗")
	fmt.Println(c.B3 + "║ " + c.Y + "SkilStak Support System, By: " + c.G + "@whitman-colm          " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.O + "This is run by the following modules:               " + c.B3 + "║")
	fmt.Println(c.B3 + "║                                                     ║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS Main  |  " + c.V + "V β-S-3.1.0                            " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS Email |  " + e.Version + "                            " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS UUID  |  " + id.Version + "                            " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS UUID  |  " + c.V + "V β-S-1.0.2                             " + c.B3 + "║")
	fmt.Println(c.B3 + "╚═════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println(c.M + "How can I help you today?")
	fmt.Println(c.R + "[A]" + c.B + " I want programming help!")
	fmt.Println(c.R + "[B]" + c.B + " Just stoppin' by.")
}
