package examples

import "embed"

//go:embed en
//go:embed es
//go:embed ignore.txt

// Content contains the example locale files
var Content embed.FS
