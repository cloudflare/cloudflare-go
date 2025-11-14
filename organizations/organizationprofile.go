// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// OrganizationProfileService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationProfileService] method instead.
type OrganizationProfileService struct {
	Options []option.RequestOption
}

// NewOrganizationProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationProfileService(opts ...option.RequestOption) (r *OrganizationProfileService) {
	r = &OrganizationProfileService{}
	r.Options = opts
	return
}

// Modify organization profile. (Currently in Closed Beta - see
// https://developers.cloudflare.com/fundamentals/organizations/)
func (r *OrganizationProfileService) Update(ctx context.Context, organizationID string, body OrganizationProfileUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/profile", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Get an organizations profile if it exists. (Currently in Closed Beta - see
// https://developers.cloudflare.com/fundamentals/organizations/)
func (r *OrganizationProfileService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *accounts.AccountProfile, err error) {
	var env OrganizationProfileGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/profile", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OrganizationProfileUpdateParams struct {
	AccountProfile accounts.AccountProfileParam `json:"account_profile,required"`
}

func (r OrganizationProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AccountProfile)
}

type OrganizationProfileGetResponseEnvelope struct {
	Errors   []interface{}                                 `json:"errors,required"`
	Messages []shared.ResponseInfo                         `json:"messages,required"`
	Result   accounts.AccountProfile                       `json:"result,required"`
	Success  OrganizationProfileGetResponseEnvelopeSuccess `json:"success,required"`
	JSON     organizationProfileGetResponseEnvelopeJSON    `json:"-"`
}

// organizationProfileGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [OrganizationProfileGetResponseEnvelope]
type organizationProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationProfileGetResponseEnvelopeSuccess bool

const (
	OrganizationProfileGetResponseEnvelopeSuccessTrue OrganizationProfileGetResponseEnvelopeSuccess = true
)

func (r OrganizationProfileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationProfileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
