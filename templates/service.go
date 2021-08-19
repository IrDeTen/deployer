package templates

var ServiceTmpl []string = []string{
	"Unit file: /etc/systemd/system/%s.service",
	"Description: %s",
	"Exec path: %s",
	"Working Directory: %s",
	"Type: %s",
	"After: %s",
	"Requires: %s",
	"Restart: %s",
	"User: %s",
	"Group: %s",
}
