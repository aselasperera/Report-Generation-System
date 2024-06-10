package funcs

import (
	"text/template"
	"time"
)

func formatDate(t time.Time) string {
	return t.Format("02-Jan-2006")
}

var FuncMap = template.FuncMap{
	"formatDate": formatDate,
}
