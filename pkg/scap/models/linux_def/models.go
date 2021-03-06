// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
package linux_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_def"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type ApparmorstatusTest struct {
	XMLName xml.Name `xml:apparmorstatus_test`

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
type ApparmorstatusObject struct {
	XMLName xml.Name `xml:apparmorstatus_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ApparmorstatusState struct {
	XMLName xml.Name `xml:apparmorstatus_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	LoadedProfilesCount *oval_def.EntityStateIntType `xml:"loaded_profiles_count"`

	EnforceModeProfilesCount *oval_def.EntityStateIntType `xml:"enforce_mode_profiles_count"`

	ComplainModeProfilesCount *oval_def.EntityStateIntType `xml:"complain_mode_profiles_count"`

	ProcessesWithProfilesCount *oval_def.EntityStateIntType `xml:"processes_with_profiles_count"`

	EnforceModeProcessesCount *oval_def.EntityStateIntType `xml:"enforce_mode_processes_count"`

	ComplainModeProcessesCount *oval_def.EntityStateIntType `xml:"complain_mode_processes_count"`

	UnconfinedProcessesWithProfilesCount *oval_def.EntityStateIntType `xml:"unconfined_processes_with_profiles_count"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type DpkginfoTest struct {
	XMLName xml.Name `xml:dpkginfo_test`

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
type DpkginfoObject struct {
	XMLName xml.Name `xml:dpkginfo_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type DpkginfoState struct {
	XMLName xml.Name `xml:dpkginfo_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Arch *oval_def.EntityStateStringType `xml:"arch"`

	Epoch *DpkginfoStateEpoch `xml:"epoch"`

	Release *DpkginfoStateRelease `xml:"release"`

	VersionElm *DpkginfoStateVersion `xml:"version"`

	Evr *DpkginfoStateEvr `xml:"evr"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type IflistenersTest struct {
	XMLName xml.Name `xml:iflisteners_test`

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
type IflistenersObject struct {
	XMLName xml.Name `xml:iflisteners_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	InterfaceName *oval_def.EntityObjectStringType `xml:"interface_name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type IflistenersState struct {
	XMLName xml.Name `xml:iflisteners_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	InterfaceName *oval_def.EntityStateStringType `xml:"interface_name"`

	Protocol *EntityStateProtocolType `xml:"protocol"`

	HwAddress *oval_def.EntityStateStringType `xml:"hw_address"`

	ProgramName *oval_def.EntityStateStringType `xml:"program_name"`

	Pid *oval_def.EntityStateIntType `xml:"pid"`

	UserId *oval_def.EntityStateIntType `xml:"user_id"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InetlisteningserversTest struct {
	XMLName xml.Name `xml:inetlisteningservers_test`

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
type InetlisteningserversObject struct {
	XMLName xml.Name `xml:inetlisteningservers_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Protocol *oval_def.EntityObjectStringType `xml:"protocol"`

	LocalAddress *oval_def.EntityObjectIPAddressStringType `xml:"local_address"`

	LocalPort *oval_def.EntityObjectIntType `xml:"local_port"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InetlisteningserversState struct {
	XMLName xml.Name `xml:inetlisteningservers_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Protocol *oval_def.EntityStateStringType `xml:"protocol"`

	LocalAddress *oval_def.EntityStateIPAddressStringType `xml:"local_address"`

	LocalPort *oval_def.EntityStateIntType `xml:"local_port"`

	LocalFullAddress *oval_def.EntityStateStringType `xml:"local_full_address"`

	ProgramName *oval_def.EntityStateStringType `xml:"program_name"`

	ForeignAddress *oval_def.EntityStateIPAddressStringType `xml:"foreign_address"`

	ForeignPort *oval_def.EntityStateIntType `xml:"foreign_port"`

	ForeignFullAddress *oval_def.EntityStateStringType `xml:"foreign_full_address"`

	Pid *oval_def.EntityStateIntType `xml:"pid"`

	UserId *oval_def.EntityStateIntType `xml:"user_id"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PartitionTest struct {
	XMLName xml.Name `xml:partition_test`

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
type PartitionObject struct {
	XMLName xml.Name `xml:partition_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	MountPoint *oval_def.EntityObjectStringType `xml:"mount_point"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PartitionState struct {
	XMLName xml.Name `xml:partition_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	MountPoint *oval_def.EntityStateStringType `xml:"mount_point"`

	Device *oval_def.EntityStateStringType `xml:"device"`

	Uuid *oval_def.EntityStateStringType `xml:"uuid"`

	FsType *oval_def.EntityStateStringType `xml:"fs_type"`

	MountOptions *oval_def.EntityStateStringType `xml:"mount_options"`

	TotalSpace *oval_def.EntityStateIntType `xml:"total_space"`

	SpaceUsed *oval_def.EntityStateIntType `xml:"space_used"`

	SpaceLeft *oval_def.EntityStateIntType `xml:"space_left"`

	SpaceLeftForUnprivilegedUsers *oval_def.EntityStateIntType `xml:"space_left_for_unprivileged_users"`

	BlockSize *oval_def.EntityStateIntType `xml:"block_size"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RpminfoTest struct {
	XMLName xml.Name `xml:rpminfo_test`

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
type RpminfoObject struct {
	XMLName xml.Name `xml:rpminfo_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *RpmInfoBehaviors `xml:"behaviors"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RpminfoState struct {
	XMLName xml.Name `xml:rpminfo_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Arch *oval_def.EntityStateStringType `xml:"arch"`

	Epoch *RpminfoStateEpoch `xml:"epoch"`

	Release *RpminfoStateRelease `xml:"release"`

	VersionElm *RpminfoStateVersion `xml:"version"`

	Evr *oval_def.EntityStateEVRStringType `xml:"evr"`

	SignatureKeyid *oval_def.EntityStateStringType `xml:"signature_keyid"`

	ExtendedName *oval_def.EntityStateStringType `xml:"extended_name"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RpmverifyTest struct {
	XMLName xml.Name `xml:rpmverify_test`

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
type RpmverifyObject struct {
	XMLName xml.Name `xml:rpmverify_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *RpmVerifyBehaviors `xml:"behaviors"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filepath *oval_def.EntityObjectStringType `xml:"filepath"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RpmverifyState struct {
	XMLName xml.Name `xml:rpmverify_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	SizeDiffers *EntityStateRpmVerifyResultType `xml:"size_differs"`

	ModeDiffers *EntityStateRpmVerifyResultType `xml:"mode_differs"`

	Md5Differs *EntityStateRpmVerifyResultType `xml:"md5_differs"`

	DeviceDiffers *EntityStateRpmVerifyResultType `xml:"device_differs"`

	LinkMismatch *EntityStateRpmVerifyResultType `xml:"link_mismatch"`

	OwnershipDiffers *EntityStateRpmVerifyResultType `xml:"ownership_differs"`

	GroupDiffers *EntityStateRpmVerifyResultType `xml:"group_differs"`

	MtimeDiffers *EntityStateRpmVerifyResultType `xml:"mtime_differs"`

	CapabilitiesDiffer *EntityStateRpmVerifyResultType `xml:"capabilities_differ"`

	ConfigurationFile *oval_def.EntityStateBoolType `xml:"configuration_file"`

	DocumentationFile *oval_def.EntityStateBoolType `xml:"documentation_file"`

	GhostFile *oval_def.EntityStateBoolType `xml:"ghost_file"`

	LicenseFile *oval_def.EntityStateBoolType `xml:"license_file"`

	ReadmeFile *oval_def.EntityStateBoolType `xml:"readme_file"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RpmverifyfileTest struct {
	XMLName xml.Name `xml:rpmverifyfile_test`

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
type RpmverifyfileObject struct {
	XMLName xml.Name `xml:rpmverifyfile_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *RpmVerifyFileBehaviors `xml:"behaviors"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Epoch *RpmverifyfileObjectEpoch `xml:"epoch"`

	VersionElm *RpmverifyfileObjectVersion `xml:"version"`

	Release *RpmverifyfileObjectRelease `xml:"release"`

	Arch *oval_def.EntityObjectStringType `xml:"arch"`

	Filepath *oval_def.EntityObjectStringType `xml:"filepath"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RpmverifyfileState struct {
	XMLName xml.Name `xml:rpmverifyfile_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Epoch *RpmverifyfileStateEpoch `xml:"epoch"`

	VersionElm *RpmverifyfileStateVersion `xml:"version"`

	Release *RpmverifyfileStateRelease `xml:"release"`

	Arch *oval_def.EntityStateStringType `xml:"arch"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	ExtendedName *oval_def.EntityStateStringType `xml:"extended_name"`

	SizeDiffers *EntityStateRpmVerifyResultType `xml:"size_differs"`

	ModeDiffers *EntityStateRpmVerifyResultType `xml:"mode_differs"`

	Md5Differs *EntityStateRpmVerifyResultType `xml:"md5_differs"`

	FiledigestDiffers *EntityStateRpmVerifyResultType `xml:"filedigest_differs"`

	DeviceDiffers *EntityStateRpmVerifyResultType `xml:"device_differs"`

	LinkMismatch *EntityStateRpmVerifyResultType `xml:"link_mismatch"`

	OwnershipDiffers *EntityStateRpmVerifyResultType `xml:"ownership_differs"`

	GroupDiffers *EntityStateRpmVerifyResultType `xml:"group_differs"`

	MtimeDiffers *EntityStateRpmVerifyResultType `xml:"mtime_differs"`

	CapabilitiesDiffer *EntityStateRpmVerifyResultType `xml:"capabilities_differ"`

	ConfigurationFile *oval_def.EntityStateBoolType `xml:"configuration_file"`

	DocumentationFile *oval_def.EntityStateBoolType `xml:"documentation_file"`

	GhostFile *oval_def.EntityStateBoolType `xml:"ghost_file"`

	LicenseFile *oval_def.EntityStateBoolType `xml:"license_file"`

	ReadmeFile *oval_def.EntityStateBoolType `xml:"readme_file"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RpmverifypackageTest struct {
	XMLName xml.Name `xml:rpmverifypackage_test`

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
type RpmverifypackageObject struct {
	XMLName xml.Name `xml:rpmverifypackage_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *RpmVerifyPackageBehaviors `xml:"behaviors"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Epoch *RpmverifypackageObjectEpoch `xml:"epoch"`

	VersionElm *RpmverifypackageObjectVersion `xml:"version"`

	Release *RpmverifypackageObjectRelease `xml:"release"`

	Arch *oval_def.EntityObjectStringType `xml:"arch"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RpmverifypackageState struct {
	XMLName xml.Name `xml:rpmverifypackage_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Epoch *RpmverifypackageStateEpoch `xml:"epoch"`

	VersionElm *RpmverifypackageStateVersion `xml:"version"`

	Release *RpmverifypackageStateRelease `xml:"release"`

	Arch *oval_def.EntityStateStringType `xml:"arch"`

	ExtendedName *oval_def.EntityStateStringType `xml:"extended_name"`

	DependencyCheckPassed *oval_def.EntityStateBoolType `xml:"dependency_check_passed"`

	DigestCheckPassed *oval_def.EntityStateBoolType `xml:"digest_check_passed"`

	VerificationScriptSuccessful *oval_def.EntityStateBoolType `xml:"verification_script_successful"`

	SignatureCheckPassed *oval_def.EntityStateBoolType `xml:"signature_check_passed"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SelinuxbooleanTest struct {
	XMLName xml.Name `xml:selinuxboolean_test`

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
type SelinuxbooleanObject struct {
	XMLName xml.Name `xml:selinuxboolean_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SelinuxbooleanState struct {
	XMLName xml.Name `xml:selinuxboolean_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	CurrentStatus *oval_def.EntityStateBoolType `xml:"current_status"`

	PendingStatus *oval_def.EntityStateBoolType `xml:"pending_status"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SelinuxsecuritycontextTest struct {
	XMLName xml.Name `xml:selinuxsecuritycontext_test`

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
type SelinuxsecuritycontextObject struct {
	XMLName xml.Name `xml:selinuxsecuritycontext_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *FileBehaviors `xml:"behaviors"`

	Filter []oval_def.Filter `xml:"filter"`

	Filepath *oval_def.EntityObjectStringType `xml:"filepath"`

	Pid *oval_def.EntityObjectIntType `xml:"pid"`

	Path *oval_def.EntityObjectStringType `xml:"path"`

	Filename *oval_def.EntityObjectStringType `xml:"filename"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SelinuxsecuritycontextState struct {
	XMLName xml.Name `xml:selinuxsecuritycontext_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Pid *oval_def.EntityStateIntType `xml:"pid"`

	User *oval_def.EntityStateStringType `xml:"user"`

	Role *oval_def.EntityStateStringType `xml:"role"`

	Type *oval_def.EntityStateStringType `xml:"type"`

	LowSensitivity *oval_def.EntityStateStringType `xml:"low_sensitivity"`

	LowCategory *oval_def.EntityStateStringType `xml:"low_category"`

	HighSensitivity *oval_def.EntityStateStringType `xml:"high_sensitivity"`

	HighCategory *oval_def.EntityStateStringType `xml:"high_category"`

	RawlowSensitivity *oval_def.EntityStateStringType `xml:"rawlow_sensitivity"`

	RawlowCategory *oval_def.EntityStateStringType `xml:"rawlow_category"`

	RawhighSensitivity *oval_def.EntityStateStringType `xml:"rawhigh_sensitivity"`

	RawhighCategory *oval_def.EntityStateStringType `xml:"rawhigh_category"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SlackwarepkginfoTest struct {
	XMLName xml.Name `xml:slackwarepkginfo_test`

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
type SlackwarepkginfoObject struct {
	XMLName xml.Name `xml:slackwarepkginfo_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SlackwarepkginfoState struct {
	XMLName xml.Name `xml:slackwarepkginfo_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	VersionElm *SlackwarepkginfoStateVersion `xml:"version"`

	Architecture *oval_def.EntityStateStringType `xml:"architecture"`

	Revision *oval_def.EntityStateStringType `xml:"revision"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SystemdunitdependencyTest struct {
	XMLName xml.Name `xml:systemdunitdependency_test`

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
type SystemdunitdependencyObject struct {
	XMLName xml.Name `xml:systemdunitdependency_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Unit *oval_def.EntityObjectStringType `xml:"unit"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SystemdunitdependencyState struct {
	XMLName xml.Name `xml:systemdunitdependency_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Unit *oval_def.EntityStateStringType `xml:"unit"`

	Dependency *oval_def.EntityStateStringType `xml:"dependency"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SystemdunitpropertyTest struct {
	XMLName xml.Name `xml:systemdunitproperty_test`

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
type SystemdunitpropertyObject struct {
	XMLName xml.Name `xml:systemdunitproperty_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Unit *oval_def.EntityObjectStringType `xml:"unit"`

	Property *oval_def.EntityObjectStringType `xml:"property"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SystemdunitpropertyState struct {
	XMLName xml.Name `xml:systemdunitproperty_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Unit *oval_def.EntityStateStringType `xml:"unit"`

	Property *oval_def.EntityStateStringType `xml:"property"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type DpkginfoStateEpoch struct {
	XMLName xml.Name `xml:epoch`
}

// Element
type DpkginfoStateRelease struct {
	XMLName xml.Name `xml:release`
}

// Element
type DpkginfoStateVersion struct {
	XMLName xml.Name `xml:version`
}

// Element
type DpkginfoStateEvr struct {
	XMLName xml.Name `xml:evr`
}

// Element
type RpminfoStateEpoch struct {
	XMLName xml.Name `xml:epoch`
}

// Element
type RpminfoStateRelease struct {
	XMLName xml.Name `xml:release`
}

// Element
type RpminfoStateVersion struct {
	XMLName xml.Name `xml:version`
}

// Element
type RpmverifyfileObjectEpoch struct {
	XMLName xml.Name `xml:epoch`
}

// Element
type RpmverifyfileObjectVersion struct {
	XMLName xml.Name `xml:version`
}

// Element
type RpmverifyfileObjectRelease struct {
	XMLName xml.Name `xml:release`
}

// Element
type RpmverifyfileStateEpoch struct {
	XMLName xml.Name `xml:epoch`
}

// Element
type RpmverifyfileStateVersion struct {
	XMLName xml.Name `xml:version`
}

// Element
type RpmverifyfileStateRelease struct {
	XMLName xml.Name `xml:release`
}

// Element
type RpmverifypackageObjectEpoch struct {
	XMLName xml.Name `xml:epoch`
}

// Element
type RpmverifypackageObjectVersion struct {
	XMLName xml.Name `xml:version`
}

// Element
type RpmverifypackageObjectRelease struct {
	XMLName xml.Name `xml:release`
}

// Element
type RpmverifypackageStateEpoch struct {
	XMLName xml.Name `xml:epoch`
}

// Element
type RpmverifypackageStateVersion struct {
	XMLName xml.Name `xml:version`
}

// Element
type RpmverifypackageStateRelease struct {
	XMLName xml.Name `xml:release`
}

// Element
type SlackwarepkginfoStateVersion struct {
	XMLName xml.Name `xml:version`
}

// XSD ComplexType declarations

type RpmInfoBehaviors struct {
	Filepaths string `xml:"filepaths,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type RpmVerifyBehaviors struct {
	Nodeps string `xml:"nodeps,attr,omitempty"`

	Nodigest string `xml:"nodigest,attr,omitempty"`

	Nofiles string `xml:"nofiles,attr,omitempty"`

	Noscripts string `xml:"noscripts,attr,omitempty"`

	Nosignature string `xml:"nosignature,attr,omitempty"`

	Nolinkto string `xml:"nolinkto,attr,omitempty"`

	Nomd5 string `xml:"nomd5,attr,omitempty"`

	Nosize string `xml:"nosize,attr,omitempty"`

	Nouser string `xml:"nouser,attr,omitempty"`

	Nogroup string `xml:"nogroup,attr,omitempty"`

	Nomtime string `xml:"nomtime,attr,omitempty"`

	Nomode string `xml:"nomode,attr,omitempty"`

	Nordev string `xml:"nordev,attr,omitempty"`

	Noconfigfiles string `xml:"noconfigfiles,attr,omitempty"`

	Noghostfiles string `xml:"noghostfiles,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type RpmVerifyFileBehaviors struct {
	Nolinkto string `xml:"nolinkto,attr,omitempty"`

	Nomd5 string `xml:"nomd5,attr,omitempty"`

	Nosize string `xml:"nosize,attr,omitempty"`

	Nouser string `xml:"nouser,attr,omitempty"`

	Nogroup string `xml:"nogroup,attr,omitempty"`

	Nomtime string `xml:"nomtime,attr,omitempty"`

	Nomode string `xml:"nomode,attr,omitempty"`

	Nordev string `xml:"nordev,attr,omitempty"`

	Noconfigfiles string `xml:"noconfigfiles,attr,omitempty"`

	Noghostfiles string `xml:"noghostfiles,attr,omitempty"`

	Nofiledigest string `xml:"nofiledigest,attr,omitempty"`

	Nocaps string `xml:"nocaps,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type RpmVerifyPackageBehaviors struct {
	Nodeps string `xml:"nodeps,attr,omitempty"`

	Nodigest string `xml:"nodigest,attr,omitempty"`

	Noscripts string `xml:"noscripts,attr,omitempty"`

	Nosignature string `xml:"nosignature,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type FileBehaviors struct {
	MaxDepth string `xml:"max_depth,attr,omitempty"`

	Recurse string `xml:"recurse,attr,omitempty"`

	RecurseDirection string `xml:"recurse_direction,attr,omitempty"`

	RecurseFileSystem string `xml:"recurse_file_system,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateRpmVerifyResultType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateProtocolType struct {
	InnerXml string `xml:",innerxml"`
}
