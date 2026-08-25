package migrations

import "embed"

// Files contains SQL migrations applied by cmd/api at startup.
//
//go:embed *.sql
var Files embed.FS
