package bz2

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"io"
	"io/ioutil"
)

func Decompress(in io.Reader) (io.Reader, error) {
	data, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}
	if len(data) >= 2 && data[0] == 'B' && data[1] == 'Z' {
		fmt.Println("Decompressing bzip2")
		return bzip2.NewReader(bytes.NewReader(data)), nil
	}
	return bytes.NewReader(data), nil
}
