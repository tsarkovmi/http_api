package handler

import (
	"encoding/json"
	"encoding/xml"

	"github.com/pelletier/go-toml/v2"
	"github.com/sirupsen/logrus"
)

type serResponse struct {
	JsonResp string `json:"JSON"`
	XmlResp  string `json:"XML"`
	TomlResp string `json:"TOML"`
}

func initSerResponse() *serResponse {
	return &serResponse{
		JsonResp: "",
		XmlResp:  "",
		TomlResp: "",
	}
}

func (resp *serResponse) serializeWorker(data interface{}) error {
	jsonCh := make(chan string)
	xmlCh := make(chan string)
	tomlCh := make(chan string)

	go func() {
		defer close(jsonCh)
		jsonData, err := json.Marshal(data)
		if err != nil {
			logrus.Printf("Ошибка сериализации JSON: %v", err)
			return
		}
		jsonCh <- string(jsonData)
	}()

	go func() {
		defer close(xmlCh)
		xmlData, err := xml.Marshal(data)
		if err != nil {
			logrus.Printf("Ошибка сериализации XML: %v", err)
			return
		}
		xmlCh <- string(xmlData)
	}()

	go func() {
		defer close(tomlCh)
		tomlData, err := toml.Marshal(data)
		if err != nil {
			logrus.Printf("Ошибка сериализации TOML: %v", err)
			return
		}
		tomlCh <- string(tomlData)
	}()

	resp.JsonResp = <-jsonCh
	resp.XmlResp = <-xmlCh
	resp.TomlResp = <-tomlCh

	return nil
}
