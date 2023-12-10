package models

import (
	"fmt"
	"strings"
)

type Get struct {
	Jobname string
	Get     string
	Trigger bool
}

func (get Get) String() string {
	return fmt.Sprintf(`
	object "get-%s" as get_%s_%s #d5e68d {
		get: %s
		trigger: %t
	}
`, get.Get, get.Jobname, get.EscapedName(), get.Get, get.Trigger)
}

func (get Get) EscapedName() string {
	return strings.ReplaceAll(get.Get, "-", "_")
}
