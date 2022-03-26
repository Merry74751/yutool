package io

import (
	"bufio"
	"github.com/Merry74751/yutool/str"
	"io"
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
	buffer := str.NewBuffer()
	for {
		s, err := i.reader.ReadString('\n')
		buffer.Append(s)
		if err == io.EOF {
			break
		}
	}
	return buffer.String()
}
