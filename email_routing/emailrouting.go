// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

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

// EmailRoutingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailRoutingService] method instead.
type EmailRoutingService struct {
	Options   []option.RequestOption
	DNS       *DNSService
	Rules     *RuleService
	Addresses *AddressService
}

// NewEmailRoutingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailRoutingService(opts ...option.RequestOption) (r *EmailRoutingService) {
	r = &EmailRoutingService{}
	r.Options = opts
	r.DNS = NewDNSService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Addresses = NewAddressService(opts...)
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *EmailRoutingService) Disable(ctx context.Context, params EmailRoutingDisableParams, opts ...option.RequestOption) (res *EmailRoutingDisableResponse, err error) {
	var env EmailRoutingDisableResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/disable", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
func (r *EmailRoutingService) Enable(ctx context.Context, params EmailRoutingEnableParams, opts ...option.RequestOption) (res *EmailRoutingEnableResponse, err error) {
	var env EmailRoutingEnableResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/enable", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about the settings for your Email Routing zone.
func (r *EmailRoutingService) Get(ctx context.Context, query EmailRoutingGetParams, opts ...option.RequestOption) (res *EmailRoutingGetResponse, err error) {
	var env EmailRoutingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingDisableResponse = interface{}

type EmailRoutingEnableResponse = interface{}

type EmailRoutingGetResponse = interface{}

type EmailRoutingDisableParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r EmailRoutingDisableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type EmailRoutingDisableResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success EmailRoutingDisableResponseEnvelopeSuccess `json:"success,required"`
	Result  EmailRoutingDisableResponse                `json:"result"`
	JSON    emailRoutingDisableResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingDisableResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingDisableResponseEnvelope]
type emailRoutingDisableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDisableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingDisableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingDisableResponseEnvelopeSuccess bool

const (
	EmailRoutingDisableResponseEnvelopeSuccessTrue EmailRoutingDisableResponseEnvelopeSuccess = true
)

func (r EmailRoutingDisableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EmailRoutingDisableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type EmailRoutingEnableParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r EmailRoutingEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type EmailRoutingEnableResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success EmailRoutingEnableResponseEnvelopeSuccess `json:"success,required"`
	Result  EmailRoutingEnableResponse                `json:"result"`
	JSON    emailRoutingEnableResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingEnableResponseEnvelopeJSON contains the JSON metadata for the struct
// [EmailRoutingEnableResponseEnvelope]
type emailRoutingEnableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingEnableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingEnableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingEnableResponseEnvelopeSuccess bool

const (
	EmailRoutingEnableResponseEnvelopeSuccessTrue EmailRoutingEnableResponseEnvelopeSuccess = true
)

func (r EmailRoutingEnableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EmailRoutingEnableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type EmailRoutingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type EmailRoutingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success EmailRoutingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  EmailRoutingGetResponse                `json:"result"`
	JSON    emailRoutingGetResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EmailRoutingGetResponseEnvelope]
type emailRoutingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingGetResponseEnvelopeSuccess bool

const (
	EmailRoutingGetResponseEnvelopeSuccessTrue EmailRoutingGetResponseEnvelopeSuccess = true
)

func (r EmailRoutingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EmailRoutingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
