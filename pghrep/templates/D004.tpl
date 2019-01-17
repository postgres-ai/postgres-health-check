# {{ .checkId }} pg_stat_statements and kcache settings #

## Observations ##
{{ if .hosts.master }}
### Master (`{{.hosts.master}}`) ###
{{ if (index (index (index .results .hosts.master) "data") "pg_stat_statements") }}
#### `pg_stat_statements` extension settings ####
Setting | Value | Unit | Type | Min value | Max value
--------|-------|------|------|-----------|-----------
{{ range $i, $setting_name := (index (index (index (index .results .hosts.master) "data") "pg_stat_statements") "_keys") }}
{{- $setting_data := (index (index (index (index $.results $.hosts.master) "data") "pg_stat_statements") $setting_name) -}}
[{{ $setting_name }}](https://postgresqlco.nf/en/doc/param/{{ $setting_name }})|{{ $setting_data.setting }}|{{ if $setting_data.unit }}{{ $setting_data.unit }} {{ end }}|{{ $setting_data.vartype }}|{{ if $setting_data.min_val }}{{ $setting_data.min_val }} {{ end }}|{{ if $setting_data.max_val }}{{ $setting_data.max_val }} {{ end }}
{{ end }}
{{- end -}}

{{ if (index (index (index .results .hosts.master) "data") "kcache") }}
#### `kcache` extension settings ####
Setting | Value | Unit | Type | Min value | Max value
--------|-------|------|------|-----------|-----------
{{ range $i, $setting_name := (index (index (index (index .results .hosts.master) "data") "kcache") "_keys") }}
{{- $setting_data := (index (index (index (index $.results $.hosts.master) "data") "kcache") $setting_name) -}}
[{{ $setting_name }}](https://postgresqlco.nf/en/doc/param/{{ $setting_name }})|{{ $setting_data.setting }}|{{ if $setting_data.unit }}{{ $setting_data.unit }} {{ end }}|{{ $setting_data.vartype }}|{{ if $setting_data.min_val }}{{ $setting_data.min_val }} {{ end }}|{{ if $setting_data.max_val }}{{ $setting_data.max_val }} {{ end }}
{{ end }}
{{- end -}}
{{ else }}
No data
{{ end }}

## Conclusions ##


## Recommendations ##
