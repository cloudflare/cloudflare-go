// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tenant_custom_nameservers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// TenantCustomNameserverService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantCustomNameserverService] method instead.
type TenantCustomNameserverService struct {
	Options []option.RequestOption
}

// NewTenantCustomNameserverService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTenantCustomNameserverService(opts ...option.RequestOption) (r *TenantCustomNameserverService) {
	r = &TenantCustomNameserverService{}
	r.Options = opts
	return
}

// Add Tenant Custom Nameserver
func (r *TenantCustomNameserverService) New(ctx context.Context, tenantTag string, body TenantCustomNameserverNewParams, opts ...option.RequestOption) (res *TenantCustomNameserverNewResponse, err error) {
	var env TenantCustomNameserverNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if tenantTag == "" {
		err = errors.New("missing required tenant_tag parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s/custom_ns", tenantTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Delete Tenant Custom Nameserver
func (r *TenantCustomNameserverService) Delete(ctx context.Context, tenantTag string, customNSID string, opts ...option.RequestOption) (res *pagination.SinglePage[string], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if tenantTag == "" {
		err = errors.New("missing required tenant_tag parameter")
		return nil, err
	}
	if customNSID == "" {
		err = errors.New("missing required custom_ns_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s/custom_ns/%s", tenantTag, customNSID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Delete Tenant Custom Nameserver
func (r *TenantCustomNameserverService) DeleteAutoPaging(ctx context.Context, tenantTag string, customNSID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[string] {
	return pagination.NewSinglePageAutoPager(r.Delete(ctx, tenantTag, customNSID, opts...))
}

// List a tenant's custom nameservers.
func (r *TenantCustomNameserverService) Get(ctx context.Context, tenantTag string, opts ...option.RequestOption) (res *pagination.SinglePage[TenantCustomNameserverGetResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if tenantTag == "" {
		err = errors.New("missing required tenant_tag parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s/custom_ns", tenantTag)
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

// List a tenant's custom nameservers.
func (r *TenantCustomNameserverService) GetAutoPaging(ctx context.Context, tenantTag string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TenantCustomNameserverGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, tenantTag, opts...))
}

// A single tenant custom nameserver.
type TenantCustomNameserverNewResponse struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []TenantCustomNameserverNewResponseDNSRecord `json:"dns_records" api:"required"`
	// The FQDN of the name server.
	NSName string `json:"ns_name" api:"required" format:"hostname"`
	// Verification status of the nameserver.
	//
	// Deprecated: deprecated
	Status TenantCustomNameserverNewResponseStatus `json:"status" api:"required"`
	// Identifier.
	ZoneTag string `json:"zone_tag" api:"required"`
	// The number of the set that this name server belongs to.
	NSSet float64                               `json:"ns_set"`
	JSON  tenantCustomNameserverNewResponseJSON `json:"-"`
}

// tenantCustomNameserverNewResponseJSON contains the JSON metadata for the struct
// [TenantCustomNameserverNewResponse]
type tenantCustomNameserverNewResponseJSON struct {
	DNSRecords  apijson.Field
	NSName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NSSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantCustomNameserverNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantCustomNameserverNewResponseJSON) RawJSON() string {
	return r.raw
}

type TenantCustomNameserverNewResponseDNSRecord struct {
	// DNS record type.
	Type TenantCustomNameserverNewResponseDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                                         `json:"value"`
	JSON  tenantCustomNameserverNewResponseDNSRecordJSON `json:"-"`
}

// tenantCustomNameserverNewResponseDNSRecordJSON contains the JSON metadata for
// the struct [TenantCustomNameserverNewResponseDNSRecord]
type tenantCustomNameserverNewResponseDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantCustomNameserverNewResponseDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantCustomNameserverNewResponseDNSRecordJSON) RawJSON() string {
	return r.raw
}

// DNS record type.
type TenantCustomNameserverNewResponseDNSRecordsType string

const (
	TenantCustomNameserverNewResponseDNSRecordsTypeA    TenantCustomNameserverNewResponseDNSRecordsType = "A"
	TenantCustomNameserverNewResponseDNSRecordsTypeAAAA TenantCustomNameserverNewResponseDNSRecordsType = "AAAA"
)

func (r TenantCustomNameserverNewResponseDNSRecordsType) IsKnown() bool {
	switch r {
	case TenantCustomNameserverNewResponseDNSRecordsTypeA, TenantCustomNameserverNewResponseDNSRecordsTypeAAAA:
		return true
	}
	return false
}

// Verification status of the nameserver.
type TenantCustomNameserverNewResponseStatus string

const (
	TenantCustomNameserverNewResponseStatusMoved    TenantCustomNameserverNewResponseStatus = "moved"
	TenantCustomNameserverNewResponseStatusPending  TenantCustomNameserverNewResponseStatus = "pending"
	TenantCustomNameserverNewResponseStatusVerified TenantCustomNameserverNewResponseStatus = "verified"
)

func (r TenantCustomNameserverNewResponseStatus) IsKnown() bool {
	switch r {
	case TenantCustomNameserverNewResponseStatusMoved, TenantCustomNameserverNewResponseStatusPending, TenantCustomNameserverNewResponseStatusVerified:
		return true
	}
	return false
}

// A single tenant custom nameserver.
type TenantCustomNameserverGetResponse struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []TenantCustomNameserverGetResponseDNSRecord `json:"dns_records" api:"required"`
	// The FQDN of the name server.
	NSName string `json:"ns_name" api:"required" format:"hostname"`
	// Verification status of the nameserver.
	//
	// Deprecated: deprecated
	Status TenantCustomNameserverGetResponseStatus `json:"status" api:"required"`
	// Identifier.
	ZoneTag string `json:"zone_tag" api:"required"`
	// The number of the set that this name server belongs to.
	NSSet float64                               `json:"ns_set"`
	JSON  tenantCustomNameserverGetResponseJSON `json:"-"`
}

// tenantCustomNameserverGetResponseJSON contains the JSON metadata for the struct
// [TenantCustomNameserverGetResponse]
type tenantCustomNameserverGetResponseJSON struct {
	DNSRecords  apijson.Field
	NSName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NSSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantCustomNameserverGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantCustomNameserverGetResponseJSON) RawJSON() string {
	return r.raw
}

type TenantCustomNameserverGetResponseDNSRecord struct {
	// DNS record type.
	Type TenantCustomNameserverGetResponseDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                                         `json:"value"`
	JSON  tenantCustomNameserverGetResponseDNSRecordJSON `json:"-"`
}

// tenantCustomNameserverGetResponseDNSRecordJSON contains the JSON metadata for
// the struct [TenantCustomNameserverGetResponseDNSRecord]
type tenantCustomNameserverGetResponseDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantCustomNameserverGetResponseDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantCustomNameserverGetResponseDNSRecordJSON) RawJSON() string {
	return r.raw
}

// DNS record type.
type TenantCustomNameserverGetResponseDNSRecordsType string

const (
	TenantCustomNameserverGetResponseDNSRecordsTypeA    TenantCustomNameserverGetResponseDNSRecordsType = "A"
	TenantCustomNameserverGetResponseDNSRecordsTypeAAAA TenantCustomNameserverGetResponseDNSRecordsType = "AAAA"
)

func (r TenantCustomNameserverGetResponseDNSRecordsType) IsKnown() bool {
	switch r {
	case TenantCustomNameserverGetResponseDNSRecordsTypeA, TenantCustomNameserverGetResponseDNSRecordsTypeAAAA:
		return true
	}
	return false
}

// Verification status of the nameserver.
type TenantCustomNameserverGetResponseStatus string

const (
	TenantCustomNameserverGetResponseStatusMoved    TenantCustomNameserverGetResponseStatus = "moved"
	TenantCustomNameserverGetResponseStatusPending  TenantCustomNameserverGetResponseStatus = "pending"
	TenantCustomNameserverGetResponseStatusVerified TenantCustomNameserverGetResponseStatus = "verified"
)

func (r TenantCustomNameserverGetResponseStatus) IsKnown() bool {
	switch r {
	case TenantCustomNameserverGetResponseStatusMoved, TenantCustomNameserverGetResponseStatusPending, TenantCustomNameserverGetResponseStatusVerified:
		return true
	}
	return false
}

type TenantCustomNameserverNewParams struct {
	// The FQDN of the name server.
	NSName param.Field[string] `json:"ns_name" api:"required" format:"hostname"`
	// The number of the set that this name server belongs to.
	NSSet param.Field[float64] `json:"ns_set"`
}

func (r TenantCustomNameserverNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantCustomNameserverNewResponseEnvelope struct {
	Errors   []TenantCustomNameserverNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []TenantCustomNameserverNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success TenantCustomNameserverNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// A single tenant custom nameserver.
	Result TenantCustomNameserverNewResponse             `json:"result"`
	JSON   tenantCustomNameserverNewResponseEnvelopeJSON `json:"-"`
}

// tenantCustomNameserverNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [TenantCustomNameserverNewResponseEnvelope]
type tenantCustomNameserverNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantCustomNameserverNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantCustomNameserverNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TenantCustomNameserverNewResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           TenantCustomNameserverNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             tenantCustomNameserverNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// tenantCustomNameserverNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [TenantCustomNameserverNewResponseEnvelopeErrors]
type tenantCustomNameserverNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TenantCustomNameserverNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantCustomNameserverNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TenantCustomNameserverNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    tenantCustomNameserverNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tenantCustomNameserverNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [TenantCustomNameserverNewResponseEnvelopeErrorsSource]
type tenantCustomNameserverNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantCustomNameserverNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantCustomNameserverNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TenantCustomNameserverNewResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           TenantCustomNameserverNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             tenantCustomNameserverNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// tenantCustomNameserverNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [TenantCustomNameserverNewResponseEnvelopeMessages]
type tenantCustomNameserverNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TenantCustomNameserverNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantCustomNameserverNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TenantCustomNameserverNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    tenantCustomNameserverNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tenantCustomNameserverNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [TenantCustomNameserverNewResponseEnvelopeMessagesSource]
type tenantCustomNameserverNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantCustomNameserverNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantCustomNameserverNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TenantCustomNameserverNewResponseEnvelopeSuccess bool

const (
	TenantCustomNameserverNewResponseEnvelopeSuccessTrue TenantCustomNameserverNewResponseEnvelopeSuccess = true
)

func (r TenantCustomNameserverNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TenantCustomNameserverNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
