// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// OrganizationDOHService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationDOHService] method instead.
type OrganizationDOHService struct {
	Options []option.RequestOption
}

// NewOrganizationDOHService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationDOHService(opts ...option.RequestOption) (r *OrganizationDOHService) {
	r = &OrganizationDOHService{}
	r.Options = opts
	return
}

// Updates the DoH settings for your Zero Trust organization.
func (r *OrganizationDOHService) Update(ctx context.Context, params OrganizationDOHUpdateParams, opts ...option.RequestOption) (res *OrganizationDOHUpdateResponse, err error) {
	var env OrganizationDOHUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/organizations/doh", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the DoH settings for your Zero Trust organization.
func (r *OrganizationDOHService) Get(ctx context.Context, query OrganizationDOHGetParams, opts ...option.RequestOption) (res *OrganizationDOHGetResponse, err error) {
	var env OrganizationDOHGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/organizations/doh", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OrganizationDOHUpdateResponse struct {
	// The ID of the service token.
	ID string `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration  string    `json:"duration"`
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// The duration the DoH JWT is valid for. Must be in the format `300ms` or `2h45m`.
	// Valid time units are: ns, us (or µs), ms, s, m, h. Note that the maximum
	// duration for this setting is the same as the key rotation period on the account.
	// Default expiration is 24h
	JWTDuration string `json:"jwt_duration"`
	// The name of the service token.
	Name      string                            `json:"name"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      organizationDOHUpdateResponseJSON `json:"-"`
}

// organizationDOHUpdateResponseJSON contains the JSON metadata for the struct
// [OrganizationDOHUpdateResponse]
type organizationDOHUpdateResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	ExpiresAt   apijson.Field
	JWTDuration apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationDOHUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDOHUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationDOHGetResponse struct {
	// The ID of the service token.
	ID string `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration  string    `json:"duration"`
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// The duration the DoH JWT is valid for. Must be in the format `300ms` or `2h45m`.
	// Valid time units are: ns, us (or µs), ms, s, m, h. Note that the maximum
	// duration for this setting is the same as the key rotation period on the account.
	JWTDuration string `json:"jwt_duration"`
	// The name of the service token.
	Name      string                         `json:"name"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      organizationDOHGetResponseJSON `json:"-"`
}

// organizationDOHGetResponseJSON contains the JSON metadata for the struct
// [OrganizationDOHGetResponse]
type organizationDOHGetResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	ExpiresAt   apijson.Field
	JWTDuration apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationDOHGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDOHGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationDOHUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The duration the DoH JWT is valid for. Must be in the format `300ms` or `2h45m`.
	// Valid time units are: ns, us (or µs), ms, s, m, h. Note that the maximum
	// duration for this setting is the same as the key rotation period on the account.
	// Default expiration is 24h
	JWTDuration param.Field[string] `json:"jwt_duration"`
	// The uuid of the service token you want to use for DoH authentication
	ServiceTokenID param.Field[string] `json:"service_token_id"`
}

func (r OrganizationDOHUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationDOHUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OrganizationDOHUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  OrganizationDOHUpdateResponse                `json:"result"`
	JSON    organizationDOHUpdateResponseEnvelopeJSON    `json:"-"`
}

// organizationDOHUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [OrganizationDOHUpdateResponseEnvelope]
type organizationDOHUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationDOHUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDOHUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OrganizationDOHUpdateResponseEnvelopeSuccess bool

const (
	OrganizationDOHUpdateResponseEnvelopeSuccessTrue OrganizationDOHUpdateResponseEnvelopeSuccess = true
)

func (r OrganizationDOHUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationDOHUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationDOHGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type OrganizationDOHGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OrganizationDOHGetResponseEnvelopeSuccess `json:"success,required"`
	Result  OrganizationDOHGetResponse                `json:"result"`
	JSON    organizationDOHGetResponseEnvelopeJSON    `json:"-"`
}

// organizationDOHGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationDOHGetResponseEnvelope]
type organizationDOHGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationDOHGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDOHGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OrganizationDOHGetResponseEnvelopeSuccess bool

const (
	OrganizationDOHGetResponseEnvelopeSuccessTrue OrganizationDOHGetResponseEnvelopeSuccess = true
)

func (r OrganizationDOHGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationDOHGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
