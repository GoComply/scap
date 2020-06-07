// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://checklists.nist.gov/xccdf/1.2
package cdf

import (
	"encoding/xml"

	"github.com/gocomply/scap/pkg/scap/models/cpe"
)

// Element
type Benchmark struct {
	XMLName xml.Name `xml:Benchmark`

	Id string `xml:"id,attr"`

	Id2 string `xml:"Id,attr"`

	Resolved string `xml:"resolved,attr"`

	Style string `xml:"style,attr"`

	StyleHref string `xml:"style-href,attr"`

	XmlLang string `xml:"lang,attr"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Title []TextType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Notice []NoticeType `xml:"notice"`

	FrontMatter []HtmlTextWithSubType `xml:"front-matter"`

	RearMatter []HtmlTextWithSubType `xml:"rear-matter"`

	Reference []ReferenceType `xml:"reference"`

	PlainText []PlainTextType `xml:"plain-text"`

	PlatformSpecification *cpe.PlatformSpecification `xml:"platform-specification"`

	Platform []CPE2IdrefType `xml:"platform"`

	Version VersionType `xml:"version"`

	Metadata []MetadataType `xml:"metadata"`

	Model []Model `xml:"model"`

	Profile []Profile `xml:"Profile"`

	Value []Value `xml:"Value"`

	TestResult []TestResult `xml:"TestResult"`

	Signature *SignatureType `xml:"signature"`

	Group []Group `xml:"Group"`

	Rule []Rule `xml:"Rule"`
}

// Element
type Status struct {
	XMLName xml.Name `xml:status`

	Date string `xml:"date,attr"`

	Text string `xml:",chardata"`
}

// Element
type Model struct {
	XMLName xml.Name `xml:model`

	System string `xml:"system,attr"`

	Param []ParamType `xml:"param"`
}

// Element
type Item struct {
	XMLName xml.Name `xml:Item`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Warning []WarningType `xml:"warning"`

	Question []TextType `xml:"question"`

	Reference []ReferenceType `xml:"reference"`

	Metadata []MetadataType `xml:"metadata"`
}

// Element
type Group struct {
	XMLName xml.Name `xml:Group`

	Value []Value `xml:"Value"`

	Signature *SignatureType `xml:"signature"`

	Group []Group `xml:"Group"`

	Rule []Rule `xml:"Rule"`
}

// Element
type Rule struct {
	XMLName xml.Name `xml:Rule`

	Ident []IdentType `xml:"ident"`

	ImpactMetric string `xml:"impact-metric"`

	ProfileNote []ProfileNoteType `xml:"profile-note"`

	Fixtext []FixTextType `xml:"fixtext"`

	Fix []FixType `xml:"fix"`

	Signature *SignatureType `xml:"signature"`

	Check []CheckType `xml:"check"`

	ComplexCheck *ComplexCheckType `xml:"complex-check"`
}

// Element
type Value struct {
	XMLName xml.Name `xml:Value`

	Match []SelStringType `xml:"match"`

	LowerBound []SelNumType `xml:"lower-bound"`

	UpperBound []SelNumType `xml:"upper-bound"`

	Choices []SelChoicesType `xml:"choices"`

	Source []UriRefType `xml:"source"`

	Signature *SignatureType `xml:"signature"`

	Value []SelStringType `xml:"value"`

	ComplexValue []SelComplexValueType `xml:"complex-value"`

	Default []SelStringType `xml:"default"`

	ComplexDefault []SelComplexValueType `xml:"complex-default"`
}

// Element
type Profile struct {
	XMLName xml.Name `xml:Profile`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Reference []ReferenceType `xml:"reference"`

	Platform []OverrideableCPE2IdrefType `xml:"platform"`

	Metadata []MetadataType `xml:"metadata"`

	Signature *SignatureType `xml:"signature"`

	Select []ProfileSelectType `xml:"select"`

	SetComplexValue []ProfileSetComplexValueType `xml:"set-complex-value"`

	SetValue []ProfileSetValueType `xml:"set-value"`

	RefineValue []ProfileRefineValueType `xml:"refine-value"`

	RefineRule []ProfileRefineRuleType `xml:"refine-rule"`
}

// Element
type TestResult struct {
	XMLName xml.Name `xml:TestResult`

	Benchmark *BenchmarkReferenceType `xml:"benchmark"`

	TailoringFile *TailoringReferenceType `xml:"tailoring-file"`

	Title []TextType `xml:"title"`

	Remark []TextType `xml:"remark"`

	Organization []string `xml:"organization"`

	Identity *IdentityType `xml:"identity"`

	Profile *IdrefType `xml:"profile"`

	Target []string `xml:"target"`

	TargetAddress []string `xml:"target-address"`

	TargetFacts *TargetFactsType `xml:"target-facts"`

	Platform []CPE2IdrefType `xml:"platform"`

	RuleResult []RuleResultType `xml:"rule-result"`

	Score []ScoreType `xml:"score"`

	Metadata []MetadataType `xml:"metadata"`

	Signature *SignatureType `xml:"signature"`

	TargetIdRef []TargetIdRefType `xml:"target-id-ref"`

	SetValue []ProfileSetValueType `xml:"set-value"`

	SetComplexValue []ProfileSetComplexValueType `xml:"set-complex-value"`
}

// Element
type Tailoring struct {
	XMLName xml.Name `xml:Tailoring`

	Benchmark *TailoringBenchmarkReferenceType `xml:"benchmark"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version TailoringVersionType `xml:"version"`

	Metadata []MetadataType `xml:"metadata"`

	Profile []Profile `xml:"Profile"`

	Signature *SignatureType `xml:"signature"`
}

