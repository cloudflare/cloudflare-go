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

// TunnelManagementService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTunnelManagementService] method
// instead.
type TunnelManagementService struct {
	Options []option.RequestOption
}

// NewTunnelManagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelManagementService(opts ...option.RequestOption) (r *TunnelManagementService) {
	r = &TunnelManagementService{}
	r.Options = opts
	return
}

// Gets a management token used to access the management resources (i.e. Streaming
// Logs) of a tunnel.
func (r *TunnelManagementService) New(ctx context.Context, tunnelID string, params TunnelManagementNewParams, opts ...option.RequestOption) (res *TunnelManagementNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelManagementNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/management", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [TunnelManagementNewResponseUnknown],
// [TunnelManagementNewResponseArray] or [shared.UnionString].
type TunnelManagementNewResponse interface {
	ImplementsTunnelManagementNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelManagementNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TunnelManagementNewResponseArray []interface{}

func (r TunnelManagementNewResponseArray) ImplementsTunnelManagementNewResponse() {}

type TunnelManagementNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string]                              `path:"account_id,required"`
	Resources param.Field[[]TunnelManagementNewParamsResource] `json:"resources,required"`
}

func (r TunnelManagementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Management resources the token will have access to.
type TunnelManagementNewParamsResource string

const (
	TunnelManagementNewParamsResourceLogs TunnelManagementNewParamsResource = "logs"
)

type TunnelManagementNewResponseEnvelope struct {
	Errors   []TunnelManagementNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelManagementNewResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelManagementNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelManagementNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelManagementNewResponseEnvelopeJSON    `json:"-"`
}

// tunnelManagementNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelManagementNewResponseEnvelope]
type tunnelManagementNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelManagementNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelManagementNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    tunnelManagementNewResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelManagementNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TunnelManagementNewResponseEnvelopeErrors]
type tunnelManagementNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelManagementNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelManagementNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    tunnelManagementNewResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelManagementNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [TunnelManagementNewResponseEnvelopeMessages]
type tunnelManagementNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelManagementNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelManagementNewResponseEnvelopeSuccess bool

const (
	TunnelManagementNewResponseEnvelopeSuccessTrue TunnelManagementNewResponseEnvelopeSuccess = true
)
