package util

import (
	"go-wafw00f/log"
	"io/ioutil"
)

func GetAllFile(pathname string) ([]string, error) {
	var s []string
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		log.Error("Read File Error")
		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}
