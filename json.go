package main

import (
	"encoding/json"
	"io/ioutil"
)

// JSONService -
type JSONService struct {
	Name string   `json:"name"`
	URL  []string `json:"url"`
}

func parseConfig(fileName string) ([]JSONService, error) {
	res := []JSONService{}
	data, err := ioutil.ReadFile(fileName)
	if nil != err {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
