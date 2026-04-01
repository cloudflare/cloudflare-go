// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
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

// Attaches a domain that routes traffic to a Worker.
func (r *DomainService) Update(ctx context.Context, params DomainUpdateParams, opts ...option.RequestOption) (res *DomainUpdateResponse, err error) {
	var env DomainUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/domains", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all domains for an account.
func (r *DomainService) List(ctx context.Context, params DomainListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DomainListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
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

// Lists all domains for an account.
func (r *DomainService) ListAutoPaging(ctx context.Context, params DomainListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DomainListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Detaches a domain from a Worker. Both the Worker and all of its previews are no
// longer routable using this domain.
func (r *DomainService) Delete(ctx context.Context, domainID string, body DomainDeleteParams, opts ...option.RequestOption) (res *DomainDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/domains/%s", body.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Gets information about a domain.
func (r *DomainService) Get(ctx context.Context, domainID string, query DomainGetParams, opts ...option.RequestOption) (res *DomainGetResponse, err error) {
	var env DomainGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/domains/%s", query.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DomainUpdateResponse struct {
	// Immutable ID of the domain.
	ID string `json:"id" api:"required"`
	// ID of the TLS certificate issued for the domain.
	CERTID string `json:"cert_id" api:"required" format:"uuid"`
	// Worker environment associated with the domain.
	//
	// Deprecated: deprecated
	Environment string `json:"environment" api:"required"`
	// Hostname of the domain. Can be either the zone apex or a subdomain of the zone.
	// Requests to this hostname will be routed to the configured Worker.
	Hostname string `json:"hostname" api:"required"`
	// Name of the Worker associated with the domain. Requests to the configured
	// hostname will be routed to this Worker.
	Service string `json:"service" api:"required"`
	// ID of the zone containing the domain hostname.
	ZoneID string `json:"zone_id" api:"required"`
	// Name of the zone containing the domain hostname.
	ZoneName string                   `json:"zone_name" api:"required"`
	JSON     domainUpdateResponseJSON `json:"-"`
}

// domainUpdateResponseJSON contains the JSON metadata for the struct
// [DomainUpdateResponse]
type domainUpdateResponseJSON struct {
	ID          apijson.Field
	CERTID      apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DomainListResponse struct {
	// Immutable ID of the domain.
	ID string `json:"id" api:"required"`
	// ID of the TLS certificate issued for the domain.
	CERTID string `json:"cert_id" api:"required" format:"uuid"`
	// Worker environment associated with the domain.
	//
	// Deprecated: deprecated
	Environment string `json:"environment" api:"required"`
	// Hostname of the domain. Can be either the zone apex or a subdomain of the zone.
	// Requests to this hostname will be routed to the configured Worker.
	Hostname string `json:"hostname" api:"required"`
	// Name of the Worker associated with the domain. Requests to the configured
	// hostname will be routed to this Worker.
	Service string `json:"service" api:"required"`
	// ID of the zone containing the domain hostname.
	ZoneID string `json:"zone_id" api:"required"`
	// Name of the zone containing the domain hostname.
	ZoneName string                 `json:"zone_name" api:"required"`
	JSON     domainListResponseJSON `json:"-"`
}

// domainListResponseJSON contains the JSON metadata for the struct
// [DomainListResponse]
type domainListResponseJSON struct {
	ID          apijson.Field
	CERTID      apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainListResponseJSON) RawJSON() string {
	return r.raw
}

type DomainDeleteResponse struct {
	Errors   []DomainDeleteResponseError   `json:"errors" api:"required"`
	Messages []DomainDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DomainDeleteResponseSuccess `json:"success" api:"required"`
	JSON    domainDeleteResponseJSON    `json:"-"`
}

// domainDeleteResponseJSON contains the JSON metadata for the struct
// [DomainDeleteResponse]
type domainDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DomainDeleteResponseError struct {
	Code             int64                            `json:"code" api:"required"`
	Message          string                           `json:"message" api:"required"`
	DocumentationURL string                           `json:"documentation_url"`
	Source           DomainDeleteResponseErrorsSource `json:"source"`
	JSON             domainDeleteResponseErrorJSON    `json:"-"`
}

// domainDeleteResponseErrorJSON contains the JSON metadata for the struct
// [DomainDeleteResponseError]
type domainDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type DomainDeleteResponseErrorsSource struct {
	Pointer string                               `json:"pointer"`
	JSON    domainDeleteResponseErrorsSourceJSON `json:"-"`
}

// domainDeleteResponseErrorsSourceJSON contains the JSON metadata for the struct
// [DomainDeleteResponseErrorsSource]
type domainDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DomainDeleteResponseMessage struct {
	Code             int64                              `json:"code" api:"required"`
	Message          string                             `json:"message" api:"required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           DomainDeleteResponseMessagesSource `json:"source"`
	JSON             domainDeleteResponseMessageJSON    `json:"-"`
}

// domainDeleteResponseMessageJSON contains the JSON metadata for the struct
// [DomainDeleteResponseMessage]
type domainDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type DomainDeleteResponseMessagesSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    domainDeleteResponseMessagesSourceJSON `json:"-"`
}

// domainDeleteResponseMessagesSourceJSON contains the JSON metadata for the struct
// [DomainDeleteResponseMessagesSource]
type domainDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DomainDeleteResponseSuccess bool

const (
	DomainDeleteResponseSuccessTrue DomainDeleteResponseSuccess = true
)

func (r DomainDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case DomainDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type DomainGetResponse struct {
	// Immutable ID of the domain.
	ID string `json:"id" api:"required"`
	// ID of the TLS certificate issued for the domain.
	CERTID string `json:"cert_id" api:"required" format:"uuid"`
	// Worker environment associated with the domain.
	//
	// Deprecated: deprecated
	Environment string `json:"environment" api:"required"`
	// Hostname of the domain. Can be either the zone apex or a subdomain of the zone.
	// Requests to this hostname will be routed to the configured Worker.
	Hostname string `json:"hostname" api:"required"`
	// Name of the Worker associated with the domain. Requests to the configured
	// hostname will be routed to this Worker.
	Service string `json:"service" api:"required"`
	// ID of the zone containing the domain hostname.
	ZoneID string `json:"zone_id" api:"required"`
	// Name of the zone containing the domain hostname.
	ZoneName string                `json:"zone_name" api:"required"`
	JSON     domainGetResponseJSON `json:"-"`
}

// domainGetResponseJSON contains the JSON metadata for the struct
// [DomainGetResponse]
type domainGetResponseJSON struct {
	ID          apijson.Field
	CERTID      apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseJSON) RawJSON() string {
	return r.raw
}

type DomainUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Hostname of the domain. Can be either the zone apex or a subdomain of the zone.
	// Requests to this hostname will be routed to the configured Worker.
	Hostname param.Field[string] `json:"hostname" api:"required"`
	// Name of the Worker associated with the domain. Requests to the configured
	// hostname will be routed to this Worker.
	Service param.Field[string] `json:"service" api:"required"`
	// Worker environment associated with the domain.
	Environment param.Field[string] `json:"environment"`
	// ID of the zone containing the domain hostname.
	ZoneID param.Field[string] `json:"zone_id"`
	// Name of the zone containing the domain hostname.
	ZoneName param.Field[string] `json:"zone_name"`
}

func (r DomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DomainUpdateResponseEnvelope struct {
	Errors   []DomainUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DomainUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   DomainUpdateResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success DomainUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    domainUpdateResponseEnvelopeJSON    `json:"-"`
}

// domainUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainUpdateResponseEnvelope]
type domainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DomainUpdateResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           DomainUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             domainUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// domainUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DomainUpdateResponseEnvelopeErrors]
type domainUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DomainUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    domainUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// domainUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DomainUpdateResponseEnvelopeErrorsSource]
type domainUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DomainUpdateResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DomainUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             domainUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// domainUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DomainUpdateResponseEnvelopeMessages]
type domainUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DomainUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    domainUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// domainUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DomainUpdateResponseEnvelopeMessagesSource]
type domainUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Worker environment associated with the domain.
	Environment param.Field[string] `query:"environment"`
	// Hostname of the domain.
	Hostname param.Field[string] `query:"hostname"`
	// Name of the Worker associated with the domain.
	Service param.Field[string] `query:"service"`
	// ID of the zone containing the domain hostname.
	ZoneID param.Field[string] `query:"zone_id"`
	// Name of the zone containing the domain hostname.
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
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DomainGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DomainGetResponseEnvelope struct {
	Errors   []DomainGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DomainGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   DomainGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success DomainGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    domainGetResponseEnvelopeJSON    `json:"-"`
}

// domainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelope]
type domainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code" api:"required"`
	Message          string                                `json:"message" api:"required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           DomainGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             domainGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// domainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelopeErrors]
type domainGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    domainGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// domainGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DomainGetResponseEnvelopeErrorsSource]
type domainGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code" api:"required"`
	Message          string                                  `json:"message" api:"required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           DomainGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             domainGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// domainGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelopeMessages]
type domainGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    domainGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// domainGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DomainGetResponseEnvelopeMessagesSource]
type domainGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
