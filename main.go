package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"

	"github.com/Jwsonic/CoreStandardsTools/bindata"
	"github.com/Jwsonic/CoreStandardsTools/corestandards"
)

func main() {
	data, err := bindata.Asset("data/ela-literacy.xml")

	if len(data) == 0 || err != nil {
		log.Fatal("Asset not found!")
	}

	var standards corestandards.XMLLearningStandards

	if err := xml.Unmarshal(data, &standards); err != nil {
		log.Fatal(err)
	}

	minimal := []corestandards.MinimalStandard{}

	for _, item := range standards.LearningStandardItems {
		minimal = append(minimal, item.MinimalStandard())
	}

	b, err := json.Marshal(map[string]interface{}{"edata": minimal})

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("standards.json", b, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
