package main

import "flag"

type parameters struct {
	configFileName      string
	serviseTestInterval int64

	webServerPort int16
	webLogin      string
	webPassword   string
}

func getParameters() *parameters {
	fileName := flag.String("c", "config.json", "config file name")
	interval := flag.Int64("i", 120, "interval for auto test service in secconds")

	port := flag.Int("p", 8080, "port for web interface")
	login := flag.String("login", "admin", "login for web interface")
	password := flag.String("password", "admin", "password for web interface")

	flag.Parse()

	conf := parameters{
		configFileName:      *fileName,
		serviseTestInterval: *interval,

		webServerPort: int16(*port),
		webLogin:      *login,
		webPassword:   *password,
	}

	return &conf
}
