// Package ui provides embedded static files (binary)
package ui

import "embed"

//go:embed "html" "static"
var Files embed.FS
