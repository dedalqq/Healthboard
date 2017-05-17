package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

var isExit chan bool

const pidFileName = "healthboard.pid"

func init() {
	isExit = make(chan bool, 1)
}

func exit() {
	sl.mutex.Lock()
	os.Remove(pidFileName)
	isExit <- true
}

func initPidFile() {
	pid := os.Getpid()
	strPid := strconv.Itoa(pid)
	ioutil.WriteFile(pidFileName, []byte(strPid), 0777)
}
