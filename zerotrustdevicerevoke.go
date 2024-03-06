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

// ZeroTrustDeviceRevokeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDeviceRevokeService]
// method instead.
type ZeroTrustDeviceRevokeService struct {
	Options []option.RequestOption
}

// NewZeroTrustDeviceRevokeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDeviceRevokeService(opts ...option.RequestOption) (r *ZeroTrustDeviceRevokeService) {
	r = &ZeroTrustDeviceRevokeService{}
	r.Options = opts
	return
}

// Revokes a list of devices.
func (r *ZeroTrustDeviceRevokeService) New(ctx context.Context, params ZeroTrustDeviceRevokeNewParams, opts ...option.RequestOption) (res *ZeroTrustDeviceRevokeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceRevokeNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/revoke", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ZeroTrustDeviceRevokeNewResponseUnknown] or
// [shared.UnionString].
type ZeroTrustDeviceRevokeNewResponse interface {
	ImplementsZeroTrustDeviceRevokeNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustDeviceRevokeNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustDeviceRevokeNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// A list of device ids to revoke.
	Body param.Field[[]string] `json:"body,required"`
}

func (r ZeroTrustDeviceRevokeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZeroTrustDeviceRevokeNewResponseEnvelope struct {
	Errors   []ZeroTrustDeviceRevokeNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceRevokeNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDeviceRevokeNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceRevokeNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceRevokeNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceRevokeNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceRevokeNewResponseEnvelope]
type zeroTrustDeviceRevokeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceRevokeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceRevokeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceRevokeNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustDeviceRevokeNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceRevokeNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceRevokeNewResponseEnvelopeErrors]
type zeroTrustDeviceRevokeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceRevokeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceRevokeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceRevokeNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustDeviceRevokeNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceRevokeNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceRevokeNewResponseEnvelopeMessages]
type zeroTrustDeviceRevokeNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceRevokeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceRevokeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceRevokeNewResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceRevokeNewResponseEnvelopeSuccessTrue ZeroTrustDeviceRevokeNewResponseEnvelopeSuccess = true
)
