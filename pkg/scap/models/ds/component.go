package ds

import (
	"github.com/gocomply/scap/pkg/scap/constants"
)

func (c *Component) DocumentType() constants.DocumentType {
	if c.Benchmark != nil {
		return constants.DocumentTypeXccdfBenchmark
	} else if c.OvalDefinitions != nil {
		return constants.DocumentTypeOvalDefinitions
	} else if c.Ocil != nil {
		return constants.DocumentTypeOcil
	} else if c.CpeList != nil {
		return constants.DocumentTypeCpeDict
	} else if c.Tailoring != nil {
		return constants.DocumentTypeXccdfTailoring
	}
	return constants.DocumentTypeUnknown
}
