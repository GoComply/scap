package oval_res_test

import (
	"testing"

	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestRFCFeedParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../examples/scap/oval/5.11/results.xml")
	if err != nil {
		t.Fatal(err)
	}
	res := doc.OvalResults
	assert.NotNil(t, res)
	assert.Equal(t, res.XMLName.Space, "http://oval.mitre.org/XMLSchema/oval-results-5")
	assert.Equal(t, res.Generator.ProductName, "cpe:/a:open-scap:oscap")
	assert.Equal(t, res.Generator.ProductVersion, "1.3.3")
	assert.Equal(t, res.Generator.Timestamp, "2020-05-27T19:50:48")

	definitions := res.OvalDefinitions
	assert.Equal(t, definitions.Generator.Timestamp, "2019-01-12T10:41:00-05:00")
	assert.Equal(t, len(definitions.Definitions.Definition), 1)

	systems := res.Results.System
	assert.Equal(t, len(systems), 1)
	defs := systems[0].Definitions.Definition
	assert.Equal(t, len(defs), 1)

	tests := systems[0].Tests.Test
	assert.Equal(t, len(tests), 1)
	assert.Equal(t, len(tests[0].Message), 0)
	assert.Equal(t, len(tests[0].TestedItem), 1)
	assert.Empty(t, tests[0].TestedItem[0].Message)
	assert.Equal(t, len(tests[0].TestedVariable), 0)

	syschar := systems[0].OvalSystemCharacteristics
	assert.Equal(t, syschar.Generator.ProductName, "cpe:/a:librescap:lscap")
	assert.Equal(t, syschar.Generator.Timestamp, "2020-05-27T19:50:48")

	assert.Equal(t, syschar.SystemInfo.OsName, "Fedora")
	assert.Equal(t, syschar.SystemInfo.OsVersion, "32 (Thirty Two)")
	assert.Equal(t, syschar.SystemInfo.Architecture, "x86_64")
	assert.Equal(t, syschar.SystemInfo.PrimaryHostName, "localhost.localdomain")

	interfaces := syschar.SystemInfo.Interfaces.Interface
	assert.Equal(t, len(interfaces), 2)
	assert.Equal(t, interfaces[0].InterfaceName, "lo")
	assert.Equal(t, interfaces[0].IpAddress, "127.0.0.1")
	assert.Equal(t, interfaces[0].MacAddress, "00:00:00:00:00:00")

	objects := syschar.CollectedObjects.Object
	assert.Equal(t, len(objects), 1)
	assert.Equal(t, len(objects[0].Message), 0)
	assert.Equal(t, len(objects[0].Reference), 1)
}
