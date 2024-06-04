// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// RiskScoringIntegrationReferenceService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRiskScoringIntegrationReferenceService] method instead.
type RiskScoringIntegrationReferenceService struct {
	Options []option.RequestOption
}

// NewRiskScoringIntegrationReferenceService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRiskScoringIntegrationReferenceService(opts ...option.RequestOption) (r *RiskScoringIntegrationReferenceService) {
	r = &RiskScoringIntegrationReferenceService{}
	r.Options = opts
	return
}

// Get risk score integration by reference id.
func (r *RiskScoringIntegrationReferenceService) Get(ctx context.Context, referenceID string, query RiskScoringIntegrationReferenceGetParams, opts ...option.RequestOption) (res *RiskScoringIntegrationReferenceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RiskScoringIntegrationReferenceGetResponseEnvelope
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if referenceID == "" {
		err = errors.New("missing required reference_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations/reference_id/%s", query.AccountID, referenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RiskScoringIntegrationReferenceGetResponse struct {
	// The id of the integration, a UUIDv4.
	ID string `json:"id,required" format:"uuid"`
	// The Cloudflare account tag.
	AccountTag string `json:"account_tag,required"`
	// Whether this integration is enabled and should export changes in risk score.
	Active bool `json:"active,required"`
	// When the integration was created in RFC3339 format.
	CreatedAt       time.Time                                                 `json:"created_at,required" format:"date-time"`
	IntegrationType RiskScoringIntegrationReferenceGetResponseIntegrationType `json:"integration_type,required"`
	// A reference ID defined by the client. Should be set to the Access-Okta IDP
	// integration ID. Useful when the risk-score integration needs to be associated
	// with a secondary asset and recalled using that ID.
	ReferenceID string `json:"reference_id,required"`
	// The base URL for the tenant. E.g. "https://tenant.okta.com"
	TenantURL string `json:"tenant_url,required"`
	// The URL for the Shared Signals Framework configuration, e.g.
	// "/.well-known/sse-configuration/{integration_uuid}/".
	// https://openid.net/specs/openid-sse-framework-1_0.html#rfc.section.6.2.1
	WellKnownURL string                                         `json:"well_known_url,required"`
	JSON         riskScoringIntegrationReferenceGetResponseJSON `json:"-"`
}

// riskScoringIntegrationReferenceGetResponseJSON contains the JSON metadata for
// the struct [RiskScoringIntegrationReferenceGetResponse]
type riskScoringIntegrationReferenceGetResponseJSON struct {
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

func (r *RiskScoringIntegrationReferenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationReferenceGetResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringIntegrationReferenceGetResponseIntegrationType string

const (
	RiskScoringIntegrationReferenceGetResponseIntegrationTypeOkta RiskScoringIntegrationReferenceGetResponseIntegrationType = "Okta"
)

func (r RiskScoringIntegrationReferenceGetResponseIntegrationType) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationReferenceGetResponseIntegrationTypeOkta:
		return true
	}
	return false
}

type RiskScoringIntegrationReferenceGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RiskScoringIntegrationReferenceGetResponseEnvelope struct {
	Errors     []shared.ResponseInfo                                        `json:"errors,required"`
	Messages   []shared.ResponseInfo                                        `json:"messages,required"`
	Success    bool                                                         `json:"success,required"`
	Result     RiskScoringIntegrationReferenceGetResponse                   `json:"result"`
	ResultInfo RiskScoringIntegrationReferenceGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       riskScoringIntegrationReferenceGetResponseEnvelopeJSON       `json:"-"`
}

// riskScoringIntegrationReferenceGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RiskScoringIntegrationReferenceGetResponseEnvelope]
type riskScoringIntegrationReferenceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringIntegrationReferenceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationReferenceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RiskScoringIntegrationReferenceGetResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                            `json:"total_count,required"`
	JSON       riskScoringIntegrationReferenceGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// riskScoringIntegrationReferenceGetResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [RiskScoringIntegrationReferenceGetResponseEnvelopeResultInfo]
type riskScoringIntegrationReferenceGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringIntegrationReferenceGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationReferenceGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
