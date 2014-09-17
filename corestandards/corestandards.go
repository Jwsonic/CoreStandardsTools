package corestandards

import (
	"encoding/xml"
	"strings"

	"github.com/kennygrant/sanitize"
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

//MinimalStandard is a cleaned up standard version for export
type MinimalStandard struct {
	Code   string   `json:"code"`
	Grades []string `json:"grades"`
	Text   string   `json:"text"`
}

//StandardType returns the standard type for the given standard
func (item XMLLearningStandardItem) StandardType() string {

	if len(item.StatementCodes) == 0 {
		return ""
	}

	parts := strings.Split(item.StatementCodes[0], ".")

	if len(parts) < 3 {
		return ""
	}

	return parts[2]
}

//MinimalStandard returns a cleaned up standard
func (item XMLLearningStandardItem) MinimalStandard() MinimalStandard {
	standard := MinimalStandard{}

	if len(item.StatementCodes) > 0 {
		standard.Code = item.StatementCodes[0]
	}
	standard.Grades = item.GradeLevels
	if len(item.Statements) > 0 {
		standard.Text = sanitize.HTML(item.Statements[0])
	}

	return standard
}
