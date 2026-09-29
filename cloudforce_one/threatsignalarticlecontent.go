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

// ThreatSignalArticleContentService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalArticleContentService] method instead.
type ThreatSignalArticleContentService struct {
	Options []option.RequestOption
}

// NewThreatSignalArticleContentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatSignalArticleContentService(opts ...option.RequestOption) (r *ThreatSignalArticleContentService) {
	r = &ThreatSignalArticleContentService{}
	r.Options = opts
	return
}

// Retrieves the stored body of a Threat Signals article as plain text or HTML.
func (r *ThreatSignalArticleContentService) Get(ctx context.Context, articleID string, params ThreatSignalArticleContentGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/html")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if articleID == "" {
		err = errors.New("missing required article_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/articles/%s/content", params.AccountID, articleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ThreatSignalArticleContentGetParams struct {
	AccountID param.Field[string]                                    `path:"account_id" api:"required"`
	Format    param.Field[ThreatSignalArticleContentGetParamsFormat] `query:"format"`
}

// URLQuery serializes [ThreatSignalArticleContentGetParams]'s query parameters as
// `url.Values`.
func (r ThreatSignalArticleContentGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatSignalArticleContentGetParamsFormat string

const (
	ThreatSignalArticleContentGetParamsFormatText ThreatSignalArticleContentGetParamsFormat = "text"
	ThreatSignalArticleContentGetParamsFormatHTML ThreatSignalArticleContentGetParamsFormat = "html"
)

func (r ThreatSignalArticleContentGetParamsFormat) IsKnown() bool {
	switch r {
	case ThreatSignalArticleContentGetParamsFormatText, ThreatSignalArticleContentGetParamsFormatHTML:
		return true
	}
	return false
}
