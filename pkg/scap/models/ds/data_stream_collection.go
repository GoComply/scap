package ds

import (
	"strings"
)

func (ds *DataStreamCollection) ScapVersion() string {
	return ds.SchematronVersion
}

func (ds *DataStreamCollection) GetComponentByRef(ref *ComponentRef) *Component {
	if strings.HasPrefix(ref.XlinkHref, "#") {
		return ds.GetComponentById(strings.TrimPrefix(ref.XlinkHref, "#"))
	}
	return nil
}

func (ds *DataStreamCollection) GetComponentById(componentId string) *Component {
	for idx := range ds.Component {
		comp := &ds.Component[idx]
		if comp.Id == componentId {
			return comp
		}
	}
	return nil
}
