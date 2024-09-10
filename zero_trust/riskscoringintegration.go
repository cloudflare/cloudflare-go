// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/pagination"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/shared"
)

// RiskScoringIntegrationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRiskScoringIntegrationService] method instead.
type RiskScoringIntegrationService struct {
	Options    []option.RequestOption
	References *RiskScoringIntegrationReferenceService
}

// NewRiskScoringIntegrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRiskScoringIntegrationService(opts ...option.RequestOption) (r *RiskScoringIntegrationService) {
	r = &RiskScoringIntegrationService{}
	r.Options = opts
	r.References = NewRiskScoringIntegrationReferenceService(opts...)
	return
}

// Create new risk score integration.
func (r *RiskScoringIntegrationService) New(ctx context.Context, params RiskScoringIntegrationNewParams, opts ...option.RequestOption) (res *RiskScoringIntegrationNewResponse, err error) {
	var env RiskScoringIntegrationNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Overwrite the reference_id, tenant_url, and active values with the ones provided
func (r *RiskScoringIntegrationService) Update(ctx context.Context, integrationID string, params RiskScoringIntegrationUpdateParams, opts ...option.RequestOption) (res *RiskScoringIntegrationUpdateResponse, err error) {
	var env RiskScoringIntegrationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations/%s", params.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all risk score integrations for the account.
func (r *RiskScoringIntegrationService) List(ctx context.Context, query RiskScoringIntegrationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RiskScoringIntegrationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// List all risk score integrations for the account.
func (r *RiskScoringIntegrationService) ListAutoPaging(ctx context.Context, query RiskScoringIntegrationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RiskScoringIntegrationListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a risk score integration.
func (r *RiskScoringIntegrationService) Delete(ctx context.Context, integrationID string, body RiskScoringIntegrationDeleteParams, opts ...option.RequestOption) (res *RiskScoringIntegrationDeleteResponse, err error) {
	var env RiskScoringIntegrationDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations/%s", body.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get risk score integration by id.
func (r *RiskScoringIntegrationService) Get(ctx context.Context, integrationID string, query RiskScoringIntegrationGetParams, opts ...option.RequestOption) (res *RiskScoringIntegrationGetResponse, err error) {
	var env RiskScoringIntegrationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations/%s", query.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RiskScoringIntegrationNewResponse struct {
	// The id of the integration, a UUIDv4.
	ID string `json:"id,required" format:"uuid"`
	// The Cloudflare account tag.
	AccountTag string `json:"account_tag,required"`
	// Whether this integration is enabled and should export changes in risk score.
	Active bool `json:"active,required"`
	// When the integration was created in RFC3339 format.
	CreatedAt       time.Time                                        `json:"created_at,required" format:"date-time"`
	IntegrationType RiskScoringIntegrationNewResponseIntegrationType `json:"integration_type,required"`
	// A reference ID defined by the client. Should be set to the Access-Okta IDP
	// integration ID. Useful when the risk-score integration needs to be associated
	// with a secondary asset and recalled using that ID.
	ReferenceID string `json:"reference_id,required"`
	// The base URL for the tenant. E.g. "https://tenant.okta.com"
	TenantURL string `json:"tenant_url,required"`
	// The URL for the Shared Signals Framework configuration, e.g.
	// "/.well-known/sse-configuration/{integration_uuid}/".
	// https://openid.net/specs/openid-sse-framework-1_0.html#rfc.section.6.2.1
	WellKnownURL string                                `json:"well_known_url,required"`
	JSON         riskScoringIntegrationNewResponseJSON `json:"-"`
}

// riskScoringIntegrationNewResponseJSON contains the JSON metadata for the struct
// [RiskScoringIntegrationNewResponse]
type riskScoringIntegrationNewResponseJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Active          apijson.Field
	CreatedAt       apijson.Field
	IntegrationType apijson.Field
	ReferenceID     apijson.Field
	TenantURL       apijson.Field
	WellKnownURL    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RiskScoringIntegrationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationNewResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringIntegrationNewResponseIntegrationType string

const (
	RiskScoringIntegrationNewResponseIntegrationTypeOkta RiskScoringIntegrationNewResponseIntegrationType = "Okta"
)

func (r RiskScoringIntegrationNewResponseIntegrationType) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationNewResponseIntegrationTypeOkta:
		return true
	}
	return false
}

type RiskScoringIntegrationUpdateResponse struct {
	// The id of the integration, a UUIDv4.
	ID string `json:"id,required" format:"uuid"`
	// The Cloudflare account tag.
	AccountTag string `json:"account_tag,required"`
	// Whether this integration is enabled and should export changes in risk score.
	Active bool `json:"active,required"`
	// When the integration was created in RFC3339 format.
	CreatedAt       time.Time                                           `json:"created_at,required" format:"date-time"`
	IntegrationType RiskScoringIntegrationUpdateResponseIntegrationType `json:"integration_type,required"`
	// A reference ID defined by the client. Should be set to the Access-Okta IDP
	// integration ID. Useful when the risk-score integration needs to be associated
	// with a secondary asset and recalled using that ID.
	ReferenceID string `json:"reference_id,required"`
	// The base URL for the tenant. E.g. "https://tenant.okta.com"
	TenantURL string `json:"tenant_url,required"`
	// The URL for the Shared Signals Framework configuration, e.g.
	// "/.well-known/sse-configuration/{integration_uuid}/".
	// https://openid.net/specs/openid-sse-framework-1_0.html#rfc.section.6.2.1
	WellKnownURL string                                   `json:"well_known_url,required"`
	JSON         riskScoringIntegrationUpdateResponseJSON `json:"-"`
}

// riskScoringIntegrationUpdateResponseJSON contains the JSON metadata for the
// struct [RiskScoringIntegrationUpdateResponse]
type riskScoringIntegrationUpdateResponseJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Active          apijson.Field
	CreatedAt       apijson.Field
	IntegrationType apijson.Field
	ReferenceID     apijson.Field
	TenantURL       apijson.Field
	WellKnownURL    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RiskScoringIntegrationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringIntegrationUpdateResponseIntegrationType string

const (
	RiskScoringIntegrationUpdateResponseIntegrationTypeOkta RiskScoringIntegrationUpdateResponseIntegrationType = "Okta"
)

func (r RiskScoringIntegrationUpdateResponseIntegrationType) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationUpdateResponseIntegrationTypeOkta:
		return true
	}
	return false
}

type RiskScoringIntegrationListResponse struct {
	// The id of the integration, a UUIDv4.
	ID string `json:"id,required" format:"uuid"`
	// The Cloudflare account tag.
	AccountTag string `json:"account_tag,required"`
	// Whether this integration is enabled and should export changes in risk score.
	Active bool `json:"active,required"`
	// When the integration was created in RFC3339 format.
	CreatedAt       time.Time                                         `json:"created_at,required" format:"date-time"`
	IntegrationType RiskScoringIntegrationListResponseIntegrationType `json:"integration_type,required"`
	// A reference ID defined by the client. Should be set to the Access-Okta IDP
	// integration ID. Useful when the risk-score integration needs to be associated
	// with a secondary asset and recalled using that ID.
	ReferenceID string `json:"reference_id,required"`
	// The base URL for the tenant. E.g. "https://tenant.okta.com"
	TenantURL string `json:"tenant_url,required"`
	// The URL for the Shared Signals Framework configuration, e.g.
	// "/.well-known/sse-configuration/{integration_uuid}/".
	// https://openid.net/specs/openid-sse-framework-1_0.html#rfc.section.6.2.1
	WellKnownURL string                                 `json:"well_known_url,required"`
	JSON         riskScoringIntegrationListResponseJSON `json:"-"`
}

// riskScoringIntegrationListResponseJSON contains the JSON metadata for the struct
// [RiskScoringIntegrationListResponse]
type riskScoringIntegrationListResponseJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Active          apijson.Field
	CreatedAt       apijson.Field
	IntegrationType apijson.Field
	ReferenceID     apijson.Field
	TenantURL       apijson.Field
	WellKnownURL    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RiskScoringIntegrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationListResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringIntegrationListResponseIntegrationType string

const (
	RiskScoringIntegrationListResponseIntegrationTypeOkta RiskScoringIntegrationListResponseIntegrationType = "Okta"
)

func (r RiskScoringIntegrationListResponseIntegrationType) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationListResponseIntegrationTypeOkta:
		return true
	}
	return false
}

