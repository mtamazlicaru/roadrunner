package http

import (
	"strings"
	"path"
)

// FsConfig describes file location and controls access to them.
type FsConfig struct {
	// Dir contains name of directory to control access to.
	Dir string

	// Forbid specifies list of file extensions which are forbidden for access.
	// Example: .php, .exe, .bat, .htaccess and etc.
	Forbid []string
}

// Forbid must return true if file extension is not allowed for the upload.
func (cfg FsConfig) Forbids(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))

	for _, v := range cfg.Forbid {
		if ext == v {
			return true
		}
	}

	return false
}
