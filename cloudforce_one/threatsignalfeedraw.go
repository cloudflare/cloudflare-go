// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ThreatSignalFeedRawService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalFeedRawService] method instead.
type ThreatSignalFeedRawService struct {
	Options []option.RequestOption
}

// NewThreatSignalFeedRawService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatSignalFeedRawService(opts ...option.RequestOption) (r *ThreatSignalFeedRawService) {
	r = &ThreatSignalFeedRawService{}
	r.Options = opts
	return
}

// Retrieves the feed document fetched by the most recent poll.
func (r *ThreatSignalFeedRawService) Get(ctx context.Context, feedID string, params ThreatSignalFeedRawGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/xml")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if feedID == "" {
		err = errors.New("missing required feed_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/feeds/%s/raw", params.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ThreatSignalFeedRawGetParams struct {
	AccountID param.Field[string]                             `path:"account_id" api:"required"`
	Format    param.Field[ThreatSignalFeedRawGetParamsFormat] `query:"format"`
}

// URLQuery serializes [ThreatSignalFeedRawGetParams]'s query parameters as
// `url.Values`.
func (r ThreatSignalFeedRawGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatSignalFeedRawGetParamsFormat string

const (
	ThreatSignalFeedRawGetParamsFormatXml ThreatSignalFeedRawGetParamsFormat = "xml"
)

func (r ThreatSignalFeedRawGetParamsFormat) IsKnown() bool {
	switch r {
	case ThreatSignalFeedRawGetParamsFormatXml:
		return true
	}
	return false
}
