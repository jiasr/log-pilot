package main

import (
	"bytes"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"log-pilot/pilot"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//日志自定义格式
type LogFormatter struct{}

//格式详情
func (s *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006.01.02 15:12:12")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	//fmt.Println(entry.Data)
	msg := fmt.Sprintf("%s [%s:%d][GOID:%d][%s] %s\n", timestamp, file, len, getGID(), strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func main() {

	template := flag.String("template", "", "Template filepath for fluentd or filebeat.")
	base := flag.String("base", "", "Directory which mount host root.")
	level := flag.String("log-level", "INFO", "Log level")

	flag.Parse()

	baseDir, err := filepath.Abs(*base)
	if err != nil {
		panic(err)
	}

	if baseDir == "/" {
		baseDir = ""
	}

	if *template == "" {
		panic("template file can not be empty")
	}
	log.SetFormatter(new(LogFormatter))
	log.SetOutput(os.Stdout)
	logLevel, err := log.ParseLevel(*level)
	if err != nil {
		panic(err)
	}
	log.SetLevel(logLevel)
	log.Info(*template)
	log.Info(*base)
	log.Info(logLevel)


	b, err := ioutil.ReadFile(*template)
	if err != nil {
		panic(err)
	}

	log.Fatal(pilot.Run(string(b), baseDir))
}
