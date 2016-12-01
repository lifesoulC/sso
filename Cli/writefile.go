package main

import (
	//"fmt"
	"log"
	"os"
)

func WriteFile(file []byte, path string) error {
	userFile := path
	M.Lock()
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		log.Printf("writefile error  Cause:%v \n", err)
		M.Unlock()
		return err
	}
	s := string(file)
	if s == "" {
		log.Printf("writefile error  Cause:%s \n", path)
		M.Unlock()
		return err
	} else {
		if s == "setnil" {
			file = []byte("[]")
		}
	}

	fout.Write(file)
	M.Unlock()
	return nil
}
