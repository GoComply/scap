package oval

func (g *GeneratorType) ScapVersion() string {
	if len(g.SchemaVersion) > 0 {
		return g.SchemaVersion[0].Text
	}
	return ""
}
