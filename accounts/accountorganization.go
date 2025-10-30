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

// AccountOrganizationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountOrganizationService] method instead.
type AccountOrganizationService struct {
	Options []option.RequestOption
}

// NewAccountOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountOrganizationService(opts ...option.RequestOption) (r *AccountOrganizationService) {
	r = &AccountOrganizationService{}
	r.Options = opts
	return
}

// Move an account within an organization hierarchy or an account outside an
// organization. (Currently in Closed Beta - see
// https://developers.cloudflare.com/fundamentals/organizations/)
func (r *AccountOrganizationService) New(ctx context.Context, params AccountOrganizationNewParams, opts ...option.RequestOption) (res *AccountOrganizationNewResponse, err error) {
	var env AccountOrganizationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/move", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccountOrganizationNewResponse struct {
	AccountID                 string                             `json:"account_id,required"`
	DestinationOrganizationID string                             `json:"destination_organization_id,required"`
	SourceOrganizationID      string                             `json:"source_organization_id,required"`
	JSON                      accountOrganizationNewResponseJSON `json:"-"`
}

// accountOrganizationNewResponseJSON contains the JSON metadata for the struct
// [AccountOrganizationNewResponse]
type accountOrganizationNewResponseJSON struct {
	AccountID                 apijson.Field
	DestinationOrganizationID apijson.Field
	SourceOrganizationID      apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AccountOrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOrganizationNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountOrganizationNewParams struct {
	AccountID                 param.Field[string] `path:"account_id,required"`
	DestinationOrganizationID param.Field[string] `json:"destination_organization_id,required"`
}

func (r AccountOrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountOrganizationNewResponseEnvelope struct {
	Errors   []interface{}                                 `json:"errors,required"`
	Messages []shared.ResponseInfo                         `json:"messages,required"`
	Result   AccountOrganizationNewResponse                `json:"result,required"`
	Success  AccountOrganizationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON     accountOrganizationNewResponseEnvelopeJSON    `json:"-"`
}

// accountOrganizationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccountOrganizationNewResponseEnvelope]
type accountOrganizationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountOrganizationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountOrganizationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccountOrganizationNewResponseEnvelopeSuccess bool

const (
	AccountOrganizationNewResponseEnvelopeSuccessTrue AccountOrganizationNewResponseEnvelopeSuccess = true
)

func (r AccountOrganizationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccountOrganizationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
