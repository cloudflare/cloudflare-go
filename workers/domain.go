// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// DomainService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	Options []option.RequestOption
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r *DomainService) {
	r = &DomainService{}
	r.Options = opts
	return
}

// Attaches a Worker to a zone and hostname.
func (r *DomainService) Update(ctx context.Context, params DomainUpdateParams, opts ...option.RequestOption) (res *Domain, err error) {
	var env DomainUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Worker Domains for an account.
func (r *DomainService) List(ctx context.Context, params DomainListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Domain], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains", params.AccountID)
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

// Lists all Worker Domains for an account.
func (r *DomainService) ListAutoPaging(ctx context.Context, params DomainListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Domain] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Detaches a Worker from a zone and hostname.
func (r *DomainService) Delete(ctx context.Context, domainID string, body DomainDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains/%s", body.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Gets a Worker domain.
func (r *DomainService) Get(ctx context.Context, domainID string, query DomainGetParams, opts ...option.RequestOption) (res *Domain, err error) {
	var env DomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains/%s", query.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Domain struct {
	// Identifer of the Worker Domain.
	ID string `json:"id"`
	// Worker environment associated with the zone and hostname.
	Environment string `json:"environment"`
	// Hostname of the Worker Domain.
	Hostname string `json:"hostname"`
	// Worker service associated with the zone and hostname.
	Service string `json:"service"`
	// Identifier of the zone.
	ZoneID string `json:"zone_id"`
	// Name of the zone.
	ZoneName string     `json:"zone_name"`
	JSON     domainJSON `json:"-"`
}

// domainJSON contains the JSON metadata for the struct [Domain]
type domainJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Domain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainJSON) RawJSON() string {
	return r.raw
}

type DomainUpdateParams struct {
	// Identifer of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Worker environment associated with the zone and hostname.
	Environment param.Field[string] `json:"environment,required"`
	// Hostname of the Worker Domain.
	Hostname param.Field[string] `json:"hostname,required"`
	// Worker service associated with the zone and hostname.
	Service param.Field[string] `json:"service,required"`
	// Identifier of the zone.
	ZoneID param.Field[string] `json:"zone_id,required"`
}

func (r DomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DomainUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DomainUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Domain                              `json:"result"`
	JSON    domainUpdateResponseEnvelopeJSON    `json:"-"`
}

// domainUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainUpdateResponseEnvelope]
type domainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainUpdateResponseEnvelopeSuccess bool

const (
	DomainUpdateResponseEnvelopeSuccessTrue DomainUpdateResponseEnvelopeSuccess = true
)

func (r DomainUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DomainListParams struct {
	// Identifer of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Worker environment associated with the zone and hostname.
	Environment param.Field[string] `query:"environment"`
	// Hostname of the Worker Domain.
	Hostname param.Field[string] `query:"hostname"`
	// Worker service associated with the zone and hostname.
	Service param.Field[string] `query:"service"`
	// Identifier of the zone.
	ZoneID param.Field[string] `query:"zone_id"`
	// Name of the zone.
	ZoneName param.Field[string] `query:"zone_name"`
}

// URLQuery serializes [DomainListParams]'s query parameters as `url.Values`.
func (r DomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DomainDeleteParams struct {
	// Identifer of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DomainGetParams struct {
	// Identifer of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DomainGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Domain                           `json:"result"`
	JSON    domainGetResponseEnvelopeJSON    `json:"-"`
}

// domainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelope]
type domainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainGetResponseEnvelopeSuccess bool

const (
	DomainGetResponseEnvelopeSuccessTrue DomainGetResponseEnvelopeSuccess = true
)

func (r DomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
