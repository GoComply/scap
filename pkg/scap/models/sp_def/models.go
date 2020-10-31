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

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpwebapplicationObject struct {
	XMLName xml.Name `xml:spwebapplication_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Webapplicationurl *oval_def.EntityObjectStringType `xml:"webapplicationurl"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpwebapplicationState struct {
	XMLName xml.Name `xml:spwebapplication_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

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

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpgroupTest struct {
	XMLName xml.Name `xml:spgroup_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpgroupObject struct {
	XMLName xml.Name `xml:spgroup_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Sitecollectionurl *oval_def.EntityObjectStringType `xml:"sitecollectionurl"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpgroupState struct {
	XMLName xml.Name `xml:spgroup_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	Gname *oval_def.EntityStateStringType `xml:"gname"`

	Autoacceptrequesttojoinleave *oval_def.EntityStateBoolType `xml:"autoacceptrequesttojoinleave"`

	Allowmemberseditmembership *oval_def.EntityStateBoolType `xml:"allowmemberseditmembership"`

	Onlyallowmembersviewmembership *oval_def.EntityStateBoolType `xml:"onlyallowmembersviewmembership"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpwebTest struct {
	XMLName xml.Name `xml:spweb_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpwebObject struct {
	XMLName xml.Name `xml:spweb_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Webcollectionurl *oval_def.EntityObjectStringType `xml:"webcollectionurl"`

	Sitecollectionurl *oval_def.EntityObjectStringType `xml:"sitecollectionurl"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpwebState struct {
	XMLName xml.Name `xml:spweb_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Webcollectionurl *oval_def.EntityStateStringType `xml:"webcollectionurl"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	Secondarysitecolladmin *oval_def.EntityStateStringType `xml:"secondarysitecolladmin"`

	Secondsitecolladminenabled *oval_def.EntityStateBoolType `xml:"secondsitecolladminenabled"`

	Allowanonymousaccess *oval_def.EntityStateBoolType `xml:"allowanonymousaccess"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SplistTest struct {
	XMLName xml.Name `xml:splist_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SplistObject struct {
	XMLName xml.Name `xml:splist_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Spsiteurl *oval_def.EntityObjectStringType `xml:"spsiteurl"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SplistState struct {
	XMLName xml.Name `xml:splist_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Spsiteurl *oval_def.EntityStateStringType `xml:"spsiteurl"`

	Irmenabled *oval_def.EntityStateBoolType `xml:"irmenabled"`

	Enableversioning *oval_def.EntityStateBoolType `xml:"enableversioning"`

	Nocrawl *oval_def.EntityStateBoolType `xml:"nocrawl"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpantivirussettingsTest struct {
	XMLName xml.Name `xml:spantivirussettings_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpantivirussettingsObject struct {
	XMLName xml.Name `xml:spantivirussettings_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Spwebservicename *oval_def.EntityObjectStringType `xml:"spwebservicename"`

	Spfarmname *oval_def.EntityObjectStringType `xml:"spfarmname"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpantivirussettingsState struct {
	XMLName xml.Name `xml:spantivirussettings_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

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

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpsiteadministrationTest struct {
	XMLName xml.Name `xml:spsiteadministration_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpsiteadministrationObject struct {
	XMLName xml.Name `xml:spsiteadministration_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Sitecollectionurl *oval_def.EntityObjectStringType `xml:"sitecollectionurl"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpsiteadministrationState struct {
	XMLName xml.Name `xml:spsiteadministration_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	Storagemaxlevel *oval_def.EntityStateIntType `xml:"storagemaxlevel"`

	Storagewarninglevel *oval_def.EntityStateIntType `xml:"storagewarninglevel"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpsiteTest struct {
	XMLName xml.Name `xml:spsite_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpsiteObject struct {
	XMLName xml.Name `xml:spsite_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Sitecollectionurl *oval_def.EntityObjectStringType `xml:"sitecollectionurl"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpsiteState struct {
	XMLName xml.Name `xml:spsite_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	Quotaname *oval_def.EntityStateStringType `xml:"quotaname"`

	Url *oval_def.EntityStateStringType `xml:"url"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpcrawlruleTest struct {
	XMLName xml.Name `xml:spcrawlrule_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpcrawlruleObject struct {
	XMLName xml.Name `xml:spcrawlrule_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Spsiteurl *oval_def.EntityObjectStringType `xml:"spsiteurl"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpcrawlruleState struct {
	XMLName xml.Name `xml:spcrawlrule_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Spsiteurl *oval_def.EntityStateStringType `xml:"spsiteurl"`

	Crawlashttp *oval_def.EntityStateBoolType `xml:"crawlashttp"`

	Enabled *oval_def.EntityStateBoolType `xml:"enabled"`

	Followcomplexurls *oval_def.EntityStateBoolType `xml:"followcomplexurls"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Priority *oval_def.EntityStateIntType `xml:"priority"`

	Suppressindexing *oval_def.EntityStateBoolType `xml:"suppressindexing"`

	Accountname *oval_def.EntityStateStringType `xml:"accountname"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpjobdefinitionTest struct {
	XMLName xml.Name `xml:spjobdefinition_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpjobdefinitionObject struct {
	XMLName xml.Name `xml:spjobdefinition_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Webappuri *oval_def.EntityObjectStringType `xml:"webappuri"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpjobdefinitionState struct {
	XMLName xml.Name `xml:spjobdefinition_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Webappuri *oval_def.EntityStateStringType `xml:"webappuri"`

	Displayname *oval_def.EntityStateStringType `xml:"displayname"`

	Isdisabled *oval_def.EntityStateBoolType `xml:"isdisabled"`

	Retry *oval_def.EntityStateBoolType `xml:"retry"`

	Title *oval_def.EntityStateStringType `xml:"title"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Spjobdefinition510Test struct {
	XMLName xml.Name `xml:spjobdefinition510_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Spjobdefinition510Object struct {
	XMLName xml.Name `xml:spjobdefinition510_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Webappuri *oval_def.EntityObjectStringType `xml:"webappuri"`

	Displayname *oval_def.EntityObjectStringType `xml:"displayname"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Spjobdefinition510State struct {
	XMLName xml.Name `xml:spjobdefinition510_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Webappuri *oval_def.EntityStateStringType `xml:"webappuri"`

	Displayname *oval_def.EntityStateStringType `xml:"displayname"`

	Isdisabled *oval_def.EntityStateBoolType `xml:"isdisabled"`

	Retry *oval_def.EntityStateBoolType `xml:"retry"`

	Title *oval_def.EntityStateStringType `xml:"title"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type BestbetTest struct {
	XMLName xml.Name `xml:bestbet_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type BestbetObject struct {
	XMLName xml.Name `xml:bestbet_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Sitecollectionurl *oval_def.EntityObjectStringType `xml:"sitecollectionurl"`

	Bestbeturl *oval_def.EntityObjectStringType `xml:"bestbeturl"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type BestbetState struct {
	XMLName xml.Name `xml:bestbet_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	Bestbeturl *oval_def.EntityStateStringType `xml:"bestbeturl"`

	Title *oval_def.EntityStateStringType `xml:"title"`

	Description *oval_def.EntityStateStringType `xml:"description"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InfopolicycollTest struct {
	XMLName xml.Name `xml:infopolicycoll_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InfopolicycollObject struct {
	XMLName xml.Name `xml:infopolicycoll_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Sitecollectionurl *oval_def.EntityObjectStringType `xml:"sitecollectionurl"`

	IdElm *oval_def.EntityObjectStringType `xml:"id"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InfopolicycollState struct {
	XMLName xml.Name `xml:infopolicycoll_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Sitecollectionurl *oval_def.EntityStateStringType `xml:"sitecollectionurl"`

	IdElm *oval_def.EntityStateStringType `xml:"id"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Description *oval_def.EntityStateStringType `xml:"description"`

	Longdescription *oval_def.EntityStateStringType `xml:"longdescription"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpdiagnosticsserviceTest struct {
	XMLName xml.Name `xml:spdiagnosticsservice_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpdiagnosticsserviceObject struct {
	XMLName xml.Name `xml:spdiagnosticsservice_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Farmname *oval_def.EntityObjectStringType `xml:"farmname"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpdiagnosticsserviceState struct {
	XMLName xml.Name `xml:spdiagnosticsservice_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Farmname *oval_def.EntityStateStringType `xml:"farmname"`

	Displayname *oval_def.EntityStateStringType `xml:"displayname"`

	Logcutinterval *oval_def.EntityStateIntType `xml:"logcutinterval"`

	Loglocation *oval_def.EntityStateStringType `xml:"loglocation"`

	Logstokeep *oval_def.EntityStateIntType `xml:"logstokeep"`

	Required *oval_def.EntityStateBoolType `xml:"required"`

	Typename *oval_def.EntityStateStringType `xml:"typename"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpdiagnosticslevelTest struct {
	XMLName xml.Name `xml:spdiagnosticslevel_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpdiagnosticslevelObject struct {
	XMLName xml.Name `xml:spdiagnosticslevel_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Farmname *oval_def.EntityObjectStringType `xml:"farmname"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SpdiagnosticslevelState struct {
	XMLName xml.Name `xml:spdiagnosticslevel_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Farmname *oval_def.EntityStateStringType `xml:"farmname"`

	Eventseverity *EntityStateEventSeverityType `xml:"eventseverity"`

	Hidden *oval_def.EntityStateBoolType `xml:"hidden"`

	Levelid *oval_def.EntityStateStringType `xml:"levelid"`

	Levelname *oval_def.EntityStateStringType `xml:"levelname"`

	Traceseverity *EntityStateTraceSeverityType `xml:"traceseverity"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SppolicyfeatureTest struct {
	XMLName xml.Name `xml:sppolicyfeature_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SppolicyfeatureObject struct {
	XMLName xml.Name `xml:sppolicyfeature_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Farmname *oval_def.EntityObjectStringType `xml:"farmname"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SppolicyfeatureState struct {
	XMLName xml.Name `xml:sppolicyfeature_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

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

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SppolicyTest struct {
	XMLName xml.Name `xml:sppolicy_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State *oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SppolicyObject struct {
	XMLName xml.Name `xml:sppolicy_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Webappuri *oval_def.EntityObjectStringType `xml:"webappuri"`

	Urlzone *EntityObjectUrlZoneType `xml:"urlzone"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SppolicyState struct {
	XMLName xml.Name `xml:sppolicy_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Webappuri *oval_def.EntityStateStringType `xml:"webappuri"`

	Urlzone *EntityStateUrlZoneType `xml:"urlzone"`

	Displayname *oval_def.EntityStateStringType `xml:"displayname"`

	Issystemuser *oval_def.EntityStateBoolType `xml:"issystemuser"`

	Username *oval_def.EntityStateStringType `xml:"username"`

	Policyroletype *EntityStatePolicyRoleType `xml:"policyroletype"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// XSD ComplexType declarations

type EntityObjectUrlZoneType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateEventSeverityType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateTraceSeverityType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStatePolicyRoleType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStatePolicyFeatureStateType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateUrlZoneType struct {
	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations
