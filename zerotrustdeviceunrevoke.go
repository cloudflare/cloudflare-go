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

// ZeroTrustDeviceUnrevokeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustDeviceUnrevokeService] method instead.
type ZeroTrustDeviceUnrevokeService struct {
	Options []option.RequestOption
}

// NewZeroTrustDeviceUnrevokeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDeviceUnrevokeService(opts ...option.RequestOption) (r *ZeroTrustDeviceUnrevokeService) {
	r = &ZeroTrustDeviceUnrevokeService{}
	r.Options = opts
	return
}

// Unrevokes a list of devices.
func (r *ZeroTrustDeviceUnrevokeService) New(ctx context.Context, params ZeroTrustDeviceUnrevokeNewParams, opts ...option.RequestOption) (res *ZeroTrustDeviceUnrevokeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceUnrevokeNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/unrevoke", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ZeroTrustDeviceUnrevokeNewResponseUnknown] or
// [shared.UnionString].
type ZeroTrustDeviceUnrevokeNewResponse interface {
	ImplementsZeroTrustDeviceUnrevokeNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustDeviceUnrevokeNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustDeviceUnrevokeNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// A list of device ids to unrevoke.
	Body param.Field[[]string] `json:"body,required"`
}

func (r ZeroTrustDeviceUnrevokeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZeroTrustDeviceUnrevokeNewResponseEnvelope struct {
	Errors   []ZeroTrustDeviceUnrevokeNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceUnrevokeNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDeviceUnrevokeNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceUnrevokeNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceUnrevokeNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceUnrevokeNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceUnrevokeNewResponseEnvelope]
type zeroTrustDeviceUnrevokeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceUnrevokeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceUnrevokeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceUnrevokeNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustDeviceUnrevokeNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceUnrevokeNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceUnrevokeNewResponseEnvelopeErrors]
type zeroTrustDeviceUnrevokeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceUnrevokeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceUnrevokeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceUnrevokeNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDeviceUnrevokeNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceUnrevokeNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceUnrevokeNewResponseEnvelopeMessages]
type zeroTrustDeviceUnrevokeNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceUnrevokeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceUnrevokeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceUnrevokeNewResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceUnrevokeNewResponseEnvelopeSuccessTrue ZeroTrustDeviceUnrevokeNewResponseEnvelopeSuccess = true
)
