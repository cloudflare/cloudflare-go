// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SiteService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteService] method instead.
type SiteService struct {
	Options []option.RequestOption
	ACLs    *SiteACLService
	LANs    *SiteLANService
	WANs    *SiteWANService
}

// NewSiteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSiteService(opts ...option.RequestOption) (r *SiteService) {
	r = &SiteService{}
	r.Options = opts
	r.ACLs = NewSiteACLService(opts...)
	r.LANs = NewSiteLANService(opts...)
	r.WANs = NewSiteWANService(opts...)
	return
}

// Creates a new Site
func (r *SiteService) New(ctx context.Context, params SiteNewParams, opts ...option.RequestOption) (res *Site, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a specific Site.
func (r *SiteService) Update(ctx context.Context, siteID string, params SiteUpdateParams, opts ...option.RequestOption) (res *Site, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Sites associated with an account. Use connector_identifier query param to
// return sites where connector_identifier matches either site.ConnectorID or
// site.SecondaryConnectorID.
func (r *SiteService) List(ctx context.Context, params SiteListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Site], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/magic/sites", params.AccountID)
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

// Lists Sites associated with an account. Use connector_identifier query param to
// return sites where connector_identifier matches either site.ConnectorID or
// site.SecondaryConnectorID.
func (r *SiteService) ListAutoPaging(ctx context.Context, params SiteListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Site] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Remove a specific Site.
func (r *SiteService) Delete(ctx context.Context, siteID string, body SiteDeleteParams, opts ...option.RequestOption) (res *Site, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", body.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific Site.
func (r *SiteService) Get(ctx context.Context, siteID string, query SiteGetParams, opts ...option.RequestOption) (res *Site, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", query.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Site struct {
	// Identifier
	ID string `json:"id"`
	// Magic WAN Connector identifier tag.
	ConnectorID string `json:"connector_id"`
	Description string `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode bool `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location SiteLocation `json:"location"`
	// The name of the site.
	Name string `json:"name"`
	// Magic WAN Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID string   `json:"secondary_connector_id"`
	JSON                 siteJSON `json:"-"`
}

// siteJSON contains the JSON metadata for the struct [Site]
type siteJSON struct {
	ID                   apijson.Field
	ConnectorID          apijson.Field
	Description          apijson.Field
	HaMode               apijson.Field
	Location             apijson.Field
	Name                 apijson.Field
	SecondaryConnectorID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Site) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteJSON) RawJSON() string {
	return r.raw
}

// Location of site in latitude and longitude.
type SiteLocation struct {
	// Latitude
	Lat string `json:"lat"`
	// Longitude
	Lon  string           `json:"lon"`
	JSON siteLocationJSON `json:"-"`
}

// siteLocationJSON contains the JSON metadata for the struct [SiteLocation]
type siteLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLocationJSON) RawJSON() string {
	return r.raw
}

// Location of site in latitude and longitude.
type SiteLocationParam struct {
	// Latitude
	Lat param.Field[string] `json:"lat"`
	// Longitude
	Lon param.Field[string] `json:"lon"`
}

func (r SiteLocationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the site.
	Name param.Field[string] `json:"name,required"`
	// Magic WAN Connector identifier tag.
	ConnectorID param.Field[string] `json:"connector_id"`
	Description param.Field[string] `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode param.Field[bool] `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location param.Field[SiteLocationParam] `json:"location"`
	// Magic WAN Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID param.Field[string] `json:"secondary_connector_id"`
}

func (r SiteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Site                  `json:"result,required"`
	// Whether the API call was successful
	Success SiteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteNewResponseEnvelopeJSON    `json:"-"`
}

// siteNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteNewResponseEnvelope]
type siteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteNewResponseEnvelopeSuccess bool

const (
	SiteNewResponseEnvelopeSuccessTrue SiteNewResponseEnvelopeSuccess = true
)

func (r SiteNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Magic WAN Connector identifier tag.
	ConnectorID param.Field[string] `json:"connector_id"`
	Description param.Field[string] `json:"description"`
	// Location of site in latitude and longitude.
	Location param.Field[SiteLocationParam] `json:"location"`
	// The name of the site.
	Name param.Field[string] `json:"name"`
	// Magic WAN Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID param.Field[string] `json:"secondary_connector_id"`
}

func (r SiteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Site                  `json:"result,required"`
	// Whether the API call was successful
	Success SiteUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteUpdateResponseEnvelopeJSON    `json:"-"`
}

// siteUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteUpdateResponseEnvelope]
type siteUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteUpdateResponseEnvelopeSuccess bool

const (
	SiteUpdateResponseEnvelopeSuccessTrue SiteUpdateResponseEnvelopeSuccess = true
)

func (r SiteUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Identifier
	ConnectorIdentifier param.Field[string] `query:"connector_identifier"`
}

// URLQuery serializes [SiteListParams]'s query parameters as `url.Values`.
func (r SiteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SiteDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Site                  `json:"result,required"`
	// Whether the API call was successful
	Success SiteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteDeleteResponseEnvelopeJSON    `json:"-"`
}

// siteDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteDeleteResponseEnvelope]
type siteDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteDeleteResponseEnvelopeSuccess bool

const (
	SiteDeleteResponseEnvelopeSuccessTrue SiteDeleteResponseEnvelopeSuccess = true
)

func (r SiteDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Site                  `json:"result,required"`
	// Whether the API call was successful
	Success SiteGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteGetResponseEnvelopeJSON    `json:"-"`
}

// siteGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteGetResponseEnvelope]
type siteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteGetResponseEnvelopeSuccess bool

const (
	SiteGetResponseEnvelopeSuccessTrue SiteGetResponseEnvelopeSuccess = true
)

func (r SiteGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
