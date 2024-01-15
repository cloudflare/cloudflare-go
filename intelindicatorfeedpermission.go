// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// IntelIndicatorFeedPermissionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewIntelIndicatorFeedPermissionService] method instead.
type IntelIndicatorFeedPermissionService struct {
	Options []option.RequestOption
}

// NewIntelIndicatorFeedPermissionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewIntelIndicatorFeedPermissionService(opts ...option.RequestOption) (r *IntelIndicatorFeedPermissionService) {
	r = &IntelIndicatorFeedPermissionService{}
	r.Options = opts
	return
}

// Grant permission to indicator feed
func (r *IntelIndicatorFeedPermissionService) Add(ctx context.Context, accountIdentifier string, body IntelIndicatorFeedPermissionAddParams, opts ...option.RequestOption) (res *IntelIndicatorFeedPermissionAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/add", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type IntelIndicatorFeedPermissionAddResponse struct {
	Errors   []IntelIndicatorFeedPermissionAddResponseError   `json:"errors"`
	Messages []IntelIndicatorFeedPermissionAddResponseMessage `json:"messages"`
	Result   IntelIndicatorFeedPermissionAddResponseResult    `json:"result"`
	// Whether the API call was successful
	Success IntelIndicatorFeedPermissionAddResponseSuccess `json:"success"`
	JSON    intelIndicatorFeedPermissionAddResponseJSON    `json:"-"`
}

// intelIndicatorFeedPermissionAddResponseJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedPermissionAddResponse]
type intelIndicatorFeedPermissionAddResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionAddResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    intelIndicatorFeedPermissionAddResponseErrorJSON `json:"-"`
}

// intelIndicatorFeedPermissionAddResponseErrorJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedPermissionAddResponseError]
type intelIndicatorFeedPermissionAddResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionAddResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionAddResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    intelIndicatorFeedPermissionAddResponseMessageJSON `json:"-"`
}

// intelIndicatorFeedPermissionAddResponseMessageJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionAddResponseMessage]
type intelIndicatorFeedPermissionAddResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionAddResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIndicatorFeedPermissionAddResponseResult struct {
	// Whether the update succeeded or not
	Success bool                                              `json:"success"`
	JSON    intelIndicatorFeedPermissionAddResponseResultJSON `json:"-"`
}

// intelIndicatorFeedPermissionAddResponseResultJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedPermissionAddResponseResult]
type intelIndicatorFeedPermissionAddResponseResultJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionAddResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIndicatorFeedPermissionAddResponseSuccess bool

const (
	IntelIndicatorFeedPermissionAddResponseSuccessTrue IntelIndicatorFeedPermissionAddResponseSuccess = true
)

type IntelIndicatorFeedPermissionAddParams struct {
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r IntelIndicatorFeedPermissionAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
