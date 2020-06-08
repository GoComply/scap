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
	item0 := items[0]
	assert.Equal(t, item0.Name, "cpe:/a:1024cms:1024_cms:0.7")
	assert.Equal(t, len(item0.Title), 1)
	assert.Equal(t, item0.Title[0].Text, "1024cms.org 1024 CMS 0.7")

	item2 := items[2]
	assert.Equal(t, len(item2.References.Reference), 2)
	assert.Equal(t, item2.References.Reference[0].Href, "http://wordpress.org/plugins/xhanch-my-twitter/changelog/")
	assert.Equal(t, item2.References.Reference[0].Text, "Xhance My Twitter Plugin page on Wordpress")
}
