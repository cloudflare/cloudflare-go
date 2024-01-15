// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEntityIPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEntityIPService] method
// instead.
type RadarEntityIPService struct {
	Options []option.RequestOption
}

// NewRadarEntityIPService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarEntityIPService(opts ...option.RequestOption) (r *RadarEntityIPService) {
	r = &RadarEntityIPService{}
	r.Options = opts
	return
}

// Get IP address information.
func (r *RadarEntityIPService) List(ctx context.Context, query RadarEntityIPListParams, opts ...option.RequestOption) (res *RadarEntityIPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/entities/ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEntityIPListResponse struct {
	Result  RadarEntityIPListResponseResult `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    radarEntityIPListResponseJSON   `json:"-"`
}

// radarEntityIPListResponseJSON contains the JSON metadata for the struct
// [RadarEntityIPListResponse]
type radarEntityIPListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityIPListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityIPListResponseResult struct {
	IP   RadarEntityIPListResponseResultIP   `json:"ip,required"`
	JSON radarEntityIPListResponseResultJSON `json:"-"`
}

// radarEntityIPListResponseResultJSON contains the JSON metadata for the struct
// [RadarEntityIPListResponseResult]
type radarEntityIPListResponseResultJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityIPListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityIPListResponseResultIP struct {
	ASN          string                                `json:"asn,required"`
	ASNLocation  string                                `json:"asnLocation,required"`
	ASNName      string                                `json:"asnName,required"`
	ASNOrgName   string                                `json:"asnOrgName,required"`
	IP           string                                `json:"ip,required"`
	IPVersion    string                                `json:"ipVersion,required"`
	Location     string                                `json:"location,required"`
	LocationName string                                `json:"locationName,required"`
	JSON         radarEntityIPListResponseResultIPJSON `json:"-"`
}

// radarEntityIPListResponseResultIPJSON contains the JSON metadata for the struct
// [RadarEntityIPListResponseResultIP]
type radarEntityIPListResponseResultIPJSON struct {
	ASN          apijson.Field
	ASNLocation  apijson.Field
	ASNName      apijson.Field
	ASNOrgName   apijson.Field
	IP           apijson.Field
	IPVersion    apijson.Field
	Location     apijson.Field
	LocationName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEntityIPListResponseResultIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityIPListParams struct {
	// IP address.
	IP param.Field[string] `query:"ip,required"`
	// Format results are returned in.
	Format param.Field[RadarEntityIPListParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityIPListParams]'s query parameters as
// `url.Values`.
func (r RadarEntityIPListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityIPListParamsFormat string

const (
	RadarEntityIPListParamsFormatJson RadarEntityIPListParamsFormat = "JSON"
	RadarEntityIPListParamsFormatCsv  RadarEntityIPListParamsFormat = "CSV"
)
