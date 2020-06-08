// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#apache
package apache_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type HttpdItem struct {
	XMLName xml.Name `xml:httpd_item`

	Message []oval.MessageType `xml:"message"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	BinaryName *oval_sc.EntityItemStringType `xml:"binary_name"`

	Version *oval_sc.EntityItemVersionType `xml:"version"`
}

// XSD ComplexType declarations
