package main

import (
	"bytes"
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	isDeflateMode = flag.Bool("d", false, "deflate file.in in file.out")
	buff          bytes.Buffer
)

func usage() {
	fmt.Printf("%s [-i] file.in file.out\n", os.Args[0])
	fmt.Printf("\tWithout flag, default mode is inflate file.in into file.out\n")
	fmt.Printf("\t-d deflate file.in in file.out\n")

}

func main() {
	flag.Parse()
	if args := flag.Args(); len(args) != 2 {
		usage()
		return
	}
	contentIn, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
	fileOut, err := os.OpenFile(flag.Arg(1), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
	if *isDeflateMode {
		w := zlib.NewWriter(fileOut)
		w.Write(contentIn)
		w.Close()
	} else {
		r, err := zlib.NewReader(bytes.NewReader(contentIn))
		if err != nil {
			fmt.Printf("error: %s", err)
			os.Exit(1)
		}
		io.Copy(fileOut, r)
	}
	fileOut.Close()

}
