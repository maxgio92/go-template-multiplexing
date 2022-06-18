package template

import (
	t "html/template"

	"github.com/maxgio92/go-template-multiplexing/pkg/matrix"
)

type TemplatePart struct {
	matrix.Column
	TemplateString  string
	MatchedVariable string
	Template        *t.Template
}
