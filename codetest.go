package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func getDrives() (r []string) {
	for _, drive := range "CDEFGH" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			d := string(drive) + ":/"
			r = append(r, d)
			f.Close()
		}
	}
	return
}

func FindFileFromExtension(extension []string, dir string, files *[]string) {
	fs, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, f := range extension {
			if strings.HasSuffix(f.Name(),ex) {
				*files = append(*files, f.Name())
			}
		}

		if f.I
	}
}