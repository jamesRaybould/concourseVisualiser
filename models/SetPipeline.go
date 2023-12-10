package models

import "fmt"

type SetPipeline struct {
	Jobname     string
	SetPipeline string `mapstructure:"set_pipeline"`
	File        string
}

func (setPipeline SetPipeline) String() string {
	return fmt.Sprintf(`
	object set_pipeline #CB67B9 {
		name: %s
		file: %s
	}
	`, setPipeline.SetPipeline, setPipeline.File)
}
