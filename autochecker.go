package main

import "time"

var autocheckTicker *time.Ticker

func initAutoChecker(sl *ServiceList, param *parameters) {
	if param.serviseTestInterval == 0 {
		println("auto checker not run")
		return
	}
	autocheckTicker = time.NewTicker(time.Duration(param.serviseTestInterval) * time.Second)
	go autoChecker(sl, autocheckTicker)
}

func autoChecker(sl *ServiceList, t *time.Ticker) {
	for {
		<-t.C
		sl.runCheckServers()
	}
}
