package scap_document

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/gocomply/scap/pkg/scap/models/cdf"
	"github.com/gocomply/scap/pkg/scap/models/cpe_dict"
	"github.com/gocomply/scap/pkg/scap/models/ds"
	"github.com/gocomply/scap/pkg/scap/models/inter"
	"github.com/gocomply/scap/pkg/scap/models/oval_def"
	"github.com/gocomply/scap/pkg/scap/models/oval_res"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

const (
	xccdfBenchmarkElement            = "Benchmark"
	cpeCpeListElement                = "cpe-list"
	ovalDefinitionsElement           = "oval_definitions"
	ovalResultsElement               = "oval_results"
	ovalSystemCharacteristicsElement = "oval_system_characteristics"
	dsDataStreamCollectionElement    = "data-stream-collection"
	ocilOcilElement                  = "ocil"
)

type Document struct {
	XMLName xml.Name `json:"-"`
	*cdf.Benchmark
	*cpe_dict.CpeList
	*oval_def.OvalDefinitions
	*oval_res.OvalResults
	*oval_sc.OvalSystemCharacteristics
	*ds.DataStreamCollection
	*inter.Ocil
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
			case dsDataStreamCollectionElement:
				var sds ds.DataStreamCollection
				if err := d.DecodeElement(&sds, &startElement); err != nil {
					return nil, err
				}
				return &Document{DataStreamCollection: &sds}, nil
			case ovalDefinitionsElement:
				var ovalDefs oval_def.OvalDefinitions
				if err := d.DecodeElement(&ovalDefs, &startElement); err != nil {
					return nil, err
				}
				return &Document{OvalDefinitions: &ovalDefs}, nil
			case ovalSystemCharacteristicsElement:
				var ovalSyschar oval_sc.OvalSystemCharacteristics
				if err := d.DecodeElement(&ovalSyschar, &startElement); err != nil {
					return nil, err
				}
				return &Document{OvalSystemCharacteristics: &ovalSyschar}, nil
			case ovalResultsElement:
				var ovalRes oval_res.OvalResults
				if err := d.DecodeElement(&ovalRes, &startElement); err != nil {
					return nil, err
				}
				return &Document{OvalResults: &ovalRes}, nil
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
			case ocilOcilElement:
				var ocil inter.Ocil
				if err := d.DecodeElement(&ocil, &startElement); err != nil {
					return nil, err
				}
				return &Document{Ocil: &ocil}, nil
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
