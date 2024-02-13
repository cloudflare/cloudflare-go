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

// AccessOrganizationRevokeUserService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccessOrganizationRevokeUserService] method instead.
type AccessOrganizationRevokeUserService struct {
	Options []option.RequestOption
}

// NewAccessOrganizationRevokeUserService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessOrganizationRevokeUserService(opts ...option.RequestOption) (r *AccessOrganizationRevokeUserService) {
	r = &AccessOrganizationRevokeUserService{}
	r.Options = opts
	return
}

// Revokes a user's access across all applications.
func (r *AccessOrganizationRevokeUserService) ZeroTrustOrganizationRevokeAllAccessTokensForAUser(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserParams, opts ...option.RequestOption) (res *AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/organizations/revoke_user", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponse bool

const (
	AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseTrue  AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponse = true
	AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseFalse AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponse = false
)

type AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserParams struct {
	// The email of the user to revoke.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelope struct {
	Result  AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponse                `json:"result"`
	Success AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelopeSuccess `json:"success"`
	JSON    accessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelopeJSON    `json:"-"`
}

// accessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelope]
type accessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelopeSuccess bool

const (
	AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelopeSuccessTrue  AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelopeSuccess = true
	AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelopeSuccessFalse AccessOrganizationRevokeUserZeroTrustOrganizationRevokeAllAccessTokensForAUserResponseEnvelopeSuccess = false
)
