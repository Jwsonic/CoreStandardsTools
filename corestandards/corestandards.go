package corestandards

import (
	"encoding/xml"
)

//XMLLearningStandardItemRefID docs go here
type XMLLearningStandardItemRefID struct {
	XMLName          xml.Name `xml:"LearningStandardItemRefId"`
	RelationshipType string   `xml:"RelationshipType,value"`
	Value            string   `xml:",chardata"`
}

//XMLStandardHierarchyLevel docs go here
type XMLStandardHierarchyLevel struct {
	XMLName     xml.Name `xml:"StandardHierarchyLevel"`
	Number      int      `xml:"number"`
	Description string   `xml:"description"`
}

//XMLLearningStandardItem docs go here
type XMLLearningStandardItem struct {
	XMLName                       xml.Name                       `xml:"LearningStandardItem"`
	RefID                         string                         `xml:"RefID,value"`
	RefURI                        string                         `xml:"RefURI"`
	LearningStandardDocumentRefID string                         `xml:"LearningStandardDocumentRefId"`
	StandardHierarchyLevel        XMLStandardHierarchyLevel      `xml:"StandardHierarchyLevel"`
	StatementCodes                []string                       `xml:"StatementCodes>StatementCode"`
	Statements                    []string                       `xml:"Statements>Statement"`
	GradeLevels                   []string                       `xml:"GradeLevels>GradeLevel"`
	RelatedLearningStandardItems  []XMLLearningStandardItemRefID `xml:"RelatedLearningStandardItems>LearningStandardItemRefId"`
}

//XMLLearningStandards docs go here
type XMLLearningStandards struct {
	XMLName               xml.Name                  `xml:"LearningStandards"`
	LearningStandardItems []XMLLearningStandardItem `xml:"LearningStandardItem"`
}
