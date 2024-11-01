package handler

import (
	"encoding/json"
	"encoding/xml"

	"github.com/pelletier/go-toml/v2"
	"github.com/sirupsen/logrus"
)

/*
Структура для респоса, который будет отправлен в качестве ответа
*/
type serResponse struct {
	JsonResp string `json:"JSON"`
	XmlResp  string `json:"XML"`
	TomlResp string `json:"TOML"`
}

/*
Конструктор инициализирующий сериализованный респонс
*/
func initSerResponse() *serResponse {
	return &serResponse{
		JsonResp: "",
		XmlResp:  "",
		TomlResp: "",
	}
}

/*
Метод сериализатор, обеспечивающая одновременную сериализацию результата
На вход поступает интерфейс, который обеспечивают работу с любыми входными данными

Для синхронизации используются каналы
Вызываются анонимные функции горутины, в которых по завершению закрываются каналы
В каждой горутине используется библиотечная сериализация данных и запись в канал
Для последующей передачи в сериализованный Респонс serResponse
*/
func (resp *serResponse) serializeWorker(data interface{}) error {
	jsonCh := make(chan string)
	xmlCh := make(chan string)
	tomlCh := make(chan string)

	errCh := make(chan error, 3)
	defer close(errCh)

	go func() {
		defer close(jsonCh)
		jsonData, err := json.Marshal(data)
		if err != nil {
			logrus.Printf("Ошибка сериализации JSON: %v", err)
			errCh <- err
			return
		}
		jsonCh <- string(jsonData)
		errCh <- nil
	}()

	go func() {
		defer close(xmlCh)
		xmlData, err := xml.Marshal(data)
		if err != nil {
			logrus.Printf("Ошибка сериализации XML: %v", err)
			errCh <- err
			return
		}
		xmlCh <- string(xmlData)
		errCh <- nil
	}()

	go func() {
		defer close(tomlCh)
		tomlData, err := toml.Marshal(data)
		if err != nil {
			logrus.Printf("Ошибка сериализации TOML: %v", err)
			errCh <- err
			return
		}
		tomlCh <- string(tomlData)
		errCh <- nil
	}()

	resp.JsonResp = <-jsonCh
	resp.XmlResp = <-xmlCh
	resp.TomlResp = <-tomlCh

	for i := 0; i < 3; i++ {
		if err := <-errCh; err != nil {
			return err
		}
	}

	return nil
}
