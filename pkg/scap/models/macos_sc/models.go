// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#macos
package macos_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type AccountinfoItem struct {
	XMLName xml.Name `xml:accountinfo_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Username *oval_sc.EntityItemStringType `xml:"username"`

	Password *oval_sc.EntityItemStringType `xml:"password"`

	Uid *oval_sc.EntityItemIntType `xml:"uid"`

	Gid *oval_sc.EntityItemIntType `xml:"gid"`

	Realname *oval_sc.EntityItemStringType `xml:"realname"`

	HomeDir *oval_sc.EntityItemStringType `xml:"home_dir"`

	LoginShell *oval_sc.EntityItemStringType `xml:"login_shell"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type AuthorizationdbItem struct {
	XMLName xml.Name `xml:authorizationdb_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	RightName *oval_sc.EntityItemStringType `xml:"right_name"`

	Xpath *oval_sc.EntityItemStringType `xml:"xpath"`

	ValueOf []oval_sc.EntityItemAnySimpleType `xml:"value_of"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type CorestorageItem struct {
	XMLName xml.Name `xml:corestorage_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Uuid oval_sc.EntityItemStringType `xml:"uuid"`

	Xpath *oval_sc.EntityItemStringType `xml:"xpath"`

	ValueOf []oval_sc.EntityItemAnySimpleType `xml:"value_of"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type DiskutilItem struct {
	XMLName xml.Name `xml:diskutil_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Device *oval_sc.EntityItemStringType `xml:"device"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Uread *EntityItemPermissionCompareType `xml:"uread"`

	Uwrite *EntityItemPermissionCompareType `xml:"uwrite"`

	Uexec *EntityItemPermissionCompareType `xml:"uexec"`

	Gread *EntityItemPermissionCompareType `xml:"gread"`

	Gwrite *EntityItemPermissionCompareType `xml:"gwrite"`

	Gexec *EntityItemPermissionCompareType `xml:"gexec"`

	Oread *EntityItemPermissionCompareType `xml:"oread"`

	Owrite *EntityItemPermissionCompareType `xml:"owrite"`

	Oexec *EntityItemPermissionCompareType `xml:"oexec"`

	UserDiffers *oval_sc.EntityItemBoolType `xml:"user_differs"`

	ActualUser *oval_sc.EntityItemIntType `xml:"actual_user"`

	ExpectedUser *oval_sc.EntityItemIntType `xml:"expected_user"`

	GroupDiffers *oval_sc.EntityItemBoolType `xml:"group_differs"`

	ActualGroup *oval_sc.EntityItemIntType `xml:"actual_group"`

	ExpectedGroup *oval_sc.EntityItemIntType `xml:"expected_group"`

	SymlinkDiffers *oval_sc.EntityItemBoolType `xml:"symlink_differs"`

	ActualSymlink *oval_sc.EntityItemStringType `xml:"actual_symlink"`

	ExpectedSymlink *oval_sc.EntityItemStringType `xml:"expected_symlink"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type GatekeeperItem struct {
	XMLName xml.Name `xml:gatekeeper_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Enabled oval_sc.EntityItemBoolType `xml:"enabled"`

	Unlabeled []oval_sc.EntityItemStringType `xml:"unlabeled"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type InetlisteningserverItem struct {
	XMLName xml.Name `xml:inetlisteningserver_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	ProgramName *oval_sc.EntityItemStringType `xml:"program_name"`

	LocalAddress *oval_sc.EntityItemIPAddressStringType `xml:"local_address"`

	LocalFullAddress *oval_sc.EntityItemStringType `xml:"local_full_address"`

	LocalPort *oval_sc.EntityItemIntType `xml:"local_port"`

	ForeignAddress *oval_sc.EntityItemIPAddressStringType `xml:"foreign_address"`

	ForeignFullAddress *oval_sc.EntityItemStringType `xml:"foreign_full_address"`

	ForeignPort *oval_sc.EntityItemStringType `xml:"foreign_port"`

	Pid *oval_sc.EntityItemIntType `xml:"pid"`

	Protocol *oval_sc.EntityItemStringType `xml:"protocol"`

	UserId *oval_sc.EntityItemStringType `xml:"user_id"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type Inetlisteningserver510Item struct {
	XMLName xml.Name `xml:inetlisteningserver510_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Protocol *oval_sc.EntityItemStringType `xml:"protocol"`

	LocalAddress *oval_sc.EntityItemIPAddressStringType `xml:"local_address"`

	LocalPort *oval_sc.EntityItemIntType `xml:"local_port"`

	LocalFullAddress *oval_sc.EntityItemStringType `xml:"local_full_address"`

	ProgramName *oval_sc.EntityItemStringType `xml:"program_name"`

	ForeignAddress *oval_sc.EntityItemIPAddressStringType `xml:"foreign_address"`

	ForeignPort *oval_sc.EntityItemIntType `xml:"foreign_port"`

	ForeignFullAddress *oval_sc.EntityItemStringType `xml:"foreign_full_address"`

	Pid *oval_sc.EntityItemIntType `xml:"pid"`

	UserId *oval_sc.EntityItemIntType `xml:"user_id"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type KeychainItem struct {
	XMLName xml.Name `xml:keychain_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Filepath oval_sc.EntityItemStringType `xml:"filepath"`

	LockOnSleep *oval_sc.EntityItemBoolType `xml:"lock_on_sleep"`

	Timeout *oval_sc.EntityItemIntType `xml:"timeout"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type LaunchdItem struct {
	XMLName xml.Name `xml:launchd_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Label oval_sc.EntityItemStringType `xml:"label"`

	Pid *oval_sc.EntityItemIntType `xml:"pid"`

	StatusElm *oval_sc.EntityItemIntType `xml:"status"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type NvramItem struct {
	XMLName xml.Name `xml:nvram_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	NvramVar *oval_sc.EntityItemStringType `xml:"nvram_var"`

	NvramValue *oval_sc.EntityItemStringType `xml:"nvram_value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PlistItem struct {
	XMLName xml.Name `xml:plist_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Key *oval_sc.EntityItemStringType `xml:"key"`

	AppId *oval_sc.EntityItemStringType `xml:"app_id"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Instance *oval_sc.EntityItemIntType `xml:"instance"`

	Type *EntityItemPlistTypeType `xml:"type"`

	Value []oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type Plist511Item struct {
	XMLName xml.Name `xml:plist511_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	AppId *oval_sc.EntityItemStringType `xml:"app_id"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Xpath *oval_sc.EntityItemStringType `xml:"xpath"`

	ValueOf []oval_sc.EntityItemAnySimpleType `xml:"value_of"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PwpolicyItem struct {
	XMLName xml.Name `xml:pwpolicy_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Username *oval_sc.EntityItemStringType `xml:"username"`

	Userpass *oval_sc.EntityItemStringType `xml:"userpass"`

	DirectoryNode *oval_sc.EntityItemStringType `xml:"directory_node"`

	MaxChars *oval_sc.EntityItemIntType `xml:"maxChars"`

	MaxFailedLoginAttempts *oval_sc.EntityItemIntType `xml:"maxFailedLoginAttempts"`

	MinChars *oval_sc.EntityItemIntType `xml:"minChars"`

	PasswordCannotBeName *oval_sc.EntityItemBoolType `xml:"passwordCannotBeName"`

	RequiresAlpha *oval_sc.EntityItemBoolType `xml:"requiresAlpha"`

	RequiresNumeric *oval_sc.EntityItemBoolType `xml:"requiresNumeric"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type Pwpolicy59Item struct {
	XMLName xml.Name `xml:pwpolicy59_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	TargetUser *oval_sc.EntityItemStringType `xml:"target_user"`

	Username *oval_sc.EntityItemStringType `xml:"username"`

	Userpass *oval_sc.EntityItemStringType `xml:"userpass"`

	DirectoryNode *oval_sc.EntityItemStringType `xml:"directory_node"`

	MaxChars *oval_sc.EntityItemIntType `xml:"maxChars"`

	MaxFailedLoginAttempts *oval_sc.EntityItemIntType `xml:"maxFailedLoginAttempts"`

	MinChars *oval_sc.EntityItemIntType `xml:"minChars"`

	PasswordCannotBeName *oval_sc.EntityItemBoolType `xml:"passwordCannotBeName"`

	RequiresAlpha *oval_sc.EntityItemBoolType `xml:"requiresAlpha"`

	RequiresNumeric *oval_sc.EntityItemBoolType `xml:"requiresNumeric"`

	MaxMinutesUntilChangePassword *oval_sc.EntityItemIntType `xml:"maxMinutesUntilChangePassword"`

	MinMinutesUntilChangePassword *oval_sc.EntityItemIntType `xml:"minMinutesUntilChangePassword"`

	RequiresMixedCase *oval_sc.EntityItemBoolType `xml:"requiresMixedCase"`

	RequiresSymbol *oval_sc.EntityItemBoolType `xml:"requiresSymbol"`

	MinutesUntilFailedLoginReset *oval_sc.EntityItemIntType `xml:"minutesUntilFailedLoginReset"`

	UsingHistory *oval_sc.EntityItemIntType `xml:"usingHistory"`

	CanModifyPasswordforSelf *oval_sc.EntityItemBoolType `xml:"canModifyPasswordforSelf"`

	UsingExpirationDate *oval_sc.EntityItemBoolType `xml:"usingExpirationDate"`

	UsingHardExpirationDate *oval_sc.EntityItemBoolType `xml:"usingHardExpirationDate"`

	ExpirationDateGMT *oval_sc.EntityItemStringType `xml:"expirationDateGMT"`

	HardExpireDateGMT *oval_sc.EntityItemStringType `xml:"hardExpireDateGMT"`

	MaxMinutesUntilDisabled *oval_sc.EntityItemIntType `xml:"maxMinutesUntilDisabled"`

	MaxMinutesOfNonUse *oval_sc.EntityItemIntType `xml:"maxMinutesOfNonUse"`

	NewPasswordRequired *oval_sc.EntityItemBoolType `xml:"newPasswordRequired"`

	NotGuessablePattern *oval_sc.EntityItemBoolType `xml:"notGuessablePattern"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type RlimitItem struct {
	XMLName xml.Name `xml:rlimit_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	CpuCurrent oval_sc.EntityItemIntType `xml:"cpu_current"`

	CpuMax oval_sc.EntityItemIntType `xml:"cpu_max"`

	FilesizeCurrent oval_sc.EntityItemIntType `xml:"filesize_current"`

	FilesizeMax oval_sc.EntityItemIntType `xml:"filesize_max"`

	DataCurrent oval_sc.EntityItemIntType `xml:"data_current"`

	DataMax oval_sc.EntityItemIntType `xml:"data_max"`

	StackCurrent oval_sc.EntityItemIntType `xml:"stack_current"`

	StackMax oval_sc.EntityItemIntType `xml:"stack_max"`

	CoreCurrent oval_sc.EntityItemIntType `xml:"core_current"`

	CoreMax oval_sc.EntityItemIntType `xml:"core_max"`

	RssCurrent oval_sc.EntityItemIntType `xml:"rss_current"`

	RssMax oval_sc.EntityItemIntType `xml:"rss_max"`

	MemlockCurrent oval_sc.EntityItemIntType `xml:"memlock_current"`

	MemlockMax oval_sc.EntityItemIntType `xml:"memlock_max"`

	MaxprocCurrent oval_sc.EntityItemIntType `xml:"maxproc_current"`

	MaxprocMax oval_sc.EntityItemIntType `xml:"maxproc_max"`

	MaxfilesCurrent oval_sc.EntityItemIntType `xml:"maxfiles_current"`

	MaxfilesMax oval_sc.EntityItemIntType `xml:"maxfiles_max"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SoftwareupdateItem struct {
	XMLName xml.Name `xml:softwareupdate_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Schedule oval_sc.EntityItemBoolType `xml:"schedule"`

	SoftwareTitle []oval_sc.EntityItemStringType `xml:"software_title"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SystemprofilerItem struct {
	XMLName xml.Name `xml:systemprofiler_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	DataType *EntityItemDataTypeType `xml:"data_type"`

	Xpath *oval_sc.EntityItemStringType `xml:"xpath"`

	ValueOf []oval_sc.EntityItemAnySimpleType `xml:"value_of"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SystemsetupItem struct {
	XMLName xml.Name `xml:systemsetup_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Timezone oval_sc.EntityItemStringType `xml:"timezone"`

	Usingnetworktime oval_sc.EntityItemBoolType `xml:"usingnetworktime"`

	Networktimeserver *oval_sc.EntityItemStringType `xml:"networktimeserver"`

	Computersleep oval_sc.EntityItemIntType `xml:"computersleep"`

	Displaysleep oval_sc.EntityItemIntType `xml:"displaysleep"`

	Harddisksleep oval_sc.EntityItemIntType `xml:"harddisksleep"`

	Wakeonmodem oval_sc.EntityItemBoolType `xml:"wakeonmodem"`

	Wakeonnetworkaccess oval_sc.EntityItemBoolType `xml:"wakeonnetworkaccess"`

	Restartfreeze oval_sc.EntityItemBoolType `xml:"restartfreeze"`

	Restartpowerfailure oval_sc.EntityItemBoolType `xml:"restartpowerfailure"`

	Allowpowerbuttontosleepcomputer oval_sc.EntityItemBoolType `xml:"allowpowerbuttontosleepcomputer"`

	Remotelogin oval_sc.EntityItemBoolType `xml:"remotelogin"`

	Remoteappleevents *oval_sc.EntityItemBoolType `xml:"remoteappleevents"`

	Computername oval_sc.EntityItemStringType `xml:"computername"`

	Localsubnetname oval_sc.EntityItemStringType `xml:"localsubnetname"`

	Startupdisk oval_sc.EntityItemStringType `xml:"startupdisk"`

	Waitforstartupafterpowerfailure oval_sc.EntityItemIntType `xml:"waitforstartupafterpowerfailure"`

	Disablekeyboardwhenenclosurelockisengaged oval_sc.EntityItemBoolType `xml:"disablekeyboardwhenenclosurelockisengaged"`

	Kernelbootarchitecturesetting oval_sc.EntityItemStringType `xml:"kernelbootarchitecturesetting"`

	Message []oval.MessageType `xml:"message"`
}

// XSD ComplexType declarations

type EntityItemDataTypeType struct {
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

type EntityItemPlistTypeType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations
