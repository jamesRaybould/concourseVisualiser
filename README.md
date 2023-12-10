# Concourse pipeline visualiser

## Getting started

```shell
./concourseVisualiser --input your_concourse_pipeline_here.yaml --output your_output_file_here.puml
```

- `--input`: default `pipeline.yaml`
- `--output`: default `pipeline.puml`

## Why?

Although being able to work with YAML means that you can have a certain degree of consistency once a file passes a certain size it becomes hard to reason about what is happening, when, in a pipeline.

This tool is to help make it a little easier to see what is happening without having to upload the pipeline to a concourse instance.

## Outputting images

To output images pass `pipeline.puml` through the PlantUML CLI, you can use the docker container like so:

```shell
docker run -v .:/data --rm -i plantuml/plantuml -o ./out pipeline.puml
```

`-v .:/data` mounts the local directory to the internal docker folder of `/data` and will enable output back to the local filesystem

## Things to improve

- Accessibility, it's been noted that images are hard for screenreaders
- Colours, currently they are completely arbitrary and provide a semi-decent contrast and a quick(er) way of distinguishing between different types of task/process in concourse. These should also be configurable and include a legend.
- Mix of string literals and templates, this was the quickest way to get going but I suspect that having actual proper templates is the 'correct' solution.
- Not all Concourse options/variables are supported, what I have works for 95% of the pipelines I work with. It needs an easier way of adding more.
- Requires another step to produce the images, in and ideal solution this tool would do all of it without an intermediate step.
- Lots of duplication, main.go in particular. I need to investigate if generics can help here (I think they can)
- Tests, probably need some actual tests writing...
