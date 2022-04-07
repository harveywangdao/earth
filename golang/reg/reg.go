package main

import (
	"fmt"
	"regexp"
)

func mm(s string) bool {
	pattern := `^\+?[1-9]\d*$`
	match, _ := regexp.MatchString(pattern, s)
	return match
}

func main() {
	fmt.Println(mm("hsjdksnksv1561v31s"))
	fmt.Println(mm("-1000000000"))
	fmt.Println(mm("11232.111"))
	fmt.Println(mm("0"))
	fmt.Println(mm("000011111"))
	fmt.Println(mm("++415641564"))
	fmt.Println(mm("+415641564"))
	fmt.Println(mm("415641564"))
}
