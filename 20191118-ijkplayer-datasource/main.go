package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	cck    int
	input  string
	output string
)

// check file exist and is not directory
func exist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir() && err == nil
}

func main() {
	flag.IntVar(&cck, "cck", 10, "key of Caesar cipher")
	flag.StringVar(&input, "i", "bee.mp4", "input file name")
	flag.StringVar(&output, "o", "bee-cc.mp4", "output file name")
	flag.Parse()

	if !exist(input) {
		fmt.Printf("input file: %s not found\n", input)
		flag.PrintDefaults()
		os.Exit(-1)
	}

	if input == output {
		fmt.Println("output file must not be same as input file")
		flag.PrintDefaults()
		os.Exit(-1)
	}
	if exist(output) {
		fmt.Printf("output file: %s already exists\n", output)
		flag.PrintDefaults()
		os.Exit(-1)
	}

	ifile, err := os.Open(input)
	defer ifile.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	ibuf := bufio.NewReader(ifile)

	ofile, err := os.Create(output)
	defer ofile.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	obuf := bufio.NewWriter(ofile)

	buf := make([]byte, 128)
	for {
		n, err := ibuf.Read(buf)
		if err == nil {
			for index := 0; index < n; index++ {
				fmt.Printf("buf %d %d\n", int(buf[index]), int(buf[index]+byte(cck)))
				buf[index] = buf[index] + byte(cck)
			}
			nw, err := obuf.Write(buf[:n])
			if err != nil || nw != n {
				fmt.Println(err.Error())
				fmt.Println("Error occurs during writing file")
				os.Exit(-1)
			}
		} else if err == io.EOF {
			break
		}
	}
	obuf.Flush()
}
