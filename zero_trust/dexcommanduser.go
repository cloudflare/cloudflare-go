// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// DEXCommandUserService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXCommandUserService] method instead.
type DEXCommandUserService struct {
	Options []option.RequestOption
}

// NewDEXCommandUserService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDEXCommandUserService(opts ...option.RequestOption) (r *DEXCommandUserService) {
	r = &DEXCommandUserService{}
	r.Options = opts
	return
}

// List users emails associated with devices with WARP client support for remote
// captures which have been connected in the last 1 hour.
func (r *DEXCommandUserService) List(ctx context.Context, params DEXCommandUserListParams, opts ...option.RequestOption) (res *DEXCommandUserListResponse, err error) {
	var env DEXCommandUserListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/commands/users", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DEXCommandUserListResponse struct {
	// List of user emails
	UserEmails []string                       `json:"userEmails"`
	JSON       dexCommandUserListResponseJSON `json:"-"`
}

// dexCommandUserListResponseJSON contains the JSON metadata for the struct
// [DEXCommandUserListResponse]
type dexCommandUserListResponseJSON struct {
	UserEmails  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandUserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandUserListResponseJSON) RawJSON() string {
	return r.raw
}

type DEXCommandUserListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// filter user emails by search
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [DEXCommandUserListParams]'s query parameters as
// `url.Values`.
func (r DEXCommandUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DEXCommandUserListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    DEXCommandUserListResponseEnvelopeSuccess    `json:"success,required"`
	Result     DEXCommandUserListResponse                   `json:"result"`
	ResultInfo DEXCommandUserListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dexCommandUserListResponseEnvelopeJSON       `json:"-"`
}

// dexCommandUserListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXCommandUserListResponseEnvelope]
type dexCommandUserListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandUserListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandUserListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXCommandUserListResponseEnvelopeSuccess bool

const (
	DEXCommandUserListResponseEnvelopeSuccessTrue DEXCommandUserListResponseEnvelopeSuccess = true
)

func (r DEXCommandUserListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXCommandUserListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DEXCommandUserListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       dexCommandUserListResponseEnvelopeResultInfoJSON `json:"-"`
}

// dexCommandUserListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DEXCommandUserListResponseEnvelopeResultInfo]
type dexCommandUserListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandUserListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandUserListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
