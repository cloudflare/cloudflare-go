// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// CfdTunnelTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCfdTunnelTokenService] method
// instead.
type CfdTunnelTokenService struct {
	Options []option.RequestOption
}

// NewCfdTunnelTokenService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCfdTunnelTokenService(opts ...option.RequestOption) (r *CfdTunnelTokenService) {
	r = &CfdTunnelTokenService{}
	r.Options = opts
	return
}

// Gets the token used to associate cloudflared with a specific tunnel.
func (r *CfdTunnelTokenService) CloudflareTunnelGetACloudflareTunnelToken(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/token", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseUnknown],
// [CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseArray] or
// [shared.UnionString].
type CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponse interface {
	ImplementsCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseArray []interface{}

func (r CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseArray) ImplementsCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponse() {
}

type CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelope struct {
	Errors   []CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeMessages `json:"messages,required"`
	Result   CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeJSON    `json:"-"`
}

// cfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelope]
type cfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeErrors struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    cfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeErrorsJSON `json:"-"`
}

// cfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeErrors]
type cfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeMessages struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    cfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeMessagesJSON `json:"-"`
}

// cfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeMessages]
type cfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeSuccess bool

const (
	CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeSuccessTrue CfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseEnvelopeSuccess = true
)
