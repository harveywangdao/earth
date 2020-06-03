package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	//"io/ioutil"
	"os"
)

func AddXlsxSheet(xlsxFile *xlsx.File, sheetName string, records [][]string) error {
	sheet, err := xlsxFile.AddSheet(sheetName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, record := range records {
		row := sheet.AddRow()

		for _, cell := range record {
			cl := row.AddCell()
			cl.Value = cell

			style := xlsx.NewStyle()
			style.Font = *xlsx.NewFont(10, "Arial")
			cl.SetStyle(style)
		}
	}

	return nil
}

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
		hdr.Name = "a.xlsx"
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
	//xlsx.SetDefaultFont(10, "Arial")
	xlsxFile := xlsx.NewFile()

	var records [][]string
	var record []string
	record = append(record, "aa", "bb", "cc")
	records = append(records, record)

	record = append(record[:0], "11", "22", "33")
	records = append(records, record)

	record = append(record[:0], "scs", "bsv vsb", "hngn")
	records = append(records, record)

	AddXlsxSheet(xlsxFile, "sheet1", records)
	records = append(records, record)
	AddXlsxSheet(xlsxFile, "sheet2", records)
	records = append(records, record)
	AddXlsxSheet(xlsxFile, "sheet3", records)
	records = append(records, record)
	AddXlsxSheet(xlsxFile, "sheet4", records)

	xlsxPath := "a.xlsx"
	xlsxFile.Save(xlsxPath)
	/*var files []string
	files = append(files, xlsxPath)
	buf, _ := compress(files)
	ioutil.WriteFile("a.tar.gz", buf.Bytes(), os.ModePerm)*/
}
