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

// V2LogoService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2LogoService] method instead.
type V2LogoService struct {
	Options []option.RequestOption
}

// NewV2LogoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2LogoService(opts ...option.RequestOption) (r *V2LogoService) {
	r = &V2LogoService{}
	r.Options = opts
	return
}

// Create a new saved brand protection logo query for visual similarity matching
func (r *V2LogoService) New(ctx context.Context, params V2LogoNewParams, opts ...option.RequestOption) (res *V2LogoNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/brand-protection/logo/queries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Delete a saved brand protection logo query. Returns 404 if the query ID doesn't
// exist.
func (r *V2LogoService) Delete(ctx context.Context, queryID string, body V2LogoDeleteParams, opts ...option.RequestOption) (res *V2LogoDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if queryID == "" {
		err = errors.New("missing required query_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/brand-protection/logo/queries/%s", body.AccountID, queryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get all saved brand protection logo queries for an account. Optionally specify
// id to get a single query. Set download=true to include base64-encoded image
// data.
func (r *V2LogoService) Get(ctx context.Context, params V2LogoGetParams, opts ...option.RequestOption) (res *[]V2LogoGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/brand-protection/logo/queries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type V2LogoNewResponse struct {
	Message string                `json:"message" api:"required"`
	Success bool                  `json:"success" api:"required"`
	QueryID int64                 `json:"query_id"`
	JSON    v2LogoNewResponseJSON `json:"-"`
}

// v2LogoNewResponseJSON contains the JSON metadata for the struct
// [V2LogoNewResponse]
type v2LogoNewResponseJSON struct {
	Message     apijson.Field
	Success     apijson.Field
	QueryID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2LogoNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2LogoNewResponseJSON) RawJSON() string {
	return r.raw
}

type V2LogoDeleteResponse struct {
	Message string                   `json:"message" api:"required"`
	Success bool                     `json:"success" api:"required"`
	JSON    v2LogoDeleteResponseJSON `json:"-"`
}

// v2LogoDeleteResponseJSON contains the JSON metadata for the struct
// [V2LogoDeleteResponse]
type v2LogoDeleteResponseJSON struct {
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2LogoDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2LogoDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type V2LogoGetResponse struct {
	ID                  int64   `json:"id" api:"required"`
	R2Path              string  `json:"r2_path" api:"required"`
	SimilarityThreshold float64 `json:"similarity_threshold" api:"required"`
	Tag                 string  `json:"tag" api:"required"`
	UploadedAt          string  `json:"uploaded_at" api:"required,nullable"`
	// MIME type of the image (only present when download=true)
	ContentType string `json:"content_type"`
	// Base64-encoded image data (only present when download=true)
	ImageData string                `json:"image_data"`
	JSON      v2LogoGetResponseJSON `json:"-"`
}

// v2LogoGetResponseJSON contains the JSON metadata for the struct
// [V2LogoGetResponse]
type v2LogoGetResponseJSON struct {
	ID                  apijson.Field
	R2Path              apijson.Field
	SimilarityThreshold apijson.Field
	Tag                 apijson.Field
	UploadedAt          apijson.Field
	ContentType         apijson.Field
	ImageData           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V2LogoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2LogoGetResponseJSON) RawJSON() string {
	return r.raw
}

type V2LogoNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Base64 encoded image data. Can include data URI prefix (e.g.,
	// 'data:image/png;base64,...') or just the base64 string.
	ImageData param.Field[string] `json:"image_data" api:"required"`
	// Minimum similarity score (0-1) required for visual matches
	SimilarityThreshold param.Field[float64] `json:"similarity_threshold" api:"required"`
	// Unique identifier for the logo query
	Tag param.Field[string] `json:"tag" api:"required"`
	// If true, search historic scanned images for matches above the similarity
	// threshold
	SearchLookback param.Field[bool] `json:"search_lookback"`
}

func (r V2LogoNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2LogoDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type V2LogoGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Optional query ID to retrieve a specific logo query
	ID param.Field[string] `query:"id"`
	// If true, include base64-encoded image data in the response
	Download param.Field[string] `query:"download"`
}

// URLQuery serializes [V2LogoGetParams]'s query parameters as `url.Values`.
func (r V2LogoGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
