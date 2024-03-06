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

// RadarEntityService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEntityService] method
// instead.
type RadarEntityService struct {
	Options   []option.RequestOption
	ASNs      *RadarEntityASNService
	Locations *RadarEntityLocationService
}

// NewRadarEntityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarEntityService(opts ...option.RequestOption) (r *RadarEntityService) {
	r = &RadarEntityService{}
	r.Options = opts
	r.ASNs = NewRadarEntityASNService(opts...)
	r.Locations = NewRadarEntityLocationService(opts...)
	return
}

// Get IP address information.
func (r *RadarEntityService) Get(ctx context.Context, query RadarEntityGetParams, opts ...option.RequestOption) (res *RadarEntityGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityGetResponseEnvelope
	path := "radar/entities/ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEntityGetResponse struct {
	IP   RadarEntityGetResponseIP   `json:"ip,required"`
	JSON radarEntityGetResponseJSON `json:"-"`
}

// radarEntityGetResponseJSON contains the JSON metadata for the struct
// [RadarEntityGetResponse]
type radarEntityGetResponseJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityGetResponseIP struct {
	ASN          string                       `json:"asn,required"`
	ASNLocation  string                       `json:"asnLocation,required"`
	ASNName      string                       `json:"asnName,required"`
	ASNOrgName   string                       `json:"asnOrgName,required"`
	IP           string                       `json:"ip,required"`
	IPVersion    string                       `json:"ipVersion,required"`
	Location     string                       `json:"location,required"`
	LocationName string                       `json:"locationName,required"`
	JSON         radarEntityGetResponseIPJSON `json:"-"`
}

// radarEntityGetResponseIPJSON contains the JSON metadata for the struct
// [RadarEntityGetResponseIP]
type radarEntityGetResponseIPJSON struct {
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

func (r *RadarEntityGetResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityGetResponseIPJSON) RawJSON() string {
	return r.raw
}

type RadarEntityGetParams struct {
	// IP address.
	IP param.Field[string] `query:"ip,required"`
	// Format results are returned in.
	Format param.Field[RadarEntityGetParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityGetParams]'s query parameters as `url.Values`.
func (r RadarEntityGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityGetParamsFormat string

const (
	RadarEntityGetParamsFormatJson RadarEntityGetParamsFormat = "JSON"
	RadarEntityGetParamsFormatCsv  RadarEntityGetParamsFormat = "CSV"
)

type RadarEntityGetResponseEnvelope struct {
	Result  RadarEntityGetResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    radarEntityGetResponseEnvelopeJSON `json:"-"`
}

// radarEntityGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarEntityGetResponseEnvelope]
type radarEntityGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
