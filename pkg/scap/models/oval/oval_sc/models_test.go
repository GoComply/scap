package oval_sc_test

import (
	"testing"

	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestRFCFeedParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../../examples/scap/oval/5.11/system-characteristics.xml")
	assert.Nil(t, err)
	defs := doc.OvalSystemCharacteristics
	assert.NotNil(t, defs)
	assert.Equal(t, defs.XMLName.Space, "http://oval.mitre.org/XMLSchema/oval-system-characteristics-5")
}
