package main

import (
	"flag"
	"io/ioutil"

	"github.com/klauspost/compress/zstd"
)

func main() {
	inputFile := flag.String("i", "input", "Name of input zstandard dictionary file")
	outputFile := flag.String("o", "output", "Name of output raw dictionary file")
	flag.Parse()

	in, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		panic(err)
	}

	dict, err := zstd.InspectDictionary(in)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(*outputFile, dict.Content(), 0644)
	if err != nil {
		panic(err)
	}
}
