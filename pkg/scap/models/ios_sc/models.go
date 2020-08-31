// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#ios
package ios_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type AclItem struct {
	XMLName xml.Name `xml:acl_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	IpVersion *EntityItemAccessListIPVersionType `xml:"ip_version"`

	Use *EntityItemAccessListUseType `xml:"use"`

	UsedIn *oval_sc.EntityItemStringType `xml:"used_in"`

	InterfaceDirection *EntityItemAccessListInterfaceDirectionType `xml:"interface_direction"`

	AclConfigLines *oval_sc.EntityItemStringType `xml:"acl_config_lines"`

	ConfigLine []oval_sc.EntityItemStringType `xml:"config_line"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type BgpneighborItem struct {
	XMLName xml.Name `xml:bgpneighbor_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Neighbor *oval_sc.EntityItemStringType `xml:"neighbor"`

	Password *oval_sc.EntityItemStringType `xml:"password"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type GlobalItem struct {
	XMLName xml.Name `xml:global_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	GlobalCommand *oval_sc.EntityItemStringType `xml:"global_command"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type InterfaceItem struct {
	XMLName xml.Name `xml:interface_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	IpDirectedBroadcastCommand *InterfaceItemIpDirectedBroadcastCommand `xml:"ip_directed_broadcast_command"`

	NoIpDirectedBroadcastCommand *oval_sc.EntityItemStringType `xml:"no_ip_directed_broadcast_command"`

	ProxyArpCommand *InterfaceItemProxyArpCommand `xml:"proxy_arp_command"`

	ShutdownCommand *InterfaceItemShutdownCommand `xml:"shutdown_command"`

	HardwareAddr *oval_sc.EntityItemStringType `xml:"hardware_addr"`

	Ipv4Address *oval_sc.EntityItemIPAddressStringType `xml:"ipv4_address"`

	Ipv6Address []oval_sc.EntityItemIPAddressStringType `xml:"ipv6_address"`

	Ipv4AccessList []oval_sc.EntityItemStringType `xml:"ipv4_access_list"`

	Ipv6AccessList []oval_sc.EntityItemStringType `xml:"ipv6_access_list"`

	CryptoMap *oval_sc.EntityItemStringType `xml:"crypto_map"`

	Ipv4UrpfCommand *oval_sc.EntityItemStringType `xml:"ipv4_urpf_command"`

	Ipv6UrpfCommand *oval_sc.EntityItemStringType `xml:"ipv6_urpf_command"`

	UrpfCommand *oval_sc.EntityItemStringType `xml:"urpf_command"`

	SwitchportTrunkEncapsulation *EntityItemTrunkEncapType `xml:"switchport_trunk_encapsulation"`

	SwitchportMode *EntityItemSwitchportModeType `xml:"switchport_mode"`

	SwitchportNativeVlan *InterfaceItemSwitchportNativeVlan `xml:"switchport_native_vlan"`

	SwitchportAccessVlan *InterfaceItemSwitchportAccessVlan `xml:"switchport_access_vlan"`

	SwitchportTrunkedVlans *oval_sc.EntityItemStringType `xml:"switchport_trunked_vlans"`

	SwitchportPrunedVlans *oval_sc.EntityItemStringType `xml:"switchport_pruned_vlans"`

	SwitchportPortSecurity *oval_sc.EntityItemStringType `xml:"switchport_port_security"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type LineItem struct {
	XMLName xml.Name `xml:line_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	ShowSubcommand *oval_sc.EntityItemStringType `xml:"show_subcommand"`

	ConfigLine *oval_sc.EntityItemStringType `xml:"config_line"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type RouterItem struct {
	XMLName xml.Name `xml:router_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Protocol *EntityItemRoutingProtocolType `xml:"protocol"`

	IdElm *oval_sc.EntityItemIntType `xml:"id"`

	Network []oval_sc.EntityItemStringType `xml:"network"`

	BgpNeighbor []oval_sc.EntityItemStringType `xml:"bgp_neighbor"`

	OspfAuthenticationArea []RouterItemOspfAuthenticationArea `xml:"ospf_authentication_area"`

	RouterConfigLines *oval_sc.EntityItemStringType `xml:"router_config_lines"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type RoutingprotocolauthintfItem struct {
	XMLName xml.Name `xml:routingprotocolauthintf_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Interface *oval_sc.EntityItemStringType `xml:"interface"`

	Protocol *EntityItemRoutingProtocolType `xml:"protocol"`

	IdElm *oval_sc.EntityItemIntType `xml:"id"`

	AuthType *EntityItemRoutingAuthTypeStringType `xml:"auth_type"`

	OspfArea *RoutingprotocolauthintfItemOspfArea `xml:"ospf_area"`

	KeyChain *oval_sc.EntityItemStringType `xml:"key_chain"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SectionItem struct {
	XMLName xml.Name `xml:section_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	SectionCommand *oval_sc.EntityItemStringType `xml:"section_command"`

	SectionConfigLines *oval_sc.EntityItemStringType `xml:"section_config_lines"`

	ConfigLine []oval_sc.EntityItemStringType `xml:"config_line"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SnmpItem struct {
	XMLName xml.Name `xml:snmp_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	AccessList *oval_sc.EntityItemStringType `xml:"access_list"`

	CommunityName *oval_sc.EntityItemStringType `xml:"community_name"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SnmpcommunityItem struct {
	XMLName xml.Name `xml:snmpcommunity_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	View *oval_sc.EntityItemStringType `xml:"view"`

	Mode *EntityItemSNMPModeStringType `xml:"mode"`

	Ipv4Acl *oval_sc.EntityItemStringType `xml:"ipv4_acl"`

	Ipv6Acl *oval_sc.EntityItemStringType `xml:"ipv6_acl"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SnmpgroupItem struct {
	XMLName xml.Name `xml:snmpgroup_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Version *EntityItemSNMPVersionStringType `xml:"version"`

	Snmpv3SecLevel *EntityItemSNMPSecLevelStringType `xml:"snmpv3_sec_level"`

	Ipv4Acl *oval_sc.EntityItemStringType `xml:"ipv4_acl"`

	Ipv6Acl *oval_sc.EntityItemStringType `xml:"ipv6_acl"`

	ReadView *oval_sc.EntityItemStringType `xml:"read_view"`

	WriteView *oval_sc.EntityItemStringType `xml:"write_view"`

	NotifyView *oval_sc.EntityItemStringType `xml:"notify_view"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SnmphostItem struct {
	XMLName xml.Name `xml:snmphost_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Host *oval_sc.EntityItemStringType `xml:"host"`

	CommunityOrUser *oval_sc.EntityItemStringType `xml:"community_or_user"`

	Version *EntityItemSNMPVersionStringType `xml:"version"`

	Snmpv3SecLevel *EntityItemSNMPSecLevelStringType `xml:"snmpv3_sec_level"`

	Traps *oval_sc.EntityItemStringType `xml:"traps"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SnmpuserItem struct {
	XMLName xml.Name `xml:snmpuser_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Group *oval_sc.EntityItemStringType `xml:"group"`

	Version *EntityItemSNMPVersionStringType `xml:"version"`

	Ipv4Acl *oval_sc.EntityItemStringType `xml:"ipv4_acl"`

	Ipv6Acl *oval_sc.EntityItemStringType `xml:"ipv6_acl"`

	Priv *EntityItemSNMPPrivStringType `xml:"priv"`

	Auth *EntityItemSNMPAuthStringType `xml:"auth"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SnmpviewItem struct {
	XMLName xml.Name `xml:snmpview_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	MibFamily *oval_sc.EntityItemStringType `xml:"mib_family"`

	Include *oval_sc.EntityItemBoolType `xml:"include"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type TclshItem struct {
	XMLName xml.Name `xml:tclsh_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Available *oval_sc.EntityItemBoolType `xml:"available"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type VersionItem struct {
	XMLName xml.Name `xml:version_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	MajorRelease *oval_sc.EntityItemStringType `xml:"major_release"`

	TrainNumber *oval_sc.EntityItemStringType `xml:"train_number"`

	MajorVersion *oval_sc.EntityItemIntType `xml:"major_version"`

	MinorVersion *oval_sc.EntityItemIntType `xml:"minor_version"`

	Release *oval_sc.EntityItemIntType `xml:"release"`

	TrainIdentifier *oval_sc.EntityItemStringType `xml:"train_identifier"`

	Rebuild *oval_sc.EntityItemIntType `xml:"rebuild"`

	Subrebuild *oval_sc.EntityItemStringType `xml:"subrebuild"`

	MainlineRebuild *oval_sc.EntityItemStringType `xml:"mainline_rebuild"`

	VersionString *oval_sc.EntityItemIOSVersionType `xml:"version_string"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type InterfaceItemIpDirectedBroadcastCommand struct {
	XMLName xml.Name `xml:ip_directed_broadcast_command`
}

// Element
type InterfaceItemProxyArpCommand struct {
	XMLName xml.Name `xml:proxy_arp_command`
}

// Element
type InterfaceItemShutdownCommand struct {
	XMLName xml.Name `xml:shutdown_command`
}

// Element
type InterfaceItemSwitchportNativeVlan struct {
	XMLName xml.Name `xml:switchport_native_vlan`
}

// Element
type InterfaceItemSwitchportAccessVlan struct {
	XMLName xml.Name `xml:switchport_access_vlan`
}

// Element
type RouterItemOspfAuthenticationArea struct {
	XMLName xml.Name `xml:ospf_authentication_area`
}

// Element
type RoutingprotocolauthintfItemOspfArea struct {
	XMLName xml.Name `xml:ospf_area`
}

// XSD ComplexType declarations

type EntityItemAccessListInterfaceDirectionType struct {
}

type EntityItemAccessListIPVersionType struct {
}

type EntityItemAccessListUseType struct {
}

type EntityItemRoutingAuthTypeStringType struct {
}

type EntityItemRoutingProtocolType struct {
}

type EntityItemSNMPVersionStringType struct {
}

type EntityItemSNMPSecLevelStringType struct {
}

type EntityItemSNMPModeStringType struct {
}

type EntityItemSNMPAuthStringType struct {
}

type EntityItemSNMPPrivStringType struct {
}

type EntityItemSwitchportModeType struct {
}

type EntityItemTrunkEncapType struct {
}
