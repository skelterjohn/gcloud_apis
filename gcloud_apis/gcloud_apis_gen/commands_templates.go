package main

import (
	"text/template"
)

var commandsSourceTemplate = template.New("commands")

func init() {
	commandsFormats := []string{commandsFormat, apiCommandsFormat}
	for _, tmpl := range commandsFormats {
		commandsSourceTemplate = template.Must(commandsSourceTemplate.Parse(tmpl))
	}
}

var commandsFormat = `{{define "all_commands"}}// generated code: commands.go
package commands

import (
  "net/http"
)

type Context struct {
  client *http.Client
}

type Command func(context Context, args ...string) (err error)

var AllMethods = map[string]Command{
{{range $id, $func := .}}  "{{$id}}": {{$func}},
{{end}}}
{{end}}`

var apiCommandsFormat = `{{define "api_commands"}}// generated code: api commands
package commands

{{range $func, $method := .}}func {{$func}}(context Context, args ...string) (err error) {
  return
}

{{end}}{{end}}`
