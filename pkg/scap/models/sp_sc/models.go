// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#sharepoint
package sp_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type SpwebapplicationItem struct {
	XMLName xml.Name `xml:spwebapplication_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Webapplicationurl *oval_sc.EntityItemStringType `xml:"webapplicationurl"`

	Allowparttopartcommunication *oval_sc.EntityItemBoolType `xml:"allowparttopartcommunication"`

	Allowaccesstowebpartcatalog *oval_sc.EntityItemBoolType `xml:"allowaccesstowebpartcatalog"`

	Blockedfileextention []oval_sc.EntityItemStringType `xml:"blockedfileextention"`

	Defaultquotatemplate *oval_sc.EntityItemStringType `xml:"defaultquotatemplate"`

	Externalworkflowparticipantsenabled *oval_sc.EntityItemBoolType `xml:"externalworkflowparticipantsenabled"`

	Recyclebinenabled *oval_sc.EntityItemBoolType `xml:"recyclebinenabled"`

	Automaticallydeleteunusedsitecollections *oval_sc.EntityItemBoolType `xml:"automaticallydeleteunusedsitecollections"`

	Selfservicesitecreationenabled *oval_sc.EntityItemBoolType `xml:"selfservicesitecreationenabled"`

	Secondstagerecyclebinquota *oval_sc.EntityItemIntType `xml:"secondstagerecyclebinquota"`

	Recyclebinretentionperiod *oval_sc.EntityItemIntType `xml:"recyclebinretentionperiod"`

	Outboundmailserverinstance *oval_sc.EntityItemStringType `xml:"outboundmailserverinstance"`

	Outboundmailsenderaddress *oval_sc.EntityItemStringType `xml:"outboundmailsenderaddress"`

	Outboundmailreplytoaddress *oval_sc.EntityItemStringType `xml:"outboundmailreplytoaddress"`

	Secvalexpires *oval_sc.EntityItemBoolType `xml:"secvalexpires"`

	Timeout *oval_sc.EntityItemIntType `xml:"timeout"`

	Isadministrationwebapplication *oval_sc.EntityItemBoolType `xml:"isadministrationwebapplication"`

	Applicationpoolname *oval_sc.EntityItemStringType `xml:"applicationpoolname"`

	Applicationpoolusername *oval_sc.EntityItemStringType `xml:"applicationpoolusername"`

	Openitems *oval_sc.EntityItemBoolType `xml:"openitems"`

	Addlistitems *oval_sc.EntityItemBoolType `xml:"addlistitems"`

	Approveitems *oval_sc.EntityItemBoolType `xml:"approveitems"`

	Deletelistitems *oval_sc.EntityItemBoolType `xml:"deletelistitems"`

	Deleteversions *oval_sc.EntityItemBoolType `xml:"deleteversions"`

	Editlistitems *oval_sc.EntityItemBoolType `xml:"editlistitems"`

	Managelists *oval_sc.EntityItemBoolType `xml:"managelists"`

	Viewversions *oval_sc.EntityItemBoolType `xml:"viewversions"`

	Viewlistitems *oval_sc.EntityItemBoolType `xml:"viewlistitems"`

	Cancelcheckout *oval_sc.EntityItemBoolType `xml:"cancelcheckout"`

	Createalerts *oval_sc.EntityItemBoolType `xml:"createalerts"`

	Viewformpages *oval_sc.EntityItemBoolType `xml:"viewformpages"`

	Viewpages *oval_sc.EntityItemBoolType `xml:"viewpages"`

	Addandcustomizepages *oval_sc.EntityItemBoolType `xml:"addandcustomizepages"`

	Applystylesheets *oval_sc.EntityItemBoolType `xml:"applystylesheets"`

	Applythemeandborder *oval_sc.EntityItemBoolType `xml:"applythemeandborder"`

	Browsedirectories *oval_sc.EntityItemBoolType `xml:"browsedirectories"`

	Browseuserinfo *oval_sc.EntityItemBoolType `xml:"browseuserinfo"`

	Creategroups *oval_sc.EntityItemBoolType `xml:"creategroups"`

	Createsscsite *oval_sc.EntityItemBoolType `xml:"createsscsite"`

	Editmyuserinfo *oval_sc.EntityItemBoolType `xml:"editmyuserinfo"`

	Enumeratepermissions *oval_sc.EntityItemBoolType `xml:"enumeratepermissions"`

	Managealerts *oval_sc.EntityItemBoolType `xml:"managealerts"`

	Managepermissions *oval_sc.EntityItemBoolType `xml:"managepermissions"`

	Managesubwebs *oval_sc.EntityItemBoolType `xml:"managesubwebs"`

	Manageweb *oval_sc.EntityItemBoolType `xml:"manageweb"`

	Open *oval_sc.EntityItemBoolType `xml:"open"`

	Useclientintegration *oval_sc.EntityItemBoolType `xml:"useclientintegration"`

	Useremoteapis *oval_sc.EntityItemBoolType `xml:"useremoteapis"`

	Viewusagedata *oval_sc.EntityItemBoolType `xml:"viewusagedata"`

	Managepersonalviews *oval_sc.EntityItemBoolType `xml:"managepersonalviews"`

	Adddelprivatewebparts *oval_sc.EntityItemBoolType `xml:"adddelprivatewebparts"`

	Updatepersonalwebparts *oval_sc.EntityItemBoolType `xml:"updatepersonalwebparts"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SpgroupItem struct {
	XMLName xml.Name `xml:spgroup_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Sitecollectionurl *oval_sc.EntityItemStringType `xml:"sitecollectionurl"`

	Gname *oval_sc.EntityItemStringType `xml:"gname"`

	Autoacceptrequesttojoinleave *oval_sc.EntityItemBoolType `xml:"autoacceptrequesttojoinleave"`

	Allowmemberseditmembership *oval_sc.EntityItemBoolType `xml:"allowmemberseditmembership"`

	Onlyallowmembersviewmembership *oval_sc.EntityItemBoolType `xml:"onlyallowmembersviewmembership"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SpwebItem struct {
	XMLName xml.Name `xml:spweb_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Webcollectionurl *oval_sc.EntityItemStringType `xml:"webcollectionurl"`

	Sitecollectionurl *oval_sc.EntityItemStringType `xml:"sitecollectionurl"`

	Secondarysitecolladmin *oval_sc.EntityItemStringType `xml:"secondarysitecolladmin"`

	Secondsitecolladminenabled *oval_sc.EntityItemBoolType `xml:"secondsitecolladminenabled"`

	Allowanonymousaccess *oval_sc.EntityItemBoolType `xml:"allowanonymousaccess"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SplistItem struct {
	XMLName xml.Name `xml:splist_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Spsiteurl *oval_sc.EntityItemStringType `xml:"spsiteurl"`

	Irmenabled *oval_sc.EntityItemBoolType `xml:"irmenabled"`

	Enableversioning *oval_sc.EntityItemBoolType `xml:"enableversioning"`

	Nocrawl *oval_sc.EntityItemBoolType `xml:"nocrawl"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SpantivirussettingsItem struct {
	XMLName xml.Name `xml:spantivirussettings_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Spwebservicename *oval_sc.EntityItemStringType `xml:"spwebservicename"`

	Spfarmname *oval_sc.EntityItemStringType `xml:"spfarmname"`

	Allowdownload *oval_sc.EntityItemBoolType `xml:"allowdownload"`

	Cleaningenabled *oval_sc.EntityItemBoolType `xml:"cleaningenabled"`

	Downloadscanenabled *oval_sc.EntityItemBoolType `xml:"downloadscanenabled"`

	Numberofthreads *oval_sc.EntityItemIntType `xml:"numberofthreads"`

	Skipsearchcrawl *oval_sc.EntityItemBoolType `xml:"skipsearchcrawl"`

	Timeout *oval_sc.EntityItemIntType `xml:"timeout"`

	Uploadscanenabled *oval_sc.EntityItemBoolType `xml:"uploadscanenabled"`

	Vendorupdatecount *oval_sc.EntityItemIntType `xml:"vendorupdatecount"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SpsiteadministrationItem struct {
	XMLName xml.Name `xml:spsiteadministration_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Sitecollectionurl *oval_sc.EntityItemStringType `xml:"sitecollectionurl"`

	Storagemaxlevel *oval_sc.EntityItemIntType `xml:"storagemaxlevel"`

	Storagewarninglevel *oval_sc.EntityItemIntType `xml:"storagewarninglevel"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SpsiteItem struct {
	XMLName xml.Name `xml:spsite_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Sitecollectionurl *oval_sc.EntityItemStringType `xml:"sitecollectionurl"`

	Quotaname *oval_sc.EntityItemStringType `xml:"quotaname"`

	Url *oval_sc.EntityItemStringType `xml:"url"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SpcrawlruleItem struct {
	XMLName xml.Name `xml:spcrawlrule_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Spsiteurl *oval_sc.EntityItemStringType `xml:"spsiteurl"`

	Crawlashttp *oval_sc.EntityItemBoolType `xml:"crawlashttp"`

	Enabled *oval_sc.EntityItemBoolType `xml:"enabled"`

	Followcomplexurls *oval_sc.EntityItemBoolType `xml:"followcomplexurls"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Priority *oval_sc.EntityItemIntType `xml:"priority"`

	Suppressindexing *oval_sc.EntityItemBoolType `xml:"suppressindexing"`

	Accountname *oval_sc.EntityItemStringType `xml:"accountname"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SpjobdefinitionItem struct {
	XMLName xml.Name `xml:spjobdefinition_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Webappuri *oval_sc.EntityItemStringType `xml:"webappuri"`

	Displayname *oval_sc.EntityItemStringType `xml:"displayname"`

	Isdisabled *oval_sc.EntityItemBoolType `xml:"isdisabled"`

	Retry *oval_sc.EntityItemBoolType `xml:"retry"`

	Title *oval_sc.EntityItemStringType `xml:"title"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type Spjobdefinition510Item struct {
	XMLName xml.Name `xml:spjobdefinition510_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Webappuri *oval_sc.EntityItemStringType `xml:"webappuri"`

	Displayname *oval_sc.EntityItemStringType `xml:"displayname"`

	Isdisabled *oval_sc.EntityItemBoolType `xml:"isdisabled"`

	Retry *oval_sc.EntityItemBoolType `xml:"retry"`

	Title *oval_sc.EntityItemStringType `xml:"title"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type BestbetItem struct {
	XMLName xml.Name `xml:bestbet_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Sitecollectionurl *oval_sc.EntityItemStringType `xml:"sitecollectionurl"`

	Bestbeturl *oval_sc.EntityItemStringType `xml:"bestbeturl"`

	Title *oval_sc.EntityItemStringType `xml:"title"`

	Description *oval_sc.EntityItemStringType `xml:"description"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type InfopolicycollItem struct {
	XMLName xml.Name `xml:infopolicycoll_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Sitecollectionurl *oval_sc.EntityItemStringType `xml:"sitecollectionurl"`

	IdElm *oval_sc.EntityItemStringType `xml:"id"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Description *oval_sc.EntityItemStringType `xml:"description"`

	Longdescription *oval_sc.EntityItemStringType `xml:"longdescription"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SpdiagnosticsserviceItem struct {
	XMLName xml.Name `xml:spdiagnosticsservice_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Farmname *oval_sc.EntityItemStringType `xml:"farmname"`

	Displayname *oval_sc.EntityItemStringType `xml:"displayname"`

	Logcutinterval *oval_sc.EntityItemIntType `xml:"logcutinterval"`

	Loglocation *oval_sc.EntityItemStringType `xml:"loglocation"`

	Logstokeep *oval_sc.EntityItemIntType `xml:"logstokeep"`

	Required *oval_sc.EntityItemBoolType `xml:"required"`

	Typename *oval_sc.EntityItemStringType `xml:"typename"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SpdiagnosticslevelItem struct {
	XMLName xml.Name `xml:spdiagnosticslevel_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Farmname *oval_sc.EntityItemStringType `xml:"farmname"`

	Eventseverity *EntityItemEventSeverityType `xml:"eventseverity"`

	Hidden *oval_sc.EntityItemBoolType `xml:"hidden"`

	Levelid *oval_sc.EntityItemStringType `xml:"levelid"`

	Levelname *oval_sc.EntityItemStringType `xml:"levelname"`

	Traceseverity *EntityItemTraceSeverityType `xml:"traceseverity"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SppolicyfeatureItem struct {
	XMLName xml.Name `xml:sppolicyfeature_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Farmname *oval_sc.EntityItemStringType `xml:"farmname"`

	Configpage *oval_sc.EntityItemStringType `xml:"configpage"`

	Defaultcustomdata *oval_sc.EntityItemStringType `xml:"defaultcustomdata"`

	Description *oval_sc.EntityItemStringType `xml:"description"`

	Globalconfigpage *oval_sc.EntityItemStringType `xml:"globalconfigpage"`

	Globalcustomdata *oval_sc.EntityItemStringType `xml:"globalcustomdata"`

	Group *oval_sc.EntityItemStringType `xml:"group"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Publisher *oval_sc.EntityItemStringType `xml:"publisher"`

	State *EntityItemPolicyFeatureStateType `xml:"state"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SppolicyItem struct {
	XMLName xml.Name `xml:sppolicy_item`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	Webappuri *oval_sc.EntityItemStringType `xml:"webappuri"`

	Urlzone *EntityItemUrlZoneType `xml:"urlzone"`

	Displayname *oval_sc.EntityItemStringType `xml:"displayname"`

	Issystemuser *oval_sc.EntityItemBoolType `xml:"issystemuser"`

	Username *oval_sc.EntityItemStringType `xml:"username"`

	Policyroletype *EntityItemPolicyRoleType `xml:"policyroletype"`

	Message []oval.MessageType `xml:"message"`
}

// XSD ComplexType declarations

type EntityItemUrlZoneType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemEventSeverityType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemTraceSeverityType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemPolicyFeatureStateType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemPolicyRoleType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations
