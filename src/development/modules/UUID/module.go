package nil

import (
	u "github.com/whitman-colm/go-1/utils/other"
	io "io/ioutil"
	"strconv"
	"strings"
)

const Version string = "\033[1;35mV ß-S-1.0.5"

func WriteUUID(fname string, data int) {
	//<<<<<<< HEAD
	bconv := /*strings.Split(*/ strconv.Itoa(data) /*, "")*/
	//Defines bconv as an array of strings from the int data, so 1020 becomes "1","0","2","0"
	//bl := []byte{}
	//makes bl an empty byte slice. A slice is golang's array with no limit on how many number
	blr := []byte(bconv)
	//define blr outside of the loop. the layer in which a constant is defined is the lowest layer it can be
	//used.
	/*for _, asc := range bconv {
		blr = append([]byte(bconv), bl)
		//appends
	}*/
	io.WriteFile(fname, blr, 0664)
	//=======
	bconv := strings.Split(strconv.Itoa(data), "")
	bl := []byte{}
	for _, asc := range bconv {
		ascb := []byte(asc)
		append(ascb, bl)
	}
	io.WriteFile(fname, bl, 0664)
	//>>>>>>> f4b01a6e7669947e6b961cb215ced408a28ee7b2
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
