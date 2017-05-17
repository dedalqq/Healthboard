
all: main.go json.go signal.go serviceList.go autochecker.go parameters.go webServer.go webClient.go system.go
	go build -o healthboard main.go json.go signal.go serviceList.go autochecker.go parameters.go webServer.go webClient.go system.go
