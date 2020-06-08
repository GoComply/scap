package cdf_test

import (
	"testing"

	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestSaniryXccdfParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../examples/scap/xccdf/1.2/test_xccdf_complex_check_nand.xccdf.xml")
	if err != nil {
		t.Fatal(err)
	}
	bench := doc.Benchmark
	assert.NotNil(t, bench)
	assert.Equal(t, bench.XMLName.Space, "http://checklists.nist.gov/xccdf/1.1")
	assert.Equal(t, bench.Id, "testing-xcccdf")
	assert.Equal(t, len(bench.Status), 1)
	assert.Equal(t, bench.Status[0].Text, "incomplete")
	assert.Empty(t, bench.Title)
	assert.Empty(t, bench.Description)
	assert.Equal(t, bench.Version.Text, "2.0")
	assert.Equal(t, bench.Version.Time, "")
	assert.Equal(t, bench.Version.Update, "")
	profiles := doc.Profile
	assert.Equal(t, len(profiles), 1)
	profile := profiles[0]
	assert.Equal(t, len(profile.Title), 1)
	assert.Equal(t, profile.Title[0].InnerXml, "Enterprise XYZ Security Guidance")
	assert.Equal(t, len(profile.Status), 1)
	assert.Equal(t, profile.Status[0].Text, "draft")
	assert.Equal(t, len(profile.Title), 1)
	assert.Equal(t, len(profile.Select), 1)
	assert.Equal(t, profile.Select[0].Idref, "def-20120006")
	assert.Equal(t, profile.Select[0].Selected, "true")
	assert.Empty(t, profile.Select[0].Remark)

	assert.Equal(t, len(doc.Rule), 1)
	rule := doc.Rule[0]
	assert.Equal(t, rule.ComplexCheck.Operator, "AND")
	assert.Equal(t, rule.ComplexCheck.Negate, "1")
	assert.Empty(t, rule.ComplexCheck.ComplexCheck)
	checks := rule.ComplexCheck.Check
	assert.Equal(t, len(checks), 3)
	assert.Equal(t, checks[0].System, "http://check-engine.test/pass")
	assert.Equal(t, len(checks[0].CheckContentRef), 1)
	assert.Equal(t, checks[0].CheckContentRef[0].Href, "file")
	assert.Equal(t, checks[0].CheckContentRef[0].Name, "oval:def:20120006")

}
