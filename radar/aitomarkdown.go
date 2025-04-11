// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// AIToMarkdownService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIToMarkdownService] method instead.
type AIToMarkdownService struct {
	Options []option.RequestOption
}

// NewAIToMarkdownService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIToMarkdownService(opts ...option.RequestOption) (r *AIToMarkdownService) {
	r = &AIToMarkdownService{}
	r.Options = opts
	return
}
