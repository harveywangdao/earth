package main

import (
	"log"
	"log/syslog"
)

// /var/log/syslog
func do1() {
	sysLog, err := syslog.Dial("", "", syslog.LOG_WARNING, "Saturday")
	if err != nil {
		log.Fatal(err)
	}

	//sysLog.Emerg("Emerg msg")
	sysLog.Alert("Alert msg")
	sysLog.Crit("Crit msg")
	sysLog.Err("Error msg")
	sysLog.Warning("Warning msg")
	sysLog.Notice("Notice msg")
	sysLog.Info("Info msg")
	sysLog.Debug("Debug msg")
}

func do2() {
	sysLog, err := syslog.Dial("unixgram", "/tmp/xiaom", syslog.LOG_WARNING, "Saturday")
	if err != nil {
		log.Fatal(err)
	}

	//sysLog.Emerg("Emerg msg")
	sysLog.Alert("Alert msg")
	sysLog.Crit("Crit msg")
	sysLog.Err("Error msg")
	sysLog.Warning("Warning msg")
	sysLog.Notice("Notice msg")
	sysLog.Info("Info msg")
	sysLog.Debug("Debug msg")
}

func main() {
	do2()
}