// XSD ComplexType declarations

type NoticeType struct {
	Id string `xml:"id,attr"`

	XmlBase string `xml:"base,attr"`

	XmlLang string `xml:"lang,attr"`

	InnerXml string `xml:",innerxml"`
}

type DcStatusType struct {
}

type PlainTextType struct {
	Id string `xml:"id,attr"`

	Text string `xml:",chardata"`
}

type ReferenceType struct {
	Href string `xml:"href,attr"`

	Override string `xml:"override,attr"`

	InnerXml string `xml:",innerxml"`
}

type SignatureType struct {
}

type MetadataType struct {
}

type ParamType struct {
	Name string `xml:"name,attr"`

	Text string `xml:",chardata"`
}

type VersionType struct {
	Time string `xml:"time,attr"`

	Update string `xml:"update,attr"`

	Text string `xml:",chardata"`
}

type TextType struct {
	XmlLang string `xml:"lang,attr"`

	Override string `xml:"override,attr"`

	Text string `xml:",chardata"`
}

type HtmlTextType struct {
	XmlLang string `xml:"lang,attr"`

	Override string `xml:"override,attr"`

	InnerXml string `xml:",innerxml"`
}

type HtmlTextWithSubType struct {
	XmlLang string `xml:"lang,attr"`

	Override string `xml:"override,attr"`

	InnerXml string `xml:",innerxml"`
}

type ProfileNoteType struct {
	XmlLang string `xml:"lang,attr"`

	Tag string `xml:"tag,attr"`

	InnerXml string `xml:",innerxml"`
}

type TextWithSubType struct {
	XmlLang string `xml:"lang,attr"`

	Override string `xml:"override,attr"`

	Sub []SubType `xml:"sub"`

	InnerXml string `xml:",innerxml"`
}

type SubType struct {
	Use string `xml:"use,attr"`
}

type IdrefType struct {
	Idref string `xml:"idref,attr"`
}

type IdrefListType struct {
	Idref string `xml:"idref,attr"`
}

type CPE2IdrefType struct {
	Idref string `xml:"idref,attr"`
}

