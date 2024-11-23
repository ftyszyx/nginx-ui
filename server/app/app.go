package app

import (
	"embed"
)

//dist/* dist/*/* src/language/* src/language/*/*
//go:embed i18n.json
var DistFS embed.FS
