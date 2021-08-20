package templates

const (
	Unit = `
[Unit]
Description={{.Description}}
{{block "After" .After}}{{range .}}{{println "After=" .}}{{"\n"}}{{end}}{{end}}
{{block "Requires" .Requires}}{{range .}}{{println "Requires=" .}}{{"\n"}}{{end}}{{end}}

[Service]
ExecStart={{.Path}}
Type={{.Type}}	
WorkingDirectory={{.WorkingDirectory}}
User={{.User}}
Group={{.Group}}

[Install]
WantedBy=multi-user.target`
)
