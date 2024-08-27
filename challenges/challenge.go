// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package challenges

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ChallengeService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChallengeService] method instead.
type ChallengeService struct {
	Options []option.RequestOption
	Widgets *WidgetService
}

// NewChallengeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChallengeService(opts ...option.RequestOption) (r *ChallengeService) {
	r = &ChallengeService{}
	r.Options = opts
	r.Widgets = NewWidgetService(opts...)
	return
}
