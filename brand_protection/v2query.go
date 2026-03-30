// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// V2QueryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2QueryService] method instead.
type V2QueryService struct {
	Options []option.RequestOption
}

// NewV2QueryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2QueryService(opts ...option.RequestOption) (r *V2QueryService) {
	r = &V2QueryService{}
	r.Options = opts
	return
}

// Get all saved brand protection queries for an account
func (r *V2QueryService) Get(ctx context.Context, params V2QueryGetParams, opts ...option.RequestOption) (res *[]V2QueryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/brand-protection/domain/queries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type V2QueryGetResponse struct {
	Created    string                       `json:"created" api:"required"`
	Parameters V2QueryGetResponseParameters `json:"parameters" api:"required,nullable"`
	QueryID    int64                        `json:"query_id" api:"required"`
	QueryTag   string                       `json:"query_tag" api:"required"`
	Scan       bool                         `json:"scan" api:"required"`
	Updated    string                       `json:"updated" api:"required"`
	JSON       v2QueryGetResponseJSON       `json:"-"`
}

// v2QueryGetResponseJSON contains the JSON metadata for the struct
// [V2QueryGetResponse]
type v2QueryGetResponseJSON struct {
	Created     apijson.Field
	Parameters  apijson.Field
	QueryID     apijson.Field
	QueryTag    apijson.Field
	Scan        apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2QueryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseJSON) RawJSON() string {
	return r.raw
}

type V2QueryGetResponseParameters struct {
	StringMatches []V2QueryGetResponseParametersStringMatch `json:"string_matches" api:"required"`
	MaxTime       string                                    `json:"max_time"`
	MinTime       string                                    `json:"min_time"`
	JSON          v2QueryGetResponseParametersJSON          `json:"-"`
}

// v2QueryGetResponseParametersJSON contains the JSON metadata for the struct
// [V2QueryGetResponseParameters]
type v2QueryGetResponseParametersJSON struct {
	StringMatches apijson.Field
	MaxTime       apijson.Field
	MinTime       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2QueryGetResponseParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseParametersJSON) RawJSON() string {
	return r.raw
}

type V2QueryGetResponseParametersStringMatch struct {
	MaxEditDistance float64                                     `json:"max_edit_distance" api:"required"`
	Pattern         string                                      `json:"pattern" api:"required"`
	JSON            v2QueryGetResponseParametersStringMatchJSON `json:"-"`
}

// v2QueryGetResponseParametersStringMatchJSON contains the JSON metadata for the
// struct [V2QueryGetResponseParametersStringMatch]
type v2QueryGetResponseParametersStringMatchJSON struct {
	MaxEditDistance apijson.Field
	Pattern         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2QueryGetResponseParametersStringMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseParametersStringMatchJSON) RawJSON() string {
	return r.raw
}

type V2QueryGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	ID        param.Field[string] `query:"id"`
}

// URLQuery serializes [V2QueryGetParams]'s query parameters as `url.Values`.
func (r V2QueryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
