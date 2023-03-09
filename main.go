package main

import (
	"fmt"
	"os"
)

func ensure_file(file string) error {
	fd, err := os.Stat(file)
	if err != nil {
		return err
	}

	if fd.IsDir() {
		return fmt.Errorf("%s is a directory", file)
	}

	fp := fd.Mode().Perm()

	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	for data[len(data)-1] == 10 {
		data = data[:len(data)-1]
	}

	data = append(data, 10)

	return os.WriteFile(file, data, fp)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ensure_newline <files...>")
		os.Exit(1)
	}

	for _, file := range os.Args[1:] {
		if err := ensure_file(file); err != nil {
			fmt.Println(err)
		}
	}
}
