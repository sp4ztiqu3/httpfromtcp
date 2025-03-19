package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Unable to open file - ", err)
		return
	}
	for {
		data := make([]byte, 8)
		_, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("Error reading contents of file ", err)
			return
		}
		fmt.Printf("read: %s\n", data)
	}
}
