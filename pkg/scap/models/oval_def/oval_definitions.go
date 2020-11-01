package oval_def

import "github.com/gocomply/scap/pkg/scap/models/oval"

func (defs *OvalDefinitions) ScapVersion() string {
	return defs.Generator.ScapVersion()
}

func (defs *OvalDefinitions) DefinitionClasses() []oval.ClassEnumeration {
	res := []oval.ClassEnumeration{}
	seen := map[oval.ClassEnumeration]struct{}{}

	for _, def := range defs.Definitions.Definition {
		if _, found := seen[def.Class]; !found {
			res = append(res, def.Class)
			seen[def.Class] = struct{}{}
		}
	}
	return res
}
