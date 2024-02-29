package daysatu

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type User struct {
	ID   string `json:"id,omitempty" xml:"id"`
	Name string `json:"name,omitempty" xml:"name"`
}

func FormatJSON(data interface{}) ([]byte, error) {
	// Menggunakan encoding/json untuk mengubah data menjadi JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	fmt.Print("apa kabard:")
	return jsonData, nil
}

func FormatYaml(data interface{}) ([]byte, error) {
	xmlData, err := xml.MarshalIndent(data, "", " ")
	if err != nil {
		return nil, err
	}
	xmlData = append([]byte(xml.Header), xmlData...)
	return xmlData, nil
}