type OverrideableCPE2IdrefType struct {
	Override string `xml:"override,attr"`
}

type ItemType struct {
	Abstract string `xml:"abstract,attr"`

	ClusterId string `xml:"cluster-id,attr"`

	Extends string `xml:"extends,attr"`

	Hidden string `xml:"hidden,attr"`

	ProhibitChanges string `xml:"prohibitChanges,attr"`

	XmlLang string `xml:"lang,attr"`

	XmlBase string `xml:"base,attr"`

	Id string `xml:"Id,attr"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Warning []WarningType `xml:"warning"`

	Question []TextType `xml:"question"`

	Reference []ReferenceType `xml:"reference"`

	Metadata []MetadataType `xml:"metadata"`
}

type SelectableItemType struct {
	Selected string `xml:"selected,attr"`

	Weight string `xml:"weight,attr"`

	Rationale []HtmlTextWithSubType `xml:"rationale"`

	Platform []OverrideableCPE2IdrefType `xml:"platform"`

	Requires []IdrefListType `xml:"requires"`

	Conflicts []IdrefType `xml:"conflicts"`
}

type GroupType struct {
	Id string `xml:"id,attr"`

	Value []Value `xml:"Value"`

	Signature *SignatureType `xml:"signature"`

	Group []Group `xml:"Group"`

	Rule []Rule `xml:"Rule"`
}

type RuleType struct {
	Id string `xml:"id,attr"`

	Role string `xml:"role,attr"`

	Severity string `xml:"severity,attr"`

	Multiple string `xml:"multiple,attr"`

	Ident []IdentType `xml:"ident"`

	ImpactMetric string `xml:"impact-metric"`

	ProfileNote []ProfileNoteType `xml:"profile-note"`

	Fixtext []FixTextType `xml:"fixtext"`

	Fix []FixType `xml:"fix"`

	Signature *SignatureType `xml:"signature"`

	Check []CheckType `xml:"check"`

	ComplexCheck *ComplexCheckType `xml:"complex-check"`
}

type IdentType struct {
	System string `xml:"system,attr"`

	Text string `xml:",chardata"`
}

type WarningType struct {
	Category string `xml:"category,attr"`

	InnerXml string `xml:",innerxml"`
}

type FixTextType struct {
	Fixref string `xml:"fixref,attr"`

	Reboot string `xml:"reboot,attr"`

	Strategy string `xml:"strategy,attr"`

	Disruption string `xml:"disruption,attr"`

	Complexity string `xml:"complexity,attr"`

	InnerXml string `xml:",innerxml"`
}

type FixType struct {
	Id string `xml:"id,attr"`

	Reboot string `xml:"reboot,attr"`

	Strategy string `xml:"strategy,attr"`

	Disruption string `xml:"disruption,attr"`

	Complexity string `xml:"complexity,attr"`

	System string `xml:"system,attr"`

	Platform string `xml:"platform,attr"`

	InnerXml string `xml:",innerxml"`
}

type InstanceFixType struct {
	Context string `xml:"context,attr"`
}

type ComplexCheckType struct {
	Operator string `xml:"operator,attr"`

	Negate string `xml:"negate,attr"`
}

type CheckType struct {
	System string `xml:"system,attr"`

	Negate string `xml:"negate,attr"`

	Id string `xml:"id,attr"`

	Selector string `xml:"selector,attr"`

	MultiCheck string `xml:"multi-check,attr"`

	XmlBase string `xml:"base,attr"`

	CheckImport []CheckImportType `xml:"check-import"`

	CheckExport []CheckExportType `xml:"check-export"`

	CheckContentRef []CheckContentRefType `xml:"check-content-ref"`

	CheckContent *CheckContentType `xml:"check-content"`
}

type CheckImportType struct {
	ImportName string `xml:"import-name,attr"`

	ImportXpath string `xml:"import-xpath,attr"`

	InnerXml string `xml:",innerxml"`
}

type CheckExportType struct {
	ValueId string `xml:"value-id,attr"`

	ExportName string `xml:"export-name,attr"`
}

type CheckContentRefType struct {
	Href string `xml:"href,attr"`

	Name string `xml:"name,attr"`
}

type CheckContentType struct {
	InnerXml string `xml:",innerxml"`
}

type ValueType struct {
	Id string `xml:"id,attr"`

	Type string `xml:"type,attr"`

	Operator string `xml:"operator,attr"`

	Interactive string `xml:"interactive,attr"`

	InterfaceHint string `xml:"interfaceHint,attr"`

	Match []SelStringType `xml:"match"`

	LowerBound []SelNumType `xml:"lower-bound"`

	UpperBound []SelNumType `xml:"upper-bound"`

	Choices []SelChoicesType `xml:"choices"`

	Source []UriRefType `xml:"source"`

	Signature *SignatureType `xml:"signature"`

	Value []SelStringType `xml:"value"`

	ComplexValue []SelComplexValueType `xml:"complex-value"`

	Default []SelStringType `xml:"default"`

	ComplexDefault []SelComplexValueType `xml:"complex-default"`
}

type ComplexValueType struct {
	Item []string `xml:"item"`
}

type SelComplexValueType struct {
	Selector string `xml:"selector,attr"`
}

type SelChoicesType struct {
	MustMatch string `xml:"mustMatch,attr"`

	Selector string `xml:"selector,attr"`
}

type SelStringType struct {
	Selector string `xml:"selector,attr"`

	Text string `xml:",chardata"`
}

type SelNumType struct {
	Selector string `xml:"selector,attr"`

	Text string `xml:",chardata"`
}

type UriRefType struct {
	Uri string `xml:"uri,attr"`
}

type ProfileType struct {
	Id string `xml:"id,attr"`

	ProhibitChanges string `xml:"prohibitChanges,attr"`

	Abstract string `xml:"abstract,attr"`

	NoteTag string `xml:"note-tag,attr"`

	Extends string `xml:"extends,attr"`

	XmlBase string `xml:"base,attr"`

	Id2 string `xml:"Id,attr"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Reference []ReferenceType `xml:"reference"`

	Platform []OverrideableCPE2IdrefType `xml:"platform"`

	Metadata []MetadataType `xml:"metadata"`

	Signature *SignatureType `xml:"signature"`

	Select []ProfileSelectType `xml:"select"`

	SetComplexValue []ProfileSetComplexValueType `xml:"set-complex-value"`

	SetValue []ProfileSetValueType `xml:"set-value"`

	RefineValue []ProfileRefineValueType `xml:"refine-value"`

	RefineRule []ProfileRefineRuleType `xml:"refine-rule"`
}

