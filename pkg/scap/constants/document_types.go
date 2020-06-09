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
	case DocumentTypeUknown:
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
		return "xccdf"
	case DocumentTypeUknown:
		fallthrough
	default:
		return "unknown"
	}
}
