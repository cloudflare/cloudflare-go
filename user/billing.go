// File generated from our OpenAPI spec by Stainless.

package user

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// BillingService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewBillingService] method instead.
type BillingService struct {
	Options []option.RequestOption
	History *BillingHistoryService
	Profile *BillingProfileService
}

// NewBillingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBillingService(opts ...option.RequestOption) (r *BillingService) {
	r = &BillingService{}
	r.Options = opts
	r.History = NewBillingHistoryService(opts...)
	r.Profile = NewBillingProfileService(opts...)
	return
}
