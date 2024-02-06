// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessAppRevokeTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessAppRevokeTokenService]
// method instead.
type AccessAppRevokeTokenService struct {
	Options []option.RequestOption
}

// NewAccessAppRevokeTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessAppRevokeTokenService(opts ...option.RequestOption) (r *AccessAppRevokeTokenService) {
	r = &AccessAppRevokeTokenService{}
	r.Options = opts
	return
}

// Revokes all tokens issued for an application.
func (r *AccessAppRevokeTokenService) AccessApplicationsRevokeServiceTokens(ctx context.Context, accountOrZone string, accountOrZoneID string, appID AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensParamsAppID, opts ...option.RequestOption) (res *AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%v/revoke_tokens", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponse = interface{}

// Identifier
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensParamsAppID interface {
	ImplementsAccessAppRevokeTokenAccessApplicationsRevokeServiceTokensParamsAppID()
}

type AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelope struct {
	Result  AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponse                `json:"result,nullable"`
	Success AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelopeSuccess `json:"success"`
	JSON    accessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelopeJSON    `json:"-"`
}

// accessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelope]
type accessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelopeSuccess bool

const (
	AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelopeSuccessTrue  AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelopeSuccess = true
	AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelopeSuccessFalse AccessAppRevokeTokenAccessApplicationsRevokeServiceTokensResponseEnvelopeSuccess = false
)