type RiskScoringIntegrationDeleteResponse = interface{}

type RiskScoringIntegrationGetResponse struct {
	// The id of the integration, a UUIDv4.
	ID string `json:"id,required" format:"uuid"`
	// The Cloudflare account tag.
	AccountTag string `json:"account_tag,required"`
	// Whether this integration is enabled and should export changes in risk score.
	Active bool `json:"active,required"`
	// When the integration was created in RFC3339 format.
	CreatedAt       time.Time                                        `json:"created_at,required" format:"date-time"`
	IntegrationType RiskScoringIntegrationGetResponseIntegrationType `json:"integration_type,required"`
	// A reference ID defined by the client. Should be set to the Access-Okta IDP
	// integration ID. Useful when the risk-score integration needs to be associated
	// with a secondary asset and recalled using that ID.
	ReferenceID string `json:"reference_id,required"`
	// The base URL for the tenant. E.g. "https://tenant.okta.com"
	TenantURL string `json:"tenant_url,required"`
	// The URL for the Shared Signals Framework configuration, e.g.
	// "/.well-known/sse-configuration/{integration_uuid}/".
	// https://openid.net/specs/openid-sse-framework-1_0.html#rfc.section.6.2.1
	WellKnownURL string                                `json:"well_known_url,required"`
	JSON         riskScoringIntegrationGetResponseJSON `json:"-"`
}

