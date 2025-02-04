//go:build tools
// +build tools

package main

import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)

//go:generate oapi-codegen -package=model -generate=types -o=model/types.go openapi.yaml
