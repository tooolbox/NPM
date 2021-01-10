package tools

// +build tools

//go:generate go build -o ../bin/oapi-codegen github.com/deepmap/oapi-codegen/cmd/oapi-codegen

import (
	_ "github.com/deepmap/oapi-codegen"
)
