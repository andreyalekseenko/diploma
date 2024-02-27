package mms

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	assert "main/internal/assertations"
	"net/http"
)

func checkMMSvalid(data MMSData) bool {

	if assert.Alpha2Map[data.Country] == "" {
		return false
	} else if !assert.CheckValueInArray(data.Provider, assert.Providers[:]) {
		return false
	}
	return true
}

func GetMMS(path string) ([]MMSData, error) {
	resp, err := http.Get(path)
	if err != nil {
		log.Println(err)
		log.Println("error has occured, when http-get response sended on ", path)
	}
	if resp.StatusCode != 200 {
		log.Println("error is occured")
		return nil, errors.New("error is occured")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return getMMSStruct(body), nil

}

func getMMSStruct(body []byte) []MMSData {
	var list []MMSData
	err := json.Unmarshal(body, &list)
	if err != nil {
		log.Println(err)
	}
	for i, v := range list {
		if !checkMMSvalid(v) {
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list
}
