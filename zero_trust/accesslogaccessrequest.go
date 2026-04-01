// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// AccessLogAccessRequestService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessLogAccessRequestService] method instead.
type AccessLogAccessRequestService struct {
	Options []option.RequestOption
}

// NewAccessLogAccessRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessLogAccessRequestService(opts ...option.RequestOption) (r *AccessLogAccessRequestService) {
	r = &AccessLogAccessRequestService{}
	r.Options = opts
	return
}

// Gets a list of Access authentication audit logs for an account.
func (r *AccessLogAccessRequestService) List(ctx context.Context, params AccessLogAccessRequestListParams, opts ...option.RequestOption) (res *[]AccessRequest, err error) {
	var env AccessLogAccessRequestListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/logs/access_requests", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AccessLogAccessRequestListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Operator for the `allowed` filter.
	AllowedOp param.Field[AccessLogAccessRequestListParamsAllowedOp] `query:"allowedOp"`
	// Operator for the `app_type` filter.
	AppTypeOp param.Field[AccessLogAccessRequestListParamsAppTypeOp] `query:"app_typeOp"`
	// Operator for the `app_uid` filter.
	AppUIDOp param.Field[AccessLogAccessRequestListParamsAppUIDOp] `query:"app_uidOp"`
	// Operator for the `country_code` filter.
	CountryCodeOp param.Field[AccessLogAccessRequestListParamsCountryCodeOp] `query:"country_codeOp"`
	// The chronological sorting order for the logs.
	Direction param.Field[AccessLogAccessRequestListParamsDirection] `query:"direction"`
	// Filter by user email. Defaults to substring matching. To force exact matching,
	// set `email_exact=true`. Example (default): `email=@example.com` returns all
	// events with that domain. Example (exact):
	// `email=user@example.com&email_exact=true` returns only that user.
	Email param.Field[string] `query:"email" format:"email"`
	// When true, `email` is matched exactly instead of substring matching.
	EmailExact param.Field[bool] `query:"email_exact"`
	// Operator for the `email` filter.
	EmailOp param.Field[AccessLogAccessRequestListParamsEmailOp] `query:"emailOp"`
	// Comma-separated list of fields to include in the response. When omitted, all
	// fields are returned.
	Fields param.Field[string] `query:"fields"`
	// Operator for the `idp` filter.
	IdPOp param.Field[AccessLogAccessRequestListParamsIdPOp] `query:"idpOp"`
	// The maximum number of log entries to retrieve.
	Limit param.Field[int64] `query:"limit"`
	// Operator for the `non_identity` filter.
	NonIdentityOp param.Field[AccessLogAccessRequestListParamsNonIdentityOp] `query:"non_identityOp"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Operator for the `ray_id` filter.
	RayIDOp param.Field[AccessLogAccessRequestListParamsRayIDOp] `query:"ray_idOp"`
	// The earliest event timestamp to query.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The latest event timestamp to query.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
	// Filter by user UUID.
	UserID param.Field[string] `query:"user_id" format:"uuid"`
	// Operator for the `user_id` filter.
	UserIDOp param.Field[AccessLogAccessRequestListParamsUserIDOp] `query:"user_idOp"`
}

// URLQuery serializes [AccessLogAccessRequestListParams]'s query parameters as
// `url.Values`.
func (r AccessLogAccessRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Operator for the `allowed` filter.
type AccessLogAccessRequestListParamsAllowedOp string

const (
	AccessLogAccessRequestListParamsAllowedOpEq  AccessLogAccessRequestListParamsAllowedOp = "eq"
	AccessLogAccessRequestListParamsAllowedOpNeq AccessLogAccessRequestListParamsAllowedOp = "neq"
)

func (r AccessLogAccessRequestListParamsAllowedOp) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsAllowedOpEq, AccessLogAccessRequestListParamsAllowedOpNeq:
		return true
	}
	return false
}

// Operator for the `app_type` filter.
type AccessLogAccessRequestListParamsAppTypeOp string

const (
	AccessLogAccessRequestListParamsAppTypeOpEq  AccessLogAccessRequestListParamsAppTypeOp = "eq"
	AccessLogAccessRequestListParamsAppTypeOpNeq AccessLogAccessRequestListParamsAppTypeOp = "neq"
)

func (r AccessLogAccessRequestListParamsAppTypeOp) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsAppTypeOpEq, AccessLogAccessRequestListParamsAppTypeOpNeq:
		return true
	}
	return false
}

// Operator for the `app_uid` filter.
type AccessLogAccessRequestListParamsAppUIDOp string

const (
	AccessLogAccessRequestListParamsAppUIDOpEq  AccessLogAccessRequestListParamsAppUIDOp = "eq"
	AccessLogAccessRequestListParamsAppUIDOpNeq AccessLogAccessRequestListParamsAppUIDOp = "neq"
)

func (r AccessLogAccessRequestListParamsAppUIDOp) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsAppUIDOpEq, AccessLogAccessRequestListParamsAppUIDOpNeq:
		return true
	}
	return false
}

// Operator for the `country_code` filter.
type AccessLogAccessRequestListParamsCountryCodeOp string

const (
	AccessLogAccessRequestListParamsCountryCodeOpEq  AccessLogAccessRequestListParamsCountryCodeOp = "eq"
	AccessLogAccessRequestListParamsCountryCodeOpNeq AccessLogAccessRequestListParamsCountryCodeOp = "neq"
)

func (r AccessLogAccessRequestListParamsCountryCodeOp) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsCountryCodeOpEq, AccessLogAccessRequestListParamsCountryCodeOpNeq:
		return true
	}
	return false
}

// The chronological sorting order for the logs.
type AccessLogAccessRequestListParamsDirection string

const (
	AccessLogAccessRequestListParamsDirectionDesc AccessLogAccessRequestListParamsDirection = "desc"
	AccessLogAccessRequestListParamsDirectionAsc  AccessLogAccessRequestListParamsDirection = "asc"
)

func (r AccessLogAccessRequestListParamsDirection) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsDirectionDesc, AccessLogAccessRequestListParamsDirectionAsc:
		return true
	}
	return false
}

// Operator for the `email` filter.
type AccessLogAccessRequestListParamsEmailOp string

const (
	AccessLogAccessRequestListParamsEmailOpEq  AccessLogAccessRequestListParamsEmailOp = "eq"
	AccessLogAccessRequestListParamsEmailOpNeq AccessLogAccessRequestListParamsEmailOp = "neq"
)

func (r AccessLogAccessRequestListParamsEmailOp) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsEmailOpEq, AccessLogAccessRequestListParamsEmailOpNeq:
		return true
	}
	return false
}

// Operator for the `idp` filter.
type AccessLogAccessRequestListParamsIdPOp string

const (
	AccessLogAccessRequestListParamsIdPOpEq  AccessLogAccessRequestListParamsIdPOp = "eq"
	AccessLogAccessRequestListParamsIdPOpNeq AccessLogAccessRequestListParamsIdPOp = "neq"
)

func (r AccessLogAccessRequestListParamsIdPOp) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsIdPOpEq, AccessLogAccessRequestListParamsIdPOpNeq:
		return true
	}
	return false
}

// Operator for the `non_identity` filter.
type AccessLogAccessRequestListParamsNonIdentityOp string

const (
	AccessLogAccessRequestListParamsNonIdentityOpEq  AccessLogAccessRequestListParamsNonIdentityOp = "eq"
	AccessLogAccessRequestListParamsNonIdentityOpNeq AccessLogAccessRequestListParamsNonIdentityOp = "neq"
)

func (r AccessLogAccessRequestListParamsNonIdentityOp) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsNonIdentityOpEq, AccessLogAccessRequestListParamsNonIdentityOpNeq:
		return true
	}
	return false
}

// Operator for the `ray_id` filter.
type AccessLogAccessRequestListParamsRayIDOp string

const (
	AccessLogAccessRequestListParamsRayIDOpEq  AccessLogAccessRequestListParamsRayIDOp = "eq"
	AccessLogAccessRequestListParamsRayIDOpNeq AccessLogAccessRequestListParamsRayIDOp = "neq"
)

func (r AccessLogAccessRequestListParamsRayIDOp) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsRayIDOpEq, AccessLogAccessRequestListParamsRayIDOpNeq:
		return true
	}
	return false
}

// Operator for the `user_id` filter.
type AccessLogAccessRequestListParamsUserIDOp string

const (
	AccessLogAccessRequestListParamsUserIDOpEq  AccessLogAccessRequestListParamsUserIDOp = "eq"
	AccessLogAccessRequestListParamsUserIDOpNeq AccessLogAccessRequestListParamsUserIDOp = "neq"
)

func (r AccessLogAccessRequestListParamsUserIDOp) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsUserIDOpEq, AccessLogAccessRequestListParamsUserIDOpNeq:
		return true
	}
	return false
}

type AccessLogAccessRequestListResponseEnvelope struct {
	Errors   []AccessLogAccessRequestListResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessLogAccessRequestListResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessLogAccessRequestListResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  []AccessRequest                                   `json:"result"`
	JSON    accessLogAccessRequestListResponseEnvelopeJSON    `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessLogAccessRequestListResponseEnvelope]
type accessLogAccessRequestListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessLogAccessRequestListResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccessLogAccessRequestListResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessLogAccessRequestListResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessLogAccessRequestListResponseEnvelopeErrors]
type accessLogAccessRequestListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessLogAccessRequestListResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accessLogAccessRequestListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessLogAccessRequestListResponseEnvelopeErrorsSource]
type accessLogAccessRequestListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessLogAccessRequestListResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccessLogAccessRequestListResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessLogAccessRequestListResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessLogAccessRequestListResponseEnvelopeMessages]
type accessLogAccessRequestListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessLogAccessRequestListResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accessLogAccessRequestListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessLogAccessRequestListResponseEnvelopeMessagesSource]
type accessLogAccessRequestListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessLogAccessRequestListResponseEnvelopeSuccess bool

const (
	AccessLogAccessRequestListResponseEnvelopeSuccessTrue AccessLogAccessRequestListResponseEnvelopeSuccess = true
)

func (r AccessLogAccessRequestListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
