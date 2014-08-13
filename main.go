package main

import (
	"encoding/xml"
	"fmt"
	"github.com/Jwsonic/CoreStandardsTools/corestandards"
	"os"
)

func main() {

	f, err := os.Open("ela-literacy.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()

	var standards corestandards.XMLLearningStandards

	if err := xml.NewDecoder(f).Decode(&standards); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(standards.LearningStandardItems[0])
}
