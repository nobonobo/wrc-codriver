//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/nobonobo/wrc-logger/easportswrc"
)

func main() {
	for k, v := range easportswrc.Stages {
		src := filepath.Join("log", strconv.FormatFloat(k, 'f', -1, 64)+".log")
		if _, err := os.Stat(src); os.IsNotExist(err) {
			continue
		}
		loc := easportswrc.Locations[v.Location]
		dir := fmt.Sprintf("%02d.%s", v.Location+1, easportswrc.LocationKeys[v.Location])
		name := fmt.Sprintf("%02d.%s", v.Stage+1, loc.Stages[v.Stage])
		dst := filepath.Join("pacenotes", dir, name+".log")
		if _, err := os.Stat(dst); os.IsExist(err) {
			continue
		}
		log.Println("move:", src, "->", dst)
		if err := os.Rename(src, dst); err != nil {
			log.Println(err)
		}
	}
}
