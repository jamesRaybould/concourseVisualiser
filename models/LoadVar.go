package models

import (
	"fmt"
	"strings"
)

type LoadVar struct {
	Key      int
	Jobname  string
	Load_var string
	File     string
}

func (loadVar LoadVar) String() string {
	return fmt.Sprintf(`
	object "%s" as load_var_%s_%s_%d #AACA2B {
		file: %s
	}`, loadVar.Load_var, loadVar.Jobname, loadVar.EscapedName(), loadVar.Key, loadVar.File)
}

func (loadVar LoadVar) EscapedName() string {
	return strings.ReplaceAll(loadVar.Load_var, "-", "_")
}
