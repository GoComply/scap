// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#freebsd
package freebsd_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type PortinfoItem struct {
	XMLName xml.Name `xml:portinfo_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr,omitempty"`

	Pkginst *oval_sc.EntityItemStringType `xml:"pkginst"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Category *oval_sc.EntityItemStringType `xml:"category"`

	Version *PortinfoItemVersion `xml:"version"`

	Vendor *oval_sc.EntityItemStringType `xml:"vendor"`

	Description *oval_sc.EntityItemStringType `xml:"description"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PortinfoItemVersion struct {
	XMLName xml.Name `xml:version`

	Datatype string `xml:"datatype,attr,omitempty"`
}

// XSD ComplexType declarations

// XSD SimpleType declarations
