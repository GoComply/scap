// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#android
package android_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_def"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type AppmanagerTest struct {
	XMLName xml.Name `xml:appmanager_test`

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
type AppmanagerObject struct {
	XMLName xml.Name `xml:appmanager_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	PackageName *oval_def.EntityObjectStringType `xml:"package_name"`

	SigningCertificate *oval_def.EntityObjectBinaryType `xml:"signing_certificate"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type AppmanagerState struct {
	XMLName xml.Name `xml:appmanager_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	ApplicationName *oval_def.EntityStateStringType `xml:"application_name"`

	Uid *oval_def.EntityStateStringType `xml:"uid"`

	Gid []oval_def.EntityStateStringType `xml:"gid"`

	PackageName *oval_def.EntityStateStringType `xml:"package_name"`

	DataDirectory *oval_def.EntityStateStringType `xml:"data_directory"`

	VersionElm *oval_def.EntityStateStringType `xml:"version"`

	CurrentStatus *oval_def.EntityStateBoolType `xml:"current_status"`

	Permission *oval_def.EntityStateStringType `xml:"permission"`

	NativeLibDir *oval_def.EntityStateStringType `xml:"native_lib_dir"`

	SigningCertificate []oval_def.EntityStateBinaryType `xml:"signing_certificate"`

	FirstInstallTime *oval_def.EntityStateIntType `xml:"first_install_time"`

	LastUpdateTime *oval_def.EntityStateIntType `xml:"last_update_time"`

	PackageFileLocation *oval_def.EntityStateStringType `xml:"package_file_location"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type BluetoothTest struct {
	XMLName xml.Name `xml:bluetooth_test`

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
type BluetoothObject struct {
	XMLName xml.Name `xml:bluetooth_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type BluetoothState struct {
	XMLName xml.Name `xml:bluetooth_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Discoverable *oval_def.EntityStateBoolType `xml:"discoverable"`

	CurrentStatus *oval_def.EntityStateBoolType `xml:"current_status"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type CameraTest struct {
	XMLName xml.Name `xml:camera_test`

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
type CameraObject struct {
	XMLName xml.Name `xml:camera_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type CameraState struct {
	XMLName xml.Name `xml:camera_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	CameraDisabledPolicy *oval_def.EntityStateBoolType `xml:"camera_disabled_policy"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type CertificateTest struct {
	XMLName xml.Name `xml:certificate_test`

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
type CertificateObject struct {
	XMLName xml.Name `xml:certificate_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type CertificateState struct {
	XMLName xml.Name `xml:certificate_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	TrustedCertificate []oval_def.EntityStateBinaryType `xml:"trusted_certificate"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type DevicesettingsTest struct {
	XMLName xml.Name `xml:devicesettings_test`

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
type DevicesettingsObject struct {
	XMLName xml.Name `xml:devicesettings_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type DevicesettingsState struct {
	XMLName xml.Name `xml:devicesettings_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	AdbEnabled *oval_def.EntityStateBoolType `xml:"adb_enabled"`

	AllowMockLocation *oval_def.EntityStateBoolType `xml:"allow_mock_location"`

	InstallNonMarketApps *oval_def.EntityStateBoolType `xml:"install_non_market_apps"`

	DeviceAdmin []oval_def.EntityStateStringType `xml:"device_admin"`

	AutoTime *oval_def.EntityStateBoolType `xml:"auto_time"`

	AutoTimeZone *oval_def.EntityStateBoolType `xml:"auto_time_zone"`

	UsbMassStorageEnabled *oval_def.EntityStateBoolType `xml:"usb_mass_storage_enabled"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type EncryptionTest struct {
	XMLName xml.Name `xml:encryption_test`

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
type EncryptionObject struct {
	XMLName xml.Name `xml:encryption_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type EncryptionState struct {
	XMLName xml.Name `xml:encryption_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	EncryptionPolicyEnabled *oval_def.EntityStateBoolType `xml:"encryption_policy_enabled"`

	EncryptionStatus *EntityStateEncryptionStatusType `xml:"encryption_status"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type LocationserviceTest struct {
	XMLName xml.Name `xml:locationservice_test`

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
type LocationserviceObject struct {
	XMLName xml.Name `xml:locationservice_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type LocationserviceState struct {
	XMLName xml.Name `xml:locationservice_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	GpsEnabled *oval_def.EntityStateBoolType `xml:"gps_enabled"`

	NetworkEnabled *oval_def.EntityStateBoolType `xml:"network_enabled"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type NetworkTest struct {
	XMLName xml.Name `xml:network_test`

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
type NetworkObject struct {
	XMLName xml.Name `xml:network_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type NetworkState struct {
	XMLName xml.Name `xml:network_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	AirplaneMode *oval_def.EntityStateBoolType `xml:"airplane_mode"`

	NfcEnabled *oval_def.EntityStateBoolType `xml:"nfc_enabled"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type PasswordTest struct {
	XMLName xml.Name `xml:password_test`

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
type PasswordObject struct {
	XMLName xml.Name `xml:password_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type PasswordState struct {
	XMLName xml.Name `xml:password_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	MaxNumFailedUserAuth *oval_def.EntityStateIntType `xml:"max_num_failed_user_auth"`

	PasswordHist *oval_def.EntityStateIntType `xml:"password_hist"`

	PasswordQuality *EntityStatePasswordQualityType `xml:"password_quality"`

	PasswordMinLength *oval_def.EntityStateIntType `xml:"password_min_length"`

	PasswordMinLetters *oval_def.EntityStateIntType `xml:"password_min_letters"`

	PasswordMinLowerCaseLetters *oval_def.EntityStateIntType `xml:"password_min_lower_case_letters"`

	PasswordMinNonLetters *oval_def.EntityStateIntType `xml:"password_min_non_letters"`

	PasswordMinNumeric *oval_def.EntityStateIntType `xml:"password_min_numeric"`

	PasswordMinSymbols *oval_def.EntityStateIntType `xml:"password_min_symbols"`

	PasswordMinUpperCaseLetters *oval_def.EntityStateIntType `xml:"password_min_upper_case_letters"`

	PasswordExpirationTimeout *oval_def.EntityStateIntType `xml:"password_expiration_timeout"`

	PasswordVisible *oval_def.EntityStateBoolType `xml:"password_visible"`

	ActivePasswordSufficient *oval_def.EntityStateBoolType `xml:"active_password_sufficient"`

	CurrentFailedPasswordAttempts *oval_def.EntityStateIntType `xml:"current_failed_password_attempts"`

	ScreenLockTimeout *oval_def.EntityStateIntType `xml:"screen_lock_timeout"`

	KeyguardDisabledFeatures *EntityStateKeyguardDisabledFeaturesType `xml:"keyguard_disabled_features"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type SystemdetailsTest struct {
	XMLName xml.Name `xml:systemdetails_test`

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
type SystemdetailsObject struct {
	XMLName xml.Name `xml:systemdetails_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type SystemdetailsState struct {
	XMLName xml.Name `xml:systemdetails_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Hardware *oval_def.EntityStateStringType `xml:"hardware"`

	Manufacturer *oval_def.EntityStateStringType `xml:"manufacturer"`

	Model *oval_def.EntityStateStringType `xml:"model"`

	Product *oval_def.EntityStateStringType `xml:"product"`

	CpuAbi *oval_def.EntityStateStringType `xml:"cpu_abi"`

	CpuAbi2 *oval_def.EntityStateStringType `xml:"cpu_abi2"`

	BuildFingerprint *oval_def.EntityStateStringType `xml:"build_fingerprint"`

	OsVersionCodeName *oval_def.EntityStateStringType `xml:"os_version_code_name"`

	OsVersionBuildNumber *oval_def.EntityStateStringType `xml:"os_version_build_number"`

	OsVersionReleaseName *oval_def.EntityStateStringType `xml:"os_version_release_name"`

	OsVersionSdkNumber *oval_def.EntityStateIntType `xml:"os_version_sdk_number"`

	HardwareKeystore *oval_def.EntityStateBoolType `xml:"hardware_keystore"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type WifiTest struct {
	XMLName xml.Name `xml:wifi_test`

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
type WifiObject struct {
	XMLName xml.Name `xml:wifi_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type WifiState struct {
	XMLName xml.Name `xml:wifi_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	WifiStatus *oval_def.EntityStateBoolType `xml:"wifi_status"`

	NetworkAvailabilityNotification *oval_def.EntityStateBoolType `xml:"network_availability_notification"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type WifinetworkTest struct {
	XMLName xml.Name `xml:wifinetwork_test`

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
type WifinetworkObject struct {
	XMLName xml.Name `xml:wifinetwork_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Ssid *oval_def.EntityObjectStringType `xml:"ssid"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type WifinetworkState struct {
	XMLName xml.Name `xml:wifinetwork_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Ssid *oval_def.EntityStateStringType `xml:"ssid"`

	Bssid *oval_def.EntityStateStringType `xml:"bssid"`

	AuthAlgorithms []EntityStateWifiAuthAlgorithmType `xml:"auth_algorithms"`

	GroupCiphers []EntityStateWifiGroupCipherType `xml:"group_ciphers"`

	KeyManagement []EntityStateWifiKeyMgmtType `xml:"key_management"`

	PairwiseCiphers []EntityStateWifiPairwiseCipherType `xml:"pairwise_ciphers"`

	Protocols []EntityStateWifiProtocolType `xml:"protocols"`

	HiddenSsid *oval_def.EntityStateBoolType `xml:"hidden_ssid"`

	NetworkId *oval_def.EntityStateIntType `xml:"network_id"`

	Priority *oval_def.EntityStateIntType `xml:"priority"`

	CurrentStatus *EntityStateWifiCurrentStatusType `xml:"current_status"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type TelephonyTest struct {
	XMLName xml.Name `xml:telephony_test`

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
type TelephonyObject struct {
	XMLName xml.Name `xml:telephony_object`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// Element
type TelephonyState struct {
	XMLName xml.Name `xml:telephony_state`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	NetworkType *EntityStateNetworkType `xml:"network_type"`

	SimCountryIso *oval_def.EntityStateStringType `xml:"sim_country_iso"`

	SimOperatorCode *oval_def.EntityStateStringType `xml:"sim_operator_code"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

// XSD ComplexType declarations

type EntityStateEncryptionStatusType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateKeyguardDisabledFeaturesType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateNetworkType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStatePasswordQualityType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateWifiAuthAlgorithmType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateWifiCurrentStatusType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateWifiGroupCipherType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateWifiKeyMgmtType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateWifiPairwiseCipherType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateWifiProtocolType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations
