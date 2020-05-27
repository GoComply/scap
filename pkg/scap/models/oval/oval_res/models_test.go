package oval_res_test

import (
	"testing"

	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestRFCFeedParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../../examples/scap/oval/5.11/results.xml")
	assert.Nil(t, err)
	res := doc.OvalResults
	assert.NotNil(t, res)
	assert.Equal(t, res.XMLName.Space, "http://oval.mitre.org/XMLSchema/oval-results-5")
}
