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
	items := cpeList.CpeItem
	assert.Equal(t, len(items), 7)
	item2 := items[2]
	assert.Equal(t, len(item2.References.Reference), 2)
}
