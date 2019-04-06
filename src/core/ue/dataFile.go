package ue

import (
	"errors"
	"io/ioutil"
	"os"
)

func recordPasswd(user string, passwd string) error {
	path := "depend/nothinghere.txt"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return errors.New("can't open file " + path)
	}

	_, err = file.WriteString(user + "," + passwd + "\n")
	if err != nil {
		return err
	}

	return nil
}


func loadHtml(key, file_name string) {
	info, err := readFile(file_name)
	if err != nil {
		return
	}
	change[key] = info
}

func readFile(file_name string) ([]byte, error) {
	fi, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	return ioutil.ReadAll(fi)
}
