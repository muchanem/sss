package main

import (
	"fmt"
	u "github.com/whitman-colm/go-1/utils/other"
	io "io/ioutil"
	"strconv"
	"strings"
)

func main() {
	test, err := strconv.Atoi(LastLine("test.txt"))
	u.QuitAtError(err)
	fmt.Println(test)
}

func LastLine(fname string) string {
	data, err := io.ReadFile(fname)
	u.QuitAtError(err)
	sdata := string(data)
	spdata := strings.Split(sdata, "\n")
	lline := spdata[len(spdata)-2]
	return lline
}
