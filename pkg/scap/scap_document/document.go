package scap_document

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/gocomply/scap/pkg/scap/models/cdf"
	"github.com/gocomply/scap/pkg/scap/models/cpe_dict"
	"github.com/gocomply/scap/pkg/scap/models/oval/oval_def"
)

const (
	xccdfBenchmarkElement  = "Benchmark"
	cpeCpeListElement      = "cpe-list"
	ovalDefinitionsElement = "oval_definitions"
)

type Document struct {
	XMLName xml.Name `json:"-"`
	*cdf.Benchmark
	*cpe_dict.CpeList
	*oval_def.OvalDefinitions
}

func ReadDocument(r io.Reader) (*Document, error) {
	d := xml.NewDecoder(r)
	for {
		token, err := d.Token()
		if err != nil || token == nil {
			return nil, fmt.Errorf("Could not decode XML: %v", err)
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			switch startElement.Name.Local {
			case ovalDefinitionsElement:
				var ovalDefs oval_def.OvalDefinitions
				if err := d.DecodeElement(&ovalDefs, &startElement); err != nil {
					return nil, err
				}
				return &Document{OvalDefinitions: &ovalDefs}, nil

			case xccdfBenchmarkElement:
				var bench cdf.Benchmark
				if err := d.DecodeElement(&bench, &startElement); err != nil {
					return nil, err
				}
				return &Document{Benchmark: &bench}, nil
			case cpeCpeListElement:
				var cpeList cpe_dict.CpeList
				if err := d.DecodeElement(&cpeList, &startElement); err != nil {
					return nil, err
				}
				return &Document{CpeList: &cpeList}, nil
			}
		}
	}
	return nil, fmt.Errorf("Could not parse input file")
}

func ReadDocumentFromFile(filepath string) (*Document, error) {
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return ReadDocument(reader)
}
