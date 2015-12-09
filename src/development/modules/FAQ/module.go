package nil

import (
	i "development/modules/other"
	"fmt"
	c "github.com/skilstak/go/colors"
	u "github.com/whitman-colm/go-1/utils/other"
	"os"
	"strings"
	//gofaq "development/modules/FAQ/list/gofaq"
)

const Version string = c.V + "V α-S-1.0.0"

func listLangs() {
	fmt.Println(c.CL + c.B2 + "Hmm... here are the available FAQs")
	u.Spacer(1)
	fmt.Println(c.B01 + "python, or py")
	fmt.Println(c.B1 + "javascript, or js")
	fmt.Println(c.B01 + "golang, or go")
	fmt.Println(c.B1 + "java")
	fmt.Println(c.B01 + "html")
	fmt.Println(c.B1 + "linux or unix")
	fmt.Println(c.B01 + "linuxpi")
	/*fmt.Println(c.B1+"")
	fmt.Println(c.B01+"")
	fmt.Println(c.B1+"")
	fmt.Println(c.B01+"")
	fmt.Println(c.B1+"")*/
	u.Go(1)
}

func Startup() {
	fmt.Println(c.B3 + "FAQ Module By " + c.G + "@whitman-colm " + c.Y + "Version" + Version + c.X)
	for {
		fmt.Println(c.B2 + "Which language does your problem relate to? (you can type \"help\" to see available languages)")
		u.Spacer(2)
		mode, _ := i.Input("FAQ", "sss")
		mode = strings.ToLower(mode)
		switch mode {
		case "go", "golang":
			gofaq.BeginSearch()
		default:
			listLangs()
		}
	}
}
