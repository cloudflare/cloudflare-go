// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRoleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRoleService] method
// instead.
type AccountRoleService struct {
	Options []option.RequestOption
}

// NewAccountRoleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRoleService(opts ...option.RequestOption) (r *AccountRoleService) {
	r = &AccountRoleService{}
	r.Options = opts
	return
}

// Get information about a specific role for an account.
func (r *AccountRoleService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *ResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/roles/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all available roles for an account.
func (r *AccountRoleService) AccountRolesListRoles(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *ResponseCollectionYgt6DzoC, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/roles", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
