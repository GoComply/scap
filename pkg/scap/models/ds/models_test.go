package ds_test

import (
	"testing"

	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestRFCFeedParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../examples/scap/ds/1.3/sds.ds.xml")
	assert.Nil(t, err)
	ds := doc.DataStreamCollection
	assert.NotNil(t, ds)
	assert.Equal(t, ds.XMLName.Space, "http://scap.nist.gov/schema/scap/source/1.2")
	assert.Equal(t, ds.Id, "scap_org.open-scap_collection_from_xccdf_test_deriving_xccdf_result_from_oval_pass.oval.xml")
	assert.Equal(t, ds.SchematronVersion, "1.3")
}
