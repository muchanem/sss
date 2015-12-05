package nil

import (
	u "github.com/whitman-colm/go-1/utils/other"
	io "io/ioutil"
	"strconv"
	"strings"
)

<<<<<<< HEAD
=======
func main() {
	x, err := strconv.Atoi(LastLine("test.txt"))


func WriteUUID(fname string) {
	err := io.WriteFile(fname)

}

>>>>>>> bdd4fdbf9e3bcc1c8ed06ad2f7c510e23fcc6551
func LastLine(fname string) int {
	data, err := io.ReadFile(fname)
	u.QuitAtError(err)
	sdata := string(data)
	spdata := strings.Split(sdata, "\n")
	lline := spdata[len(spdata)-2]
	return strconv.Atoi(lline)
}
