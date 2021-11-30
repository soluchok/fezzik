package render

var OperationsTemplate string = `
package {{ .PackageName }}

import (
	"context"

	"github.com/inigolabs/fezzik/client"
)

{{ $root := . }}

{{- range $operation := .Operations }}

var {{ $operation.Name }}Operation string = ~~
{{ $operation.Operation }}
~~

{{ if len $operation.Inputs }}
type {{ $operation.Name }}InputArgs struct {
{{- range $val := $operation.Inputs }}
{{ pascal $val.Name }} {{ $val.Type}} ~~json:"{{ $val.Name }}"~~
{{- end }}
}
{{ end }}

type {{ $operation.Name }}Response struct {
{{ $operation.ResponseType }}
}

{{ if len $operation.Inputs }}
func (c *gqlclient) {{ $operation.Name }}(ctx context.Context, input *{{ $operation.Name }}InputArgs) (
	*{{ $operation.Name }}Response, error) {
{{ else }}
func (c *gqlclient) {{ $operation.Name }}(ctx context.Context) (*{{ $operation.Name }}Response, error) {
{{ end }}

	gqlreq := &client.GQLRequest{
		OperationName: "{{ $operation.Name}}",
		Query: {{ $operation.Name }}Operation,
		Variables: map[string]interface{}{
			{{- range $val := $operation.Inputs }}	
			"{{ $val.Name }}" : input.{{ pascal $val.Name }},
			{{- end }}
		},
	}

	var gqldata {{ $operation.Name }}Response
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

{{- end }}
`
