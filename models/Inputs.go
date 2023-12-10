package models

import (
	"fmt"
)

type Inputs struct {
	Name string
}

func (input Inputs) String() string {
	return fmt.Sprintf(`\t\t- %s`, input.Name)
}
