// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// SubdomainService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubdomainService] method instead.
type SubdomainService struct {
	Options []option.RequestOption
	DNS     *SubdomainDNSService
}

// NewSubdomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubdomainService(opts ...option.RequestOption) (r *SubdomainService) {
	r = &SubdomainService{}
	r.Options = opts
	r.DNS = NewSubdomainDNSService(opts...)
	return
}

// Creates a new sending subdomain or re-enables sending on an existing subdomain
// that had it disabled. If zone-level Email Sending has not been enabled yet, the
// zone flag is automatically set when the entitlement is present.
func (r *SubdomainService) New(ctx context.Context, params SubdomainNewParams, opts ...option.RequestOption) (res *SubdomainNewResponse, err error) {
	var env SubdomainNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/sending/subdomains", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all sending-enabled subdomains for the zone.
func (r *SubdomainService) List(ctx context.Context, query SubdomainListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SubdomainListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/sending/subdomains", query.ZoneID)
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

// Lists all sending-enabled subdomains for the zone.
func (r *SubdomainService) ListAutoPaging(ctx context.Context, query SubdomainListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SubdomainListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Disables sending on a subdomain and removes its DNS records. If routing is still
// active on the subdomain, only sending is disabled.
func (r *SubdomainService) Delete(ctx context.Context, subdomainID string, body SubdomainDeleteParams, opts ...option.RequestOption) (res *SubdomainDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if subdomainID == "" {
		err = errors.New("missing required subdomain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/sending/subdomains/%s", body.ZoneID, subdomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Gets information for a specific sending subdomain.
func (r *SubdomainService) Get(ctx context.Context, subdomainID string, query SubdomainGetParams, opts ...option.RequestOption) (res *SubdomainGetResponse, err error) {
	var env SubdomainGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if subdomainID == "" {
		err = errors.New("missing required subdomain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/sending/subdomains/%s", query.ZoneID, subdomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SubdomainNewResponse struct {
	// Whether Email Sending is enabled on this subdomain.
	Enabled bool `json:"enabled" api:"required"`
	// The subdomain domain name.
	Name string `json:"name" api:"required"`
	// Sending subdomain identifier.
	Tag string `json:"tag" api:"required"`
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The DKIM selector used for email signing.
	DKIMSelector string `json:"dkim_selector"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// The return-path domain used for bounce handling.
	ReturnPathDomain string                   `json:"return_path_domain"`
	JSON             subdomainNewResponseJSON `json:"-"`
}

// subdomainNewResponseJSON contains the JSON metadata for the struct
// [SubdomainNewResponse]
type subdomainNewResponseJSON struct {
	Enabled          apijson.Field
	Name             apijson.Field
	Tag              apijson.Field
	Created          apijson.Field
	DKIMSelector     apijson.Field
	Modified         apijson.Field
	ReturnPathDomain apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainNewResponseJSON) RawJSON() string {
	return r.raw
}

type SubdomainListResponse struct {
	// Whether Email Sending is enabled on this subdomain.
	Enabled bool `json:"enabled" api:"required"`
	// The subdomain domain name.
	Name string `json:"name" api:"required"`
	// Sending subdomain identifier.
	Tag string `json:"tag" api:"required"`
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The DKIM selector used for email signing.
	DKIMSelector string `json:"dkim_selector"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// The return-path domain used for bounce handling.
	ReturnPathDomain string                    `json:"return_path_domain"`
	JSON             subdomainListResponseJSON `json:"-"`
}

// subdomainListResponseJSON contains the JSON metadata for the struct
// [SubdomainListResponse]
type subdomainListResponseJSON struct {
	Enabled          apijson.Field
	Name             apijson.Field
	Tag              apijson.Field
	Created          apijson.Field
	DKIMSelector     apijson.Field
	Modified         apijson.Field
	ReturnPathDomain apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainListResponseJSON) RawJSON() string {
	return r.raw
}

type SubdomainDeleteResponse struct {
	Errors   []SubdomainDeleteResponseError   `json:"errors" api:"required"`
	Messages []SubdomainDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SubdomainDeleteResponseSuccess `json:"success" api:"required"`
	JSON    subdomainDeleteResponseJSON    `json:"-"`
}

// subdomainDeleteResponseJSON contains the JSON metadata for the struct
// [SubdomainDeleteResponse]
type subdomainDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SubdomainDeleteResponseError struct {
	Code             int64                               `json:"code" api:"required"`
	Message          string                              `json:"message" api:"required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           SubdomainDeleteResponseErrorsSource `json:"source"`
	JSON             subdomainDeleteResponseErrorJSON    `json:"-"`
}

// subdomainDeleteResponseErrorJSON contains the JSON metadata for the struct
// [SubdomainDeleteResponseError]
type subdomainDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type SubdomainDeleteResponseErrorsSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    subdomainDeleteResponseErrorsSourceJSON `json:"-"`
}

// subdomainDeleteResponseErrorsSourceJSON contains the JSON metadata for the
// struct [SubdomainDeleteResponseErrorsSource]
type subdomainDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SubdomainDeleteResponseMessage struct {
	Code             int64                                 `json:"code" api:"required"`
	Message          string                                `json:"message" api:"required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           SubdomainDeleteResponseMessagesSource `json:"source"`
	JSON             subdomainDeleteResponseMessageJSON    `json:"-"`
}

// subdomainDeleteResponseMessageJSON contains the JSON metadata for the struct
// [SubdomainDeleteResponseMessage]
type subdomainDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type SubdomainDeleteResponseMessagesSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    subdomainDeleteResponseMessagesSourceJSON `json:"-"`
}

// subdomainDeleteResponseMessagesSourceJSON contains the JSON metadata for the
// struct [SubdomainDeleteResponseMessagesSource]
type subdomainDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SubdomainDeleteResponseSuccess bool

const (
	SubdomainDeleteResponseSuccessTrue SubdomainDeleteResponseSuccess = true
)

func (r SubdomainDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case SubdomainDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type SubdomainGetResponse struct {
	// Whether Email Sending is enabled on this subdomain.
	Enabled bool `json:"enabled" api:"required"`
	// The subdomain domain name.
	Name string `json:"name" api:"required"`
	// Sending subdomain identifier.
	Tag string `json:"tag" api:"required"`
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The DKIM selector used for email signing.
	DKIMSelector string `json:"dkim_selector"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// The return-path domain used for bounce handling.
	ReturnPathDomain string                   `json:"return_path_domain"`
	JSON             subdomainGetResponseJSON `json:"-"`
}

// subdomainGetResponseJSON contains the JSON metadata for the struct
// [SubdomainGetResponse]
type subdomainGetResponseJSON struct {
	Enabled          apijson.Field
	Name             apijson.Field
	Tag              apijson.Field
	Created          apijson.Field
	DKIMSelector     apijson.Field
	Modified         apijson.Field
	ReturnPathDomain apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type SubdomainNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The subdomain name. Must be within the zone.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r SubdomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubdomainNewResponseEnvelope struct {
	Errors   []SubdomainNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SubdomainNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SubdomainNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SubdomainNewResponse                `json:"result"`
	JSON    subdomainNewResponseEnvelopeJSON    `json:"-"`
}

// subdomainNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubdomainNewResponseEnvelope]
type subdomainNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SubdomainNewResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           SubdomainNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             subdomainNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// subdomainNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SubdomainNewResponseEnvelopeErrors]
type subdomainNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SubdomainNewResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    subdomainNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// subdomainNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SubdomainNewResponseEnvelopeErrorsSource]
type subdomainNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SubdomainNewResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           SubdomainNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             subdomainNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// subdomainNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SubdomainNewResponseEnvelopeMessages]
type subdomainNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SubdomainNewResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    subdomainNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// subdomainNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SubdomainNewResponseEnvelopeMessagesSource]
type subdomainNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SubdomainNewResponseEnvelopeSuccess bool

const (
	SubdomainNewResponseEnvelopeSuccessTrue SubdomainNewResponseEnvelopeSuccess = true
)

func (r SubdomainNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubdomainNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubdomainListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type SubdomainDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type SubdomainGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type SubdomainGetResponseEnvelope struct {
	Errors   []SubdomainGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SubdomainGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SubdomainGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SubdomainGetResponse                `json:"result"`
	JSON    subdomainGetResponseEnvelopeJSON    `json:"-"`
}

// subdomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubdomainGetResponseEnvelope]
type subdomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SubdomainGetResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           SubdomainGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             subdomainGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// subdomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SubdomainGetResponseEnvelopeErrors]
type subdomainGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SubdomainGetResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    subdomainGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// subdomainGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SubdomainGetResponseEnvelopeErrorsSource]
type subdomainGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SubdomainGetResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           SubdomainGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             subdomainGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// subdomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SubdomainGetResponseEnvelopeMessages]
type subdomainGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SubdomainGetResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    subdomainGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// subdomainGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SubdomainGetResponseEnvelopeMessagesSource]
type subdomainGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SubdomainGetResponseEnvelopeSuccess bool

const (
	SubdomainGetResponseEnvelopeSuccessTrue SubdomainGetResponseEnvelopeSuccess = true
)

func (r SubdomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubdomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
