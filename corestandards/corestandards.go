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

//XMLRelatedLearningStandardItems docs go here
type XMLRelatedLearningStandardItems struct {
	XMLName xml.Name                       `xml:"RelatedLearningStandardItems"`
	Items   []XMLLearningStandardItemRefID `xml:"XMLLearningStandardItemRefId"`
}

//XMLGradeLevels docs go here
type XMLGradeLevels struct {
	XMLName xml.Name `xml:"GradeLevels"`
	Levels  []string `xml:"GradeLevel"`
}

//XMLStatements docs go here
type XMLStatements struct {
	XMLName    xml.Name `xml:"Statements"`
	Statements []string `xml:"Statement"`
}

//XMLStatementCodes docs go here
type XMLStatementCodes struct {
	XMLName        xml.Name `xml:"StatementCodes"`
	StatementCodes []string `xml:"StatementCode"`
}

//XMLStandardHierarchyLevel docs go here
type XMLStandardHierarchyLevel struct {
	XMLName     xml.Name `xml:"StandardHierarchyLevel"`
	Number      int      `xml:"number"`
	Description string   `xml:"description"`
}

//XMLLearningStandardItem docs go here
type XMLLearningStandardItem struct {
	XMLName                       xml.Name                        `xml:"LearningStandardItem"`
	RefID                         string                          `xml:"RefID,value"`
	RefURI                        string                          `xml:"RefURI"`
	LearningStandardDocumentRefID string                          `xml:"LearningStandardDocumentRefId"`
	StandardHierarchyLevel        XMLStandardHierarchyLevel       `xml:"StandardHierarchyLevel"`
	StatementCodes                XMLStatementCodes               `xml:"StatementCodes"`
	Statements                    XMLStatements                   `xml:"Statements"`
	GradeLevels                   XMLGradeLevels                  `xml:"GradeLevels"`
	RelatedLearningStandardItems  XMLRelatedLearningStandardItems `xml:"RelatedLearningStandardItems"`
}

//XMLLearningStandards docs go here
type XMLLearningStandards struct {
	XMLName               xml.Name                  `xml:"LearningStandards"`
	LearningStandardItems []XMLLearningStandardItem `xml:"LearningStandardItem"`
}
