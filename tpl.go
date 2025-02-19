package tpl

import (
	"github.com/Masterminds/sprig/v3"

	"embed"
	"html/template"
	"log/slog"
	"net/http"
)

type Render struct {
	fs   embed.FS
	name string
}

func NewRender(fs embed.FS, name string) *Render {
	return &Render{fs, name}
}

func (r *Render) Render(w http.ResponseWriter, tmpl string, data any) {
	t := template.New(r.name).Funcs(sprig.FuncMap())

	t, err := t.ParseFS(
		r.fs, "templates/layout.gotmpl", "templates/"+tmpl+".gotmpl")
	if err != nil {
		slog.Error("failed to parse/load template", slog.Any("err", err))
		return
	}

	if err := t.ExecuteTemplate(w, "layout", data); err != nil {
		slog.Error("failed to execute template", slog.Any("err", err))
		return
	}
}
