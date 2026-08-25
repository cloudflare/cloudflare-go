// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// TransformerVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransformerVersionService] method instead.
type TransformerVersionService struct {
	Options []option.RequestOption
}

// NewTransformerVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransformerVersionService(opts ...option.RequestOption) (r *TransformerVersionService) {
	r = &TransformerVersionService{}
	r.Options = opts
	return
}

// Returns version metadata for a transformer, newest first. Each version
// corresponds to a SQL query update.
func (r *TransformerVersionService) List(ctx context.Context, transformerID int64, params TransformerVersionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[TransformerVersionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/logpush/transformers/%v/versions", params.AccountID, transformerID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns version metadata for a transformer, newest first. Each version
// corresponds to a SQL query update.
func (r *TransformerVersionService) ListAutoPaging(ctx context.Context, transformerID int64, params TransformerVersionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TransformerVersionListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, transformerID, params, opts...))
}

type TransformerVersionListResponse struct {
	// Unique identifier for this version.
	ID int64 `json:"id"`
	// When this version was created (RFC 3339).
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Sequential version number.
	Version int64                              `json:"version"`
	JSON    transformerVersionListResponseJSON `json:"-"`
}

// transformerVersionListResponseJSON contains the JSON metadata for the struct
// [TransformerVersionListResponse]
type transformerVersionListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerVersionListResponseJSON) RawJSON() string {
	return r.raw
}

type TransformerVersionListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Maximum number of versions to return.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [TransformerVersionListParams]'s query parameters as
// `url.Values`.
func (r TransformerVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
