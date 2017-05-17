package main

var sl ServiceList

func main() {

	conf := getParameters()

	if conf.serviseTestInterval < 0 {
		println("interval must be a positive number")
		return
	}

	print("init... ")
	sl = ServiceList{}

	sl.updateConfig(conf.configFileName)

	initSignalsHandler(&sl, conf)
	initAutoChecker(&sl, conf)
	initWebServer(&sl, conf)
	initPidFile()
	println("[ ok ]")
	sl.runCheckServers()

	<-isExit
}
