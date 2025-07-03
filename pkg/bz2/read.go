package bz2

import (
	"bytes"
	"compress/bzip2"
	"io"
)

func Decompress(in io.Reader) (io.Reader, bool, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, false, err
	}
	if len(data) >= 2 && data[0] == 'B' && data[1] == 'Z' {
		return bzip2.NewReader(bytes.NewReader(data)), true, nil
	}
	return bytes.NewReader(data), false, nil
}
