package main

import (
	"fmt"
	"os"
)

func fileLen(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	stats, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return int(stats.Size()), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./program file_path")
	}

	size, err := fileLen(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("file size:", size)
}
