package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

func main() {
	b := new(bytes.Buffer)
	cv := csv.NewWriter(b)

	var record []string
	record = append(record, "1,1")
	record = append(record, "2，2")
	record = append(record, "3,,,3")
	record = append(record, "好好")
	cv.Write(record)

	record = record[0:0]
	record = append(record, "aa")
	record = append(record, "bb")
	record = append(record, "cc")
	record = append(record, "解决")
	cv.Write(record)
	cv.Flush()

	fmt.Println(b.String())
}
