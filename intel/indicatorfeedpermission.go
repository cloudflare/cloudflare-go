// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// IndicatorFeedPermissionService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndicatorFeedPermissionService] method instead.
type IndicatorFeedPermissionService struct {
	Options []option.RequestOption
}

// NewIndicatorFeedPermissionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIndicatorFeedPermissionService(opts ...option.RequestOption) (r *IndicatorFeedPermissionService) {
	r = &IndicatorFeedPermissionService{}
	r.Options = opts
	return
}

// Grant permission to indicator feed
func (r *IndicatorFeedPermissionService) New(ctx context.Context, params IndicatorFeedPermissionNewParams, opts ...option.RequestOption) (res *IndicatorFeedPermissionNewResponse, err error) {
	var env IndicatorFeedPermissionNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/add", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List indicator feed permissions
func (r *IndicatorFeedPermissionService) List(ctx context.Context, query IndicatorFeedPermissionListParams, opts ...option.RequestOption) (res *[]IndicatorFeedPermissionListResponse, err error) {
	var env IndicatorFeedPermissionListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/view", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revoke permission to indicator feed
func (r *IndicatorFeedPermissionService) Delete(ctx context.Context, params IndicatorFeedPermissionDeleteParams, opts ...option.RequestOption) (res *IndicatorFeedPermissionDeleteResponse, err error) {
	var env IndicatorFeedPermissionDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/remove", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IndicatorFeedPermissionNewResponse struct {
	// Whether the update succeeded or not
	Success bool                                   `json:"success"`
	JSON    indicatorFeedPermissionNewResponseJSON `json:"-"`
}

// indicatorFeedPermissionNewResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedPermissionNewResponse]
type indicatorFeedPermissionNewResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionNewResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionListResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The description of the example test
	Description string `json:"description"`
	// Whether the indicator feed can be attributed to a provider
	IsAttributable bool `json:"is_attributable"`
	// Whether the indicator feed can be downloaded
	IsDownloadable bool `json:"is_downloadable"`
	// Whether the indicator feed is exposed to customers
	IsPublic bool `json:"is_public"`
	// The name of the indicator feed
	Name string                                  `json:"name"`
	JSON indicatorFeedPermissionListResponseJSON `json:"-"`
}

// indicatorFeedPermissionListResponseJSON contains the JSON metadata for the
// struct [IndicatorFeedPermissionListResponse]
type indicatorFeedPermissionListResponseJSON struct {
	ID             apijson.Field
	Description    apijson.Field
	IsAttributable apijson.Field
	IsDownloadable apijson.Field
	IsPublic       apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IndicatorFeedPermissionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionListResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionDeleteResponse struct {
	// Whether the update succeeded or not
	Success bool                                      `json:"success"`
	JSON    indicatorFeedPermissionDeleteResponseJSON `json:"-"`
}

// indicatorFeedPermissionDeleteResponseJSON contains the JSON metadata for the
// struct [IndicatorFeedPermissionDeleteResponse]
type indicatorFeedPermissionDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r IndicatorFeedPermissionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndicatorFeedPermissionNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IndicatorFeedPermissionNewResponseEnvelopeSuccess `json:"success,required"`
	Result  IndicatorFeedPermissionNewResponse                `json:"result"`
	JSON    indicatorFeedPermissionNewResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedPermissionNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [IndicatorFeedPermissionNewResponseEnvelope]
type indicatorFeedPermissionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedPermissionNewResponseEnvelopeSuccess bool

const (
	IndicatorFeedPermissionNewResponseEnvelopeSuccessTrue IndicatorFeedPermissionNewResponseEnvelopeSuccess = true
)

func (r IndicatorFeedPermissionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedPermissionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndicatorFeedPermissionListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedPermissionListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IndicatorFeedPermissionListResponseEnvelopeSuccess `json:"success,required"`
	Result  []IndicatorFeedPermissionListResponse              `json:"result"`
	JSON    indicatorFeedPermissionListResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedPermissionListResponseEnvelopeJSON contains the JSON metadata for
// the struct [IndicatorFeedPermissionListResponseEnvelope]
type indicatorFeedPermissionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedPermissionListResponseEnvelopeSuccess bool

const (
	IndicatorFeedPermissionListResponseEnvelopeSuccessTrue IndicatorFeedPermissionListResponseEnvelopeSuccess = true
)

func (r IndicatorFeedPermissionListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedPermissionListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndicatorFeedPermissionDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r IndicatorFeedPermissionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndicatorFeedPermissionDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IndicatorFeedPermissionDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  IndicatorFeedPermissionDeleteResponse                `json:"result"`
	JSON    indicatorFeedPermissionDeleteResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedPermissionDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [IndicatorFeedPermissionDeleteResponseEnvelope]
type indicatorFeedPermissionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedPermissionDeleteResponseEnvelopeSuccess bool

const (
	IndicatorFeedPermissionDeleteResponseEnvelopeSuccessTrue IndicatorFeedPermissionDeleteResponseEnvelopeSuccess = true
)

func (r IndicatorFeedPermissionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedPermissionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
