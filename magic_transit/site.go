// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SiteService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSiteService] method instead.
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
func (r *SiteService) New(ctx context.Context, params SiteNewParams, opts ...option.RequestOption) (res *SiteNewResponse, err error) {
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
func (r *SiteService) Update(ctx context.Context, siteID string, params SiteUpdateParams, opts ...option.RequestOption) (res *SiteUpdateResponse, err error) {
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
func (r *SiteService) List(ctx context.Context, params SiteListParams, opts ...option.RequestOption) (res *SiteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a specific Site.
func (r *SiteService) Delete(ctx context.Context, siteID string, params SiteDeleteParams, opts ...option.RequestOption) (res *SiteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific Site.
func (r *SiteService) Get(ctx context.Context, siteID string, query SiteGetParams, opts ...option.RequestOption) (res *SiteGetResponse, err error) {
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

type SiteNewResponse struct {
	Site SiteNewResponseSite `json:"site"`
	JSON siteNewResponseJSON `json:"-"`
}

// siteNewResponseJSON contains the JSON metadata for the struct [SiteNewResponse]
type siteNewResponseJSON struct {
	Site        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteNewResponseJSON) RawJSON() string {
	return r.raw
}

type SiteNewResponseSite struct {
	// Identifier
	ID string `json:"id"`
	// Magic WAN Connector identifier tag.
	ConnectorID string `json:"connector_id"`
	Description string `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode bool `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location SiteNewResponseSiteLocation `json:"location"`
	// The name of the site.
	Name string `json:"name"`
	// Magic WAN Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID string                  `json:"secondary_connector_id"`
	JSON                 siteNewResponseSiteJSON `json:"-"`
}

// siteNewResponseSiteJSON contains the JSON metadata for the struct
// [SiteNewResponseSite]
type siteNewResponseSiteJSON struct {
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

func (r *SiteNewResponseSite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteNewResponseSiteJSON) RawJSON() string {
	return r.raw
}

// Location of site in latitude and longitude.
type SiteNewResponseSiteLocation struct {
	// Latitude
	Lat string `json:"lat"`
	// Longitude
	Lon  string                          `json:"lon"`
	JSON siteNewResponseSiteLocationJSON `json:"-"`
}

// siteNewResponseSiteLocationJSON contains the JSON metadata for the struct
// [SiteNewResponseSiteLocation]
type siteNewResponseSiteLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteNewResponseSiteLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteNewResponseSiteLocationJSON) RawJSON() string {
	return r.raw
}

type SiteUpdateResponse struct {
	Site SiteUpdateResponseSite `json:"site"`
	JSON siteUpdateResponseJSON `json:"-"`
}

// siteUpdateResponseJSON contains the JSON metadata for the struct
// [SiteUpdateResponse]
type siteUpdateResponseJSON struct {
	Site        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SiteUpdateResponseSite struct {
	// Identifier
	ID string `json:"id"`
	// Magic WAN Connector identifier tag.
	ConnectorID string `json:"connector_id"`
	Description string `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode bool `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location SiteUpdateResponseSiteLocation `json:"location"`
	// The name of the site.
	Name string `json:"name"`
	// Magic WAN Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID string                     `json:"secondary_connector_id"`
	JSON                 siteUpdateResponseSiteJSON `json:"-"`
}

// siteUpdateResponseSiteJSON contains the JSON metadata for the struct
// [SiteUpdateResponseSite]
type siteUpdateResponseSiteJSON struct {
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

func (r *SiteUpdateResponseSite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteUpdateResponseSiteJSON) RawJSON() string {
	return r.raw
}

// Location of site in latitude and longitude.
type SiteUpdateResponseSiteLocation struct {
	// Latitude
	Lat string `json:"lat"`
	// Longitude
	Lon  string                             `json:"lon"`
	JSON siteUpdateResponseSiteLocationJSON `json:"-"`
}

// siteUpdateResponseSiteLocationJSON contains the JSON metadata for the struct
// [SiteUpdateResponseSiteLocation]
type siteUpdateResponseSiteLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteUpdateResponseSiteLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteUpdateResponseSiteLocationJSON) RawJSON() string {
	return r.raw
}

type SiteListResponse struct {
	Sites []SiteListResponseSite `json:"sites"`
	JSON  siteListResponseJSON   `json:"-"`
}

// siteListResponseJSON contains the JSON metadata for the struct
// [SiteListResponse]
type siteListResponseJSON struct {
	Sites       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteListResponseJSON) RawJSON() string {
	return r.raw
}

type SiteListResponseSite struct {
	// Identifier
	ID string `json:"id"`
	// Magic WAN Connector identifier tag.
	ConnectorID string `json:"connector_id"`
	Description string `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode bool `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location SiteListResponseSitesLocation `json:"location"`
	// The name of the site.
	Name string `json:"name"`
	// Magic WAN Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID string                   `json:"secondary_connector_id"`
	JSON                 siteListResponseSiteJSON `json:"-"`
}

// siteListResponseSiteJSON contains the JSON metadata for the struct
// [SiteListResponseSite]
type siteListResponseSiteJSON struct {
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

func (r *SiteListResponseSite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteListResponseSiteJSON) RawJSON() string {
	return r.raw
}

// Location of site in latitude and longitude.
type SiteListResponseSitesLocation struct {
	// Latitude
	Lat string `json:"lat"`
	// Longitude
	Lon  string                            `json:"lon"`
	JSON siteListResponseSitesLocationJSON `json:"-"`
}

// siteListResponseSitesLocationJSON contains the JSON metadata for the struct
// [SiteListResponseSitesLocation]
type siteListResponseSitesLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteListResponseSitesLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteListResponseSitesLocationJSON) RawJSON() string {
	return r.raw
}

type SiteDeleteResponse struct {
	Deleted     bool                          `json:"deleted"`
	DeletedSite SiteDeleteResponseDeletedSite `json:"deleted_site"`
	JSON        siteDeleteResponseJSON        `json:"-"`
}

// siteDeleteResponseJSON contains the JSON metadata for the struct
// [SiteDeleteResponse]
type siteDeleteResponseJSON struct {
	Deleted     apijson.Field
	DeletedSite apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SiteDeleteResponseDeletedSite struct {
	// Identifier
	ID string `json:"id"`
	// Magic WAN Connector identifier tag.
	ConnectorID string `json:"connector_id"`
	Description string `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode bool `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location SiteDeleteResponseDeletedSiteLocation `json:"location"`
	// The name of the site.
	Name string `json:"name"`
	// Magic WAN Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID string                            `json:"secondary_connector_id"`
	JSON                 siteDeleteResponseDeletedSiteJSON `json:"-"`
}

// siteDeleteResponseDeletedSiteJSON contains the JSON metadata for the struct
// [SiteDeleteResponseDeletedSite]
type siteDeleteResponseDeletedSiteJSON struct {
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

func (r *SiteDeleteResponseDeletedSite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteDeleteResponseDeletedSiteJSON) RawJSON() string {
	return r.raw
}

// Location of site in latitude and longitude.
type SiteDeleteResponseDeletedSiteLocation struct {
	// Latitude
	Lat string `json:"lat"`
	// Longitude
	Lon  string                                    `json:"lon"`
	JSON siteDeleteResponseDeletedSiteLocationJSON `json:"-"`
}

// siteDeleteResponseDeletedSiteLocationJSON contains the JSON metadata for the
// struct [SiteDeleteResponseDeletedSiteLocation]
type siteDeleteResponseDeletedSiteLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteDeleteResponseDeletedSiteLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteDeleteResponseDeletedSiteLocationJSON) RawJSON() string {
	return r.raw
}

type SiteGetResponse struct {
	Site SiteGetResponseSite `json:"site"`
	JSON siteGetResponseJSON `json:"-"`
}

// siteGetResponseJSON contains the JSON metadata for the struct [SiteGetResponse]
type siteGetResponseJSON struct {
	Site        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteGetResponseJSON) RawJSON() string {
	return r.raw
}

type SiteGetResponseSite struct {
	// Identifier
	ID string `json:"id"`
	// Magic WAN Connector identifier tag.
	ConnectorID string `json:"connector_id"`
	Description string `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode bool `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location SiteGetResponseSiteLocation `json:"location"`
	// The name of the site.
	Name string `json:"name"`
	// Magic WAN Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID string                  `json:"secondary_connector_id"`
	JSON                 siteGetResponseSiteJSON `json:"-"`
}

// siteGetResponseSiteJSON contains the JSON metadata for the struct
// [SiteGetResponseSite]
type siteGetResponseSiteJSON struct {
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

func (r *SiteGetResponseSite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteGetResponseSiteJSON) RawJSON() string {
	return r.raw
}

// Location of site in latitude and longitude.
type SiteGetResponseSiteLocation struct {
	// Latitude
	Lat string `json:"lat"`
	// Longitude
	Lon  string                          `json:"lon"`
	JSON siteGetResponseSiteLocationJSON `json:"-"`
}

// siteGetResponseSiteLocationJSON contains the JSON metadata for the struct
// [SiteGetResponseSiteLocation]
type siteGetResponseSiteLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteGetResponseSiteLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteGetResponseSiteLocationJSON) RawJSON() string {
	return r.raw
}

type SiteNewParams struct {
	// Identifier
	AccountID param.Field[string]            `path:"account_id,required"`
	Site      param.Field[SiteNewParamsSite] `json:"site"`
}

func (r SiteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteNewParamsSite struct {
	// The name of the site.
	Name param.Field[string] `json:"name,required"`
	// Magic WAN Connector identifier tag.
	ConnectorID param.Field[string] `json:"connector_id"`
	Description param.Field[string] `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode param.Field[bool] `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location param.Field[SiteNewParamsSiteLocation] `json:"location"`
	// Magic WAN Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID param.Field[string] `json:"secondary_connector_id"`
}

func (r SiteNewParamsSite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Location of site in latitude and longitude.
type SiteNewParamsSiteLocation struct {
	// Latitude
	Lat param.Field[string] `json:"lat"`
	// Longitude
	Lon param.Field[string] `json:"lon"`
}

func (r SiteNewParamsSiteLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteNewResponse       `json:"result,required"`
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
	AccountID param.Field[string]               `path:"account_id,required"`
	Site      param.Field[SiteUpdateParamsSite] `json:"site"`
}

func (r SiteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteUpdateParamsSite struct {
	// Magic WAN Connector identifier tag.
	ConnectorID param.Field[string] `json:"connector_id"`
	Description param.Field[string] `json:"description"`
	// Location of site in latitude and longitude.
	Location param.Field[SiteUpdateParamsSiteLocation] `json:"location"`
	// The name of the site.
	Name param.Field[string] `json:"name"`
	// Magic WAN Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID param.Field[string] `json:"secondary_connector_id"`
}

func (r SiteUpdateParamsSite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Location of site in latitude and longitude.
type SiteUpdateParamsSiteLocation struct {
	// Latitude
	Lat param.Field[string] `json:"lat"`
	// Longitude
	Lon param.Field[string] `json:"lon"`
}

func (r SiteUpdateParamsSiteLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteUpdateResponse    `json:"result,required"`
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

type SiteListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteListResponse      `json:"result,required"`
	// Whether the API call was successful
	Success SiteListResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteListResponseEnvelopeJSON    `json:"-"`
}

// siteListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteListResponseEnvelope]
type siteListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteListResponseEnvelopeSuccess bool

const (
	SiteListResponseEnvelopeSuccessTrue SiteListResponseEnvelopeSuccess = true
)

func (r SiteListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r SiteDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SiteDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteDeleteResponse    `json:"result,required"`
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
	Result   SiteGetResponse       `json:"result,required"`
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
