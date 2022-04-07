package main

import (
	"log"
	"log/syslog"
)

// /var/log/syslog
func main() {
	sysLog, err := syslog.Dial("", "", syslog.LOG_ERR, "Saturday")
	if err != nil {
		log.Fatal(err)
	}

	sysLog.Emerg("Hello world!")
}
