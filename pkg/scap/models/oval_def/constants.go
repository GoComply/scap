package oval_def

type DefinitionClass string

const (
	DefinitionClassUnknow     = ""
	DefinitionClassCompliance = "compliance"
	DefinitionClassInventory  = "inventory"
	DefinitionClassPatch      = "patch"
)
