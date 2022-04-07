package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(filepath.Dir("/aaa/bbb/ccc/ddd/ff"))
	fmt.Println(filepath.Dir("./aaa/bbb/ccc/ddd/ff"))
	fmt.Println(filepath.Join("aaa/bbb/ccc/ddd", "ff"))
	fmt.Println(filepath.Base("aaa/bbb/ccc/ddd.dd"))

	fmt.Println(filepath.Ext("aaa/bbb/ccc/ddd.dd"))
	fmt.Println(filepath.Ext("aaa/bbb/ccc/ddd."))
	fmt.Println(filepath.Ext("aaa/bbb/ccc/ddd"))

	fmt.Println(strings.TrimLeft("041341412", "0"))
	fmt.Println(strings.TrimLeft("0041341412", "0"))

	fmt.Println(cutFileName("ascasd_2.kkj"))
	fmt.Println(cutFileName("ascasd.kkj"))
	fmt.Println(cutFileName("ascasd_02.kkj"))
	fmt.Println(cutFileName("ascasd_002.kkj"))

	s := "ascasd_002.kkj"
	cutFileName(s)
	fmt.Println(s)

	s1 := "d006877067770847"

	fmt.Println(aa(s1))
	fmt.Println(s1)

	s2 := "1234"
	fmt.Println(s2[:])

	fmt.Println(getFileNumber("2.kkj"))
	fmt.Println(getFileNumber("2."))
	fmt.Println(getFileNumber("4"))
	fmt.Println(getFileNumber("4d.png"))

}

func cutFileName(filename string) string {
	pos := strings.Index(filename, "_")
	if pos == -1 {
		return "1" + filepath.Ext(filename)
	}

	name := strings.TrimLeft(filename[pos:], "_")

	return strings.TrimLeft(name, "0")
}

func aa(s1 string) string {
	s1 = strings.Replace(s1, "0", "9", -1)
	return s1
}

func getFileNumber(s string) int {
	_, file := filepath.Split(s)
	if file == "" {
		return 0
	}

	pos := strings.Index(file, ".")
	if pos != -1 {
		file = file[:pos]
	}

	num, err := strconv.Atoi(file)
	if err != nil {
		return 0
	}

	return num
}
