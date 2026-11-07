// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// V2LogoMatchService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2LogoMatchService] method instead.
type V2LogoMatchService struct {
	Options []option.RequestOption
}

// NewV2LogoMatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2LogoMatchService(opts ...option.RequestOption) (r *V2LogoMatchService) {
	r = &V2LogoMatchService{}
	r.Options = opts
	return
}

// Lists images that matched a saved logo query, most recently matched first by
// default. Set `download=true` to include each matched image as base64 data.
func (r *V2LogoMatchService) Get(ctx context.Context, params V2LogoMatchGetParams, opts ...option.RequestOption) (res *V2LogoMatchGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/brand-protection/logo/matches", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type V2LogoMatchGetResponse struct {
	Matches []V2LogoMatchGetResponseMatch `json:"matches" api:"required"`
	Total   int64                         `json:"total" api:"required"`
	JSON    v2LogoMatchGetResponseJSON    `json:"-"`
}

// v2LogoMatchGetResponseJSON contains the JSON metadata for the struct
// [V2LogoMatchGetResponse]
type v2LogoMatchGetResponseJSON struct {
	Matches     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2LogoMatchGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2LogoMatchGetResponseJSON) RawJSON() string {
	return r.raw
}

type V2LogoMatchGetResponseMatch struct {
	// ID of the logo match.
	ID int64 `json:"id" api:"required"`
	// Domain where the matching image was found, if known.
	Domain string `json:"domain" api:"required,nullable"`
	// When the match was recorded.
	MatchedAt string `json:"matched_at" api:"required,nullable"`
	// ID of the logo query that produced this match.
	QueryID int64 `json:"query_id" api:"required"`
	// Registrar of `domain`, if known.
	Registrar string `json:"registrar" api:"required,nullable"`
	// Visual similarity between the matched image and the query logo. Higher is
	// closer.
	SimilarityScore float64 `json:"similarity_score" api:"required"`
	URLScanID       string  `json:"url_scan_id" api:"required,nullable"`
	// MIME type of the matched image (only present when download=true)
	ContentType string `json:"content_type"`
	// Base64-encoded matched image (only present when download=true)
	ImageData string                          `json:"image_data"`
	JSON      v2LogoMatchGetResponseMatchJSON `json:"-"`
}

// v2LogoMatchGetResponseMatchJSON contains the JSON metadata for the struct
// [V2LogoMatchGetResponseMatch]
type v2LogoMatchGetResponseMatchJSON struct {
	ID              apijson.Field
	Domain          apijson.Field
	MatchedAt       apijson.Field
	QueryID         apijson.Field
	Registrar       apijson.Field
	SimilarityScore apijson.Field
	URLScanID       apijson.Field
	ContentType     apijson.Field
	ImageData       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V2LogoMatchGetResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2LogoMatchGetResponseMatchJSON) RawJSON() string {
	return r.raw
}

type V2LogoMatchGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// ID of the logo query whose matches to list.
	QueryID param.Field[string] `query:"query_id" api:"required"`
	// Set to `true` to include each matched image as base64 data in `image_data`, with
	// its MIME type in `content_type`.
	Download param.Field[string] `query:"download"`
	// Maximum number of matches to return. Defaults to 50.
	Limit param.Field[string] `query:"limit"`
	// Number of matches to skip, for pagination. Defaults to 0.
	Offset param.Field[string] `query:"offset"`
	// Sort direction: `asc` or `desc`. Defaults to `desc`.
	Order param.Field[V2LogoMatchGetParamsOrder] `query:"order"`
	// Field to sort by: `matchedAt`, `domain`, `similarityScore`, or `registrar`.
	// Defaults to `matchedAt`.
	OrderBy param.Field[V2LogoMatchGetParamsOrderBy] `query:"orderBy"`
}

// URLQuery serializes [V2LogoMatchGetParams]'s query parameters as `url.Values`.
func (r V2LogoMatchGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction: `asc` or `desc`. Defaults to `desc`.
type V2LogoMatchGetParamsOrder string

const (
	V2LogoMatchGetParamsOrderAsc  V2LogoMatchGetParamsOrder = "asc"
	V2LogoMatchGetParamsOrderDesc V2LogoMatchGetParamsOrder = "desc"
)

func (r V2LogoMatchGetParamsOrder) IsKnown() bool {
	switch r {
	case V2LogoMatchGetParamsOrderAsc, V2LogoMatchGetParamsOrderDesc:
		return true
	}
	return false
}

// Field to sort by: `matchedAt`, `domain`, `similarityScore`, or `registrar`.
// Defaults to `matchedAt`.
type V2LogoMatchGetParamsOrderBy string

const (
	V2LogoMatchGetParamsOrderByMatchedAt       V2LogoMatchGetParamsOrderBy = "matchedAt"
	V2LogoMatchGetParamsOrderByDomain          V2LogoMatchGetParamsOrderBy = "domain"
	V2LogoMatchGetParamsOrderBySimilarityScore V2LogoMatchGetParamsOrderBy = "similarityScore"
	V2LogoMatchGetParamsOrderByRegistrar       V2LogoMatchGetParamsOrderBy = "registrar"
)

func (r V2LogoMatchGetParamsOrderBy) IsKnown() bool {
	switch r {
	case V2LogoMatchGetParamsOrderByMatchedAt, V2LogoMatchGetParamsOrderByDomain, V2LogoMatchGetParamsOrderBySimilarityScore, V2LogoMatchGetParamsOrderByRegistrar:
		return true
	}
	return false
}
