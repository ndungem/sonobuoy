/*
Copyright 2018 Heptio Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	_ "embed"
	"strings"
	"text/template"
)

// TemplateFuncs exports (currently singular) functions to be used inside the template
var (
	templateFuncs = map[string]interface{}{
		"indent": func(i int, input string) string {
			split := strings.Split(input, "\n")
			ident := "\n" + strings.Repeat(" ", i)
			// Don't indent the first line, it's already indented in the template
			return strings.Join(split, ident)
		},
	}

	//go:embed gen.tmpl.yaml
	templateDoc string

	// genManifest is the template for the `sonobuoy gen` output
	genManifest = newTemplate("manifest", templateDoc)
)

// newTemplate declares a new template that already has templateFuncs in scope
func newTemplate(name, tmpl string) *template.Template {
	return template.Must(template.New(name).Funcs(templateFuncs).Parse(tmpl))
}
