package main

import (
	"fmt"
	"os"
	"io"
	"github.com/okabe-yuya/png-analysis/types"
)

const (
	HEADER_SIZE = 8
	IHDR_SIZE = 25
	IEND_SIZE = 12
)

func main() {
	png, err := os.Open("./profile.png")
	if err != nil {
		fmt.Println("[ERROR] Failed Open file: ", err)
		os.Exit(1)
	}
	defer png.Close()
	signature, err := ReadBinary(png, HEADER_SIZE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ihdr, err := ReadBinary(png, IHDR_SIZE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileInfo, err := png.Stat()
	if err != nil {
		fmt.Println("[ERROR] Failed get file infomation: ", err)
		os.Exit(1)
	}
	fileSize := fileInfo.Size()
	offset := fileSize - (HEADER_SIZE + IHDR_SIZE + IEND_SIZE)
	idat, err := ReadBinary(png, int(offset))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	iend, err := ReadBinary(png, IEND_SIZE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := ReadBinary(png, 1); err != io.EOF {
		fmt.Println("[ERROR] Occured something error: ", err)
		os.Exit(1)
	}

	s, err := types.New(signature, ihdr, idat, iend)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(s)
}

func ReadBinary(file *os.File, x int) ([]byte, error) {
	data := make([]byte, x)
	if _, err := file.Read(data); err != nil {
		return nil, err
	}
	return data, nil
}

func debugger(header string, data []byte) {
	fmt.Printf("%s: ", header)
	for _, v := range data {
		fmt.Printf("%#x,", v)
	}
	fmt.Printf("\n")
}

// func IntTo16x(i int) {}