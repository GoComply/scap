// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
package sol_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_def"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type FacetTest struct {
	XMLName xml.Name `xml:facet_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FacetObject struct {
	XMLName xml.Name `xml:facet_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Path *oval_def.EntityObjectStringType `xml:"path"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FacetState struct {
	XMLName xml.Name `xml:facet_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Value *oval_def.EntityStateBoolType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ImageTest struct {
	XMLName xml.Name `xml:image_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ImageObject struct {
	XMLName xml.Name `xml:image_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Path *oval_def.EntityObjectStringType `xml:"path"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ImageState struct {
	XMLName xml.Name `xml:image_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type IsainfoTest struct {
	XMLName xml.Name `xml:isainfo_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type IsainfoObject struct {
	XMLName xml.Name `xml:isainfo_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type IsainfoState struct {
	XMLName xml.Name `xml:isainfo_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Bits *oval_def.EntityStateIntType `xml:"bits"`

	KernelIsa *oval_def.EntityStateStringType `xml:"kernel_isa"`

	ApplicationIsa *oval_def.EntityStateStringType `xml:"application_isa"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type NddTest struct {
	XMLName xml.Name `xml:ndd_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type NddObject struct {
	XMLName xml.Name `xml:ndd_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Device *oval_def.EntityObjectStringType `xml:"device"`

	Parameter *oval_def.EntityObjectStringType `xml:"parameter"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type NddState struct {
	XMLName xml.Name `xml:ndd_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Device *oval_def.EntityStateStringType `xml:"device"`

	Instance *oval_def.EntityStateIntType `xml:"instance"`

	Parameter *oval_def.EntityStateStringType `xml:"parameter"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackageTest struct {
	XMLName xml.Name `xml:package_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackageObject struct {
	XMLName xml.Name `xml:package_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Pkginst *oval_def.EntityObjectStringType `xml:"pkginst"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackageState struct {
	XMLName xml.Name `xml:package_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Pkginst *oval_def.EntityStateStringType `xml:"pkginst"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Category *oval_def.EntityStateStringType `xml:"category"`

	VersionElm *oval_def.EntityStateStringType `xml:"version"`

	Vendor *oval_def.EntityStateStringType `xml:"vendor"`

	Description *oval_def.EntityStateStringType `xml:"description"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Package511Test struct {
	XMLName xml.Name `xml:package511_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Package511Object struct {
	XMLName xml.Name `xml:package511_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Publisher *oval_def.EntityObjectStringType `xml:"publisher"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	VersionElm *oval_def.EntityObjectVersionType `xml:"version"`

	Timestamp *oval_def.EntityObjectStringType `xml:"timestamp"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Package511State struct {
	XMLName xml.Name `xml:package511_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Publisher *oval_def.EntityStateStringType `xml:"publisher"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	VersionElm *oval_def.EntityStateVersionType `xml:"version"`

	Timestamp *oval_def.EntityStateStringType `xml:"timestamp"`

	Fmri *oval_def.EntityStateStringType `xml:"fmri"`

	Summary *oval_def.EntityStateStringType `xml:"summary"`

	Description *oval_def.EntityStateStringType `xml:"description"`

	Category *oval_def.EntityStateStringType `xml:"category"`

	UpdatesAvailable *oval_def.EntityStateBoolType `xml:"updates_available"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackageavoidlistTest struct {
	XMLName xml.Name `xml:packageavoidlist_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackageavoidlistObject struct {
	XMLName xml.Name `xml:packageavoidlist_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackageavoidlistState struct {
	XMLName xml.Name `xml:packageavoidlist_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Fmri *oval_def.EntityStateStringType `xml:"fmri"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackagecheckTest struct {
	XMLName xml.Name `xml:packagecheck_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackagecheckObject struct {
	XMLName xml.Name `xml:packagecheck_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *PackageCheckBehaviors `xml:"behaviors"`

	Pkginst *oval_def.EntityObjectStringType `xml:"pkginst"`

	Filepath oval_def.EntityObjectStringType `xml:"filepath"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackagecheckState struct {
	XMLName xml.Name `xml:packagecheck_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Pkginst *oval_def.EntityStateStringType `xml:"pkginst"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	ChecksumDiffers *oval_def.EntityStateBoolType `xml:"checksum_differs"`

	SizeDiffers *oval_def.EntityStateBoolType `xml:"size_differs"`

	MtimeDiffers *oval_def.EntityStateBoolType `xml:"mtime_differs"`

	Uread *EntityStatePermissionCompareType `xml:"uread"`

	Uwrite *EntityStatePermissionCompareType `xml:"uwrite"`

	Uexec *EntityStatePermissionCompareType `xml:"uexec"`

	Gread *EntityStatePermissionCompareType `xml:"gread"`

	Gwrite *EntityStatePermissionCompareType `xml:"gwrite"`

	Gexec *EntityStatePermissionCompareType `xml:"gexec"`

	Oread *EntityStatePermissionCompareType `xml:"oread"`

	Owrite *EntityStatePermissionCompareType `xml:"owrite"`

	Oexec *EntityStatePermissionCompareType `xml:"oexec"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackagefreezelistTest struct {
	XMLName xml.Name `xml:packagefreezelist_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackagefreezelistObject struct {
	XMLName xml.Name `xml:packagefreezelist_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackagefreezelistState struct {
	XMLName xml.Name `xml:packagefreezelist_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Fmri *oval_def.EntityStateStringType `xml:"fmri"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackagepublisherTest struct {
	XMLName xml.Name `xml:packagepublisher_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackagepublisherObject struct {
	XMLName xml.Name `xml:packagepublisher_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Type *EntityObjectPublisherTypeType `xml:"type"`

	OriginUri *oval_def.EntityObjectStringType `xml:"origin_uri"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PackagepublisherState struct {
	XMLName xml.Name `xml:packagepublisher_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Type *EntityStatePublisherTypeType `xml:"type"`

	OriginUri *oval_def.EntityStateStringType `xml:"origin_uri"`

	Alias *oval_def.EntityStateStringType `xml:"alias"`

	SslKey *oval_def.EntityStateStringType `xml:"ssl_key"`

	SslCert *oval_def.EntityStateStringType `xml:"ssl_cert"`

	ClientUuid *EntityStateClientUUIDType `xml:"client_uuid"`

	CatalogUpdated *oval_def.EntityStateIntType `xml:"catalog_updated"`

	Enabled *oval_def.EntityStateBoolType `xml:"enabled"`

	Order *oval_def.EntityStateIntType `xml:"order"`

	Properties *oval_def.EntityStateRecordType `xml:"properties"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Patch54Test struct {
	XMLName xml.Name `xml:patch54_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PatchTest struct {
	XMLName xml.Name `xml:patch_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Patch54Object struct {
	XMLName xml.Name `xml:patch54_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *PatchBehaviors `xml:"behaviors"`

	Base *oval_def.EntityObjectIntType `xml:"base"`

	VersionElm *oval_def.EntityObjectIntType `xml:"version"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PatchObject struct {
	XMLName xml.Name `xml:patch_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Base *oval_def.EntityObjectIntType `xml:"base"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PatchState struct {
	XMLName xml.Name `xml:patch_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Base *oval_def.EntityStateIntType `xml:"base"`

	VersionElm *oval_def.EntityStateIntType `xml:"version"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SmfTest struct {
	XMLName xml.Name `xml:smf_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SmfObject struct {
	XMLName xml.Name `xml:smf_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Fmri *oval_def.EntityObjectStringType `xml:"fmri"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SmfState struct {
	XMLName xml.Name `xml:smf_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Fmri *oval_def.EntityStateStringType `xml:"fmri"`

	ServiceName *oval_def.EntityStateStringType `xml:"service_name"`

	ServiceState *EntityStateSmfServiceStateType `xml:"service_state"`

	Protocol *oval_def.EntityStateStringType `xml:"protocol"`

	ServerExecutable *oval_def.EntityStateStringType `xml:"server_executable"`

	ServerArguements *oval_def.EntityStateStringType `xml:"server_arguements"`

	ExecAsUser *oval_def.EntityStateStringType `xml:"exec_as_user"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SmfpropertyTest struct {
	XMLName xml.Name `xml:smfproperty_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SmfpropertyObject struct {
	XMLName xml.Name `xml:smfproperty_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Service *oval_def.EntityObjectStringType `xml:"service"`

	Instance *oval_def.EntityObjectStringType `xml:"instance"`

	Property *oval_def.EntityObjectStringType `xml:"property"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SmfpropertyState struct {
	XMLName xml.Name `xml:smfproperty_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Service *oval_def.EntityStateStringType `xml:"service"`

	Instance *oval_def.EntityStateStringType `xml:"instance"`

	Property *oval_def.EntityStateStringType `xml:"property"`

	Fmri *oval_def.EntityStateStringType `xml:"fmri"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VariantTest struct {
	XMLName xml.Name `xml:variant_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VariantObject struct {
	XMLName xml.Name `xml:variant_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Path *oval_def.EntityObjectStringType `xml:"path"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VariantState struct {
	XMLName xml.Name `xml:variant_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VirtualizationinfoTest struct {
	XMLName xml.Name `xml:virtualizationinfo_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VirtualizationinfoObject struct {
	XMLName xml.Name `xml:virtualizationinfo_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VirtualizationinfoState struct {
	XMLName xml.Name `xml:virtualizationinfo_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Current *oval_def.EntityStateStringType `xml:"current"`

	Supported *EntityStateV12NEnvType `xml:"supported"`

	Parent *EntityStateV12NEnvType `xml:"parent"`

	LdomRole *EntityStateLDOMRoleType `xml:"ldom-role"`

	Properties *oval_def.EntityStateRecordType `xml:"properties"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// XSD ComplexType declarations

type PackageCheckBehaviors struct {
	FileattributesOnly string `xml:"fileattributes_only,attr,omitempty"`

	FilecontentsOnly string `xml:"filecontents_only,attr,omitempty"`

	NoVolatileeditable string `xml:"no_volatileeditable,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type PatchBehaviors struct {
	Supersedence string `xml:"supersedence,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectPublisherTypeType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateClientUUIDType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStatePermissionCompareType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStatePublisherTypeType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateSmfServiceStateType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateV12NEnvType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateLDOMRoleType struct {
	InnerXml string `xml:",innerxml"`
}
