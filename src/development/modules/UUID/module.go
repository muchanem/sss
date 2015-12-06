package nil

import (
	u "github.com/whitman-colm/go-1/utils/other"
	io "io/ioutil"
	"strconv"
	"strings"
)

const Version string = "\033[1;35mV ß-S-1.0.4"

func WriteUUID(fname string, data int) {
	bconv := strings.Split(stringconv.Itoa(data), "")
	bl := []byte{}
	for _, asc := range bconv {
		blr := append(byte(bconv), bl)
	}
	io.WriteFile(fname, blr, 0664)
}

func LastLine(fname string) int {
	data, err := io.ReadFile(fname)
	if err != nil {
		u.QuitAtError(err)
	}
	sdata := string(data)
	spdata := strings.Split(sdata, "\n")
	lline := spdata[len(spdata)-2]
	uuid, err := strconv.Atoi(lline)
	if err != nil {
		u.QuitAtError(err)
	}
	return uuid
}