type ProfileSelectType struct {
	Idref string `xml:"idref,attr"`

	Selected string `xml:"selected,attr"`

	Remark []TextType `xml:"remark"`
}

type ProfileSetValueType struct {
	Idref string `xml:"idref,attr"`

	Text string `xml:",chardata"`
}

type ProfileSetComplexValueType struct {
	Idref string `xml:"idref,attr"`
}

type ProfileRefineValueType struct {
	Idref string `xml:"idref,attr"`

	Selector string `xml:"selector,attr"`

	Operator string `xml:"operator,attr"`

	Remark []TextType `xml:"remark"`
}

type ProfileRefineRuleType struct {
	Idref string `xml:"idref,attr"`

	Weight string `xml:"weight,attr"`

	Selector string `xml:"selector,attr"`

	Severity string `xml:"severity,attr"`

	Role string `xml:"role,attr"`

	Remark []TextType `xml:"remark"`
}

type TestResultType struct {
	Id string `xml:"id,attr"`

	StartTime string `xml:"start-time,attr"`

	EndTime string `xml:"end-time,attr"`

	TestSystem string `xml:"test-system,attr"`

	Version string `xml:"version,attr"`

	Id2 string `xml:"Id,attr"`

	Benchmark *BenchmarkReferenceType `xml:"benchmark"`

	TailoringFile *TailoringReferenceType `xml:"tailoring-file"`

	Title []TextType `xml:"title"`

	Remark []TextType `xml:"remark"`

	Organization []string `xml:"organization"`

	Identity *IdentityType `xml:"identity"`

	Profile *IdrefType `xml:"profile"`

	Target []string `xml:"target"`

	TargetAddress []string `xml:"target-address"`

	TargetFacts *TargetFactsType `xml:"target-facts"`

	Platform []CPE2IdrefType `xml:"platform"`

	RuleResult []RuleResultType `xml:"rule-result"`

	Score []ScoreType `xml:"score"`

	Metadata []MetadataType `xml:"metadata"`

	Signature *SignatureType `xml:"signature"`

	TargetIdRef []TargetIdRefType `xml:"target-id-ref"`

	SetValue []ProfileSetValueType `xml:"set-value"`

	SetComplexValue []ProfileSetComplexValueType `xml:"set-complex-value"`
}

