// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5
package oval_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type OvalSystemCharacteristics struct {
	XMLName xml.Name `xml:oval_system_characteristics`

	Generator oval.GeneratorType `xml:"generator"`

	SystemInfo SystemInfoType `xml:"system_info"`

	CollectedObjects *CollectedObjectsType `xml:"collected_objects"`

	SystemData *SystemDataType `xml:"system_data"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`
}

// Element
type Item struct {
	XMLName xml.Name `xml:item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Message []oval.MessageType `xml:",any"`
}

// XSD ComplexType declarations

type SystemInfoType struct {
	XMLName xml.Name

	OsName string `xml:"os_name"`

	OsVersion string `xml:"os_version"`

	Architecture string `xml:"architecture"`

	PrimaryHostName string `xml:"primary_host_name"`

	Interfaces InterfacesType `xml:"interfaces"`

	InnerXml string `xml:",innerxml"`
}

type InterfacesType struct {
	XMLName xml.Name

	Interface []InterfaceType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type InterfaceType struct {
	XMLName xml.Name

	InterfaceName string `xml:"interface_name"`

	IpAddress string `xml:"ip_address"`

	MacAddress string `xml:"mac_address"`

	InnerXml string `xml:",innerxml"`
}

type CollectedObjectsType struct {
	XMLName xml.Name

	Object []ObjectType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type ObjectType struct {
	XMLName xml.Name

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	VariableInstance int `xml:"variable_instance,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Flag FlagEnumeration `xml:"flag,attr"`

	Message []oval.MessageType `xml:"message"`

	VariableValue []VariableValueType `xml:"variable_value"`

	Reference []ReferenceType `xml:"reference"`

	InnerXml string `xml:",innerxml"`
}

type VariableValueType struct {
	XMLName xml.Name

	VariableId oval.VariableIDPattern `xml:"variable_id,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type ReferenceType struct {
	XMLName xml.Name

	ItemRef oval.ItemIDPattern `xml:"item_ref,attr"`

	InnerXml string `xml:",innerxml"`
}

type SystemDataType struct {
	XMLName xml.Name

	Item []ItemType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type ItemType struct {
	XMLName xml.Name

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Message []oval.MessageType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemSimpleBaseType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type EntityItemComplexBaseType struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

type EntityItemIPAddressType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemIPAddressStringType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemAnySimpleType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemBinaryType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemBoolType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemFloatType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemIntType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemStringType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemRecordType struct {
	XMLName xml.Name

	Field []EntityItemFieldType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemFieldType struct {
	XMLName xml.Name

	Name string `xml:"name,attr"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type EntityItemVersionType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemFilesetRevisionType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemIOSVersionType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemEVRStringType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemDebianEVRStringType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations

type FlagEnumeration string

const FlagEnumerationError FlagEnumeration = "error"

const FlagEnumerationComplete FlagEnumeration = "complete"

const FlagEnumerationIncomplete FlagEnumeration = "incomplete"

const FlagEnumerationDoesNotExist FlagEnumeration = "does not exist"

const FlagEnumerationNotCollected FlagEnumeration = "not collected"

const FlagEnumerationNotApplicable FlagEnumeration = "not applicable"

type StatusEnumeration string

const StatusEnumerationError StatusEnumeration = "error"

const StatusEnumerationExists StatusEnumeration = "exists"

const StatusEnumerationDoesNotExist StatusEnumeration = "does not exist"

const StatusEnumerationNotCollected StatusEnumeration = "not collected"
