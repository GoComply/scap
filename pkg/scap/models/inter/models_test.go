package inter_test

import (
	"testing"

	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestSanityOcilParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../examples/scap/ocil/2.0/ssg-rhel8-ocil.xml")
	if err != nil {
		t.Fatal(err)
	}
	ocil := doc.Ocil
	assert.NotNil(t, ocil)
	assert.Equal(t, ocil.XMLName.Space, "http://scap.nist.gov/schema/ocil/2.0")
	assert.Equal(t, ocil.Generator.ProductName, "xccdf-create-ocil.xslt from SCAP Security Guide")
	assert.Equal(t, ocil.Generator.ProductVersion, "ssg: 0.1.51")
	assert.Equal(t, ocil.Generator.Timestamp, "2020-05-25T17:34:46+02:00")

	questionnaires := ocil.Questionnaires.Questionnaire
	assert.Equal(t, len(questionnaires), 1121)
}
