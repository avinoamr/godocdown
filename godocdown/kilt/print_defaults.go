// This file was AUTOMATICALLY GENERATED by kilt-import (smuggol) from github.com/robertkrimen/kilt

package kilt

import (
	"bufio"
	"bytes"
	Flag "flag"
	"fmt"
	"os"
	"strings"
)

// PrintDefaults
func (self Kilt) PrintDefaults(flag *Flag.FlagSet) {
	PrintDefaults(flag)
}

func PrintDefaults(flag *Flag.FlagSet) {
	var into bytes.Buffer
	flag.SetOutput(&into)
	flag.PrintDefaults()
	outfrom := bufio.NewReader(&into)
	for {
		line, err := outfrom.ReadString('\n')
		if err != nil {
			break
		}
		if strings.HasSuffix(line, ": \x00\n") {
			continue
		}
		fmt.Fprint(os.Stderr, line)
	}
}
