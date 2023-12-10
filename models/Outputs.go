package models

import "fmt"

type Outputs struct {
	Name string
}

func (output Outputs) String() string {
	return fmt.Sprintf(`\t\t- %s`, output.Name)
}
