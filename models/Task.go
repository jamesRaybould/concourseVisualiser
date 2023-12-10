package models

import (
	"fmt"
	"strings"
)

type Task struct {
	Jobname string
	Task    string
	Image   string
	Privileged bool
	Timeout string
	Attempts int
	Config  Config
	
}

func (task Task) EscapedName() string {
	return strings.ReplaceAll(task.Task, "-", "_")
}

func (task Task) String() string {
	return fmt.Sprintf(`
	object "%s" as %s_%s #e8871e {
		image: %s
		privileged: %t
		config:
			%s
	}
`, task.Task, task.Jobname, task.EscapedName(), task.Image, task.Privileged, task.Config)
}
