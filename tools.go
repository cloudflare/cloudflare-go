//go:build tools
// +build tools

package tools

//go:generate go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
//go:generate go install github.com/bflad/tfproviderlint/cmd/tfproviderlintx
//go:generate go install github.com/client9/misspell/cmd/misspell
//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go install github.com/hashicorp/go-changelog/cmd/changelog-build
//go:generate go install github.com/google/go-github/github
//go:generate go install golang.org/x/oauth2

import (
	_ "github.com/bflad/tfproviderlint/cmd/tfproviderlintx"
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/go-github/github"
	_ "github.com/hashicorp/go-changelog/cmd/changelog-build"
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	_ "golang.org/x/oauth2"
)
