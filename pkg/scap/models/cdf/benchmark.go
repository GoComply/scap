package cdf

import (
	"strings"
)

const xmlnsPrefix = "http://checklists.nist.gov/xccdf/"

func (bench *Benchmark) ScapVersion() string {
	xmlns := bench.XMLName.Space
	if strings.HasPrefix(xmlns, xmlnsPrefix) {
		return strings.TrimPrefix(xmlns, xmlnsPrefix)
	}
	return ""
}
