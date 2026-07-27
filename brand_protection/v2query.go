// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
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
func (r *V2QueryService) Get(ctx context.Context, params V2QueryGetParams, opts ...option.RequestOption) (res *V2QueryGetResponseUnion, err error) {
	var env V2QueryGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/brand-protection/domain/queries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Union satisfied by [V2QueryGetResponseArray] or [V2QueryGetResponseObject].
type V2QueryGetResponseUnion interface {
	implementsV2QueryGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2QueryGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2QueryGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2QueryGetResponseObject{}),
		},
	)
}

type V2QueryGetResponseArray []V2QueryGetResponseArrayItem

func (r V2QueryGetResponseArray) implementsV2QueryGetResponseUnion() {}

type V2QueryGetResponseArrayItem struct {
	Created    string                            `json:"created" api:"required"`
	Parameters V2QueryGetResponseArrayParameters `json:"parameters" api:"required,nullable"`
	QueryID    int64                             `json:"query_id" api:"required"`
	QueryTag   string                            `json:"query_tag" api:"required"`
	Scan       bool                              `json:"scan" api:"required"`
	Updated    string                            `json:"updated" api:"required"`
	JSON       v2QueryGetResponseArrayItemJSON   `json:"-"`
}

// v2QueryGetResponseArrayItemJSON contains the JSON metadata for the struct
// [V2QueryGetResponseArrayItem]
type v2QueryGetResponseArrayItemJSON struct {
	Created     apijson.Field
	Parameters  apijson.Field
	QueryID     apijson.Field
	QueryTag    apijson.Field
	Scan        apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2QueryGetResponseArrayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseArrayItemJSON) RawJSON() string {
	return r.raw
}

type V2QueryGetResponseArrayParameters struct {
	StringMatches []V2QueryGetResponseArrayParametersStringMatch `json:"string_matches" api:"required"`
	MaxTime       string                                         `json:"max_time"`
	MinTime       string                                         `json:"min_time"`
	JSON          v2QueryGetResponseArrayParametersJSON          `json:"-"`
}

// v2QueryGetResponseArrayParametersJSON contains the JSON metadata for the struct
// [V2QueryGetResponseArrayParameters]
type v2QueryGetResponseArrayParametersJSON struct {
	StringMatches apijson.Field
	MaxTime       apijson.Field
	MinTime       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2QueryGetResponseArrayParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseArrayParametersJSON) RawJSON() string {
	return r.raw
}

type V2QueryGetResponseArrayParametersStringMatch struct {
	Pattern string                                           `json:"pattern" api:"required"`
	JSON    v2QueryGetResponseArrayParametersStringMatchJSON `json:"-"`
}

// v2QueryGetResponseArrayParametersStringMatchJSON contains the JSON metadata for
// the struct [V2QueryGetResponseArrayParametersStringMatch]
type v2QueryGetResponseArrayParametersStringMatchJSON struct {
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2QueryGetResponseArrayParametersStringMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseArrayParametersStringMatchJSON) RawJSON() string {
	return r.raw
}

type V2QueryGetResponseObject struct {
	Created    string                             `json:"created" api:"required"`
	Parameters V2QueryGetResponseObjectParameters `json:"parameters" api:"required,nullable"`
	QueryID    int64                              `json:"query_id" api:"required"`
	QueryTag   string                             `json:"query_tag" api:"required"`
	Scan       bool                               `json:"scan" api:"required"`
	Updated    string                             `json:"updated" api:"required"`
	JSON       v2QueryGetResponseObjectJSON       `json:"-"`
}

// v2QueryGetResponseObjectJSON contains the JSON metadata for the struct
// [V2QueryGetResponseObject]
type v2QueryGetResponseObjectJSON struct {
	Created     apijson.Field
	Parameters  apijson.Field
	QueryID     apijson.Field
	QueryTag    apijson.Field
	Scan        apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2QueryGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r V2QueryGetResponseObject) implementsV2QueryGetResponseUnion() {}

