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

// AccountAccessOrganizationRevokeUserService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAccessOrganizationRevokeUserService] method instead.
type AccountAccessOrganizationRevokeUserService struct {
	Options []option.RequestOption
}

// NewAccountAccessOrganizationRevokeUserService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountAccessOrganizationRevokeUserService(opts ...option.RequestOption) (r *AccountAccessOrganizationRevokeUserService) {
	r = &AccountAccessOrganizationRevokeUserService{}
	r.Options = opts
	return
}

// Revokes a user's access across all applications.
func (r *AccountAccessOrganizationRevokeUserService) ZeroTrustOrganizationRevokeAllAccessTokensForAUser(ctx context.Context, identifier interface{}, body AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserParams, opts ...option.RequestOption) (res *AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/organizations/revoke_user", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponse struct {
	Result  AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseResult  `json:"result"`
	Success AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseSuccess `json:"success"`
	JSON    accountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseJSON    `json:"-"`
}

// accountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponse]
type accountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseResult bool

const (
	AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseResultTrue  AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseResult = true
	AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseResultFalse AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseResult = false
)

type AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseSuccess bool

const (
	AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseSuccessTrue  AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseSuccess = true
	AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseSuccessFalse AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseSuccess = false
)

type AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserParams struct {
	// The email of the user to revoke.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
