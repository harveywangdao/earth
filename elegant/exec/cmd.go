package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	//touch -d "2020-02-12 08:00:00" a.txt
	//find /home/thomas/log/ -type f -mtime +7 -exec rm -rf {} \;
	cmd := exec.Command("find", "/home/thomas/log/", "-type", "f", "-mtime", "+7", "-exec", "rm", "-rf", "{}", ";")
	//find . -type f -mtime +80 -exec ls -l {} \;
	//cmd := exec.Command("find", ".", "-type", "f", "-mtime", "+80", "-exec", "ls", "-l", "{}", ";")
	//cmd := exec.Command("ls", "-l", "-a", "-h")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out.String())
}
