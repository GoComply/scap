// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios
package apple_ios_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_def"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type GlobalrestrictionsTest struct {
	XMLName xml.Name `xml:globalrestrictions_test`

	Id oval.TestIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	CheckExistence oval.ExistenceEnumeration `xml:"check_existence,attr,omitempty"`

	Check oval.CheckEnumeration `xml:"check,attr"`

	StateOperator oval.OperatorEnumeration `xml:"state_operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type GlobalrestrictionsObject struct {
	XMLName xml.Name `xml:globalrestrictions_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type GlobalrestrictionsState struct {
	XMLName xml.Name `xml:globalrestrictions_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	AllowAccountModification *oval_def.EntityStateBoolType `xml:"allow_account_modification"`

	AllowAirdrop *oval_def.EntityStateBoolType `xml:"allow_airdrop"`

	AllowAppCellularDataModification *oval_def.EntityStateBoolType `xml:"allow_app_cellular_data_modification"`

	AllowAppInstallation *oval_def.EntityStateBoolType `xml:"allow_app_installation"`

	AllowAssistant *oval_def.EntityStateBoolType `xml:"allow_assistant"`

	AllowAssistantUserGeneratedContent *oval_def.EntityStateBoolType `xml:"allow_assistant_user_generated_content"`

	AllowAssistantWhileLocked *oval_def.EntityStateBoolType `xml:"allow_assistant_while_locked"`

	AllowBookstore *oval_def.EntityStateBoolType `xml:"allow_bookstore"`

	AllowBookstoreErotica *oval_def.EntityStateBoolType `xml:"allow_bookstore_erotica"`

	AllowCamera *oval_def.EntityStateBoolType `xml:"allow_camera"`

	AllowCloudBackup *oval_def.EntityStateBoolType `xml:"allow_cloud_backup"`

	AllowCloudDocumentSync *oval_def.EntityStateBoolType `xml:"allow_cloud_document_sync"`

	AllowCloudKeychainSync *oval_def.EntityStateBoolType `xml:"allow_cloud_keychain_sync"`

	AllowDiagnosticSubmission *oval_def.EntityStateBoolType `xml:"allow_diagnostic_submission"`

	AllowExplicitContent *oval_def.EntityStateBoolType `xml:"allow_explicit_content"`

	AllowFindMyFriendsModification *oval_def.EntityStateBoolType `xml:"allow_find_my_friends_modification"`

	AllowFingerprintForUnlock *oval_def.EntityStateBoolType `xml:"allow_fingerprint_for_unlock"`

	AllowGameCenter *oval_def.EntityStateBoolType `xml:"allow_game_center"`

	AllowHostPairing *oval_def.EntityStateBoolType `xml:"allow_host_pairing"`

	AllowLockScreenControlCenter *oval_def.EntityStateBoolType `xml:"allow_lock_screen_control_center"`

	AllowLockScreenNotificationsView *oval_def.EntityStateBoolType `xml:"allow_lock_screen_notifications_view"`

	AllowLockScreenTodayView *oval_def.EntityStateBoolType `xml:"allow_lock_screen_today_view"`

	AllowOpenFromManagedToUnmanaged *oval_def.EntityStateBoolType `xml:"allow_open_from_managed_to_unmanaged"`

	AllowOpenFromUnmanagedToManaged *oval_def.EntityStateBoolType `xml:"allow_open_from_unmanaged_to_managed"`

	AllowOtaPkiUpdates *oval_def.EntityStateBoolType `xml:"allow_ota_pki_updates"`

	AllowPassbookWhileLocked *oval_def.EntityStateBoolType `xml:"allow_passbook_while_locked"`

	AllowPhotoStream *oval_def.EntityStateBoolType `xml:"allow_photo_stream"`

	AllowSafari *oval_def.EntityStateBoolType `xml:"allow_safari"`

	AllowScreenShot *oval_def.EntityStateBoolType `xml:"allow_screen_shot"`

	AllowSharedStream *oval_def.EntityStateBoolType `xml:"allow_shared_stream"`

	AllowUiConfigurationProfileInstallation *oval_def.EntityStateBoolType `xml:"allow_ui_configuration_profile_installation"`

	AllowUntrustedTlsPrompt *oval_def.EntityStateBoolType `xml:"allow_untrusted_tls_prompt"`

	AllowVoiceDialing *oval_def.EntityStateBoolType `xml:"allow_voice_dialing"`

	AllowYoutube *oval_def.EntityStateBoolType `xml:"allow_youtube"`

	AllowItunes *oval_def.EntityStateBoolType `xml:"allow_itunes"`

	AutonomousSingleAppModePermittedAppids *oval_def.EntityStateStringType `xml:"autonomous_single_app_mode_permitted_appids"`

	ForceEncryptedBackup *oval_def.EntityStateBoolType `xml:"force_encrypted_backup"`

	ForceItunesStorePasswordEntry *oval_def.EntityStateBoolType `xml:"force_itunes_store_password_entry"`

	ForceLimitAdTracking *oval_def.EntityStateBoolType `xml:"force_limit_ad_tracking"`

	SafariAllowAutoFill *oval_def.EntityStateBoolType `xml:"safari_allow_auto_fill"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type PasscodepolicyTest struct {
	XMLName xml.Name `xml:passcodepolicy_test`

	Id oval.TestIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	CheckExistence oval.ExistenceEnumeration `xml:"check_existence,attr,omitempty"`

	Check oval.CheckEnumeration `xml:"check,attr"`

	StateOperator oval.OperatorEnumeration `xml:"state_operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type PasscodepolicyObject struct {
	XMLName xml.Name `xml:passcodepolicy_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type PasscodepolicyState struct {
	XMLName xml.Name `xml:passcodepolicy_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	AllowSimple *oval_def.EntityStateBoolType `xml:"allow_simple"`

	ForcePin *oval_def.EntityStateBoolType `xml:"force_pin"`

	MaxFailedAttempts *oval_def.EntityStateIntType `xml:"max_failed_attempts"`

	MaxInactivity *oval_def.EntityStateIntType `xml:"max_inactivity"`

	MaxPinAgeInDays *oval_def.EntityStateIntType `xml:"max_pin_age_in_days"`

	MinComplexChars *oval_def.EntityStateIntType `xml:"min_complex_chars"`

	MinLength *oval_def.EntityStateIntType `xml:"min_length"`

	RequireAlphanumeric *oval_def.EntityStateBoolType `xml:"require_alphanumeric"`

	PinHistory *oval_def.EntityStateIntType `xml:"pin_history"`

	MaxGracePeriod *oval_def.EntityStateIntType `xml:"max_grace_period"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type ProfileTest struct {
	XMLName xml.Name `xml:profile_test`

	Id oval.TestIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	CheckExistence oval.ExistenceEnumeration `xml:"check_existence,attr,omitempty"`

	Check oval.CheckEnumeration `xml:"check,attr"`

	StateOperator oval.OperatorEnumeration `xml:"state_operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type ProfileObject struct {
	XMLName xml.Name `xml:profile_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Identifier *oval_def.EntityObjectStringType `xml:"identifier"`

	Uuid *oval_def.EntityObjectStringType `xml:"uuid"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type ProfileState struct {
	XMLName xml.Name `xml:profile_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	HasRemovalPasscode *oval_def.EntityStateBoolType `xml:"has_removal_passcode"`

	IsEncrypted *oval_def.EntityStateBoolType `xml:"is_encrypted"`

	Payload *oval_def.EntityStateRecordType `xml:"payload"`

	Description *oval_def.EntityStateStringType `xml:"description"`

	DisplayName *oval_def.EntityStateStringType `xml:"display_name"`

	Identifier *oval_def.EntityStateStringType `xml:"identifier"`

	Organization *oval_def.EntityStateStringType `xml:"organization"`

	RemovalDisallowed *oval_def.EntityStateBoolType `xml:"removal_disallowed"`

	Uuid *oval_def.EntityStateStringType `xml:"uuid"`

	VersionElm *oval_def.EntityStateIntType `xml:"version"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// XSD ComplexType declarations

// XSD SimpleType declarations
