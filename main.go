package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Unable to open file - ", err)
		return
	}
	curLine := ""
	for {
		data := make([]byte, 8)
		_, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading contents of file ", err)
			return
		}
		splits := strings.Split(string(data[:]), "\n")
		curLine += splits[0]
		if len(splits) > 1 {
			fmt.Printf("read: %s\n", curLine)
			curLine = ""
			curLine += splits[1]
		}
	}
	if len(curLine) > 0 {
		fmt.Printf("read: %s", curLine)
	}
}
