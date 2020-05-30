package oval_sc_test

import (
	"testing"

	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestSanityOvalSyscharParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../examples/scap/oval/5.11/system-characteristics.xml")
	if err != nil {
		t.Fatal(err)
	}
	syschar := doc.OvalSystemCharacteristics
	assert.NotNil(t, syschar)
	assert.Equal(t, syschar.XMLName.Space, "http://oval.mitre.org/XMLSchema/oval-system-characteristics-5")
	assert.Equal(t, syschar.Generator.ProductName, "OVAL Definition Interpreter")
	assert.Equal(t, syschar.Generator.ProductVersion, "5.5 Build: 4")
	assert.Equal(t, syschar.Generator.Timestamp, "2009-04-10T13:04:10")

	assert.Equal(t, syschar.SystemInfo.OsName, "unknown Professional Service Pack 1")
	assert.Equal(t, syschar.SystemInfo.OsVersion, "6.0.6001")
	assert.Equal(t, syschar.SystemInfo.Architecture, "INTEL32")
	assert.Equal(t, syschar.SystemInfo.PrimaryHostName, "tarpit.g2-inc.net")

	interfaces := syschar.SystemInfo.Interfaces.Interface
	assert.Equal(t, len(interfaces), 3)
	assert.Equal(t, interfaces[0].InterfaceName, "Broadcom NetXtreme 57xx Gigabit Controller")
	assert.Equal(t, interfaces[0].IpAddress, "10.1.5.14")
	assert.Equal(t, interfaces[0].MacAddress, "00-22-19-21-75-83")

	objects := syschar.CollectedObjects.Object
	assert.Equal(t, len(objects), 223)
	assert.Equal(t, len(objects[0].Message), 1)
}
