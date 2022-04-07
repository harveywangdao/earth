package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func IpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

func IntToIp(ipInt int) string {
	ipSegs := make([]string, 4)
	var length int = len(ipSegs)
	for i := 0; i < length; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[length-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}

	buffer := bytes.NewBufferString("")
	for i := 0; i < length; i++ {
		buffer.WriteString(ipSegs[i])
		if i < length-1 {
			buffer.WriteString(".")
		}
	}

	return buffer.String()
}

func IpToInt32(ipstring string) int32 {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int32 = 0
	var pos uint32 = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | int32(tempInt)
		pos -= 8
	}
	return ipInt
}

func main() {
	fmt.Println(IpToInt32("255.255.255.255"))
	fmt.Println(IpToInt32("0.0.0.0"))
	fmt.Println(IpToInt("255.255.255.255"))
	fmt.Println(IpToInt("0.0.0.0"))
	fmt.Println(IntToIp(16819456))
}
