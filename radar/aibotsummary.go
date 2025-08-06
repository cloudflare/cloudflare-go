// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AIBotSummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIBotSummaryService] method instead.
type AIBotSummaryService struct {
	Options []option.RequestOption
}

// NewAIBotSummaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIBotSummaryService(opts ...option.RequestOption) (r *AIBotSummaryService) {
	r = &AIBotSummaryService{}
	r.Options = opts
	return
}
