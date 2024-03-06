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

// ZeroTrustTunnelManagementService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustTunnelManagementService] method instead.
type ZeroTrustTunnelManagementService struct {
	Options []option.RequestOption
}

// NewZeroTrustTunnelManagementService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustTunnelManagementService(opts ...option.RequestOption) (r *ZeroTrustTunnelManagementService) {
	r = &ZeroTrustTunnelManagementService{}
	r.Options = opts
	return
}

// Gets a management token used to access the management resources (i.e. Streaming
// Logs) of a tunnel.
func (r *ZeroTrustTunnelManagementService) New(ctx context.Context, tunnelID string, params ZeroTrustTunnelManagementNewParams, opts ...option.RequestOption) (res *ZeroTrustTunnelManagementNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustTunnelManagementNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/management", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ZeroTrustTunnelManagementNewResponseUnknown],
// [ZeroTrustTunnelManagementNewResponseArray] or [shared.UnionString].
type ZeroTrustTunnelManagementNewResponse interface {
	ImplementsZeroTrustTunnelManagementNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustTunnelManagementNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustTunnelManagementNewResponseArray []interface{}

func (r ZeroTrustTunnelManagementNewResponseArray) ImplementsZeroTrustTunnelManagementNewResponse() {}

type ZeroTrustTunnelManagementNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string]                                       `path:"account_id,required"`
	Resources param.Field[[]ZeroTrustTunnelManagementNewParamsResource] `json:"resources,required"`
}

func (r ZeroTrustTunnelManagementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Management resources the token will have access to.
type ZeroTrustTunnelManagementNewParamsResource string

const (
	ZeroTrustTunnelManagementNewParamsResourceLogs ZeroTrustTunnelManagementNewParamsResource = "logs"
)

type ZeroTrustTunnelManagementNewResponseEnvelope struct {
	Errors   []ZeroTrustTunnelManagementNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustTunnelManagementNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustTunnelManagementNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustTunnelManagementNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustTunnelManagementNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustTunnelManagementNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustTunnelManagementNewResponseEnvelope]
type zeroTrustTunnelManagementNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelManagementNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelManagementNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelManagementNewResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustTunnelManagementNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustTunnelManagementNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelManagementNewResponseEnvelopeErrors]
type zeroTrustTunnelManagementNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelManagementNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelManagementNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustTunnelManagementNewResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustTunnelManagementNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustTunnelManagementNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustTunnelManagementNewResponseEnvelopeMessages]
type zeroTrustTunnelManagementNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTunnelManagementNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTunnelManagementNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustTunnelManagementNewResponseEnvelopeSuccess bool

const (
	ZeroTrustTunnelManagementNewResponseEnvelopeSuccessTrue ZeroTrustTunnelManagementNewResponseEnvelopeSuccess = true
)
