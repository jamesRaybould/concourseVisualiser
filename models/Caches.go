package models

import (
	"fmt"
)

type Caches struct {
	Path string
}

func (caches Caches) String() string {
	return fmt.Sprintf(`\t\t- path: %s`, caches.Path)
}
