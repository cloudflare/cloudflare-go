// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// TokenPermissionGroupService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTokenPermissionGroupService]
// method instead.
type TokenPermissionGroupService struct {
	Options []option.RequestOption
}

// NewTokenPermissionGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTokenPermissionGroupService(opts ...option.RequestOption) (r *TokenPermissionGroupService) {
	r = &TokenPermissionGroupService{}
	r.Options = opts
	return
}

// Find all available permission groups.
func (r *TokenPermissionGroupService) List(ctx context.Context, opts ...option.RequestOption) (res *[]TokenPermissionGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TokenPermissionGroupListResponseEnvelope
	path := "user/tokens/permission_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TokenPermissionGroupListResponse = interface{}

type TokenPermissionGroupListResponseEnvelope struct {
	Errors   []TokenPermissionGroupListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenPermissionGroupListResponseEnvelopeMessages `json:"messages,required"`
	Result   []TokenPermissionGroupListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    TokenPermissionGroupListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo TokenPermissionGroupListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       tokenPermissionGroupListResponseEnvelopeJSON       `json:"-"`
}

// tokenPermissionGroupListResponseEnvelopeJSON contains the JSON metadata for the
// struct [TokenPermissionGroupListResponseEnvelope]
type tokenPermissionGroupListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPermissionGroupListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenPermissionGroupListResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    tokenPermissionGroupListResponseEnvelopeErrorsJSON `json:"-"`
}

// tokenPermissionGroupListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [TokenPermissionGroupListResponseEnvelopeErrors]
type tokenPermissionGroupListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPermissionGroupListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenPermissionGroupListResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    tokenPermissionGroupListResponseEnvelopeMessagesJSON `json:"-"`
}

// tokenPermissionGroupListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [TokenPermissionGroupListResponseEnvelopeMessages]
type tokenPermissionGroupListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPermissionGroupListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TokenPermissionGroupListResponseEnvelopeSuccess bool

const (
	TokenPermissionGroupListResponseEnvelopeSuccessTrue TokenPermissionGroupListResponseEnvelopeSuccess = true
)

type TokenPermissionGroupListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       tokenPermissionGroupListResponseEnvelopeResultInfoJSON `json:"-"`
}

// tokenPermissionGroupListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [TokenPermissionGroupListResponseEnvelopeResultInfo]
type tokenPermissionGroupListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPermissionGroupListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
