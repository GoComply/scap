// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#hpux
package hpux_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type GetconfItem struct {
	XMLName xml.Name `xml:getconf_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	ParameterName *oval_sc.EntityItemStringType `xml:"parameter_name"`

	Pathname *oval_sc.EntityItemStringType `xml:"pathname"`

	Output *oval_sc.EntityItemAnySimpleType `xml:"output"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type NddItem struct {
	XMLName xml.Name `xml:ndd_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Device *oval_sc.EntityItemStringType `xml:"device"`

	Parameter *oval_sc.EntityItemStringType `xml:"parameter"`

	Value *oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PatchItem struct {
	XMLName xml.Name `xml:patch_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	PatchName *oval_sc.EntityItemStringType `xml:"patch_name"`

	Swtype *oval_sc.EntityItemStringType `xml:"swtype"`

	AreaPatched *oval_sc.EntityItemStringType `xml:"area_patched"`

	PatchBase *oval_sc.EntityItemStringType `xml:"patch_base"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SwlistItem struct {
	XMLName xml.Name `xml:swlist_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Swlist *oval_sc.EntityItemStringType `xml:"swlist"`

	Bundle *oval_sc.EntityItemStringType `xml:"bundle"`

	Fileset *oval_sc.EntityItemStringType `xml:"fileset"`

	Version *SwlistItemVersion `xml:"version"`

	Title *oval_sc.EntityItemStringType `xml:"title"`

	Vendor *oval_sc.EntityItemStringType `xml:"vendor"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type TrustedItem struct {
	XMLName xml.Name `xml:trusted_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Username *oval_sc.EntityItemStringType `xml:"username"`

	Uid *oval_sc.EntityItemIntType `xml:"uid"`

	Password *oval_sc.EntityItemStringType `xml:"password"`

	AccountOwner *oval_sc.EntityItemIntType `xml:"account_owner"`

	BootAuth *oval_sc.EntityItemStringType `xml:"boot_auth"`

	AuditId *oval_sc.EntityItemStringType `xml:"audit_id"`

	AuditFlag *oval_sc.EntityItemStringType `xml:"audit_flag"`

	PwChangeMin *oval_sc.EntityItemStringType `xml:"pw_change_min"`

	PwMaxSize *oval_sc.EntityItemIntType `xml:"pw_max_size"`

	PwExpiration *oval_sc.EntityItemIntType `xml:"pw_expiration"`

	PwLife *oval_sc.EntityItemStringType `xml:"pw_life"`

	PwChangeS *oval_sc.EntityItemStringType `xml:"pw_change_s"`

	PwChangeU *oval_sc.EntityItemStringType `xml:"pw_change_u"`

	AcctExpire *oval_sc.EntityItemIntType `xml:"acct_expire"`

	MaxLlogin *oval_sc.EntityItemStringType `xml:"max_llogin"`

	ExpWarning *oval_sc.EntityItemIntType `xml:"exp_warning"`

	UsrChgPw *oval_sc.EntityItemStringType `xml:"usr_chg_pw"`

	GenPw *oval_sc.EntityItemStringType `xml:"gen_pw"`

	PwRestrict *oval_sc.EntityItemStringType `xml:"pw_restrict"`

	PwNull *oval_sc.EntityItemStringType `xml:"pw_null"`

	PwGenChar *oval_sc.EntityItemStringType `xml:"pw_gen_char"`

	PwGenLet *oval_sc.EntityItemStringType `xml:"pw_gen_let"`

	LoginTime *oval_sc.EntityItemStringType `xml:"login_time"`

	PwChanger *oval_sc.EntityItemIntType `xml:"pw_changer"`

	LoginTimeS *oval_sc.EntityItemStringType `xml:"login_time_s"`

	LoginTimeU *oval_sc.EntityItemStringType `xml:"login_time_u"`

	LoginTtyS *oval_sc.EntityItemStringType `xml:"login_tty_s"`

	LoginTtyU *oval_sc.EntityItemStringType `xml:"login_tty_u"`

	NumULogins *oval_sc.EntityItemIntType `xml:"num_u_logins"`

	MaxULogins *oval_sc.EntityItemIntType `xml:"max_u_logins"`

	LockFlag *oval_sc.EntityItemBoolType `xml:"lock_flag"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SwlistItemVersion struct {
	XMLName xml.Name `xml:version`

	Datatype string `xml:"datatype,attr,omitempty"`
}

// XSD ComplexType declarations
