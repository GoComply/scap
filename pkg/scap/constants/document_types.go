package constants

type DocumentType int

const (
	DocumentTypeUknown DocumentType = iota
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
