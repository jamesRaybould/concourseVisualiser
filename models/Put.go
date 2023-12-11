package models

import (
	"fmt"
	"strings"
)

type Put struct {
	Key      int
	Jobname  string
	Put      string
	Inputs   string
	No_get   bool
	Resource string
	Params   map[string]string
}

func (put Put) String() string {
	return fmt.Sprintf(`
	object "%s" as put_%s_%s_%d #CB67B9 {
		resource: %s
		name: %s
		inputs: %s
		no_get: %t
	}`, put.Put, put.Jobname, put.EscapedName(), put.Key, put.Resource, put.Put, put.Inputs, put.No_get)
}

func (put Put) EscapedName() string {
	return strings.ReplaceAll(put.Put, "-", "_")
}
