package oval_def_test

import (
	"testing"

	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestSanityOvalDefParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../examples/scap/oval/5.11/comment.xml")
	if err != nil {
		t.Fatal(err)
	}
	defs := doc.OvalDefinitions
	assert.NotNil(t, defs)
	assert.Equal(t, defs.XMLName.Space, "http://oval.mitre.org/XMLSchema/oval-definitions-5")
	assert.Equal(t, defs.Generator.Timestamp, "2019-01-12T10:41:00-05:00")

	definitions := defs.Definitions.Definition
	assert.Equal(t, len(definitions), 1)
	assert.Equal(t, definitions[0].Metadata.Title, "Test presence /etc/passwd")
	assert.Equal(t, definitions[0].Metadata.Description, "Test presence /etc/passwd")
}

func TestOvalDefsScapVersion(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../examples/scap/oval/5.11/comment.xml")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "5.11", doc.ScapVersion())
}
