// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *DNSSettingService) Edit(ctx context.Context, params DNSSettingEditParams, opts ...option.RequestOption) (res *DNSSetting, err error) {
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
func (r *DNSSettingService) Get(ctx context.Context, query DNSSettingGetParams, opts ...option.RequestOption) (res *DNSSetting, err error) {
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

type DNSSetting struct {
	// Settings determining the nameservers through which the zone should be available.
	Nameservers Nameserver     `json:"nameservers"`
	JSON        dnsSettingJSON `json:"-"`
}

// dnsSettingJSON contains the JSON metadata for the struct [DNSSetting]
type dnsSettingJSON struct {
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingJSON) RawJSON() string {
	return r.raw
}

// Settings determining the nameservers through which the zone should be available.
type Nameserver struct {
	// Nameserver type
	Type NameserverType `json:"type,required"`
	JSON nameserverJSON `json:"-"`
}

// nameserverJSON contains the JSON metadata for the struct [Nameserver]
type nameserverJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Nameserver) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nameserverJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type NameserverType string

const (
	NameserverTypeCloudflareStandard      NameserverType = "cloudflare.standard"
	NameserverTypeCloudflareFoundationDNS NameserverType = "cloudflare.foundation_dns"
)

func (r NameserverType) IsKnown() bool {
	switch r {
	case NameserverTypeCloudflareStandard, NameserverTypeCloudflareFoundationDNS:
		return true
	}
	return false
}

// Settings determining the nameservers through which the zone should be available.
type NameserverParam struct {
	// Nameserver type
	Type param.Field[NameserverType] `json:"type,required"`
}

func (r NameserverParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSSettingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers param.Field[NameserverParam] `json:"nameservers"`
}

func (r DNSSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSSettingEditResponseEnvelope struct {
	Errors      []shared.ResponseInfo `json:"errors,required"`
	Messages    []shared.ResponseInfo `json:"messages,required"`
	Nameservers interface{}           `json:"nameservers,required"`
	Result      DNSSetting            `json:"result,required"`
	// Whether the API call was successful
	Success DNSSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsSettingEditResponseEnvelopeJSON    `json:"-"`
}

// dnsSettingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSettingEditResponseEnvelope]
type dnsSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Nameservers apijson.Field
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
	Errors      []shared.ResponseInfo `json:"errors,required"`
	Messages    []shared.ResponseInfo `json:"messages,required"`
	Nameservers interface{}           `json:"nameservers,required"`
	Result      DNSSetting            `json:"result,required"`
	// Whether the API call was successful
	Success DNSSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsSettingGetResponseEnvelopeJSON    `json:"-"`
}

// dnsSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSettingGetResponseEnvelope]
type dnsSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Nameservers apijson.Field
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
