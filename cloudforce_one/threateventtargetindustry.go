// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

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

// ThreatEventTargetIndustryService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventTargetIndustryService] method instead.
type ThreatEventTargetIndustryService struct {
	Options []option.RequestOption
}

// NewThreatEventTargetIndustryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatEventTargetIndustryService(opts ...option.RequestOption) (r *ThreatEventTargetIndustryService) {
	r = &ThreatEventTargetIndustryService{}
	r.Options = opts
	return
}

// Lists target industries across multiple datasets
func (r *ThreatEventTargetIndustryService) List(ctx context.Context, params ThreatEventTargetIndustryListParams, opts ...option.RequestOption) (res *ThreatEventTargetIndustryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/targetIndustries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type ThreatEventTargetIndustryListResponse struct {
	Items ThreatEventTargetIndustryListResponseItems `json:"items,required"`
	Type  string                                     `json:"type,required"`
	JSON  threatEventTargetIndustryListResponseJSON  `json:"-"`
}

// threatEventTargetIndustryListResponseJSON contains the JSON metadata for the
// struct [ThreatEventTargetIndustryListResponse]
type threatEventTargetIndustryListResponseJSON struct {
	Items       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTargetIndustryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTargetIndustryListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTargetIndustryListResponseItems struct {
	Type string                                         `json:"type,required"`
	JSON threatEventTargetIndustryListResponseItemsJSON `json:"-"`
}

// threatEventTargetIndustryListResponseItemsJSON contains the JSON metadata for
// the struct [ThreatEventTargetIndustryListResponseItems]
type threatEventTargetIndustryListResponseItemsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTargetIndustryListResponseItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTargetIndustryListResponseItemsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTargetIndustryListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Array of dataset IDs to query target industries from. If not provided, uses the
	// default dataset.
	DatasetIDs param.Field[[]string] `query:"datasetIds"`
}

// URLQuery serializes [ThreatEventTargetIndustryListParams]'s query parameters as
// `url.Values`.
func (r ThreatEventTargetIndustryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
