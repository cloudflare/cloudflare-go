// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// StreamTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamTokenService] method
// instead.
type StreamTokenService struct {
	Options []option.RequestOption
}

// NewStreamTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamTokenService(opts ...option.RequestOption) (r *StreamTokenService) {
	r = &StreamTokenService{}
	r.Options = opts
	return
}

// Creates a signed URL token for a video. If a body is not provided in the
// request, a token is created with default values.
func (r *StreamTokenService) New(ctx context.Context, identifier string, params StreamTokenNewParams, opts ...option.RequestOption) (res *StreamTokenNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamTokenNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/token", params.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamTokenNewResponse struct {
	// The signed token used with the signed URLs feature.
	Token string                     `json:"token"`
	JSON  streamTokenNewResponseJSON `json:"-"`
}

// streamTokenNewResponseJSON contains the JSON metadata for the struct
// [StreamTokenNewResponse]
type streamTokenNewResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamTokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamTokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type StreamTokenNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The optional ID of a Stream signing key. If present, the `pem` field is also
	// required.
	ID param.Field[string] `json:"id"`
	// The optional list of access rule constraints on the token. Access can be blocked
	// or allowed based on an IP, IP range, or by country. Access rules are evaluated
	// from first to last. If a rule matches, the associated action is applied and no
	// further rules are evaluated.
	AccessRules param.Field[[]StreamTokenNewParamsAccessRule] `json:"accessRules"`
	// The optional boolean value that enables using signed tokens to access MP4
	// download links for a video.
	Downloadable param.Field[bool] `json:"downloadable"`
	// The optional unix epoch timestamp that specficies the time after a token is not
	// accepted. The maximum time specification is 24 hours from issuing time. If this
	// field is not set, the default is one hour after issuing.
	Exp param.Field[int64] `json:"exp"`
	// The optional unix epoch timestamp that specifies the time before a the token is
	// not accepted. If this field is not set, the default is one hour before issuing.
	Nbf param.Field[int64] `json:"nbf"`
	// The optional base64 encoded private key in PEM format associated with a Stream
	// signing key. If present, the `id` field is also required.
	Pem param.Field[string] `json:"pem"`
}

func (r StreamTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines rules for fine-grained control over content than signed URL tokens
// alone. Access rules primarily make tokens conditionally valid based on user
// information. Access Rules are specified on token payloads as the `accessRules`
// property containing an array of Rule objects.
type StreamTokenNewParamsAccessRule struct {
	// The action to take when a request matches a rule. If the action is `block`, the
	// signed token blocks views for viewers matching the rule.
	Action param.Field[StreamTokenNewParamsAccessRulesAction] `json:"action"`
	// An array of 2-letter country codes in ISO 3166-1 Alpha-2 format used to match
	// requests.
	Country param.Field[[]string] `json:"country"`
	// An array of IPv4 or IPV6 addresses or CIDRs used to match requests.
	IP param.Field[[]string] `json:"ip"`
	// Lists available rule types to match for requests. An `any` type matches all
	// requests and can be used as a wildcard to apply default actions after other
	// rules.
	Type param.Field[StreamTokenNewParamsAccessRulesType] `json:"type"`
}

func (r StreamTokenNewParamsAccessRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when a request matches a rule. If the action is `block`, the
// signed token blocks views for viewers matching the rule.
type StreamTokenNewParamsAccessRulesAction string

const (
	StreamTokenNewParamsAccessRulesActionAllow StreamTokenNewParamsAccessRulesAction = "allow"
	StreamTokenNewParamsAccessRulesActionBlock StreamTokenNewParamsAccessRulesAction = "block"
)

// Lists available rule types to match for requests. An `any` type matches all
// requests and can be used as a wildcard to apply default actions after other
// rules.
type StreamTokenNewParamsAccessRulesType string

const (
	StreamTokenNewParamsAccessRulesTypeAny            StreamTokenNewParamsAccessRulesType = "any"
	StreamTokenNewParamsAccessRulesTypeIPSrc          StreamTokenNewParamsAccessRulesType = "ip.src"
	StreamTokenNewParamsAccessRulesTypeIPGeoipCountry StreamTokenNewParamsAccessRulesType = "ip.geoip.country"
)

type StreamTokenNewResponseEnvelope struct {
	Errors   []StreamTokenNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamTokenNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamTokenNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamTokenNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamTokenNewResponseEnvelopeJSON    `json:"-"`
}

// streamTokenNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamTokenNewResponseEnvelope]
type streamTokenNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamTokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamTokenNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamTokenNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    streamTokenNewResponseEnvelopeErrorsJSON `json:"-"`
}

// streamTokenNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamTokenNewResponseEnvelopeErrors]
type streamTokenNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamTokenNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamTokenNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamTokenNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    streamTokenNewResponseEnvelopeMessagesJSON `json:"-"`
}

// streamTokenNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamTokenNewResponseEnvelopeMessages]
type streamTokenNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamTokenNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamTokenNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamTokenNewResponseEnvelopeSuccess bool

const (
	StreamTokenNewResponseEnvelopeSuccessTrue StreamTokenNewResponseEnvelopeSuccess = true
)
