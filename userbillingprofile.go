// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserBillingProfileService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserBillingProfileService] method
// instead.
type UserBillingProfileService struct {
	Options []option.RequestOption
}

// NewUserBillingProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserBillingProfileService(opts ...option.RequestOption) (r *UserBillingProfileService) {
	r = &UserBillingProfileService{}
	r.Options = opts
	return
}

// Accesses your billing profile object.
func (r *UserBillingProfileService) UserBillingProfileBillingProfileDetails(ctx context.Context, opts ...option.RequestOption) (res *BillingResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/billing/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
