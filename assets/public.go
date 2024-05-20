package assets

import (
	"embed"
)

//go:embed all:public/*
var PublicFiles embed.FS
