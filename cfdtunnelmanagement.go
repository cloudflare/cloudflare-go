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

// CfdTunnelManagementService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCfdTunnelManagementService]
// method instead.
type CfdTunnelManagementService struct {
	Options []option.RequestOption
}

// NewCfdTunnelManagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCfdTunnelManagementService(opts ...option.RequestOption) (r *CfdTunnelManagementService) {
	r = &CfdTunnelManagementService{}
	r.Options = opts
	return
}

// Gets a management token used to access the management resources (i.e. Streaming
// Logs) of a tunnel.
func (r *CfdTunnelManagementService) New(ctx context.Context, accountID string, tunnelID string, body CfdTunnelManagementNewParams, opts ...option.RequestOption) (res *CfdTunnelManagementNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfdTunnelManagementNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/management", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [CfdTunnelManagementNewResponseUnknown],
// [CfdTunnelManagementNewResponseArray] or [shared.UnionString].
type CfdTunnelManagementNewResponse interface {
	ImplementsCfdTunnelManagementNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CfdTunnelManagementNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CfdTunnelManagementNewResponseArray []interface{}

func (r CfdTunnelManagementNewResponseArray) ImplementsCfdTunnelManagementNewResponse() {}

type CfdTunnelManagementNewParams struct {
	Resources param.Field[[]CfdTunnelManagementNewParamsResource] `json:"resources,required"`
}

func (r CfdTunnelManagementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Management resources the token will have access to.
type CfdTunnelManagementNewParamsResource string

const (
	CfdTunnelManagementNewParamsResourceLogs CfdTunnelManagementNewParamsResource = "logs"
)

type CfdTunnelManagementNewResponseEnvelope struct {
	Errors   []CfdTunnelManagementNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CfdTunnelManagementNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CfdTunnelManagementNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CfdTunnelManagementNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfdTunnelManagementNewResponseEnvelopeJSON    `json:"-"`
}

// cfdTunnelManagementNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CfdTunnelManagementNewResponseEnvelope]
type cfdTunnelManagementNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelManagementNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelManagementNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    cfdTunnelManagementNewResponseEnvelopeErrorsJSON `json:"-"`
}

// cfdTunnelManagementNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CfdTunnelManagementNewResponseEnvelopeErrors]
type cfdTunnelManagementNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelManagementNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CfdTunnelManagementNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    cfdTunnelManagementNewResponseEnvelopeMessagesJSON `json:"-"`
}

// cfdTunnelManagementNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CfdTunnelManagementNewResponseEnvelopeMessages]
type cfdTunnelManagementNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfdTunnelManagementNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CfdTunnelManagementNewResponseEnvelopeSuccess bool

const (
	CfdTunnelManagementNewResponseEnvelopeSuccessTrue CfdTunnelManagementNewResponseEnvelopeSuccess = true
)
