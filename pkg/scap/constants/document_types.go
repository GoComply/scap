package constants

type DocumentType int

const (
	DocumentTypeUnknown DocumentType = iota
	DocumentTypeCpeDict
	DocumentTypeOcil
	DocumentTypeOvalDefinitions
	DocumentTypeOvalSyschar
	DocumentTypeOvalResults
	DocumentTypeSourceDataStream
	DocumentTypeXccdfBenchmark
	DocumentTypeXccdfTailoring
)

func (typ DocumentType) String() string {
	switch typ {
	case DocumentTypeCpeDict:
		return "CPE Dictionary"
	case DocumentTypeOcil:
		return "OCIL Questionaire"
	case DocumentTypeOvalDefinitions:
		return "OVAL Definitions"
	case DocumentTypeOvalSyschar:
		return "OVAL System Characteristics"
	case DocumentTypeOvalResults:
		return "OVAL Results"
	case DocumentTypeSourceDataStream:
		return "SCAP Source DataStream"
	case DocumentTypeXccdfBenchmark:
		return "XCCDF Benchmark"
	case DocumentTypeXccdfTailoring:
		return "XCCDF Tailoring"
	case DocumentTypeUnknown:
		fallthrough
	default:
		return "Uknown Document"
	}
}

func (typ DocumentType) ShortName() string {
	switch typ {
	case DocumentTypeCpeDict:
		return "cpe"
	case DocumentTypeOcil:
		return "ocil"
	case DocumentTypeOvalDefinitions:
		fallthrough
	case DocumentTypeOvalSyschar:
		fallthrough
	case DocumentTypeOvalResults:
		return "oval"
	case DocumentTypeSourceDataStream:
		return "sds"
	case DocumentTypeXccdfBenchmark:
		fallthrough
	case DocumentTypeXccdfTailoring:
		return "xccdf"
	case DocumentTypeUnknown:
		fallthrough
	default:
		return "unknown"
	}
}
