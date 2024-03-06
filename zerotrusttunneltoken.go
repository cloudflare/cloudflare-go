// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZeroTrustTunnelTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustTunnelTokenService]
// method instead.
type ZeroTrustTunnelTokenService struct {
	Options []option.RequestOption
}

// NewZeroTrustTunnelTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustTunnelTokenService(opts ...option.RequestOption) (r *ZeroTrustTunnelTokenService) {
	r = &ZeroTrustTunnelTokenService{}
	r.Options = opts
	return
}

// Gets the token used to associate cloudflared with a specific tunnel.
func (r *ZeroTrustTunnelTokenService) Get(ctx context.Context, tunnelID string, query ZeroTrustTunnelTokenGetParams, opts ...option.RequestOption) (res *ZeroTrustTunnelTokenGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustTunnelTokenGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/token", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ZeroTrustTunnelTokenGetResponseUnknown],
// [ZeroTrustTunnelTokenGetResponseArray] or [shared.UnionString].
type ZeroTrustTunnelTokenGetResponse interface {
	ImplementsZeroTrustTunnelTokenGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustTunnelTokenGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustTunnelTokenGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustTunnelTokenGetResponseArray []interface{}

func (r ZeroTrustTunnelTokenGetResponseArray) ImplementsZeroTrustTunnelTokenGetResponse() {}

type ZeroTrustTunnelTokenGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustTunnelTokenGetResponseEnvelope struct {
	Errors   []ZeroTrustTunnelTokenGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustTunnelTokenGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustTunnelTokenGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustTunnelTokenGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustTunnelTokenGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustTunnelTokenGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustTunnelTokenGetResponseEnvelope]
type zeroTrustTunnelTokenGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelTokenGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelTokenGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelTokenGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustTunnelTokenGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustTunnelTokenGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustTunnelTokenGetResponseEnvelopeErrors]
type zeroTrustTunnelTokenGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelTokenGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelTokenGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelTokenGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustTunnelTokenGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustTunnelTokenGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustTunnelTokenGetResponseEnvelopeMessages]
type zeroTrustTunnelTokenGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelTokenGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelTokenGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustTunnelTokenGetResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelTokenGetResponseEnvelopeSuccessTrue ZeroTrustTunnelTokenGetResponseEnvelopeSuccess = true
)
