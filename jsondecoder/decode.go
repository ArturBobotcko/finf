package jsondecoder

import (
	"embed"
	"encoding/json"
	"io"
	"log"
)

//go:embed extensions.json
var jsonData embed.FS

// FileInfo represents an entry in extensions.json file,
// which contains all files extensions in JSON format
type FileInfo struct {
	Descriptions []string `json:"desriptions"`
}

// Unmarsharls extensions.json using json.Unmarshal()
func unmarshalJson(data []byte) interface{} {
	var unmarshaledData interface{}

	err := json.Unmarshal(data, &unmarshaledData)
	if err != nil {
		log.Fatal(err)
	}

	return unmarshaledData
}

// convert unmarshaled extensions.json into map,
// which contains description of all file extensions
func decodeFileExtensionsJson(data []byte) map[string]FileInfo {
	unmarshaledData := unmarshalJson(data)

	rawData, ok := unmarshaledData.(map[string]interface{})
	if !ok {
		log.Fatal("Failed to convert JSON data to map[string]interface{}")
	}

	dataMap := make(map[string]FileInfo)

	for ext, info := range rawData {
		infoMap, ok := info.(map[string]interface{})
		if !ok {
			log.Printf("Skipping invalid entry for extension %q\n", ext)
			continue
		}

		descriptionsRaw, ok := infoMap["descriptions"].([]interface{})
		if !ok {
			log.Printf("No descriptions available for extension %q\n", ext)
			continue
		}

		descriptions := make([]string, len(descriptionsRaw))
		for i, desc := range descriptionsRaw {
			if descStr, ok := desc.(string); ok {
				descriptions[i] = descStr
			} else {
				log.Printf("Skipping invalid description for extension %q\n", ext)
			}
		}

		dataMap[ext] = FileInfo{Descriptions: descriptions}
	}

	return dataMap
}

// returns a map of all file extensions with their descriptions
func LoadFileExtensions() map[string]FileInfo {
	file, err := jsonData.Open("extensions.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return decodeFileExtensionsJson(data)
}
