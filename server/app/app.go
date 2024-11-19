package app

import (
	"embed"
)

//dist/* dist/*/*
//go:embed i18n.json src/language/* src/language/*/*
var DistFS embed.FS
