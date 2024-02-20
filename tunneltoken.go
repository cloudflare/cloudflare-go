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

// TunnelTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTunnelTokenService] method
// instead.
type TunnelTokenService struct {
	Options []option.RequestOption
}

// NewTunnelTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTunnelTokenService(opts ...option.RequestOption) (r *TunnelTokenService) {
	r = &TunnelTokenService{}
	r.Options = opts
	return
}

// Gets the token used to associate cloudflared with a specific tunnel.
func (r *TunnelTokenService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *TunnelTokenGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelTokenGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/token", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [TunnelTokenGetResponseUnknown],
// [TunnelTokenGetResponseArray] or [shared.UnionString].
type TunnelTokenGetResponse interface {
	ImplementsTunnelTokenGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelTokenGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TunnelTokenGetResponseArray []interface{}

func (r TunnelTokenGetResponseArray) ImplementsTunnelTokenGetResponse() {}

type TunnelTokenGetResponseEnvelope struct {
	Errors   []TunnelTokenGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelTokenGetResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelTokenGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelTokenGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelTokenGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelTokenGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TunnelTokenGetResponseEnvelope]
type tunnelTokenGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelTokenGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelTokenGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    tunnelTokenGetResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelTokenGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TunnelTokenGetResponseEnvelopeErrors]
type tunnelTokenGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelTokenGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelTokenGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    tunnelTokenGetResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelTokenGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TunnelTokenGetResponseEnvelopeMessages]
type tunnelTokenGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelTokenGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelTokenGetResponseEnvelopeSuccess bool

const (
	TunnelTokenGetResponseEnvelopeSuccessTrue TunnelTokenGetResponseEnvelopeSuccess = true
)
