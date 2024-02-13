// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarVerifiedBotTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarVerifiedBotTopService]
// method instead.
type RadarVerifiedBotTopService struct {
	Options    []option.RequestOption
	Bots       *RadarVerifiedBotTopBotService
	Categories *RadarVerifiedBotTopCategoryService
}

// NewRadarVerifiedBotTopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarVerifiedBotTopService(opts ...option.RequestOption) (r *RadarVerifiedBotTopService) {
	r = &RadarVerifiedBotTopService{}
	r.Options = opts
	r.Bots = NewRadarVerifiedBotTopBotService(opts...)
	r.Categories = NewRadarVerifiedBotTopCategoryService(opts...)
	return
}
