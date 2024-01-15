// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountStreamTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamTokenService] method
// instead.
type AccountStreamTokenService struct {
	Options []option.RequestOption
}

// NewAccountStreamTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamTokenService(opts ...option.RequestOption) (r *AccountStreamTokenService) {
	r = &AccountStreamTokenService{}
	r.Options = opts
	return
}

// Creates a signed URL token for a video. If a body is not provided in the
// request, a token is created with default values.
func (r *AccountStreamTokenService) StreamVideosNewSignedURLTokensForVideos(ctx context.Context, accountIdentifier string, identifier string, body AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParams, opts ...option.RequestOption) (res *AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/token", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponse struct {
	Errors   []AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseError   `json:"errors"`
	Messages []AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseMessage `json:"messages"`
	Result   AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseSuccess `json:"success"`
	JSON    accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseJSON    `json:"-"`
}

// accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseJSON contains
// the JSON metadata for the struct
// [AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponse]
type accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseErrorJSON `json:"-"`
}

// accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseError]
type accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseMessageJSON `json:"-"`
}

// accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseMessage]
type accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseResult struct {
	// The signed token used with the signed URLs feature.
	Token string                                                                      `json:"token"`
	JSON  accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseResultJSON `json:"-"`
}

// accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseResultJSON
// contains the JSON metadata for the struct
// [AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseResult]
type accountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseResultJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseSuccess bool

const (
	AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseSuccessTrue AccountStreamTokenStreamVideosNewSignedURLTokensForVideosResponseSuccess = true
)

type AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParams struct {
	// The optional ID of a Stream signing key. If present, the `pem` field is also
	// required.
	ID param.Field[string] `json:"id"`
	// The optional list of access rule constraints on the token. Access can be blocked
	// or allowed based on an IP, IP range, or by country. Access rules are evaluated
	// from first to last. If a rule matches, the associated action is applied and no
	// further rules are evaluated.
	AccessRules param.Field[[]AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRule] `json:"accessRules"`
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

func (r AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines rules for fine-grained control over content than signed URL tokens
// alone. Access rules primarily make tokens conditionally valid based on user
// information. Access Rules are specified on token payloads as the `accessRules`
// property containing an array of Rule objects.
type AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRule struct {
	// The action to take when a request matches a rule. If the action is `block`, the
	// signed token blocks views for viewers matching the rule.
	Action param.Field[AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesAction] `json:"action"`
	// An array of 2-letter country codes in ISO 3166-1 Alpha-2 format used to match
	// requests.
	Country param.Field[[]string] `json:"country"`
	// An array of IPv4 or IPV6 addresses or CIDRs used to match requests.
	IP param.Field[[]string] `json:"ip"`
	// Lists available rule types to match for requests. An `any` type matches all
	// requests and can be used as a wildcard to apply default actions after other
	// rules.
	Type param.Field[AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesType] `json:"type"`
}

func (r AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when a request matches a rule. If the action is `block`, the
// signed token blocks views for viewers matching the rule.
type AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesAction string

const (
	AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesActionAllow AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesAction = "allow"
	AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesActionBlock AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesAction = "block"
)

// Lists available rule types to match for requests. An `any` type matches all
// requests and can be used as a wildcard to apply default actions after other
// rules.
type AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesType string

const (
	AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesTypeAny            AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesType = "any"
	AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesTypeIPSrc          AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesType = "ip.src"
	AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesTypeIPGeoipCountry AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesType = "ip.geoip.country"
)
