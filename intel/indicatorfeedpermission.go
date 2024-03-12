// File generated from our OpenAPI spec by Stainless.

package intel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// IndicatorFeedPermissionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewIndicatorFeedPermissionService] method instead.
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
func (r *IndicatorFeedPermissionService) New(ctx context.Context, params IndicatorFeedPermissionNewParams, opts ...option.RequestOption) (res *IntelPermissionsUpdate, err error) {
	opts = append(r.Options[:], opts...)
	var env IndicatorFeedPermissionNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/add", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List indicator feed permissions
func (r *IndicatorFeedPermissionService) List(ctx context.Context, query IndicatorFeedPermissionListParams, opts ...option.RequestOption) (res *[]IntelPermissionListItem, err error) {
	opts = append(r.Options[:], opts...)
	var env IndicatorFeedPermissionListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/view", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revoke permission to indicator feed
func (r *IndicatorFeedPermissionService) Delete(ctx context.Context, params IndicatorFeedPermissionDeleteParams, opts ...option.RequestOption) (res *IntelPermissionsUpdate, err error) {
	opts = append(r.Options[:], opts...)
	var env IndicatorFeedPermissionDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/remove", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelPermissionListItem struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The description of the example test
	Description string `json:"description"`
	// The name of the indicator feed
	Name string                      `json:"name"`
	JSON intelPermissionListItemJSON `json:"-"`
}

// intelPermissionListItemJSON contains the JSON metadata for the struct
// [IntelPermissionListItem]
type intelPermissionListItemJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelPermissionListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPermissionListItemJSON) RawJSON() string {
	return r.raw
}

type IntelPermissionsUpdate struct {
	// Whether the update succeeded or not
	Success bool                       `json:"success"`
	JSON    intelPermissionsUpdateJSON `json:"-"`
}

// intelPermissionsUpdateJSON contains the JSON metadata for the struct
// [IntelPermissionsUpdate]
type intelPermissionsUpdateJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelPermissionsUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPermissionsUpdateJSON) RawJSON() string {
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
	Errors   []IndicatorFeedPermissionNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedPermissionNewResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelPermissionsUpdate                               `json:"result,required"`
	// Whether the API call was successful
	Success IndicatorFeedPermissionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    indicatorFeedPermissionNewResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedPermissionNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [IndicatorFeedPermissionNewResponseEnvelope]
type indicatorFeedPermissionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    indicatorFeedPermissionNewResponseEnvelopeErrorsJSON `json:"-"`
}

// indicatorFeedPermissionNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [IndicatorFeedPermissionNewResponseEnvelopeErrors]
type indicatorFeedPermissionNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    indicatorFeedPermissionNewResponseEnvelopeMessagesJSON `json:"-"`
}

// indicatorFeedPermissionNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [IndicatorFeedPermissionNewResponseEnvelopeMessages]
type indicatorFeedPermissionNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedPermissionNewResponseEnvelopeSuccess bool

const (
	IndicatorFeedPermissionNewResponseEnvelopeSuccessTrue IndicatorFeedPermissionNewResponseEnvelopeSuccess = true
)

type IndicatorFeedPermissionListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedPermissionListResponseEnvelope struct {
	Errors   []IndicatorFeedPermissionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedPermissionListResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelPermissionListItem                             `json:"result,required"`
	// Whether the API call was successful
	Success IndicatorFeedPermissionListResponseEnvelopeSuccess `json:"success,required"`
	JSON    indicatorFeedPermissionListResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedPermissionListResponseEnvelopeJSON contains the JSON metadata for
// the struct [IndicatorFeedPermissionListResponseEnvelope]
type indicatorFeedPermissionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    indicatorFeedPermissionListResponseEnvelopeErrorsJSON `json:"-"`
}

// indicatorFeedPermissionListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [IndicatorFeedPermissionListResponseEnvelopeErrors]
type indicatorFeedPermissionListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    indicatorFeedPermissionListResponseEnvelopeMessagesJSON `json:"-"`
}

// indicatorFeedPermissionListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [IndicatorFeedPermissionListResponseEnvelopeMessages]
type indicatorFeedPermissionListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedPermissionListResponseEnvelopeSuccess bool

const (
	IndicatorFeedPermissionListResponseEnvelopeSuccessTrue IndicatorFeedPermissionListResponseEnvelopeSuccess = true
)

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
	Errors   []IndicatorFeedPermissionDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedPermissionDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelPermissionsUpdate                                  `json:"result,required"`
	// Whether the API call was successful
	Success IndicatorFeedPermissionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    indicatorFeedPermissionDeleteResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedPermissionDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [IndicatorFeedPermissionDeleteResponseEnvelope]
type indicatorFeedPermissionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionDeleteResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    indicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// indicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [IndicatorFeedPermissionDeleteResponseEnvelopeErrors]
type indicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionDeleteResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    indicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// indicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [IndicatorFeedPermissionDeleteResponseEnvelopeMessages]
type indicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedPermissionDeleteResponseEnvelopeSuccess bool

const (
	IndicatorFeedPermissionDeleteResponseEnvelopeSuccessTrue IndicatorFeedPermissionDeleteResponseEnvelopeSuccess = true
)
