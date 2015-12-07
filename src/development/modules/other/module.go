package input

import (
	"bufio"
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
	"os/user"
	"strings"
)

func Input(dir string, program string) (t string, err error) {
	usr, _ := user.Current()
	fmt.Print(c.B1 + usr.Username + c.B01 + "@" + c.B00 + program + ":" + c.Y + dir + c.C + "$ " + c.X)
	reader := bufio.NewReader(os.Stdin)
	t, err = reader.ReadString('\n')
	if err != nil {
		return t, err
	}
	t = strings.TrimSpace(t)
	return t, nil
}
