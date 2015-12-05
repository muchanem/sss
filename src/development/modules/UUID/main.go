package nil

import (
	"fmt"
	u "github.com/whitman-colm/go-1/utils/other"
	io "io/ioutil"
	"strconv"
	"strings"
)

func main() {
	x, err := strconv.Atoi(LastLine("test.txt"))


func WriteUUID(fname string) {
	err := io.WriteFile(fname)

}

func LastLine(fname string) int {
	data, err := io.ReadFile(fname)
	u.QuitAtError(err)
	sdata := string(data)
	spdata := strings.Split(sdata, "\n")
	lline := spdata[len(spdata)-2]
	return strconv.Atoi(lline)
}
