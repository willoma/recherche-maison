// Package static provides handlers for static files
package static

import (
	_ "embed"
	"log/slog"
	"net/http"
)

//go:embed script.js
var script []byte

//go:embed style.css
var style []byte

// ServeScript serves the JavaScript file
func ServeScript(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	if _, err := w.Write(script); err != nil {
		slog.Error("Failed to write script.js", "error", err)
	}
}

// ServeStyle serves the CSS file
func ServeStyle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	if _, err := w.Write(style); err != nil {
		slog.Error("Failed to write style.css", "error", err)
	}
}
