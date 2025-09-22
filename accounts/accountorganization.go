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
// organization.
func (r *AccountOrganizationService) New(ctx context.Context, params AccountOrganizationNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/move", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type AccountOrganizationNewParams struct {
	AccountID                 param.Field[string] `path:"account_id,required"`
	DestinationOrganizationID param.Field[string] `json:"destination_organization_id,required"`
}

func (r AccountOrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
