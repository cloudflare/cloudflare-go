// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *IntelIndicatorFeedPermissionService) New(ctx context.Context, params IntelIndicatorFeedPermissionNewParams, opts ...option.RequestOption) (res *IntelIndicatorFeedPermissionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedPermissionNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/add", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List indicator feed permissions
func (r *IntelIndicatorFeedPermissionService) List(ctx context.Context, query IntelIndicatorFeedPermissionListParams, opts ...option.RequestOption) (res *[]IntelIndicatorFeedPermissionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedPermissionListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/view", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revoke permission to indicator feed
func (r *IntelIndicatorFeedPermissionService) Delete(ctx context.Context, params IntelIndicatorFeedPermissionDeleteParams, opts ...option.RequestOption) (res *IntelIndicatorFeedPermissionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIndicatorFeedPermissionDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/remove", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelIndicatorFeedPermissionNewResponse struct {
	// Whether the update succeeded or not
	Success bool                                        `json:"success"`
	JSON    intelIndicatorFeedPermissionNewResponseJSON `json:"-"`
}

// intelIndicatorFeedPermissionNewResponseJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedPermissionNewResponse]
type intelIndicatorFeedPermissionNewResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionNewResponseJSON) RawJSON() string {
	return r.raw
}

type IntelIndicatorFeedPermissionListResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The description of the example test
	Description string `json:"description"`
	// The name of the indicator feed
	Name string                                       `json:"name"`
	JSON intelIndicatorFeedPermissionListResponseJSON `json:"-"`
}

// intelIndicatorFeedPermissionListResponseJSON contains the JSON metadata for the
// struct [IntelIndicatorFeedPermissionListResponse]
type intelIndicatorFeedPermissionListResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionListResponseJSON) RawJSON() string {
	return r.raw
}

type IntelIndicatorFeedPermissionDeleteResponse struct {
	// Whether the update succeeded or not
	Success bool                                           `json:"success"`
	JSON    intelIndicatorFeedPermissionDeleteResponseJSON `json:"-"`
}

// intelIndicatorFeedPermissionDeleteResponseJSON contains the JSON metadata for
// the struct [IntelIndicatorFeedPermissionDeleteResponse]
type intelIndicatorFeedPermissionDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type IntelIndicatorFeedPermissionNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r IntelIndicatorFeedPermissionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelIndicatorFeedPermissionNewResponseEnvelope struct {
	Errors   []IntelIndicatorFeedPermissionNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIndicatorFeedPermissionNewResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelIndicatorFeedPermissionNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelIndicatorFeedPermissionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelIndicatorFeedPermissionNewResponseEnvelopeJSON    `json:"-"`
}

// intelIndicatorFeedPermissionNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionNewResponseEnvelope]
type intelIndicatorFeedPermissionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IntelIndicatorFeedPermissionNewResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    intelIndicatorFeedPermissionNewResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIndicatorFeedPermissionNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [IntelIndicatorFeedPermissionNewResponseEnvelopeErrors]
type intelIndicatorFeedPermissionNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IntelIndicatorFeedPermissionNewResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    intelIndicatorFeedPermissionNewResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIndicatorFeedPermissionNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [IntelIndicatorFeedPermissionNewResponseEnvelopeMessages]
type intelIndicatorFeedPermissionNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IntelIndicatorFeedPermissionNewResponseEnvelopeSuccess bool

const (
	IntelIndicatorFeedPermissionNewResponseEnvelopeSuccessTrue IntelIndicatorFeedPermissionNewResponseEnvelopeSuccess = true
)

type IntelIndicatorFeedPermissionListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IntelIndicatorFeedPermissionListResponseEnvelope struct {
	Errors   []IntelIndicatorFeedPermissionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIndicatorFeedPermissionListResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelIndicatorFeedPermissionListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success IntelIndicatorFeedPermissionListResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelIndicatorFeedPermissionListResponseEnvelopeJSON    `json:"-"`
}

// intelIndicatorFeedPermissionListResponseEnvelopeJSON contains the JSON metadata
// for the struct [IntelIndicatorFeedPermissionListResponseEnvelope]
type intelIndicatorFeedPermissionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IntelIndicatorFeedPermissionListResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    intelIndicatorFeedPermissionListResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIndicatorFeedPermissionListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [IntelIndicatorFeedPermissionListResponseEnvelopeErrors]
type intelIndicatorFeedPermissionListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IntelIndicatorFeedPermissionListResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    intelIndicatorFeedPermissionListResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIndicatorFeedPermissionListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [IntelIndicatorFeedPermissionListResponseEnvelopeMessages]
type intelIndicatorFeedPermissionListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IntelIndicatorFeedPermissionListResponseEnvelopeSuccess bool

const (
	IntelIndicatorFeedPermissionListResponseEnvelopeSuccessTrue IntelIndicatorFeedPermissionListResponseEnvelopeSuccess = true
)

type IntelIndicatorFeedPermissionDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r IntelIndicatorFeedPermissionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelIndicatorFeedPermissionDeleteResponseEnvelope struct {
	Errors   []IntelIndicatorFeedPermissionDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIndicatorFeedPermissionDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelIndicatorFeedPermissionDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelIndicatorFeedPermissionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelIndicatorFeedPermissionDeleteResponseEnvelopeJSON    `json:"-"`
}

// intelIndicatorFeedPermissionDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [IntelIndicatorFeedPermissionDeleteResponseEnvelope]
type intelIndicatorFeedPermissionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IntelIndicatorFeedPermissionDeleteResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    intelIndicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIndicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [IntelIndicatorFeedPermissionDeleteResponseEnvelopeErrors]
type intelIndicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IntelIndicatorFeedPermissionDeleteResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    intelIndicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIndicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [IntelIndicatorFeedPermissionDeleteResponseEnvelopeMessages]
type intelIndicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIndicatorFeedPermissionDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelIndicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IntelIndicatorFeedPermissionDeleteResponseEnvelopeSuccess bool

const (
	IntelIndicatorFeedPermissionDeleteResponseEnvelopeSuccessTrue IntelIndicatorFeedPermissionDeleteResponseEnvelopeSuccess = true
)
