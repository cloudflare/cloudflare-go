// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// FirewallReverseDNSService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFirewallReverseDNSService] method instead.
type FirewallReverseDNSService struct {
	Options []option.RequestOption
}

// NewFirewallReverseDNSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallReverseDNSService(opts ...option.RequestOption) (r *FirewallReverseDNSService) {
	r = &FirewallReverseDNSService{}
	r.Options = opts
	return
}

// Update reverse DNS configuration (PTR records) for a DNS Firewall cluster
func (r *FirewallReverseDNSService) Edit(ctx context.Context, dnsFirewallID string, params FirewallReverseDNSEditParams, opts ...option.RequestOption) (res *FirewallReverseDNSEditResponse, err error) {
	var env FirewallReverseDNSEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/reverse_dns", params.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show reverse DNS configuration (PTR records) for a DNS Firewall cluster
func (r *FirewallReverseDNSService) Get(ctx context.Context, dnsFirewallID string, query FirewallReverseDNSGetParams, opts ...option.RequestOption) (res *FirewallReverseDNSGetResponse, err error) {
	var env FirewallReverseDNSGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/reverse_dns", query.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallReverseDNSEditResponse struct {
	// Map of cluster IP addresses to PTR record contents
	PTR  map[string]string                  `json:"ptr,required"`
	JSON firewallReverseDNSEditResponseJSON `json:"-"`
}

// firewallReverseDNSEditResponseJSON contains the JSON metadata for the struct
// [FirewallReverseDNSEditResponse]
type firewallReverseDNSEditResponseJSON struct {
	PTR         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallReverseDNSEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallReverseDNSEditResponseJSON) RawJSON() string {
	return r.raw
}

type FirewallReverseDNSGetResponse struct {
	// Map of cluster IP addresses to PTR record contents
	PTR  map[string]string                 `json:"ptr,required"`
	JSON firewallReverseDNSGetResponseJSON `json:"-"`
}

// firewallReverseDNSGetResponseJSON contains the JSON metadata for the struct
// [FirewallReverseDNSGetResponse]
type firewallReverseDNSGetResponseJSON struct {
	PTR         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallReverseDNSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallReverseDNSGetResponseJSON) RawJSON() string {
	return r.raw
}

type FirewallReverseDNSEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Map of cluster IP addresses to PTR record contents
	PTR param.Field[map[string]string] `json:"ptr"`
}

func (r FirewallReverseDNSEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallReverseDNSEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success FirewallReverseDNSEditResponseEnvelopeSuccess `json:"success,required"`
	Result  FirewallReverseDNSEditResponse                `json:"result"`
	JSON    firewallReverseDNSEditResponseEnvelopeJSON    `json:"-"`
}

// firewallReverseDNSEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallReverseDNSEditResponseEnvelope]
type firewallReverseDNSEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallReverseDNSEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallReverseDNSEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallReverseDNSEditResponseEnvelopeSuccess bool

const (
	FirewallReverseDNSEditResponseEnvelopeSuccessTrue FirewallReverseDNSEditResponseEnvelopeSuccess = true
)

func (r FirewallReverseDNSEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallReverseDNSEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FirewallReverseDNSGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type FirewallReverseDNSGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success FirewallReverseDNSGetResponseEnvelopeSuccess `json:"success,required"`
	Result  FirewallReverseDNSGetResponse                `json:"result"`
	JSON    firewallReverseDNSGetResponseEnvelopeJSON    `json:"-"`
}

// firewallReverseDNSGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallReverseDNSGetResponseEnvelope]
type firewallReverseDNSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallReverseDNSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallReverseDNSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallReverseDNSGetResponseEnvelopeSuccess bool

const (
	FirewallReverseDNSGetResponseEnvelopeSuccessTrue FirewallReverseDNSGetResponseEnvelopeSuccess = true
)

func (r FirewallReverseDNSGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallReverseDNSGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
