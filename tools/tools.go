//go:build tools
// +build tools

package tools

//go:generate go install golang.org/x/tools/gopls@latest
//go:generate go install github.com/orijtech/structslop/cmd/structslop@latest
//go:generate go install github.com/ramya-rao-a/go-outline@latest
//go:generate go install github.com/cweill/gotests/gotests@latest
//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
//go:generate go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
//go:generate go install github.com/go-delve/delve/cmd/dlv@latest
//go:generate go install golang.org/x/lint/golint@latest

import (
	_ "github.com/cweill/gotests/gotests"
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/orijtech/structslop/cmd/structslop"
	_ "github.com/ramya-rao-a/go-outline"
	_ "github.com/uudashr/gopkgs/v2/cmd/gopkgs"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/gopls"
)
