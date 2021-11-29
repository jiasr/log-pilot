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

/*
   判断文件或文件夹是否存在
   如果返回的错误为nil,说明文件或文件夹存在
   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
   如果返回的错误为其它类型,则不确定是否在存在
*/
func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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

	var docker_sock_location  = "/var/run/docker.sock"
	usedockertuntime, err := PathExists(docker_sock_location)
	log.Info(usedockertuntime)
	if err != nil {
		log.Info("PathExists(%s),err(%v)\n", docker_sock_location, err)
	}else{

	}


	if !usedockertuntime{
		*template = strings.Replace(*template,"fluentd.tpl","fluentdcd.tpl",-1)
		log.Debug(*template)
	}

	b, err := ioutil.ReadFile(*template)
	if err != nil {
		panic(err)
	}


	log.Info("usedockertuntime====>%v",usedockertuntime)
	log.Fatal(pilot.Run(string(b), baseDir,usedockertuntime))
}
