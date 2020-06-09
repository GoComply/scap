package ds

func (ds *DataStreamCollection) GetComponentById(componentId string) *Component {
	for idx, _ := range ds.Component {
		comp := &ds.Component[idx]
		if comp.Id == componentId {
			return comp
		}
	}
	return nil
}
