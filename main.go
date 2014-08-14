package main

import (
	"encoding/xml"
	"fmt"
	"github.com/Jwsonic/CoreStandardsTools/bindata"
	"github.com/Jwsonic/CoreStandardsTools/corestandards"
	"os"
)

func main() {
	data, err := bindata.Asset("data/ela-literacy.xml")

	if len(data) == 0 || err != nil {
		fmt.Println("Asset not found!")
		os.Exit(1)
	}

	var standards corestandards.XMLLearningStandards

	if err := xml.Unmarshal(data, &standards); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, item := range standards.LearningStandardItems {
		if len(item.RelatedLearningStandardItems) == 0 {
			fmt.Println("yep")
		}

	}
}
