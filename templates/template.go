package templates

const ProxyTemplate = `// Code generated by proxygen. DO NOT EDIT.
package {{ .PackageName }}

import (
    proxygenInterceptors "github.com/panagiotisptr/proxygen/interceptor"
    proxygenCaster "github.com/panagiotisptr/proxygen/caster"

    {{range $import := .Imports}}
    {{ if $import.Used }} {{ $import.Alias }} "{{ $import.Path }}" {{ end }}
    {{- end }}
)

type {{ .Name }} struct {
	Implementation {{ .ImplementationType }}
	Interceptors   proxygenInterceptors.InterceptorChain
}

var _ {{ .ImplementationType }} = (*{{ .Name }})(nil)

{{- range $method := .Methods }}

func (this *{{ $.Name }}) {{ $method.Name }}(
{{- if gt (len $method.Params) 0 -}}
{{range $idx, $param := $method.Params }}
   arg{{ $idx }} {{ $param }},
{{- end}}
{{end -}}
) {{ if ne (len $method.Rets) 0 }}(
{{- range $ret := $method.Rets }}
   {{ $ret }},
{{- end}}
) {{end}}{
    {{if ne (len $method.Rets) 0 -}}
    rets := this.Interceptors.Apply(
    {{- else -}}
    this.Interceptors.Apply(
    {{- end}}
        []interface{}{
        {{- if gt (len $method.Params) 0 -}}
        {{range $idx, $param := $method.Params }}
           arg{{ $idx }},
        {{- end}}
        {{end -}}
        },
        "{{ $method.Name }}",
        func(args []interface{}) []interface{} {
            {{if ne (len $method.Rets) 0 -}}
            {{range $idx, $ret := $method.Rets -}}
            {{- if ne $idx 0 -}}
            ,
            res{{ $idx }}
            {{- else -}}
            res{{ $idx }}
            {{- end }}
            {{- end}} := this.Implementation.{{ $method.Name }}(
                {{- if gt (len $method.Params) 0 -}}
                {{range $idx, $param := $method.Params }}
                   args[{{ $idx }}].({{ $param }}),
                {{- end}}
                {{end -}}
            )
            {{- else -}}
            this.Implementation.{{ $method.Name }}(
                {{- if gt (len $method.Params) 0 -}}
                {{range $idx, $param := $method.Params }}
                   args[{{ $idx }}].({{ $param }}),
                {{- end}}
                {{end -}}
            )
            {{- end}}
        {{if eq (len $method.Rets) 0}}
            return []interface{}{}
        {{- else}}
            return []interface{}{
            {{- range $idx, $ret := $method.Rets }}
                res{{ $idx }},
            {{- end}}
            }
        {{- end}}
        },
    )

    {{if ne (len $method.Rets) 0 -}}
    return {{range $idx, $ret := $method.Rets -}}
        {{- if ne $idx 0 -}}
        ,
        proxygenCaster.Cast[{{ $ret }}](rets[{{ $idx }}])
        {{- else -}}
        proxygenCaster.Cast[{{ $ret }}](rets[{{ $idx }}])
        {{- end -}} 
    {{end}}
    {{- end}}
}
{{- end}}`