// riskScoringIntegrationGetResponseJSON contains the JSON metadata for the struct
// [RiskScoringIntegrationGetResponse]
type riskScoringIntegrationGetResponseJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Active          apijson.Field
	CreatedAt       apijson.Field
	IntegrationType apijson.Field
	ReferenceID     apijson.Field
	TenantURL       apijson.Field
	WellKnownURL    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RiskScoringIntegrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationGetResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringIntegrationGetResponseIntegrationType string

const (
	RiskScoringIntegrationGetResponseIntegrationTypeOkta RiskScoringIntegrationGetResponseIntegrationType = "Okta"
)

func (r RiskScoringIntegrationGetResponseIntegrationType) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationGetResponseIntegrationTypeOkta:
		return true
	}
	return false
}

type RiskScoringIntegrationNewParams struct {
	AccountID       param.Field[string]                                         `path:"account_id,required"`
	IntegrationType param.Field[RiskScoringIntegrationNewParamsIntegrationType] `json:"integration_type,required"`
	// The base url of the tenant, e.g. "https://tenant.okta.com"
	TenantURL param.Field[string] `json:"tenant_url,required" format:"uri"`
	// A reference id that can be supplied by the client. Currently this should be set
	// to the Access-Okta IDP ID (a UUIDv4).
	// https://developers.cloudflare.com/api/operations/access-identity-providers-get-an-access-identity-provider
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r RiskScoringIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RiskScoringIntegrationNewParamsIntegrationType string

const (
	RiskScoringIntegrationNewParamsIntegrationTypeOkta RiskScoringIntegrationNewParamsIntegrationType = "Okta"
)

func (r RiskScoringIntegrationNewParamsIntegrationType) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationNewParamsIntegrationTypeOkta:
		return true
	}
	return false
}

type RiskScoringIntegrationNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RiskScoringIntegrationNewResponseEnvelopeSuccess `json:"success,required"`
	Result  RiskScoringIntegrationNewResponse                `json:"result"`
	JSON    riskScoringIntegrationNewResponseEnvelopeJSON    `json:"-"`
}

// riskScoringIntegrationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [RiskScoringIntegrationNewResponseEnvelope]
type riskScoringIntegrationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringIntegrationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RiskScoringIntegrationNewResponseEnvelopeSuccess bool

const (
	RiskScoringIntegrationNewResponseEnvelopeSuccessTrue RiskScoringIntegrationNewResponseEnvelopeSuccess = true
)

func (r RiskScoringIntegrationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RiskScoringIntegrationUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether this integration is enabled. If disabled, no risk changes will be
	// exported to the third-party.
	Active param.Field[bool] `json:"active,required"`
	// The base url of the tenant, e.g. "https://tenant.okta.com"
	TenantURL param.Field[string] `json:"tenant_url,required" format:"uri"`
	// A reference id that can be supplied by the client. Currently this should be set
	// to the Access-Okta IDP ID (a UUIDv4).
	// https://developers.cloudflare.com/api/operations/access-identity-providers-get-an-access-identity-provider
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r RiskScoringIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RiskScoringIntegrationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RiskScoringIntegrationUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  RiskScoringIntegrationUpdateResponse                `json:"result"`
	JSON    riskScoringIntegrationUpdateResponseEnvelopeJSON    `json:"-"`
}

// riskScoringIntegrationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [RiskScoringIntegrationUpdateResponseEnvelope]
type riskScoringIntegrationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringIntegrationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RiskScoringIntegrationUpdateResponseEnvelopeSuccess bool

const (
	RiskScoringIntegrationUpdateResponseEnvelopeSuccessTrue RiskScoringIntegrationUpdateResponseEnvelopeSuccess = true
)

func (r RiskScoringIntegrationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RiskScoringIntegrationListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RiskScoringIntegrationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RiskScoringIntegrationDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RiskScoringIntegrationDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  RiskScoringIntegrationDeleteResponse                `json:"result,nullable"`
	JSON    riskScoringIntegrationDeleteResponseEnvelopeJSON    `json:"-"`
}

// riskScoringIntegrationDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [RiskScoringIntegrationDeleteResponseEnvelope]
type riskScoringIntegrationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringIntegrationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RiskScoringIntegrationDeleteResponseEnvelopeSuccess bool

const (
	RiskScoringIntegrationDeleteResponseEnvelopeSuccessTrue RiskScoringIntegrationDeleteResponseEnvelopeSuccess = true
)

func (r RiskScoringIntegrationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RiskScoringIntegrationGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RiskScoringIntegrationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RiskScoringIntegrationGetResponseEnvelopeSuccess `json:"success,required"`
	Result  RiskScoringIntegrationGetResponse                `json:"result"`
	JSON    riskScoringIntegrationGetResponseEnvelopeJSON    `json:"-"`
}

// riskScoringIntegrationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RiskScoringIntegrationGetResponseEnvelope]
type riskScoringIntegrationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringIntegrationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RiskScoringIntegrationGetResponseEnvelopeSuccess bool

const (
	RiskScoringIntegrationGetResponseEnvelopeSuccessTrue RiskScoringIntegrationGetResponseEnvelopeSuccess = true
)

func (r RiskScoringIntegrationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
