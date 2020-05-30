package cpe_dict_test

import (
	"testing"

	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestSimpleParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../examples/scap/cpe/2.3/official-cpe-dictionary_v2.3.xml")
	if err != nil {
		t.Fatal(err)
	}
	cpeList := doc.CpeList
	assert.NotNil(t, cpeList)
	assert.Equal(t, cpeList.XMLName.Space, "http://cpe.mitre.org/dictionary/2.0")
}
