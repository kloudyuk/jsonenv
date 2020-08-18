package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readStdin() ([]byte, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return nil, fmt.Errorf("Error reading input")
	}

	reader := bufio.NewReader(os.Stdin)
	var output []byte

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		output = append(output, b)
	}
	return output, nil
}

func readFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {
	var in []byte
	var err error
	if len(os.Args) > 1 {
		in, err = readFile(os.Args[1])
	} else {
		in, err = readStdin()
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var data map[string]interface{}
	if err := json.Unmarshal(in, &data); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for k, v := range data {
		fmt.Printf("export %s=%s\n", k, v)
	}
}