type V2QueryGetResponseObjectParameters struct {
	StringMatches []V2QueryGetResponseObjectParametersStringMatch `json:"string_matches" api:"required"`
	MaxTime       string                                          `json:"max_time"`
	MinTime       string                                          `json:"min_time"`
	JSON          v2QueryGetResponseObjectParametersJSON          `json:"-"`
}

// v2QueryGetResponseObjectParametersJSON contains the JSON metadata for the struct
// [V2QueryGetResponseObjectParameters]
type v2QueryGetResponseObjectParametersJSON struct {
	StringMatches apijson.Field
	MaxTime       apijson.Field
	MinTime       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2QueryGetResponseObjectParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseObjectParametersJSON) RawJSON() string {
	return r.raw
}

type V2QueryGetResponseObjectParametersStringMatch struct {
	Pattern string                                            `json:"pattern" api:"required"`
	JSON    v2QueryGetResponseObjectParametersStringMatchJSON `json:"-"`
}

// v2QueryGetResponseObjectParametersStringMatchJSON contains the JSON metadata for
// the struct [V2QueryGetResponseObjectParametersStringMatch]
type v2QueryGetResponseObjectParametersStringMatchJSON struct {
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2QueryGetResponseObjectParametersStringMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseObjectParametersStringMatchJSON) RawJSON() string {
	return r.raw
}

type V2QueryGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	ID        param.Field[string] `query:"id"`
	// Optional page number for paginated list requests. Defaults to 1 when only
	// per_page is supplied. Omit page and per_page to preserve the legacy full-list
	// response.
	Page param.Field[int64] `query:"page"`
	// Optional number of queries per page for paginated list requests. Defaults to 100
	// when only page is supplied. Maximum 100. Omit page and per_page to preserve the
	// legacy full-list response.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [V2QueryGetParams]'s query parameters as `url.Values`.
func (r V2QueryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type V2QueryGetResponseEnvelope struct {
	Errors   []V2QueryGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []V2QueryGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   V2QueryGetResponseUnion              `json:"result" api:"required"`
	Success  bool                                 `json:"success" api:"required"`
	// Present on paginated list responses when page or per_page is supplied.
	ResultInfo V2QueryGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       v2QueryGetResponseEnvelopeJSON       `json:"-"`
}

// v2QueryGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [V2QueryGetResponseEnvelope]
type v2QueryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2QueryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V2QueryGetResponseEnvelopeErrors struct {
	Message string                               `json:"message" api:"required"`
	Code    V2QueryGetResponseEnvelopeErrorsCode `json:"code"`
	JSON    v2QueryGetResponseEnvelopeErrorsJSON `json:"-"`
}

// v2QueryGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V2QueryGetResponseEnvelopeErrors]
type v2QueryGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2QueryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type V2QueryGetResponseEnvelopeErrorsCode interface {
	ImplementsV2QueryGetResponseEnvelopeErrorsCode()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2QueryGetResponseEnvelopeErrorsCode)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type V2QueryGetResponseEnvelopeMessages struct {
	Message string                                 `json:"message" api:"required"`
	Code    V2QueryGetResponseEnvelopeMessagesCode `json:"code"`
	JSON    v2QueryGetResponseEnvelopeMessagesJSON `json:"-"`
}

// v2QueryGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [V2QueryGetResponseEnvelopeMessages]
type v2QueryGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2QueryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type V2QueryGetResponseEnvelopeMessagesCode interface {
	ImplementsV2QueryGetResponseEnvelopeMessagesCode()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2QueryGetResponseEnvelopeMessagesCode)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

// Present on paginated list responses when page or per_page is supplied.
type V2QueryGetResponseEnvelopeResultInfo struct {
	Count      int64                                    `json:"count" api:"required"`
	Page       int64                                    `json:"page" api:"required"`
	PerPage    int64                                    `json:"per_page" api:"required"`
	TotalCount int64                                    `json:"total_count" api:"required"`
	JSON       v2QueryGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// v2QueryGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [V2QueryGetResponseEnvelopeResultInfo]
type v2QueryGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2QueryGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QueryGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
