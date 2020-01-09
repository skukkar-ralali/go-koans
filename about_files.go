package go_koans

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func aboutFiles() {
	filename := "about_files.go"
	contents, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(contents), "\n")
	assert(lines[5] == fmt.Sprintf("	\"strings\"")) // handling files is too trivial
}
