package cpe_dict

import (
	"fmt"
	"strings"
)

const xmlnsPrefix = "http://cpe.mitre.org/dictionary/"

func (clist *CpeList) ScapVersion() string {
	if clist.Generator != nil {
		return fmt.Sprintf("%.1f", clist.Generator.SchemaVersion)
	}
	xmlns := clist.XMLName.Space
	if strings.HasPrefix(xmlns, xmlnsPrefix) {
		return strings.TrimPrefix(xmlns, xmlnsPrefix)
	}
	return ""
}
