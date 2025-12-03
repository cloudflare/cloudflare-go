// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// AccountProfileService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountProfileService] method instead.
type AccountProfileService struct {
	Options []option.RequestOption
}

// NewAccountProfileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountProfileService(opts ...option.RequestOption) (r *AccountProfileService) {
	r = &AccountProfileService{}
	r.Options = opts
	return
}

// Modify account profile
func (r *AccountProfileService) Update(ctx context.Context, params AccountProfileUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/profile", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// Get account profile
func (r *AccountProfileService) Get(ctx context.Context, query AccountProfileGetParams, opts ...option.RequestOption) (res *AccountProfile, err error) {
	var env AccountProfileGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/profile", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccountProfile struct {
	BusinessAddress  string             `json:"business_address,required"`
	BusinessEmail    string             `json:"business_email,required"`
	BusinessName     string             `json:"business_name,required"`
	BusinessPhone    string             `json:"business_phone,required"`
	ExternalMetadata string             `json:"external_metadata,required"`
	JSON             accountProfileJSON `json:"-"`
}

// accountProfileJSON contains the JSON metadata for the struct [AccountProfile]
type accountProfileJSON struct {
	BusinessAddress  apijson.Field
	BusinessEmail    apijson.Field
	BusinessName     apijson.Field
	BusinessPhone    apijson.Field
	ExternalMetadata apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountProfileJSON) RawJSON() string {
	return r.raw
}

type AccountProfileParam struct {
	BusinessAddress  param.Field[string] `json:"business_address,required"`
	BusinessEmail    param.Field[string] `json:"business_email,required"`
	BusinessName     param.Field[string] `json:"business_name,required"`
	BusinessPhone    param.Field[string] `json:"business_phone,required"`
	ExternalMetadata param.Field[string] `json:"external_metadata,required"`
}

func (r AccountProfileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountProfileUpdateParams struct {
	AccountID      param.Field[string] `path:"account_id,required"`
	AccountProfile AccountProfileParam `json:"account_profile,required"`
}

func (r AccountProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AccountProfile)
}

type AccountProfileGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccountProfileGetResponseEnvelope struct {
	Errors   []interface{}                            `json:"errors,required"`
	Messages []shared.ResponseInfo                    `json:"messages,required"`
	Result   AccountProfile                           `json:"result,required"`
	Success  AccountProfileGetResponseEnvelopeSuccess `json:"success,required"`
	JSON     accountProfileGetResponseEnvelopeJSON    `json:"-"`
}

// accountProfileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountProfileGetResponseEnvelope]
type accountProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccountProfileGetResponseEnvelopeSuccess bool

const (
	AccountProfileGetResponseEnvelopeSuccessTrue AccountProfileGetResponseEnvelopeSuccess = true
)

func (r AccountProfileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccountProfileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
