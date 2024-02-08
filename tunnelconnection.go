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

// TunnelConnectionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTunnelConnectionService] method
// instead.
type TunnelConnectionService struct {
	Options []option.RequestOption
}

// NewTunnelConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelConnectionService(opts ...option.RequestOption) (r *TunnelConnectionService) {
	r = &TunnelConnectionService{}
	r.Options = opts
	return
}

// Removes connections that are in a disconnected or pending reconnect state. We
// recommend running this command after shutting down a tunnel.
func (r *TunnelConnectionService) Delete(ctx context.Context, accountID string, tunnelID string, body TunnelConnectionDeleteParams, opts ...option.RequestOption) (res *TunnelConnectionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelConnectionDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [TunnelConnectionDeleteResponseUnknown],
// [TunnelConnectionDeleteResponseArray] or [shared.UnionString].
type TunnelConnectionDeleteResponse interface {
	ImplementsTunnelConnectionDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelConnectionDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TunnelConnectionDeleteResponseArray []interface{}

func (r TunnelConnectionDeleteResponseArray) ImplementsTunnelConnectionDeleteResponse() {}

type TunnelConnectionDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r TunnelConnectionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type TunnelConnectionDeleteResponseEnvelope struct {
	Errors   []TunnelConnectionDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelConnectionDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelConnectionDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelConnectionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelConnectionDeleteResponseEnvelopeJSON    `json:"-"`
}

// tunnelConnectionDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelConnectionDeleteResponseEnvelope]
type tunnelConnectionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    tunnelConnectionDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelConnectionDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [TunnelConnectionDeleteResponseEnvelopeErrors]
type tunnelConnectionDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    tunnelConnectionDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelConnectionDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [TunnelConnectionDeleteResponseEnvelopeMessages]
type tunnelConnectionDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelConnectionDeleteResponseEnvelopeSuccess bool

const (
	TunnelConnectionDeleteResponseEnvelopeSuccessTrue TunnelConnectionDeleteResponseEnvelopeSuccess = true
)
