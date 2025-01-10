// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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

// SubdomainService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubdomainService] method instead.
type SubdomainService struct {
	Options []option.RequestOption
}

// NewSubdomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubdomainService(opts ...option.RequestOption) (r *SubdomainService) {
	r = &SubdomainService{}
	r.Options = opts
	return
}

// Creates a Workers subdomain for an account.
func (r *SubdomainService) Update(ctx context.Context, params SubdomainUpdateParams, opts ...option.RequestOption) (res *SubdomainUpdateResponse, err error) {
	var env SubdomainUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/subdomain", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a Workers subdomain for an account.
func (r *SubdomainService) Get(ctx context.Context, query SubdomainGetParams, opts ...option.RequestOption) (res *SubdomainGetResponse, err error) {
	var env SubdomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/subdomain", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SubdomainUpdateResponse struct {
	Subdomain string                      `json:"subdomain"`
	JSON      subdomainUpdateResponseJSON `json:"-"`
}

// subdomainUpdateResponseJSON contains the JSON metadata for the struct
// [SubdomainUpdateResponse]
type subdomainUpdateResponseJSON struct {
	Subdomain   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SubdomainGetResponse struct {
	Subdomain string                   `json:"subdomain"`
	JSON      subdomainGetResponseJSON `json:"-"`
}

// subdomainGetResponseJSON contains the JSON metadata for the struct
// [SubdomainGetResponse]
type subdomainGetResponseJSON struct {
	Subdomain   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type SubdomainUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Subdomain param.Field[string] `json:"subdomain"`
}

func (r SubdomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubdomainUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SubdomainUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  SubdomainUpdateResponse                `json:"result"`
	JSON    subdomainUpdateResponseEnvelopeJSON    `json:"-"`
}

// subdomainUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubdomainUpdateResponseEnvelope]
type subdomainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubdomainUpdateResponseEnvelopeSuccess bool

const (
	SubdomainUpdateResponseEnvelopeSuccessTrue SubdomainUpdateResponseEnvelopeSuccess = true
)

func (r SubdomainUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubdomainUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubdomainGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SubdomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SubdomainGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SubdomainGetResponse                `json:"result"`
	JSON    subdomainGetResponseEnvelopeJSON    `json:"-"`
}

// subdomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubdomainGetResponseEnvelope]
type subdomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubdomainGetResponseEnvelopeSuccess bool

const (
	SubdomainGetResponseEnvelopeSuccessTrue SubdomainGetResponseEnvelopeSuccess = true
)

func (r SubdomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubdomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
