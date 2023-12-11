package main

import (
	"concourseVisualiser/models"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

type Pipeline struct {
	Name      string
	Jobs      []Job
	Resources []models.Resource
}

type Job struct {
	Name string
	Plan []map[string]interface{}
}

type InParallel struct {
	InParallel []map[string]interface{} `mapstructure:"in_parallel"`
	Key        int
}

func ParsePipeline(pipelineFile string) (*Pipeline, error) {
	// Read the Concourse CI pipeline YAML file
	file, err := os.Open(pipelineFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the Concourse CI pipeline YAML file into a Pipeline struct
	var pipeline Pipeline
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&pipeline)
	if err != nil {
		return nil, err
	}

	return &pipeline, nil
}

func parseTasks(key int, task map[string]interface{}, jobName string, jobOrder *[]string, err error, f *os.File) {
	if task["task"] != nil {
		var result models.Task
		mapstructure.Decode(task, &result)
		if err != nil {
			log.Fatal(err)
		}
		result.Jobname = jobName
		f.WriteString(result.String())
		if jobOrder != nil {
			*jobOrder = append(*jobOrder, jobName+"_"+result.EscapedName())
		}
	}

	if task["get"] != nil {
		var result models.Get
		mapstructure.Decode(task, &result)
		if err != nil {
			log.Fatal(err)
		}
		result.Jobname = jobName

		f.WriteString(result.String())
		if jobOrder != nil {
			*jobOrder = append(*jobOrder, "get_"+jobName+"_"+result.EscapedName())
		}
	}

	if task["put"] != nil {
		var result models.Put
		mapstructure.Decode(task, &result)
		if err != nil {
			log.Fatal(err)
		}
		result.Jobname = jobName
		result.Key = key

		f.WriteString(result.String())
		if jobOrder != nil {
			*jobOrder = append(*jobOrder, fmt.Sprintf("put_%s_%s_%d", jobName, result.EscapedName(), key))
		}
	}

	if task["load_var"] != nil {
		var result models.LoadVar
		mapstructure.Decode(task, &result)
		if err != nil {
			log.Fatal(err)
		}
		result.Jobname = jobName
		result.Key = key

		f.WriteString(result.String())
		if jobOrder != nil {
			*jobOrder = append(*jobOrder, fmt.Sprintf("load_var_%s_%s_%d", jobName, result.EscapedName(), key))
		}
	}

	if task["in_parallel"] != nil {
		var result InParallel
		mapstructure.Decode(task, &result)
		if err != nil {
			log.Fatal(err)
		}
		f.WriteString(fmt.Sprintf("\n\trectangle in_parallel_%s_%d #47a025 [\n\t== In Parallel\n\t{{\n\tskinparam backgroundcolor transparent", jobName, key))

		innerJobOrder := []string{}

		for innerKey, innerTask := range result.InParallel {
			parseTasks(innerKey, innerTask, jobName, &innerJobOrder, err, f)
		}

		for key := range innerJobOrder {
			if (key + 1) == len(innerJobOrder) {
				f.WriteString("\n")
				break
			}

			f.WriteString(fmt.Sprintf(`
	%s -[hidden] %s`, innerJobOrder[key], innerJobOrder[key+1]))
		}
		*jobOrder = append(*jobOrder, fmt.Sprintf("in_parallel_%s_%d", jobName, key))
		f.WriteString("\n\t}}\n\t]")
	}

	if task["set_pipeline"] != nil {
		var result models.SetPipeline
		mapstructure.Decode(task, &result)
		if err != nil {
			log.Fatal(err)
		}
		result.Jobname = jobName
		f.WriteString(result.String())
		*jobOrder = append(*jobOrder, "set_pipeline_"+jobName)
	}

}

func main() {
	// Parse the Concourse CI pipeline YAML file
	inputFile := flag.String("input", "pipeline.yaml", "input file, a well formed ConcourseCI pipeline file")
	outputFile := flag.String("output", "pipeline.puml", "outfile file, filename to output the results")

	flag.Parse()

	pipeline, err := ParsePipeline(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	f, _ := os.Create(*outputFile)
	defer f.Close()

	f.WriteString(`@startuml resources

skinparam arrowThickness 3
skinparam linetype ortho
skinparam tabSize 3

allowmixing

`)

	f.WriteString("rectangle Resources #3891a6 {\n")
	for _, resource := range pipeline.Resources {
		f.WriteString(resource.String())
	}
	f.WriteString("}\n")
	f.WriteString("@enduml\n")
	for _, job := range pipeline.Jobs {

		f.WriteString(fmt.Sprintf(`@startuml pipeline-%s

skinparam arrowThickness 3
skinparam linetype ortho
skinparam tabSize 3

allowmixing
`, job.Name))

		jobOrder := []string{}
		escapedJobName := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(job.Name, "-", "_"), "((", "_"), "))", "_")
		f.WriteString(fmt.Sprintf(`
rectangle "%s" as %s #F5CDA3 {
`, job.Name, escapedJobName))

		for key, task := range job.Plan {
			parseTasks(key, task, escapedJobName, &jobOrder, err, f)
		}

		for key := range jobOrder {
			if (key + 1) == len(jobOrder) {
				break
			}

			f.WriteString(fmt.Sprintf(`
			%s --> %s`, jobOrder[key], jobOrder[key+1]))
		}
		f.WriteString("\n}\n")

		f.WriteString("@enduml\n")
	}

	f.Sync()
}
