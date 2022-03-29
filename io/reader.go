package io

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

type rw struct {
	filename string
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func ReadFile(filename string) rw {
	i := rw{filename: filename}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Printf("open file %s error: %s", filename, err)
	}
	i.reader = bufio.NewReader(file)
	return i
}

func (i rw) ReadAll() string {
	bytes, err := ioutil.ReadAll(i.reader)
	if err != nil {
		log.Printf("readfile error: %s", err)
	}
	return string(bytes)
}
