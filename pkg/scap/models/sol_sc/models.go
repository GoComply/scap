// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#solaris
package sol_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type FacetItem struct {
	XMLName xml.Name `xml:facet_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Value *oval_sc.EntityItemBoolType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type ImageItem struct {
	XMLName xml.Name `xml:image_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Value []oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type IsainfoItem struct {
	XMLName xml.Name `xml:isainfo_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Bits *oval_sc.EntityItemIntType `xml:"bits"`

	KernelIsa *oval_sc.EntityItemStringType `xml:"kernel_isa"`

	ApplicationIsa *oval_sc.EntityItemStringType `xml:"application_isa"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type NddItem struct {
	XMLName xml.Name `xml:ndd_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Device *oval_sc.EntityItemStringType `xml:"device"`

	Instance *oval_sc.EntityItemIntType `xml:"instance"`

	Parameter *oval_sc.EntityItemStringType `xml:"parameter"`

	Value *oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PackageItem struct {
	XMLName xml.Name `xml:package_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Pkginst *oval_sc.EntityItemStringType `xml:"pkginst"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Category *oval_sc.EntityItemStringType `xml:"category"`

	Version *oval_sc.EntityItemStringType `xml:"version"`

	Vendor *oval_sc.EntityItemStringType `xml:"vendor"`

	Description *oval_sc.EntityItemStringType `xml:"description"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type Package511Item struct {
	XMLName xml.Name `xml:package511_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Publisher *oval_sc.EntityItemStringType `xml:"publisher"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Version *oval_sc.EntityItemVersionType `xml:"version"`

	Timestamp *oval_sc.EntityItemStringType `xml:"timestamp"`

	Fmri *oval_sc.EntityItemStringType `xml:"fmri"`

	Summary *oval_sc.EntityItemStringType `xml:"summary"`

	Description *oval_sc.EntityItemStringType `xml:"description"`

	Category *oval_sc.EntityItemStringType `xml:"category"`

	UpdatesAvailable *oval_sc.EntityItemBoolType `xml:"updates_available"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PackageavoidlistItem struct {
	XMLName xml.Name `xml:packageavoidlist_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Fmri *oval_sc.EntityItemStringType `xml:"fmri"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PackagecheckItem struct {
	XMLName xml.Name `xml:packagecheck_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Pkginst *oval_sc.EntityItemStringType `xml:"pkginst"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	ChecksumDiffers *oval_sc.EntityItemBoolType `xml:"checksum_differs"`

	SizeDiffers *oval_sc.EntityItemBoolType `xml:"size_differs"`

	MtimeDiffers *oval_sc.EntityItemBoolType `xml:"mtime_differs"`

	Uread *EntityItemPermissionCompareType `xml:"uread"`

	Uwrite *EntityItemPermissionCompareType `xml:"uwrite"`

	Uexec *EntityItemPermissionCompareType `xml:"uexec"`

	Gread *EntityItemPermissionCompareType `xml:"gread"`

	Gwrite *EntityItemPermissionCompareType `xml:"gwrite"`

	Gexec *EntityItemPermissionCompareType `xml:"gexec"`

	Oread *EntityItemPermissionCompareType `xml:"oread"`

	Owrite *EntityItemPermissionCompareType `xml:"owrite"`

	Oexec *EntityItemPermissionCompareType `xml:"oexec"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PackagefreezelistItem struct {
	XMLName xml.Name `xml:packagefreezelist_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Fmri *oval_sc.EntityItemStringType `xml:"fmri"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PackagepublisherItem struct {
	XMLName xml.Name `xml:packagepublisher_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Type *EntityItemPublisherTypeType `xml:"type"`

	OriginUri *oval_sc.EntityItemStringType `xml:"origin_uri"`

	Alias *oval_sc.EntityItemStringType `xml:"alias"`

	SslKey *oval_sc.EntityItemStringType `xml:"ssl_key"`

	SslCert *oval_sc.EntityItemStringType `xml:"ssl_cert"`

	ClientUuid *EntityItemClientUUIDType `xml:"client_uuid"`

	CatalogUpdated *oval_sc.EntityItemIntType `xml:"catalog_updated"`

	Enabled *oval_sc.EntityItemBoolType `xml:"enabled"`

	Order *oval_sc.EntityItemIntType `xml:"order"`

	Properties *oval_sc.EntityItemRecordType `xml:"properties"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PatchItem struct {
	XMLName xml.Name `xml:patch_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Base *oval_sc.EntityItemIntType `xml:"base"`

	Version *oval_sc.EntityItemIntType `xml:"version"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SmfItem struct {
	XMLName xml.Name `xml:smf_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Fmri *oval_sc.EntityItemStringType `xml:"fmri"`

	ServiceName *oval_sc.EntityItemStringType `xml:"service_name"`

	ServiceState *EntityItemSmfServiceStateType `xml:"service_state"`

	Protocol []oval_sc.EntityItemStringType `xml:"protocol"`

	ServerExecutable *oval_sc.EntityItemStringType `xml:"server_executable"`

	ServerArguements *oval_sc.EntityItemStringType `xml:"server_arguements"`

	ExecAsUser *oval_sc.EntityItemStringType `xml:"exec_as_user"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SmfpropertyItem struct {
	XMLName xml.Name `xml:smfproperty_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Service *oval_sc.EntityItemStringType `xml:"service"`

	Instance *oval_sc.EntityItemStringType `xml:"instance"`

	Property *oval_sc.EntityItemStringType `xml:"property"`

	Fmri *oval_sc.EntityItemStringType `xml:"fmri"`

	Value *oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type VariantItem struct {
	XMLName xml.Name `xml:variant_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Value []oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type VirtualizationinfoItem struct {
	XMLName xml.Name `xml:virtualizationinfo_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Current *oval_sc.EntityItemStringType `xml:"current"`

	Supported []EntityItemV12NEnvType `xml:"supported"`

	Parent *EntityItemV12NEnvType `xml:"parent"`

	LdomRole []EntityItemLDOMRoleType `xml:"ldom-role"`

	Properties *oval_sc.EntityItemRecordType `xml:"properties"`

	Message []oval.MessageType `xml:"message"`
}

// XSD ComplexType declarations

type EntityItemClientUUIDType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemPermissionCompareType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemPublisherTypeType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemSmfServiceStateType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemV12NEnvType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemLDOMRoleType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations
