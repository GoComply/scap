// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
package hpux_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_def"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type GetconfTest struct {
	XMLName xml.Name `xml:getconf_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type GetconfObject struct {
	XMLName xml.Name `xml:getconf_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type GetconfState struct {
	XMLName xml.Name `xml:getconf_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	ParameterName *oval_def.EntityStateStringType `xml:"parameter_name"`

	Pathname *oval_def.EntityStateStringType `xml:"pathname"`

	Output *oval_def.EntityStateAnySimpleType `xml:"output"`
}

// Element
type NddTest struct {
	XMLName xml.Name `xml:ndd_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type NddObject struct {
	XMLName xml.Name `xml:ndd_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type NddState struct {
	XMLName xml.Name `xml:ndd_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Device *oval_def.EntityStateStringType `xml:"device"`

	Parameter *oval_def.EntityStateStringType `xml:"parameter"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`
}

// Element
type Patch53Test struct {
	XMLName xml.Name `xml:patch53_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type Patch53Object struct {
	XMLName xml.Name `xml:patch53_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type Patch53State struct {
	XMLName xml.Name `xml:patch53_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Swtype *oval_def.EntityStateStringType `xml:"swtype"`

	AreaPatched *oval_def.EntityStateStringType `xml:"area_patched"`

	PatchBase *oval_def.EntityStateStringType `xml:"patch_base"`
}

// Element
type PatchTest struct {
	XMLName xml.Name `xml:patch_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type PatchObject struct {
	XMLName xml.Name `xml:patch_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`

	PatchName oval_def.EntityObjectStringType `xml:"patch_name"`
}

// Element
type PatchState struct {
	XMLName xml.Name `xml:patch_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	PatchName *oval_def.EntityStateStringType `xml:"patch_name"`

	Swtype *oval_def.EntityStateStringType `xml:"swtype"`

	AreaPatched *oval_def.EntityStateStringType `xml:"area_patched"`

	PatchBase *oval_def.EntityStateStringType `xml:"patch_base"`
}

// Element
type SwlistTest struct {
	XMLName xml.Name `xml:swlist_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type SwlistObject struct {
	XMLName xml.Name `xml:swlist_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SwlistState struct {
	XMLName xml.Name `xml:swlist_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Swlist *oval_def.EntityStateStringType `xml:"swlist"`

	Bundle *oval_def.EntityStateStringType `xml:"bundle"`

	Fileset *oval_def.EntityStateStringType `xml:"fileset"`

	Version *SwlistStateVersion `xml:"version"`

	Title *oval_def.EntityStateStringType `xml:"title"`

	Vendor *oval_def.EntityStateStringType `xml:"vendor"`
}

// Element
type TrustedTest struct {
	XMLName xml.Name `xml:trusted_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type TrustedObject struct {
	XMLName xml.Name `xml:trusted_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type TrustedState struct {
	XMLName xml.Name `xml:trusted_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Username *oval_def.EntityStateStringType `xml:"username"`

	Uid *oval_def.EntityStateIntType `xml:"uid"`

	Password *oval_def.EntityStateStringType `xml:"password"`

	AccountOwner *oval_def.EntityStateIntType `xml:"account_owner"`

	BootAuth *oval_def.EntityStateStringType `xml:"boot_auth"`

	AuditId *oval_def.EntityStateStringType `xml:"audit_id"`

	AuditFlag *oval_def.EntityStateStringType `xml:"audit_flag"`

	PwChangeMin *oval_def.EntityStateStringType `xml:"pw_change_min"`

	PwMaxSize *oval_def.EntityStateIntType `xml:"pw_max_size"`

	PwExpiration *oval_def.EntityStateIntType `xml:"pw_expiration"`

	PwLife *oval_def.EntityStateStringType `xml:"pw_life"`

	PwChangeS *oval_def.EntityStateStringType `xml:"pw_change_s"`

	PwChangeU *oval_def.EntityStateStringType `xml:"pw_change_u"`

	AcctExpire *oval_def.EntityStateIntType `xml:"acct_expire"`

	MaxLlogin *oval_def.EntityStateStringType `xml:"max_llogin"`

	ExpWarning *oval_def.EntityStateIntType `xml:"exp_warning"`

	UsrChgPw *oval_def.EntityStateStringType `xml:"usr_chg_pw"`

	GenPw *oval_def.EntityStateStringType `xml:"gen_pw"`

	PwRestrict *oval_def.EntityStateStringType `xml:"pw_restrict"`

	PwNull *oval_def.EntityStateStringType `xml:"pw_null"`

	PwGenChar *oval_def.EntityStateStringType `xml:"pw_gen_char"`

	PwGenLet *oval_def.EntityStateStringType `xml:"pw_gen_let"`

	LoginTime *oval_def.EntityStateStringType `xml:"login_time"`

	PwChanger *oval_def.EntityStateIntType `xml:"pw_changer"`

	LoginTimeS *oval_def.EntityStateStringType `xml:"login_time_s"`

	LoginTimeU *oval_def.EntityStateStringType `xml:"login_time_u"`

	LoginTtyS *oval_def.EntityStateStringType `xml:"login_tty_s"`

	LoginTtyU *oval_def.EntityStateStringType `xml:"login_tty_u"`

	NumULogins *oval_def.EntityStateIntType `xml:"num_u_logins"`

	MaxULogins *oval_def.EntityStateIntType `xml:"max_u_logins"`

	LockFlag *oval_def.EntityStateBoolType `xml:"lock_flag"`
}

// Element
type SwlistStateVersion struct {
	XMLName xml.Name `xml:version`
}

// XSD ComplexType declarations

type Patch53Behaviors struct {
	Supersedence string `xml:"supersedence,attr"`
}
