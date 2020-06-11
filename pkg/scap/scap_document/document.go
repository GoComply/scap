package scap_document

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/gocomply/scap/pkg/scap/constants"
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
	XMLName xml.Name               `json:"-"`
	Type    constants.DocumentType `json:"-"`
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
				return &Document{DataStreamCollection: &sds, Type: constants.DocumentTypeSourceDataStream}, nil
			case ovalDefinitionsElement:
				var ovalDefs oval_def.OvalDefinitions
				if err := d.DecodeElement(&ovalDefs, &startElement); err != nil {
					return nil, err
				}
				return &Document{OvalDefinitions: &ovalDefs, Type: constants.DocumentTypeOvalDefinitions}, nil
			case ovalSystemCharacteristicsElement:
				var ovalSyschar oval_sc.OvalSystemCharacteristics
				if err := d.DecodeElement(&ovalSyschar, &startElement); err != nil {
					return nil, err
				}
				return &Document{OvalSystemCharacteristics: &ovalSyschar, Type: constants.DocumentTypeOvalSyschar}, nil
			case ovalResultsElement:
				var ovalRes oval_res.OvalResults
				if err := d.DecodeElement(&ovalRes, &startElement); err != nil {
					return nil, err
				}
				return &Document{OvalResults: &ovalRes, Type: constants.DocumentTypeOvalResults}, nil
			case xccdfBenchmarkElement:
				var bench cdf.Benchmark
				if err := d.DecodeElement(&bench, &startElement); err != nil {
					return nil, err
				}
				return &Document{Benchmark: &bench, Type: constants.DocumentTypeXccdfBenchmark}, nil
			case cpeCpeListElement:
				var cpeList cpe_dict.CpeList
				if err := d.DecodeElement(&cpeList, &startElement); err != nil {
					return nil, err
				}
				return &Document{CpeList: &cpeList, Type: constants.DocumentTypeCpeDict}, nil
			case ocilOcilElement:
				var ocil inter.Ocil
				if err := d.DecodeElement(&ocil, &startElement); err != nil {
					return nil, err
				}
				return &Document{Ocil: &ocil, Type: constants.DocumentTypeOcil}, nil
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

func (doc *Document) Xmlns() string {
	switch doc.Type {
	case constants.DocumentTypeXccdfBenchmark:
		return doc.Benchmark.XMLName.Space
	case constants.DocumentTypeCpeDict:
		return doc.CpeList.XMLName.Space
	case constants.DocumentTypeOvalDefinitions:
		return doc.OvalDefinitions.XMLName.Space
	case constants.DocumentTypeOvalResults:
		return doc.OvalResults.XMLName.Space
	case constants.DocumentTypeOvalSyschar:
		return doc.OvalSystemCharacteristics.XMLName.Space
	case constants.DocumentTypeOcil:
		return doc.Ocil.XMLName.Space
	case constants.DocumentTypeSourceDataStream:
		return doc.DataStreamCollection.XMLName.Space
	}
	return ""
}

// This comment is self explanatory.
func (doc *Document) ScapVersion() string {
	switch doc.Type {
	case constants.DocumentTypeXccdfBenchmark:
		return doc.Benchmark.ScapVersion()
	case constants.DocumentTypeCpeDict:
		return doc.CpeList.ScapVersion()
	case constants.DocumentTypeOvalDefinitions:
		return doc.OvalDefinitions.ScapVersion()
	case constants.DocumentTypeOvalResults:
		return doc.OvalResults.ScapVersion()
	case constants.DocumentTypeOvalSyschar:
		return doc.OvalSystemCharacteristics.ScapVersion()
	case constants.DocumentTypeOcil:
		return doc.Ocil.ScapVersion()
	case constants.DocumentTypeSourceDataStream:
		return doc.DataStreamCollection.ScapVersion()
	}
	return ""
}
