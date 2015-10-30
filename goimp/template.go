package goimp

import (
	"html/template"
)

var tmpl = template.Must(template.New("index").Parse(`<doctype html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	<meta name="go-import" content="{{.ImportRoot}} {{.VCS}} {{.VCSRoot}}">
	<meta http-equiv="refresh" content="0;
		url=https://gowalker.org/{{.ImportRoot}}{{.Suffix}}">
</head>
<body>
One day, code will readable as literature.
</body>
</html>
`))

type data struct {
	ImportRoot string
	VCS        string
	VCSRoot    string
	Suffix     string
}
