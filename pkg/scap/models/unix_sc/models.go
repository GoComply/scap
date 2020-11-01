// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#unix
package unix_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type DnscacheItem struct {
	XMLName xml.Name `xml:dnscache_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	DomainName *oval_sc.EntityItemStringType `xml:"domain_name"`

	Ttl *oval_sc.EntityItemIntType `xml:"ttl"`

	IpAddress []oval_sc.EntityItemIPAddressStringType `xml:"ip_address"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type FileItem struct {
	XMLName xml.Name `xml:file_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Filename *oval_sc.EntityItemStringType `xml:"filename"`

	Type *oval_sc.EntityItemStringType `xml:"type"`

	GroupId *FileItemGroupId `xml:"group_id"`

	UserId *FileItemUserId `xml:"user_id"`

	ATime *FileItemATime `xml:"a_time"`

	CTime *FileItemCTime `xml:"c_time"`

	MTime *FileItemMTime `xml:"m_time"`

	Size *oval_sc.EntityItemIntType `xml:"size"`

	Suid *oval_sc.EntityItemBoolType `xml:"suid"`

	Sgid *oval_sc.EntityItemBoolType `xml:"sgid"`

	Sticky *oval_sc.EntityItemBoolType `xml:"sticky"`

	Uread *oval_sc.EntityItemBoolType `xml:"uread"`

	Uwrite *oval_sc.EntityItemBoolType `xml:"uwrite"`

	Uexec *oval_sc.EntityItemBoolType `xml:"uexec"`

	Gread *oval_sc.EntityItemBoolType `xml:"gread"`

	Gwrite *oval_sc.EntityItemBoolType `xml:"gwrite"`

	Gexec *oval_sc.EntityItemBoolType `xml:"gexec"`

	Oread *oval_sc.EntityItemBoolType `xml:"oread"`

	Owrite *oval_sc.EntityItemBoolType `xml:"owrite"`

	Oexec *oval_sc.EntityItemBoolType `xml:"oexec"`

	HasExtendedAcl *oval_sc.EntityItemBoolType `xml:"has_extended_acl"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type FileextendedattributeItem struct {
	XMLName xml.Name `xml:fileextendedattribute_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Filename *oval_sc.EntityItemStringType `xml:"filename"`

	AttributeName *oval_sc.EntityItemStringType `xml:"attribute_name"`

	Value *oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type GconfItem struct {
	XMLName xml.Name `xml:gconf_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Key *oval_sc.EntityItemStringType `xml:"key"`

	Source *oval_sc.EntityItemStringType `xml:"source"`

	Type *EntityItemGconfTypeType `xml:"type"`

	IsWritable *oval_sc.EntityItemBoolType `xml:"is_writable"`

	ModUser *oval_sc.EntityItemStringType `xml:"mod_user"`

	ModTime *oval_sc.EntityItemIntType `xml:"mod_time"`

	IsDefault *oval_sc.EntityItemBoolType `xml:"is_default"`

	Value []oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type InetdItem struct {
	XMLName xml.Name `xml:inetd_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Protocol *oval_sc.EntityItemStringType `xml:"protocol"`

	ServiceName *oval_sc.EntityItemStringType `xml:"service_name"`

	ServerProgram *oval_sc.EntityItemStringType `xml:"server_program"`

	ServerArguments *oval_sc.EntityItemStringType `xml:"server_arguments"`

	EndpointType *EntityItemEndpointType `xml:"endpoint_type"`

	ExecAsUser *oval_sc.EntityItemStringType `xml:"exec_as_user"`

	WaitStatus *EntityItemWaitStatusType `xml:"wait_status"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type InterfaceItem struct {
	XMLName xml.Name `xml:interface_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Type *EntityItemInterfaceType `xml:"type"`

	HardwareAddr *oval_sc.EntityItemStringType `xml:"hardware_addr"`

	InetAddr *oval_sc.EntityItemIPAddressStringType `xml:"inet_addr"`

	BroadcastAddr *oval_sc.EntityItemIPAddressStringType `xml:"broadcast_addr"`

	Netmask *oval_sc.EntityItemIPAddressStringType `xml:"netmask"`

	Flag []oval_sc.EntityItemStringType `xml:"flag"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PasswordItem struct {
	XMLName xml.Name `xml:password_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Username *oval_sc.EntityItemStringType `xml:"username"`

	Password *oval_sc.EntityItemStringType `xml:"password"`

	UserId *PasswordItemUserId `xml:"user_id"`

	GroupId *PasswordItemGroupId `xml:"group_id"`

	Gcos *oval_sc.EntityItemStringType `xml:"gcos"`

	HomeDir *oval_sc.EntityItemStringType `xml:"home_dir"`

	LoginShell *oval_sc.EntityItemStringType `xml:"login_shell"`

	LastLogin *oval_sc.EntityItemIntType `xml:"last_login"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type ProcessItem struct {
	XMLName xml.Name `xml:process_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Command *oval_sc.EntityItemStringType `xml:"command"`

	ExecTime *oval_sc.EntityItemStringType `xml:"exec_time"`

	Pid *oval_sc.EntityItemIntType `xml:"pid"`

	Ppid *oval_sc.EntityItemIntType `xml:"ppid"`

	Priority *oval_sc.EntityItemIntType `xml:"priority"`

	Ruid *oval_sc.EntityItemIntType `xml:"ruid"`

	SchedulingClass *oval_sc.EntityItemStringType `xml:"scheduling_class"`

	StartTime *oval_sc.EntityItemStringType `xml:"start_time"`

	Tty *oval_sc.EntityItemStringType `xml:"tty"`

	UserId *oval_sc.EntityItemIntType `xml:"user_id"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type Process58Item struct {
	XMLName xml.Name `xml:process58_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	CommandLine *oval_sc.EntityItemStringType `xml:"command_line"`

	ExecTime *oval_sc.EntityItemStringType `xml:"exec_time"`

	Pid *oval_sc.EntityItemIntType `xml:"pid"`

	Ppid *oval_sc.EntityItemIntType `xml:"ppid"`

	Priority *oval_sc.EntityItemIntType `xml:"priority"`

	Ruid *oval_sc.EntityItemIntType `xml:"ruid"`

	SchedulingClass *oval_sc.EntityItemStringType `xml:"scheduling_class"`

	StartTime *oval_sc.EntityItemStringType `xml:"start_time"`

	Tty *oval_sc.EntityItemStringType `xml:"tty"`

	UserId *oval_sc.EntityItemIntType `xml:"user_id"`

	ExecShield *oval_sc.EntityItemBoolType `xml:"exec_shield"`

	Loginuid *oval_sc.EntityItemIntType `xml:"loginuid"`

	PosixCapability []EntityItemCapabilityType `xml:"posix_capability"`

	SelinuxDomainLabel *oval_sc.EntityItemStringType `xml:"selinux_domain_label"`

	SessionId *oval_sc.EntityItemIntType `xml:"session_id"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type RoutingtableItem struct {
	XMLName xml.Name `xml:routingtable_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Destination *oval_sc.EntityItemIPAddressType `xml:"destination"`

	Gateway *oval_sc.EntityItemIPAddressType `xml:"gateway"`

	Flags []EntityItemRoutingTableFlagsType `xml:"flags"`

	InterfaceName *oval_sc.EntityItemStringType `xml:"interface_name"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type RunlevelItem struct {
	XMLName xml.Name `xml:runlevel_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	ServiceName *oval_sc.EntityItemStringType `xml:"service_name"`

	Runlevel *oval_sc.EntityItemStringType `xml:"runlevel"`

	Start *oval_sc.EntityItemBoolType `xml:"start"`

	Kill *oval_sc.EntityItemBoolType `xml:"kill"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SccsItem struct {
	XMLName xml.Name `xml:sccs_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Filename *oval_sc.EntityItemStringType `xml:"filename"`

	ModuleName *oval_sc.EntityItemStringType `xml:"module_name"`

	ModuleType *oval_sc.EntityItemStringType `xml:"module_type"`

	Release *oval_sc.EntityItemStringType `xml:"release"`

	Level *oval_sc.EntityItemStringType `xml:"level"`

	Branch *oval_sc.EntityItemStringType `xml:"branch"`

	Sequence *oval_sc.EntityItemStringType `xml:"sequence"`

	WhatString *oval_sc.EntityItemStringType `xml:"what_string"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type ShadowItem struct {
	XMLName xml.Name `xml:shadow_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Username *oval_sc.EntityItemStringType `xml:"username"`

	Password *oval_sc.EntityItemStringType `xml:"password"`

	ChgLst *ShadowItemChgLst `xml:"chg_lst"`

	ChgAllow *ShadowItemChgAllow `xml:"chg_allow"`

	ChgReq *ShadowItemChgReq `xml:"chg_req"`

	ExpWarn *ShadowItemExpWarn `xml:"exp_warn"`

	ExpInact *ShadowItemExpInact `xml:"exp_inact"`

	ExpDate *ShadowItemExpDate `xml:"exp_date"`

	Flag *ShadowItemFlag `xml:"flag"`

	EncryptMethod *EntityItemEncryptMethodType `xml:"encrypt_method"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SymlinkItem struct {
	XMLName xml.Name `xml:symlink_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Filepath oval_sc.EntityItemStringType `xml:"filepath"`

	CanonicalPath oval_sc.EntityItemStringType `xml:"canonical_path"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SysctlItem struct {
	XMLName xml.Name `xml:sysctl_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Value []oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type UnameItem struct {
	XMLName xml.Name `xml:uname_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	MachineClass *oval_sc.EntityItemStringType `xml:"machine_class"`

	NodeName *oval_sc.EntityItemStringType `xml:"node_name"`

	OsName *oval_sc.EntityItemStringType `xml:"os_name"`

	OsRelease *oval_sc.EntityItemStringType `xml:"os_release"`

	OsVersion *oval_sc.EntityItemStringType `xml:"os_version"`

	ProcessorType *oval_sc.EntityItemStringType `xml:"processor_type"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type XinetdItem struct {
	XMLName xml.Name `xml:xinetd_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Protocol *oval_sc.EntityItemStringType `xml:"protocol"`

	ServiceName *oval_sc.EntityItemStringType `xml:"service_name"`

	Flags []oval_sc.EntityItemStringType `xml:"flags"`

	NoAccess []oval_sc.EntityItemStringType `xml:"no_access"`

	OnlyFrom []oval_sc.EntityItemIPAddressStringType `xml:"only_from"`

	Port *oval_sc.EntityItemIntType `xml:"port"`

	Server *oval_sc.EntityItemStringType `xml:"server"`

	ServerArguments *oval_sc.EntityItemStringType `xml:"server_arguments"`

	SocketType *oval_sc.EntityItemStringType `xml:"socket_type"`

	Type []EntityItemXinetdTypeStatusType `xml:"type"`

	User *oval_sc.EntityItemStringType `xml:"user"`

	Wait *oval_sc.EntityItemBoolType `xml:"wait"`

	Disabled *oval_sc.EntityItemBoolType `xml:"disabled"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type FileItemGroupId struct {
	XMLName xml.Name `xml:group_id`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type FileItemUserId struct {
	XMLName xml.Name `xml:user_id`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type FileItemATime struct {
	XMLName xml.Name `xml:a_time`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type FileItemCTime struct {
	XMLName xml.Name `xml:c_time`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type FileItemMTime struct {
	XMLName xml.Name `xml:m_time`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type PasswordItemUserId struct {
	XMLName xml.Name `xml:user_id`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type PasswordItemGroupId struct {
	XMLName xml.Name `xml:group_id`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type ShadowItemChgLst struct {
	XMLName xml.Name `xml:chg_lst`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type ShadowItemChgAllow struct {
	XMLName xml.Name `xml:chg_allow`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type ShadowItemChgReq struct {
	XMLName xml.Name `xml:chg_req`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type ShadowItemExpWarn struct {
	XMLName xml.Name `xml:exp_warn`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type ShadowItemExpInact struct {
	XMLName xml.Name `xml:exp_inact`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type ShadowItemExpDate struct {
	XMLName xml.Name `xml:exp_date`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// Element
type ShadowItemFlag struct {
	XMLName xml.Name `xml:flag`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`
}

// XSD ComplexType declarations

type EntityItemCapabilityType struct {
	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemEndpointType struct {
	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemGconfTypeType struct {
	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemRoutingTableFlagsType struct {
	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemXinetdTypeStatusType struct {
	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemWaitStatusType struct {
	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemEncryptMethodType struct {
	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemInterfaceType struct {
	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations
