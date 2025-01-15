// web/static_embed.go
package web

import (
	"embed"
	"io/fs"
	"net/http"
)

// Embed all files in the static directory
//
//go:embed static/**/*
var staticFiles embed.FS

// ServeStaticFiles will return a handler for serving embedded static files
func ServeStaticFiles() http.FileSystem {
	stripped, err := fs.Sub(staticFiles, "static")
	if err != nil {
		panic("failed to create file system from embedded static files: " + err.Error())
	}
	return http.FS(stripped)
}
