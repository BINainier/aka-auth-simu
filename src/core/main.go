package main

import (
	"core/hss"
	"core/mme"
	"core/service"
	"core/ue"
	"flag"
	"os"
	"path/filepath"
)

var (
	path = flag.String("path", "", "HTML directory path")
)

func main() {
	flag.Parse()
	p := genPath()
	go ue.Run("8082", p)
	go mme.Run("8083")
	go hss.Run("8084")
	go sv.Run("8085")

	select {}
}

func genPath() string {
	flag.Parse()

	if len(*path) > 0 {
		return *path
	} else if workingDir, err := os.Getwd(); err == nil {
		path := filepath.Join(workingDir, "html")
		return path
	}

	return ""
}

