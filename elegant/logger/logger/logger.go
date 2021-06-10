package logger

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/harveywangdao/ants/util"
	"google.golang.org/grpc/metadata"
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

	levelReverseMap = map[string]int{
		"DEBUG": DEBUG,
		"INFO":  INFO,
		"WARN":  WARN,
		"ERROR": ERROR,
		"FATAL": FATAL,
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

func (l *Logger) Output(depth, level int, ctxStr, msg string) {
	if l.level > level {
		return
	}

	// 2021-06-07 11:15:19.048707 hostname DEBUG /go/src/apple/main.go:12 msg
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = "???"
		line = 0
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString(time.Now().Format("2006-01-02 15:04:05.999999"))
	buf.WriteByte(' ')

	buf.WriteString(l.hostname)
	buf.WriteByte(' ')

	buf.WriteString(levelMap[level])
	buf.WriteByte(' ')

	buf.WriteString(file)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))
	buf.WriteByte(' ')

	if ctxStr != "" {
		buf.WriteString(ctxStr)
		buf.WriteByte(' ')
	}

	buf.WriteString(msg)
	if msg[len(msg)-1] != '\n' {
		buf.WriteByte('\n')
	}

	l.write(buf.Bytes())
}

func (l *Logger) Debug(v ...interface{}) {
	l.Output(2, DEBUG, "", fmt.Sprintln(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(2, DEBUG, "", fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.Output(2, INFO, "", fmt.Sprintln(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(2, INFO, "", fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.Output(2, WARN, "", fmt.Sprintln(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(2, WARN, "", fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.Output(2, ERROR, "", fmt.Sprintln(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(2, ERROR, "", fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Output(2, FATAL, "", fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(2, FATAL, "", fmt.Sprintf(format, v...))
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
	if logname == "" {
		logname = filepath.Base(os.Args[0])
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

func do2() {
	logger := NewLogger(NewFileOutput("", "scheduler"), DEBUG)
	err := fmt.Errorf("ABCDEF")

	logger.Info("abcdef", 123456, err, err, "oitcfc", "512112")
	logger.Infof("%s %d %v", "abcdef", 123456, err)
}

var (
	std *Logger
)

func init() {
	level, ok := levelReverseMap[strings.ToUpper(os.Getenv("LOG_LEVEL"))]
	if !ok {
		level = INFO
	}

	if os.Getenv("LOG_TYPE") == "file" {
		std = NewLogger(NewFileOutput(os.Getenv("LOG_DIR"), os.Getenv("LOG_NAME")), level)
	} else {
		std = NewLogger(os.Stderr, level)
	}
}

func Debug(v ...interface{}) {
	std.Output(2, DEBUG, "", fmt.Sprintln(v...))
}

func Debugf(format string, v ...interface{}) {
	std.Output(2, DEBUG, "", fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	std.Output(2, INFO, "", fmt.Sprintln(v...))
}

func Infof(format string, v ...interface{}) {
	std.Output(2, INFO, "", fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	std.Output(2, WARN, "", fmt.Sprintln(v...))
}

func Warnf(format string, v ...interface{}) {
	std.Output(2, WARN, "", fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	std.Output(2, ERROR, "", fmt.Sprintln(v...))
}

func Errorf(format string, v ...interface{}) {
	std.Output(2, ERROR, "", fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	std.Output(2, FATAL, "", fmt.Sprintln(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	std.Output(2, FATAL, "", fmt.Sprintf(format, v...))
	os.Exit(1)
}

const (
	CtxKey = "ctx-key"

	XTraceId = "x-trace-id"
	XUserId  = "x-user-id"
)

type LoggerWithContext struct {
	md metadata.MD
}

func CreateLogContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		out := metadata.MD{}
		for k, v := range md {
			if strings.HasPrefix(k, "x-") {
				out[k] = append(out[k], v...)
			}
		}
		md = out
		if vs := md.Get(XTraceId); len(vs) == 0 {
			md.Set(XTraceId, util.GetUUID())
		}
	} else {
		md = metadata.Pairs(XTraceId, util.GetUUID())
	}

	return context.WithValue(ctx, CtxKey, md)
}

func CreateOutgoingContext(ctx context.Context) context.Context {
	md, ok := ctx.Value(CtxKey).(metadata.MD)
	if !ok {
		return metadata.NewOutgoingContext(ctx, metadata.MD{})
	}
	return metadata.NewOutgoingContext(ctx, md.Copy())
}

func With(ctx context.Context) *LoggerWithContext {
	lwc := &LoggerWithContext{}
	md, ok := ctx.Value(CtxKey).(metadata.MD)
	if !ok {
		return lwc
	}
	lwc.md = md.Copy()
	return lwc
}

func (l *LoggerWithContext) kvstr() string {
	// key1: value1 key2: value2
	buf := bytes.NewBuffer(nil)
	for k, vs := range l.md {
		if len(vs) > 0 && strings.HasPrefix(k, "x-") {
			if buf.Len() > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('[')
			buf.WriteString(k)
			buf.WriteByte(':')
			buf.WriteString(vs[0])
			buf.WriteByte(']')
		}
	}
	return buf.String()
}

func (l *LoggerWithContext) Debug(v ...interface{}) {
	std.Output(2, DEBUG, l.kvstr(), fmt.Sprintln(v...))
}

func (l *LoggerWithContext) Debugf(format string, v ...interface{}) {
	std.Output(2, DEBUG, l.kvstr(), fmt.Sprintf(format, v...))
}

func (l *LoggerWithContext) Info(v ...interface{}) {
	std.Output(2, INFO, l.kvstr(), fmt.Sprintln(v...))
}

func (l *LoggerWithContext) Infof(format string, v ...interface{}) {
	std.Output(2, INFO, l.kvstr(), fmt.Sprintf(format, v...))
}

func (l *LoggerWithContext) Warn(v ...interface{}) {
	std.Output(2, WARN, l.kvstr(), fmt.Sprintln(v...))
}

func (l *LoggerWithContext) Warnf(format string, v ...interface{}) {
	std.Output(2, WARN, l.kvstr(), fmt.Sprintf(format, v...))
}

func (l *LoggerWithContext) Error(v ...interface{}) {
	std.Output(2, ERROR, l.kvstr(), fmt.Sprintln(v...))
}

func (l *LoggerWithContext) Errorf(format string, v ...interface{}) {
	std.Output(2, ERROR, l.kvstr(), fmt.Sprintf(format, v...))
}

func (l *LoggerWithContext) Fatal(v ...interface{}) {
	std.Output(2, FATAL, l.kvstr(), fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *LoggerWithContext) Fatalf(format string, v ...interface{}) {
	std.Output(2, FATAL, l.kvstr(), fmt.Sprintf(format, v...))
	os.Exit(1)
}
