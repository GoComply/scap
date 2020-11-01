// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-common-5
package oval

import (
	"encoding/xml"
)

// Element
type DeprecatedInfo struct {
	XMLName xml.Name `xml:deprecated_info`

	Version string `xml:"version"`

	Reason string `xml:"reason"`

	Comment string `xml:"comment"`
}

// Element
type ElementMapping struct {
	XMLName xml.Name `xml:element_mapping`

	Test ElementMapItemType `xml:"test"`

	Object *ElementMapItemType `xml:"object"`

	State *ElementMapItemType `xml:"state"`

	Item *ElementMapItemType `xml:"item"`
}

// Element
type Notes struct {
	XMLName xml.Name `xml:notes`

	Note []string `xml:"note"`
}

// XSD ComplexType declarations

type ElementMapType struct {
	Test ElementMapItemType `xml:"test"`

	Object *ElementMapItemType `xml:"object"`

	State *ElementMapItemType `xml:"state"`

	Item *ElementMapItemType `xml:"item"`

	InnerXml string `xml:",innerxml"`
}

type ElementMapItemType struct {
	TargetNamespace string `xml:"target_namespace,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type DeprecatedInfoType struct {
	Version string `xml:"version"`

	Reason string `xml:"reason"`

	Comment string `xml:"comment"`

	InnerXml string `xml:",innerxml"`
}

type GeneratorType struct {
	ProductName string `xml:"product_name"`

	ProductVersion string `xml:"product_version"`

	SchemaVersion []SchemaVersionType `xml:"schema_version"`

	Timestamp string `xml:"timestamp"`

	InnerXml string `xml:",innerxml"`
}

type SchemaVersionType struct {
	Platform string `xml:"platform,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type MessageType struct {
	Level string `xml:"level,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type NotesType struct {
	Note []string `xml:"note"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations

type CheckEnumeration string

const CheckEnumerationAll CheckEnumeration = "all"

const CheckEnumerationAtLeastOne CheckEnumeration = "at least one"

const CheckEnumerationNoneExist CheckEnumeration = "none exist"

const CheckEnumerationNoneSatisfy CheckEnumeration = "none satisfy"

const CheckEnumerationOnlyOne CheckEnumeration = "only one"

type ClassEnumeration string

const ClassEnumerationCompliance ClassEnumeration = "compliance"

const ClassEnumerationInventory ClassEnumeration = "inventory"

const ClassEnumerationMiscellaneous ClassEnumeration = "miscellaneous"

const ClassEnumerationPatch ClassEnumeration = "patch"

const ClassEnumerationVulnerability ClassEnumeration = "vulnerability"

type SimpleDatatypeEnumeration string

const SimpleDatatypeEnumerationBinary SimpleDatatypeEnumeration = "binary"

const SimpleDatatypeEnumerationBoolean SimpleDatatypeEnumeration = "boolean"

const SimpleDatatypeEnumerationEvrString SimpleDatatypeEnumeration = "evr_string"

const SimpleDatatypeEnumerationDebianEvrString SimpleDatatypeEnumeration = "debian_evr_string"

const SimpleDatatypeEnumerationFilesetRevision SimpleDatatypeEnumeration = "fileset_revision"

const SimpleDatatypeEnumerationFloat SimpleDatatypeEnumeration = "float"

const SimpleDatatypeEnumerationIosVersion SimpleDatatypeEnumeration = "ios_version"

const SimpleDatatypeEnumerationInt SimpleDatatypeEnumeration = "int"

const SimpleDatatypeEnumerationIpv4Address SimpleDatatypeEnumeration = "ipv4_address"

const SimpleDatatypeEnumerationIpv6Address SimpleDatatypeEnumeration = "ipv6_address"

const SimpleDatatypeEnumerationString SimpleDatatypeEnumeration = "string"

const SimpleDatatypeEnumerationVersion SimpleDatatypeEnumeration = "version"

type ComplexDatatypeEnumeration string

const ComplexDatatypeEnumerationRecord ComplexDatatypeEnumeration = "record"

type DatatypeEnumeration string

type ExistenceEnumeration string

const ExistenceEnumerationAllExist ExistenceEnumeration = "all_exist"

const ExistenceEnumerationAnyExist ExistenceEnumeration = "any_exist"

const ExistenceEnumerationAtLeastOneExists ExistenceEnumeration = "at_least_one_exists"

const ExistenceEnumerationNoneExist ExistenceEnumeration = "none_exist"

const ExistenceEnumerationOnlyOneExists ExistenceEnumeration = "only_one_exists"

type FamilyEnumeration string

const FamilyEnumerationAndroid FamilyEnumeration = "android"

const FamilyEnumerationAsa FamilyEnumeration = "asa"

const FamilyEnumerationAppleIos FamilyEnumeration = "apple_ios"

const FamilyEnumerationCatos FamilyEnumeration = "catos"

const FamilyEnumerationIos FamilyEnumeration = "ios"

const FamilyEnumerationIosxe FamilyEnumeration = "iosxe"

const FamilyEnumerationJunos FamilyEnumeration = "junos"

const FamilyEnumerationMacos FamilyEnumeration = "macos"

const FamilyEnumerationPixos FamilyEnumeration = "pixos"

const FamilyEnumerationUndefined FamilyEnumeration = "undefined"

const FamilyEnumerationUnix FamilyEnumeration = "unix"

const FamilyEnumerationVmwareInfrastructure FamilyEnumeration = "vmware_infrastructure"

const FamilyEnumerationWindows FamilyEnumeration = "windows"

type MessageLevelEnumeration string

const MessageLevelEnumerationDebug MessageLevelEnumeration = "debug"

const MessageLevelEnumerationError MessageLevelEnumeration = "error"

const MessageLevelEnumerationFatal MessageLevelEnumeration = "fatal"

const MessageLevelEnumerationInfo MessageLevelEnumeration = "info"

const MessageLevelEnumerationWarning MessageLevelEnumeration = "warning"

type OperationEnumeration string

const OperationEnumerationEquals OperationEnumeration = "equals"

const OperationEnumerationNotEqual OperationEnumeration = "not equal"

const OperationEnumerationCaseInsensitiveEquals OperationEnumeration = "case insensitive equals"

const OperationEnumerationCaseInsensitiveNotEqual OperationEnumeration = "case insensitive not equal"

const OperationEnumerationGreaterThan OperationEnumeration = "greater than"

const OperationEnumerationLessThan OperationEnumeration = "less than"

const OperationEnumerationGreaterThanOrEqual OperationEnumeration = "greater than or equal"

const OperationEnumerationLessThanOrEqual OperationEnumeration = "less than or equal"

const OperationEnumerationBitwiseAnd OperationEnumeration = "bitwise and"

const OperationEnumerationBitwiseOr OperationEnumeration = "bitwise or"

const OperationEnumerationPatternMatch OperationEnumeration = "pattern match"

const OperationEnumerationSubsetOf OperationEnumeration = "subset of"

const OperationEnumerationSupersetOf OperationEnumeration = "superset of"

type OperatorEnumeration string

const OperatorEnumerationAnd OperatorEnumeration = "AND"

const OperatorEnumerationOne OperatorEnumeration = "ONE"

const OperatorEnumerationOr OperatorEnumeration = "OR"

const OperatorEnumerationXor OperatorEnumeration = "XOR"

type DefinitionIDPattern string

type ObjectIDPattern string

type StateIDPattern string

type TestIDPattern string

type VariableIDPattern string

type ItemIDPattern string

type SchemaVersionPattern string

type EmptyStringType string

type NonEmptyStringType string
