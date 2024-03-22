// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DNSSettingService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDNSSettingService] method instead.
type DNSSettingService struct {
	Options []option.RequestOption
}

// NewDNSSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSSettingService(opts ...option.RequestOption) (r *DNSSettingService) {
	r = &DNSSettingService{}
	r.Options = opts
	return
}

// Update DNS settings for a zone
func (r *DNSSettingService) Edit(ctx context.Context, params DNSSettingEditParams, opts ...option.RequestOption) (res *DNSSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSSettingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show DNS settings for a zone
func (r *DNSSettingService) Get(ctx context.Context, query DNSSettingGetParams, opts ...option.RequestOption) (res *DNSSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSSettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSSettingEditResponse struct {
	// Settings determining the nameservers through which the zone should be available.
	Nameservers DNSSettingEditResponseNameservers `json:"nameservers"`
	JSON        dnsSettingEditResponseJSON        `json:"-"`
}

// dnsSettingEditResponseJSON contains the JSON metadata for the struct
// [DNSSettingEditResponse]
type dnsSettingEditResponseJSON struct {
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// Settings determining the nameservers through which the zone should be available.
type DNSSettingEditResponseNameservers struct {
	// Nameserver type
	Type DNSSettingEditResponseNameserversType `json:"type,required"`
	JSON dnsSettingEditResponseNameserversJSON `json:"-"`
}

// dnsSettingEditResponseNameserversJSON contains the JSON metadata for the struct
// [DNSSettingEditResponseNameservers]
type dnsSettingEditResponseNameserversJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingEditResponseNameservers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingEditResponseNameserversJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type DNSSettingEditResponseNameserversType string

const (
	DNSSettingEditResponseNameserversTypeCloudflareStandard      DNSSettingEditResponseNameserversType = "cloudflare.standard"
	DNSSettingEditResponseNameserversTypeCloudflareFoundationDNS DNSSettingEditResponseNameserversType = "cloudflare.foundation_dns"
)

func (r DNSSettingEditResponseNameserversType) IsKnown() bool {
	switch r {
	case DNSSettingEditResponseNameserversTypeCloudflareStandard, DNSSettingEditResponseNameserversTypeCloudflareFoundationDNS:
		return true
	}
	return false
}

type DNSSettingGetResponse struct {
	// Settings determining the nameservers through which the zone should be available.
	Nameservers DNSSettingGetResponseNameservers `json:"nameservers"`
	JSON        dnsSettingGetResponseJSON        `json:"-"`
}

// dnsSettingGetResponseJSON contains the JSON metadata for the struct
// [DNSSettingGetResponse]
type dnsSettingGetResponseJSON struct {
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// Settings determining the nameservers through which the zone should be available.
type DNSSettingGetResponseNameservers struct {
	// Nameserver type
	Type DNSSettingGetResponseNameserversType `json:"type,required"`
	JSON dnsSettingGetResponseNameserversJSON `json:"-"`
}

// dnsSettingGetResponseNameserversJSON contains the JSON metadata for the struct
// [DNSSettingGetResponseNameservers]
type dnsSettingGetResponseNameserversJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingGetResponseNameservers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingGetResponseNameserversJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type DNSSettingGetResponseNameserversType string

const (
	DNSSettingGetResponseNameserversTypeCloudflareStandard      DNSSettingGetResponseNameserversType = "cloudflare.standard"
	DNSSettingGetResponseNameserversTypeCloudflareFoundationDNS DNSSettingGetResponseNameserversType = "cloudflare.foundation_dns"
)

func (r DNSSettingGetResponseNameserversType) IsKnown() bool {
	switch r {
	case DNSSettingGetResponseNameserversTypeCloudflareStandard, DNSSettingGetResponseNameserversTypeCloudflareFoundationDNS:
		return true
	}
	return false
}

type DNSSettingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers param.Field[DNSSettingEditParamsNameservers] `json:"nameservers"`
}

func (r DNSSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings determining the nameservers through which the zone should be available.
type DNSSettingEditParamsNameservers struct {
	// Nameserver type
	Type param.Field[DNSSettingEditParamsNameserversType] `json:"type,required"`
}

func (r DNSSettingEditParamsNameservers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Nameserver type
type DNSSettingEditParamsNameserversType string

const (
	DNSSettingEditParamsNameserversTypeCloudflareStandard      DNSSettingEditParamsNameserversType = "cloudflare.standard"
	DNSSettingEditParamsNameserversTypeCloudflareFoundationDNS DNSSettingEditParamsNameserversType = "cloudflare.foundation_dns"
)

func (r DNSSettingEditParamsNameserversType) IsKnown() bool {
	switch r {
	case DNSSettingEditParamsNameserversTypeCloudflareStandard, DNSSettingEditParamsNameserversTypeCloudflareFoundationDNS:
		return true
	}
	return false
}

type DNSSettingEditResponseEnvelope struct {
	Errors   []DNSSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSSettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSSettingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsSettingEditResponseEnvelopeJSON    `json:"-"`
}

// dnsSettingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSettingEditResponseEnvelope]
type dnsSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSSettingEditResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dnsSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsSettingEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSSettingEditResponseEnvelopeErrors]
type dnsSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSSettingEditResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dnsSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsSettingEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSSettingEditResponseEnvelopeMessages]
type dnsSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSSettingEditResponseEnvelopeSuccess bool

const (
	DNSSettingEditResponseEnvelopeSuccessTrue DNSSettingEditResponseEnvelopeSuccess = true
)

func (r DNSSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSSettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSSettingGetResponseEnvelope struct {
	Errors   []DNSSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsSettingGetResponseEnvelopeJSON    `json:"-"`
}

// dnsSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSettingGetResponseEnvelope]
type dnsSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSSettingGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dnsSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSSettingGetResponseEnvelopeErrors]
type dnsSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSSettingGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dnsSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSSettingGetResponseEnvelopeMessages]
type dnsSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSSettingGetResponseEnvelopeSuccess bool

const (
	DNSSettingGetResponseEnvelopeSuccessTrue DNSSettingGetResponseEnvelopeSuccess = true
)

func (r DNSSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
