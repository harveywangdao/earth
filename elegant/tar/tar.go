package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func compress(files []string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	gw := gzip.NewWriter(buf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		info, _ := f.Stat()
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return nil, err
		}
		hdr.Name = hdr.Name + ".hhh"
		err = tw.WriteHeader(hdr)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(tw, f)
		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}

func main() {
	var files []string
	files = append(files, "file1.txt", "file2.txt")

	data, err := compress(files)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(data.Bytes())

	ioutil.WriteFile("a.tar.gz", data.Bytes(), os.ModePerm)
}
