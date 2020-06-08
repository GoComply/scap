// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
package sp_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_def"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type SpwebapplicationTest struct {
	XMLName xml.Name `xml:spwebapplication_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SpwebapplicationObject struct {
	XMLName xml.Name `xml:spwebapplication_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SpwebapplicationState struct {
	XMLName xml.Name `xml:spwebapplication_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Webapplicationurl *oval_def.EntityStateStringType `xml:"webapplicationurl"`

	Allowparttopartcommunication *oval_def.EntityStateBoolType `xml:"allowparttopartcommunication"`

	Allowaccesstowebpartcatalog *oval_def.EntityStateBoolType `xml:"allowaccesstowebpartcatalog"`

	Blockedfileextention *oval_def.EntityStateStringType `xml:"blockedfileextention"`

	Defaultquotatemplate *oval_def.EntityStateStringType `xml:"defaultquotatemplate"`

	Externalworkflowparticipantsenabled *oval_def.EntityStateBoolType `xml:"externalworkflowparticipantsenabled"`

	Recyclebinenabled *oval_def.EntityStateBoolType `xml:"recyclebinenabled"`

	Automaticallydeleteunusedsitecollections *oval_def.EntityStateBoolType `xml:"automaticallydeleteunusedsitecollections"`

	Selfservicesitecreationenabled *oval_def.EntityStateBoolType `xml:"selfservicesitecreationenabled"`

	Secondstagerecyclebinquota *oval_def.EntityStateIntType `xml:"secondstagerecyclebinquota"`

	Recyclebinretentionperiod *oval_def.EntityStateIntType `xml:"recyclebinretentionperiod"`

	Outboundmailserverinstance *oval_def.EntityStateStringType `xml:"outboundmailserverinstance"`

	Outboundmailsenderaddress *oval_def.EntityStateStringType `xml:"outboundmailsenderaddress"`

	Outboundmailreplytoaddress *oval_def.EntityStateStringType `xml:"outboundmailreplytoaddress"`

	Secvalexpires *oval_def.EntityStateBoolType `xml:"secvalexpires"`

	Timeout *oval_def.EntityStateIntType `xml:"timeout"`

	Isadministrationwebapplication *oval_def.EntityStateBoolType `xml:"isadministrationwebapplication"`

	Applicationpoolname *oval_def.EntityStateStringType `xml:"applicationpoolname"`

	Applicationpoolusername *oval_def.EntityStateStringType `xml:"applicationpoolusername"`

	Openitems *oval_def.EntityStateBoolType `xml:"openitems"`

	Addlistitems *oval_def.EntityStateBoolType `xml:"addlistitems"`

	Approveitems *oval_def.EntityStateBoolType `xml:"approveitems"`

	Deletelistitems *oval_def.EntityStateBoolType `xml:"deletelistitems"`

	Deleteversions *oval_def.EntityStateBoolType `xml:"deleteversions"`

	Editlistitems *oval_def.EntityStateBoolType `xml:"editlistitems"`

	Managelists *oval_def.EntityStateBoolType `xml:"managelists"`

	Viewversions *oval_def.EntityStateBoolType `xml:"viewversions"`

	Viewlistitems *oval_def.EntityStateBoolType `xml:"viewlistitems"`

	Cancelcheckout *oval_def.EntityStateBoolType `xml:"cancelcheckout"`

	Createalerts *oval_def.EntityStateBoolType `xml:"createalerts"`

	Viewformpages *oval_def.EntityStateBoolType `xml:"viewformpages"`

	Viewpages *oval_def.EntityStateBoolType `xml:"viewpages"`

	Addandcustomizepages *oval_def.EntityStateBoolType `xml:"addandcustomizepages"`

	Applystylesheets *oval_def.EntityStateBoolType `xml:"applystylesheets"`

	Applythemeandborder *oval_def.EntityStateBoolType `xml:"applythemeandborder"`

	Browsedirectories *oval_def.EntityStateBoolType `xml:"browsedirectories"`

	Browseuserinfo *oval_def.EntityStateBoolType `xml:"browseuserinfo"`

	Creategroups *oval_def.EntityStateBoolType `xml:"creategroups"`

	Createsscsite *oval_def.EntityStateBoolType `xml:"createsscsite"`

	Editmyuserinfo *oval_def.EntityStateBoolType `xml:"editmyuserinfo"`

	Enumeratepermissions *oval_def.EntityStateBoolType `xml:"enumeratepermissions"`

	Managealerts *oval_def.EntityStateBoolType `xml:"managealerts"`

	Managepermissions *oval_def.EntityStateBoolType `xml:"managepermissions"`

	Managesubwebs *oval_def.EntityStateBoolType `xml:"managesubwebs"`

	Manageweb *oval_def.EntityStateBoolType `xml:"manageweb"`

	Open *oval_def.EntityStateBoolType `xml:"open"`

	Useclientintegration *oval_def.EntityStateBoolType `xml:"useclientintegration"`

	Useremoteapis *oval_def.EntityStateBoolType `xml:"useremoteapis"`

	Viewusagedata *oval_def.EntityStateBoolType `xml:"viewusagedata"`

	Managepersonalviews *oval_def.EntityStateBoolType `xml:"managepersonalviews"`

	Adddelprivatewebparts *oval_def.EntityStateBoolType `xml:"adddelprivatewebparts"`

	Updatepersonalwebparts *oval_def.EntityStateBoolType `xml:"updatepersonalwebparts"`
}

// Element
type SpgroupTest struct {
	XMLName xml.Name `xml:spgroup_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SpgroupObject struct {
	XMLName xml.Name `xml:spgroup_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SpgroupState struct {
	XMLName xml.Name `xml:spgroup_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	Gname *oval_def.EntityStateStringType `xml:"gname"`

	Autoacceptrequesttojoinleave *oval_def.EntityStateBoolType `xml:"autoacceptrequesttojoinleave"`

	Allowmemberseditmembership *oval_def.EntityStateBoolType `xml:"allowmemberseditmembership"`

	Onlyallowmembersviewmembership *oval_def.EntityStateBoolType `xml:"onlyallowmembersviewmembership"`
}

// Element
type SpwebTest struct {
	XMLName xml.Name `xml:spweb_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SpwebObject struct {
	XMLName xml.Name `xml:spweb_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SpwebState struct {
	XMLName xml.Name `xml:spweb_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Webcollectionurl *oval_def.EntityStateStringType `xml:"webcollectionurl"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	Secondarysitecolladmin *oval_def.EntityStateStringType `xml:"secondarysitecolladmin"`

	Secondsitecolladminenabled *oval_def.EntityStateBoolType `xml:"secondsitecolladminenabled"`

	Allowanonymousaccess *oval_def.EntityStateBoolType `xml:"allowanonymousaccess"`
}

// Element
type SplistTest struct {
	XMLName xml.Name `xml:splist_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SplistObject struct {
	XMLName xml.Name `xml:splist_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SplistState struct {
	XMLName xml.Name `xml:splist_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Spsiteurl *oval_def.EntityStateStringType `xml:"spsiteurl"`

	Irmenabled *oval_def.EntityStateBoolType `xml:"irmenabled"`

	Enableversioning *oval_def.EntityStateBoolType `xml:"enableversioning"`

	Nocrawl *oval_def.EntityStateBoolType `xml:"nocrawl"`
}

// Element
type SpantivirussettingsTest struct {
	XMLName xml.Name `xml:spantivirussettings_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SpantivirussettingsObject struct {
	XMLName xml.Name `xml:spantivirussettings_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SpantivirussettingsState struct {
	XMLName xml.Name `xml:spantivirussettings_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Spwebservicename *oval_def.EntityStateStringType `xml:"spwebservicename"`

	Spfarmname *oval_def.EntityStateStringType `xml:"spfarmname"`

	Allowdownload *oval_def.EntityStateBoolType `xml:"allowdownload"`

	Cleaningenabled *oval_def.EntityStateBoolType `xml:"cleaningenabled"`

	Downloadscanenabled *oval_def.EntityStateBoolType `xml:"downloadscanenabled"`

	Numberofthreads *oval_def.EntityStateIntType `xml:"numberofthreads"`

	Skipsearchcrawl *oval_def.EntityStateBoolType `xml:"skipsearchcrawl"`

	Timeout *oval_def.EntityStateIntType `xml:"timeout"`

	Uploadscanenabled *oval_def.EntityStateBoolType `xml:"uploadscanenabled"`

	Vendorupdatecount *oval_def.EntityStateIntType `xml:"vendorupdatecount"`
}

// Element
type SpsiteadministrationTest struct {
	XMLName xml.Name `xml:spsiteadministration_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SpsiteadministrationObject struct {
	XMLName xml.Name `xml:spsiteadministration_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SpsiteadministrationState struct {
	XMLName xml.Name `xml:spsiteadministration_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	Storagemaxlevel *oval_def.EntityStateIntType `xml:"storagemaxlevel"`

	Storagewarninglevel *oval_def.EntityStateIntType `xml:"storagewarninglevel"`
}

// Element
type SpsiteTest struct {
	XMLName xml.Name `xml:spsite_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SpsiteObject struct {
	XMLName xml.Name `xml:spsite_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SpsiteState struct {
	XMLName xml.Name `xml:spsite_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	Quotaname *oval_def.EntityStateStringType `xml:"quotaname"`

	Url *oval_def.EntityStateStringType `xml:"url"`
}

// Element
type SpcrawlruleTest struct {
	XMLName xml.Name `xml:spcrawlrule_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SpcrawlruleObject struct {
	XMLName xml.Name `xml:spcrawlrule_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SpcrawlruleState struct {
	XMLName xml.Name `xml:spcrawlrule_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Spsiteurl *oval_def.EntityStateStringType `xml:"spsiteurl"`

	Crawlashttp *oval_def.EntityStateBoolType `xml:"crawlashttp"`

	Enabled *oval_def.EntityStateBoolType `xml:"enabled"`

	Followcomplexurls *oval_def.EntityStateBoolType `xml:"followcomplexurls"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Priority *oval_def.EntityStateIntType `xml:"priority"`

	Suppressindexing *oval_def.EntityStateBoolType `xml:"suppressindexing"`

	Accountname *oval_def.EntityStateStringType `xml:"accountname"`
}

// Element
type SpjobdefinitionTest struct {
	XMLName xml.Name `xml:spjobdefinition_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SpjobdefinitionObject struct {
	XMLName xml.Name `xml:spjobdefinition_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SpjobdefinitionState struct {
	XMLName xml.Name `xml:spjobdefinition_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Webappuri *oval_def.EntityStateStringType `xml:"webappuri"`

	Displayname *oval_def.EntityStateStringType `xml:"displayname"`

	Isdisabled *oval_def.EntityStateBoolType `xml:"isdisabled"`

	Retry *oval_def.EntityStateBoolType `xml:"retry"`

	Title *oval_def.EntityStateStringType `xml:"title"`
}

// Element
type Spjobdefinition510Test struct {
	XMLName xml.Name `xml:spjobdefinition510_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type Spjobdefinition510Object struct {
	XMLName xml.Name `xml:spjobdefinition510_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type Spjobdefinition510State struct {
	XMLName xml.Name `xml:spjobdefinition510_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Webappuri *oval_def.EntityStateStringType `xml:"webappuri"`

	Displayname *oval_def.EntityStateStringType `xml:"displayname"`

	Isdisabled *oval_def.EntityStateBoolType `xml:"isdisabled"`

	Retry *oval_def.EntityStateBoolType `xml:"retry"`

	Title *oval_def.EntityStateStringType `xml:"title"`
}

// Element
type BestbetTest struct {
	XMLName xml.Name `xml:bestbet_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type BestbetObject struct {
	XMLName xml.Name `xml:bestbet_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type BestbetState struct {
	XMLName xml.Name `xml:bestbet_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	Bestbeturl *oval_def.EntityStateStringType `xml:"bestbeturl"`

	Title *oval_def.EntityStateStringType `xml:"title"`

	Description *oval_def.EntityStateStringType `xml:"description"`
}

// Element
type InfopolicycollTest struct {
	XMLName xml.Name `xml:infopolicycoll_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type InfopolicycollObject struct {
	XMLName xml.Name `xml:infopolicycoll_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type InfopolicycollState struct {
	XMLName xml.Name `xml:infopolicycoll_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	Id *oval_def.EntityStateStringType `xml:"id"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Description *oval_def.EntityStateStringType `xml:"description"`

	Longdescription *oval_def.EntityStateStringType `xml:"longdescription"`
}

// Element
type SpdiagnosticsserviceTest struct {
	XMLName xml.Name `xml:spdiagnosticsservice_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SpdiagnosticsserviceObject struct {
	XMLName xml.Name `xml:spdiagnosticsservice_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SpdiagnosticsserviceState struct {
	XMLName xml.Name `xml:spdiagnosticsservice_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Farmname *oval_def.EntityStateStringType `xml:"farmname"`

	Displayname *oval_def.EntityStateStringType `xml:"displayname"`

	Logcutinterval *oval_def.EntityStateIntType `xml:"logcutinterval"`

	Loglocation *oval_def.EntityStateStringType `xml:"loglocation"`

	Logstokeep *oval_def.EntityStateIntType `xml:"logstokeep"`

	Required *oval_def.EntityStateBoolType `xml:"required"`

	Typename *oval_def.EntityStateStringType `xml:"typename"`
}

// Element
type SpdiagnosticslevelTest struct {
	XMLName xml.Name `xml:spdiagnosticslevel_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SpdiagnosticslevelObject struct {
	XMLName xml.Name `xml:spdiagnosticslevel_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SpdiagnosticslevelState struct {
	XMLName xml.Name `xml:spdiagnosticslevel_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Farmname *oval_def.EntityStateStringType `xml:"farmname"`

	Eventseverity *EntityStateEventSeverityType `xml:"eventseverity"`

	Hidden *oval_def.EntityStateBoolType `xml:"hidden"`

	Levelid *oval_def.EntityStateStringType `xml:"levelid"`

	Levelname *oval_def.EntityStateStringType `xml:"levelname"`

	Traceseverity *EntityStateTraceSeverityType `xml:"traceseverity"`
}

// Element
type SppolicyfeatureTest struct {
	XMLName xml.Name `xml:sppolicyfeature_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SppolicyfeatureObject struct {
	XMLName xml.Name `xml:sppolicyfeature_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SppolicyfeatureState struct {
	XMLName xml.Name `xml:sppolicyfeature_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Farmname *oval_def.EntityStateStringType `xml:"farmname"`

	Configpage *oval_def.EntityStateStringType `xml:"configpage"`

	Defaultcustomdata *oval_def.EntityStateStringType `xml:"defaultcustomdata"`

	Description *oval_def.EntityStateStringType `xml:"description"`

	Globalconfigpage *oval_def.EntityStateStringType `xml:"globalconfigpage"`

	Globalcustomdata *oval_def.EntityStateStringType `xml:"globalcustomdata"`

	Group *oval_def.EntityStateStringType `xml:"group"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Publisher *oval_def.EntityStateStringType `xml:"publisher"`

	State *EntityStatePolicyFeatureStateType `xml:"state"`
}

// Element
type SppolicyTest struct {
	XMLName xml.Name `xml:sppolicy_test`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`
}

// Element
type SppolicyObject struct {
	XMLName xml.Name `xml:sppolicy_object`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Set oval_def.Set `xml:"set"`
}

// Element
type SppolicyState struct {
	XMLName xml.Name `xml:sppolicy_state`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	Webappuri *oval_def.EntityStateStringType `xml:"webappuri"`

	Urlzone *EntityStateUrlZoneType `xml:"urlzone"`

	Displayname *oval_def.EntityStateStringType `xml:"displayname"`

	Issystemuser *oval_def.EntityStateBoolType `xml:"issystemuser"`

	Username *oval_def.EntityStateStringType `xml:"username"`

	Policyroletype *EntityStatePolicyRoleType `xml:"policyroletype"`
}

// XSD ComplexType declarations

type EntityObjectUrlZoneType struct {
}

type EntityStateEventSeverityType struct {
}

type EntityStateTraceSeverityType struct {
}

type EntityStatePolicyRoleType struct {
}

type EntityStatePolicyFeatureStateType struct {
}

type EntityStateUrlZoneType struct {
}
