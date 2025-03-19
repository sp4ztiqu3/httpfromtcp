package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func getLinesChannel(file io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func(ch chan string, file io.ReadCloser) {
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
				ch <- curLine
				curLine = ""
				curLine += splits[1]
			}
		}
		if curLine != "" {
			ch <- curLine
		}

		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file - ", err)
		}
		close(ch)
	}(ch, file)

	return ch
}

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Unable to open file - ", err)
		return
	}
	ch := getLinesChannel(file)
	for line := range ch {
		fmt.Printf("read: %s\n", line)
	}
}
