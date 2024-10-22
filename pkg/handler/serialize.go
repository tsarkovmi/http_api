package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/BurntSushi/toml"

	httpapi "github.com/tsarkovmi/http_api"
)

func serializeWorker(worker httpapi.Worker) (string, error) {
	jsonCh := make(chan string)
	xmlCh := make(chan string)
	tomlCh := make(chan string)

	go func() {
		jsonData, err := json.Marshal(worker)
		if err != nil {
			close(jsonCh)
			return
		}
		jsonCh <- string(jsonData)
	}()

	go func() {
		xmlData, err := xml.Marshal(worker)
		if err != nil {
			close(xmlCh)
			return
		}
		xmlCh <- string(xmlData)
	}()
	go func() {
		tomlData, err := toml.Marshal(worker)
		if err != nil {
			close(tomlCh)
			return
		}
		tomlCh <- string(tomlData)
	}()

	jsonResult := <-jsonCh
	xmlResult := <-xmlCh
	tomlResult := <-tomlCh

	result := fmt.Sprintf("JSON:\n%s\n\nXML:\n%s\n\nTOML:\n%s\n\n", jsonResult, xmlResult, tomlResult)
	return result, nil
}
