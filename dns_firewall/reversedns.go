// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ReverseDNSService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReverseDNSService] method instead.
type ReverseDNSService struct {
	Options []option.RequestOption
}

// NewReverseDNSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReverseDNSService(opts ...option.RequestOption) (r *ReverseDNSService) {
	r = &ReverseDNSService{}
	r.Options = opts
	return
}

// Update reverse DNS configuration (PTR records) for a DNS Firewall cluster
func (r *ReverseDNSService) Edit(ctx context.Context, dnsFirewallID string, params ReverseDNSEditParams, opts ...option.RequestOption) (res *ReverseDNSEditResponse, err error) {
	var env ReverseDNSEditResponseEnvelope
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
func (r *ReverseDNSService) Get(ctx context.Context, dnsFirewallID string, query ReverseDNSGetParams, opts ...option.RequestOption) (res *ReverseDNSGetResponse, err error) {
	var env ReverseDNSGetResponseEnvelope
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

type ReverseDNSEditResponse struct {
	// Map of cluster IP addresses to PTR record contents
	PTR  map[string]string          `json:"ptr,required"`
	JSON reverseDNSEditResponseJSON `json:"-"`
}

// reverseDNSEditResponseJSON contains the JSON metadata for the struct
// [ReverseDNSEditResponse]
type reverseDNSEditResponseJSON struct {
	PTR         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSEditResponseJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSGetResponse struct {
	// Map of cluster IP addresses to PTR record contents
	PTR  map[string]string         `json:"ptr,required"`
	JSON reverseDNSGetResponseJSON `json:"-"`
}

// reverseDNSGetResponseJSON contains the JSON metadata for the struct
// [ReverseDNSGetResponse]
type reverseDNSGetResponseJSON struct {
	PTR         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSGetResponseJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Map of cluster IP addresses to PTR record contents
	PTR param.Field[map[string]string] `json:"ptr"`
}

func (r ReverseDNSEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReverseDNSEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ReverseDNSEditResponseEnvelopeSuccess `json:"success,required"`
	Result  ReverseDNSEditResponse                `json:"result"`
	JSON    reverseDNSEditResponseEnvelopeJSON    `json:"-"`
}

// reverseDNSEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ReverseDNSEditResponseEnvelope]
type reverseDNSEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ReverseDNSEditResponseEnvelopeSuccess bool

const (
	ReverseDNSEditResponseEnvelopeSuccessTrue ReverseDNSEditResponseEnvelopeSuccess = true
)

func (r ReverseDNSEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ReverseDNSEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ReverseDNSGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ReverseDNSGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ReverseDNSGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ReverseDNSGetResponse                `json:"result"`
	JSON    reverseDNSGetResponseEnvelopeJSON    `json:"-"`
}

// reverseDNSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ReverseDNSGetResponseEnvelope]
type reverseDNSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ReverseDNSGetResponseEnvelopeSuccess bool

const (
	ReverseDNSGetResponseEnvelopeSuccessTrue ReverseDNSGetResponseEnvelopeSuccess = true
)

func (r ReverseDNSGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ReverseDNSGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
