{{range .configList}}
<source>
  @type tail
  tag docker.{{ $.containerId }}.{{ .Name }}
  path {{ .HostDir }}/{{ .File }}
  path_key tailed_path
  <parse>
  {{if .Stdout}}
  @type json
  {{else}}
      @type nonex
      time_format = %Y-%m-%dT%H:%M:%S.%NZ
  {{end}}
  {{ $time_key := "" }}
  {{if .FormatConfig}}
  {{range $key, $value := .FormatConfig}}
  {{ $key }} {{ $value }}
  {{end}}
  {{end}}
  {{ if .EstimateTime }}
  estimate_current_event true
  {{end}}
  keep_time_key true
  </parse>
  read_from_head true
  pos_file /var/log/fluentd_pos/{{ $.containerId }}.{{ .Name }}.pos
</source>

<filter docker.{{ $.containerId }}.{{ .Name }}>
  @type record_transformer
  enable_ruby true
  <record>
    log_file "${record['tailed_path']}"
    host "#{Socket.gethostname}"
    {{range $key, $value := .Tags}}
    {{ $key }} {{ $value }}
    {{end}}

    {{if eq $.output "elasticsearch"}}
    _target {{if .Target}}{{.Target}}-${time.strftime('%Y.%m.%d')}{{else}}{{ .Name }}-${time.strftime('%Y.%m.%d')}{{end}}
    {{else}}
    _target {{if .Target}}{{.Target}}{{else}}{{ .Name }}{{end}}
    {{end}}

    {{range $key, $value := $.container}}
    {{ $key }} {{ $value }}
    {{end}}
  </record>
</filter>
<filter docker.{{ $.containerId }}.{{ .Name }}>
    @type parser
    format /^(?<Timestamp>\S+ \S+) (?<Pid>\d+) (?<log_level>\S+) (?<python_module>\S+) (\[(?<global_id>\S+) (req-(?<request_id>\S+) (?<user_id>\S+) (?<tenant_id>\S+) (?<domain_id>\S+) (?<user_domain>\S+) (?<project_domain>\S+)|-)\])? (?<Payload>.*)?$/
    time_format dd%d/%b/%Y:%H:%M:%S %z
    key_name log
    reserve_data true
    emit_invalid_record_to_error false
  </filter>
{{end}}
