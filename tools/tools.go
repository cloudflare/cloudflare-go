//go:build tools
// +build tools

package tools

//go:generate go install github.com/orijtech/structslop/cmd/structslop

import (
	_ "github.com/orijtech/structslop/cmd/structslop"
)