type BenchmarkReferenceType struct {
	Href string `xml:"href,attr"`

	Id string `xml:"id,attr"`
}

type ScoreType struct {
	System string `xml:"system,attr"`

	Maximum string `xml:"maximum,attr"`

	Text string `xml:",chardata"`
}

type TargetFactsType struct {
	Fact []FactType `xml:"fact"`
}

type TargetIdRefType struct {
	System string `xml:"system,attr"`

	Href string `xml:"href,attr"`

	Name string `xml:"name,attr"`
}

type IdentityType struct {
	Authenticated string `xml:"authenticated,attr"`

	Privileged string `xml:"privileged,attr"`

	Text string `xml:",chardata"`
}

type FactType struct {
	Name string `xml:"name,attr"`

	Type string `xml:"type,attr"`

	Text string `xml:",chardata"`
}

type TailoringReferenceType struct {
	Href string `xml:"href,attr"`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Time string `xml:"time,attr"`
}

type RuleResultType struct {
	Idref string `xml:"idref,attr"`

	Role string `xml:"role,attr"`

	Severity string `xml:"severity,attr"`

	Time string `xml:"time,attr"`

	Version string `xml:"version,attr"`

	Weight string `xml:"weight,attr"`

	Result string `xml:"result"`

	Override []OverrideType `xml:"override"`

	Ident []IdentType `xml:"ident"`

	Metadata []MetadataType `xml:"metadata"`

	Message []MessageType `xml:"message"`

	Instance []InstanceResultType `xml:"instance"`

	Fix []FixType `xml:"fix"`

	Check []CheckType `xml:"check"`

	ComplexCheck *ComplexCheckType `xml:"complex-check"`
}

type InstanceResultType struct {
	Context string `xml:"context,attr"`

	ParentContext string `xml:"parentContext,attr"`

	Text string `xml:",chardata"`
}

type OverrideType struct {
	Time string `xml:"time,attr"`

	Authority string `xml:"authority,attr"`

	OldResult string `xml:"old-result"`

	NewResult string `xml:"new-result"`

	Remark TextType `xml:"remark"`
}

type MessageType struct {
	Severity string `xml:"severity,attr"`

	Text string `xml:",chardata"`
}

type TailoringType struct {
	Id string `xml:"id,attr"`

	Id2 string `xml:"Id,attr"`

	Benchmark *TailoringBenchmarkReferenceType `xml:"benchmark"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version TailoringVersionType `xml:"version"`

	Metadata []MetadataType `xml:"metadata"`

	Profile []Profile `xml:"Profile"`

	Signature *SignatureType `xml:"signature"`
}

type TailoringBenchmarkReferenceType struct {
	Version string `xml:"version,attr"`
}

type TailoringVersionType struct {
	Time string `xml:"time,attr"`

	Text string `xml:",chardata"`
}
