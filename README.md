# tpl

A small wrapper around `html/template`. This is not optimized for performance, but optimized for using this library to write somewhat decent code.

## Example

Put your files into a `templates/` directory.

The only required file is a `layout.gotmpl` with appropriate blocks for the `layout` and `content`.

```go
{{ define "layout" }}
<!DOCTYPE html>
<html lang="en">
<body>
  <main class="container">
    {{ template "content" . }}
  </main>
</body>
</html>
{{ end }}
```

Initialize:

```go
var (
	//go:embed templates/*
	tmplFS embed.FS

)
renderer := tpl.NewRender(tmplFS, "name")
```

... and pass it into your router/handlers.

Define a struct like this and assign data to it so your templates have something to show.

```go
type tplData struct {
	Data any
} 
```

Output a template from your handler:

```go
// filename.gotmpl
tpl.Render(w, "filename", tplData{
	Data: allData,
})
```