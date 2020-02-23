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
			
		}
	}
}