package inter

import (
	"fmt"
)

func (ocil *Ocil) ScapVersion() string {
	return fmt.Sprintf("%.1f", ocil.Generator.SchemaVersion)
}
