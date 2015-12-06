package nil

import (
	u "github.com/whitman-colm/go-1/utils/other"
	io "io/ioutil"
	"strconv"
	"strings"
)

const Version string = "\033[1;35mV ß-S-1.0.2"

func WriteUUID(fname string, data []byte) {
	io.WriteFile(fname, data, 0664)
}

func LastLine(fname string) int {
	data, err := io.ReadFile(fname)
	u.QuitAtError(err)
	sdata := string(data)
	spdata := strings.Split(sdata, "\n")
	lline := spdata[len(spdata)-2]
	uuid, err := strconv.Atoi(lline)
	u.QuitAtError(err)
	return uuid
}
