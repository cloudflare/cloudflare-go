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

// AccountIntelIndicatorFeedPermissionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountIntelIndicatorFeedPermissionService] method instead.
type AccountIntelIndicatorFeedPermissionService struct {
	Options []option.RequestOption
}

// NewAccountIntelIndicatorFeedPermissionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountIntelIndicatorFeedPermissionService(opts ...option.RequestOption) (r *AccountIntelIndicatorFeedPermissionService) {
	r = &AccountIntelIndicatorFeedPermissionService{}
	r.Options = opts
	return
}

// Revoke permission to indicator feed
func (r *AccountIntelIndicatorFeedPermissionService) Remove(ctx context.Context, accountIdentifier string, body AccountIntelIndicatorFeedPermissionRemoveParams, opts ...option.RequestOption) (res *AccountIntelIndicatorFeedPermissionRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/remove", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List indicator feed permissions
func (r *AccountIntelIndicatorFeedPermissionService) View(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountIntelIndicatorFeedPermissionViewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/view", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountIntelIndicatorFeedPermissionRemoveResponse struct {
	Errors   []AccountIntelIndicatorFeedPermissionRemoveResponseError   `json:"errors"`
	Messages []AccountIntelIndicatorFeedPermissionRemoveResponseMessage `json:"messages"`
	Result   AccountIntelIndicatorFeedPermissionRemoveResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountIntelIndicatorFeedPermissionRemoveResponseSuccess `json:"success"`
	JSON    accountIntelIndicatorFeedPermissionRemoveResponseJSON    `json:"-"`
}

// accountIntelIndicatorFeedPermissionRemoveResponseJSON contains the JSON metadata
// for the struct [AccountIntelIndicatorFeedPermissionRemoveResponse]
type accountIntelIndicatorFeedPermissionRemoveResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedPermissionRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIndicatorFeedPermissionRemoveResponseError struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountIntelIndicatorFeedPermissionRemoveResponseErrorJSON `json:"-"`
}

// accountIntelIndicatorFeedPermissionRemoveResponseErrorJSON contains the JSON
// metadata for the struct [AccountIntelIndicatorFeedPermissionRemoveResponseError]
type accountIntelIndicatorFeedPermissionRemoveResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedPermissionRemoveResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIndicatorFeedPermissionRemoveResponseMessage struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accountIntelIndicatorFeedPermissionRemoveResponseMessageJSON `json:"-"`
}

// accountIntelIndicatorFeedPermissionRemoveResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountIntelIndicatorFeedPermissionRemoveResponseMessage]
type accountIntelIndicatorFeedPermissionRemoveResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedPermissionRemoveResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIndicatorFeedPermissionRemoveResponseResult struct {
	// Whether the update succeeded or not
	Success bool                                                        `json:"success"`
	JSON    accountIntelIndicatorFeedPermissionRemoveResponseResultJSON `json:"-"`
}

// accountIntelIndicatorFeedPermissionRemoveResponseResultJSON contains the JSON
// metadata for the struct
// [AccountIntelIndicatorFeedPermissionRemoveResponseResult]
type accountIntelIndicatorFeedPermissionRemoveResponseResultJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedPermissionRemoveResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountIntelIndicatorFeedPermissionRemoveResponseSuccess bool

const (
	AccountIntelIndicatorFeedPermissionRemoveResponseSuccessTrue AccountIntelIndicatorFeedPermissionRemoveResponseSuccess = true
)

type AccountIntelIndicatorFeedPermissionViewResponse struct {
	Errors   []AccountIntelIndicatorFeedPermissionViewResponseError   `json:"errors"`
	Messages []AccountIntelIndicatorFeedPermissionViewResponseMessage `json:"messages"`
	Result   []AccountIntelIndicatorFeedPermissionViewResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountIntelIndicatorFeedPermissionViewResponseSuccess `json:"success"`
	JSON    accountIntelIndicatorFeedPermissionViewResponseJSON    `json:"-"`
}

// accountIntelIndicatorFeedPermissionViewResponseJSON contains the JSON metadata
// for the struct [AccountIntelIndicatorFeedPermissionViewResponse]
type accountIntelIndicatorFeedPermissionViewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedPermissionViewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIndicatorFeedPermissionViewResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountIntelIndicatorFeedPermissionViewResponseErrorJSON `json:"-"`
}

// accountIntelIndicatorFeedPermissionViewResponseErrorJSON contains the JSON
// metadata for the struct [AccountIntelIndicatorFeedPermissionViewResponseError]
type accountIntelIndicatorFeedPermissionViewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedPermissionViewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIndicatorFeedPermissionViewResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountIntelIndicatorFeedPermissionViewResponseMessageJSON `json:"-"`
}

// accountIntelIndicatorFeedPermissionViewResponseMessageJSON contains the JSON
// metadata for the struct [AccountIntelIndicatorFeedPermissionViewResponseMessage]
type accountIntelIndicatorFeedPermissionViewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedPermissionViewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIndicatorFeedPermissionViewResponseResult struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The description of the example test
	Description string `json:"description"`
	// The name of the indicator feed
	Name string                                                    `json:"name"`
	JSON accountIntelIndicatorFeedPermissionViewResponseResultJSON `json:"-"`
}

// accountIntelIndicatorFeedPermissionViewResponseResultJSON contains the JSON
// metadata for the struct [AccountIntelIndicatorFeedPermissionViewResponseResult]
type accountIntelIndicatorFeedPermissionViewResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedPermissionViewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountIntelIndicatorFeedPermissionViewResponseSuccess bool

const (
	AccountIntelIndicatorFeedPermissionViewResponseSuccessTrue AccountIntelIndicatorFeedPermissionViewResponseSuccess = true
)

type AccountIntelIndicatorFeedPermissionRemoveParams struct {
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r AccountIntelIndicatorFeedPermissionRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
