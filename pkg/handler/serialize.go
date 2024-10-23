package handler

import (
	"encoding/json"
	"encoding/xml"

	"github.com/BurntSushi/toml"

	httpapi "github.com/tsarkovmi/http_api"
)

type serResponse struct {
	JsonResp []string `json:"JSON"`
	XmlResp  []string `json:"XML"`
	TomlResp []string `json:"TOML"`
}

func initSerResponse() *serResponse {
	return &serResponse{
		JsonResp: []string{},
		XmlResp:  []string{},
		TomlResp: []string{},
	}
}

func (resp *serResponse) serializeWorker(worker httpapi.Worker) error {

	jsonCh := make(chan string)
	xmlCh := make(chan string)
	tomlCh := make(chan string)

	go func() error {
		jsonData, err := json.Marshal(worker)
		if err != nil {
			close(jsonCh)
			return err
		}
		jsonCh <- string(jsonData)
		return nil
	}()

	go func() error {
		xmlData, err := xml.Marshal(worker)
		if err != nil {
			close(xmlCh)
			return err
		}
		xmlCh <- string(xmlData)
		return nil
	}()
	go func() error {
		tomlData, err := toml.Marshal(worker)
		if err != nil {
			close(tomlCh)
			return err
		}
		tomlCh <- string(tomlData)
		return nil
	}()

	resp.JsonResp = append(resp.JsonResp, <-jsonCh)
	resp.XmlResp = append(resp.XmlResp, <-xmlCh)
	resp.TomlResp = append(resp.TomlResp, <-tomlCh)

	return nil
}
