package global

{{- if .HasGlobal }}

import "github.com/switfs/switfs-oms/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}