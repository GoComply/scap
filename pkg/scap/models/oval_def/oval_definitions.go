package oval_def

func (defs *OvalDefinitions) ScapVersion() string {
	return defs.Generator.ScapVersion()
}

func (defs *OvalDefinitions) DefinitionClasses() []DefinitionClass {
	res := []DefinitionClass{}
	seen := map[string]struct{}{}

	for _, def := range defs.Definitions.Definition {
		if _, found := seen[string(def.Class)]; !found {
			res = append(res, DefinitionClass(def.Class))
			seen[string(def.Class)] = struct{}{}
		}
	}
	return res
}
