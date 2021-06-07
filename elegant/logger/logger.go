package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

var (
	levelMap = map[int]string{
		DEBUG: "DEBUG",
		INFO:  "INFO",
		WARN:  "WARN",
		ERROR: "ERROR",
		FATAL: "FATAL",
	}
)

type Logger struct {
	mu       sync.Mutex
	out      io.Writer
	level    int
	hostname string
}

func NewLogger(out io.Writer, level int) *Logger {
	_, ok := levelMap[level]
	if !ok {
		level = INFO
	}
	hn, _ := os.Hostname()
	return &Logger{
		out:      out,
		level:    level,
		hostname: hn,
	}
}

func (l *Logger) write(data []byte) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out.Write(data)
}

func (l *Logger) Output(depth, level int, msg string) {
	// 2021-06-07 11:15:19.048707 hostname DEBUG /go/src/apple/main.go:12 msg
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = "???"
		line = 0
	}

	s := fmt.Sprintf("%s %s %s %s:%d %s", time.Now().Format("2006-01-02 15:04:05.999999"), l.hostname, levelMap[level], file, line, msg)
	l.write([]byte(s))
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level > DEBUG {
		return
	}
	l.Output(2, DEBUG, fmt.Sprintln(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level > DEBUG {
		return
	}
	l.Output(2, DEBUG, fmt.Sprintf(format+"\n", v...))
}

func (l *Logger) Info(v ...interface{}) {
	if l.level > INFO {
		return
	}
	l.Output(2, INFO, fmt.Sprintln(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level > INFO {
		return
	}
	l.Output(2, INFO, fmt.Sprintf(format+"\n", v...))
}

func (l *Logger) Warn(v ...interface{}) {
	if l.level > WARN {
		return
	}
	l.Output(2, WARN, fmt.Sprintln(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level > WARN {
		return
	}
	l.Output(2, WARN, fmt.Sprintf(format+"\n", v...))
}

func (l *Logger) Error(v ...interface{}) {
	if l.level > ERROR {
		return
	}
	l.Output(2, ERROR, fmt.Sprintln(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level > ERROR {
		return
	}
	l.Output(2, ERROR, fmt.Sprintf(format+"\n", v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.level > FATAL {
		return
	}
	l.Output(2, FATAL, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.level > FATAL {
		return
	}
	l.Output(2, FATAL, fmt.Sprintf(format+"\n", v...))
	os.Exit(1)
}

type FileOutput struct {
	wr       *os.File
	logname  string
	dir      string
	date     string
	hostname string
}

func NewFileOutput(dir, logname string) *FileOutput {
	if dir == "" {
		dir = "log"
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	hn, _ := os.Hostname()
	fo := &FileOutput{
		dir:      dir,
		logname:  logname,
		hostname: hn,
	}
	return fo
}

func (f *FileOutput) rotate() {
	date := time.Now().Format("2006-01-02")
	if f.date != date {
		// hostname-name-2006-01-02.log
		logname := fmt.Sprintf("%s-%s-%s.log", f.hostname, f.logname, date)
		logpath := filepath.Join(f.dir, logname)

		file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Println(err)
			return
		}
		if f.wr != nil {
			f.wr.Close()
		}
		f.wr = file
		f.date = date
	}
}

func (f *FileOutput) Write(data []byte) (int, error) {
	f.rotate()
	return f.wr.Write(data)
}

func do1() {
	logger := NewLogger(os.Stderr, DEBUG)
	err := fmt.Errorf("ABCDEF")

	logger.Debug("abcdef", 123456, err, err, "oitcfc", "512112")
	logger.Debugf("%s %d %v", "abcdef", 123456, err)

	logger.Info("abcdef", 123456, err, err, "oitcfc", "512112")
	logger.Infof("%s %d %v", "abcdef", 123456, err)

	logger.Warn("abcdef", 123456, err, err, "oitcfc", "512112")
	logger.Warnf("%s %d %v", "abcdef", 123456, err)

	logger.Error("abcdef", 123456, err, err, "oitcfc", "512112")
	logger.Errorf("%s %d %v", "abcdef", 123456, err)

	logger.Fatal("abcdef", 123456, err, err, "oitcfc", "512112")
	logger.Fatalf("%s %d %v", "abcdef", 123456, err)

	log.SetFlags(log.Llongfile | log.LstdFlags | log.Lmicroseconds)
	log.Println("abcdef", 123456, err, err, "oitcfc", "512112")
}

func main() {
	logger := NewLogger(NewFileOutput("", "scheduler"), DEBUG)

	err := fmt.Errorf("ABCDEF")
	logger.Info("abcdef", 123456, err, err, "oitcfc", "512112")
	logger.Infof("%s %d %v", "abcdef", 123456, err)
}
