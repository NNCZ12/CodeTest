package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getDrives(drive string) (r []string) {
	f, err := os.Open(string(drive) + ":\\")
	if err == nil {
		d := string(drive) + ":/"
		r = append(r, d)
		f.Close()
	}
	return
}

func findFileFromExtension(extension []string, dir string, files *[]string) {
	fs, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, f := range fs {
			for _, ex := range extension {
				if strings.HasSuffix(f.Name(), ex) {
					*files = append(*files, dir+"/"+f.Name())
				}
			}

			if f.IsDir() {
				path := dir + "/" + f.Name()
				findFileFromExtension(extension, path, files)
			}
		}
	}
}

func Createtxt(a []string) {
	file, _ := os.Create("output.txt")
	for _, o := range a {
		file.WriteString(o)
		file.WriteString("\n")
	}
}

func main() {
	var inputdrives string
	fmt.Print("Input Your Drives : ")
	fmt.Scan(&inputdrives)
	drives := getDrives(inputdrives)
	files := []string{}
	for _, d := range drives {
		findFileFromExtension([]string{".jpg"}, d, &files)
	}

	Createtxt(files)

}
