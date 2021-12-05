// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#pixos
package pixos_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type LineItem struct {
	XMLName xml.Name `xml:"line_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	ShowSubcommand *oval_sc.EntityItemStringType `xml:"show_subcommand"`

	ConfigLine *oval_sc.EntityItemStringType `xml:"config_line"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type VersionItem struct {
	XMLName xml.Name `xml:"version_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	PixRelease *oval_sc.EntityItemStringType `xml:"pix_release"`

	PixMajorRelease *oval_sc.EntityItemVersionType `xml:"pix_major_release"`

	PixMinorRelease *oval_sc.EntityItemVersionType `xml:"pix_minor_release"`

	PixBuild *oval_sc.EntityItemIntType `xml:"pix_build"`

	Message []oval.MessageType `xml:"message"`
}

// XSD ComplexType declarations

// XSD SimpleType declarations
