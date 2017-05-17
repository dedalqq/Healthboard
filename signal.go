package main

import (
	"os"
	"os/signal"
	"syscall"
)

var sigsChan chan os.Signal

func initSignalsHandler(sl *ServiceList, param *parameters) {
	sigsChan = make(chan os.Signal, 1)
	signal.Notify(sigsChan, syscall.SIGHUP, syscall.SIGINT)
	go signalsHandler(sigsChan, sl, param)
}

func signalsHandler(sh chan os.Signal, sl *ServiceList, param *parameters) {
	for {
		sig := <-sh
		if sig == syscall.SIGHUP {
			println("update config")
			sl.updateConfig(param.configFileName)
		}
		if sig == syscall.SIGINT {
			println("shutdown")
			exit()
		}
	}
}
