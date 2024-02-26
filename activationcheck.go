// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ActivationCheckService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewActivationCheckService] method
// instead.
type ActivationCheckService struct {
	Options []option.RequestOption
}

// NewActivationCheckService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActivationCheckService(opts ...option.RequestOption) (r *ActivationCheckService) {
	r = &ActivationCheckService{}
	r.Options = opts
	return
}

// Triggeres a new activation check for a PENDING Zone. This can be triggered every
// 5 min for paygo/ent customers, every hour for FREE Zones.
func (r *ActivationCheckService) Update(ctx context.Context, body ActivationCheckUpdateParams, opts ...option.RequestOption) (res *ActivationCheckUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ActivationCheckUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/activation_check", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ActivationCheckUpdateResponse struct {
	// Identifier
	ID   string                            `json:"id"`
	JSON activationCheckUpdateResponseJSON `json:"-"`
}

// activationCheckUpdateResponseJSON contains the JSON metadata for the struct
// [ActivationCheckUpdateResponse]
type activationCheckUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ActivationCheckUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ActivationCheckUpdateResponseEnvelope struct {
	Errors   []ActivationCheckUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ActivationCheckUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ActivationCheckUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ActivationCheckUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    activationCheckUpdateResponseEnvelopeJSON    `json:"-"`
}

// activationCheckUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ActivationCheckUpdateResponseEnvelope]
type activationCheckUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ActivationCheckUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    activationCheckUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// activationCheckUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ActivationCheckUpdateResponseEnvelopeErrors]
type activationCheckUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ActivationCheckUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    activationCheckUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// activationCheckUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ActivationCheckUpdateResponseEnvelopeMessages]
type activationCheckUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ActivationCheckUpdateResponseEnvelopeSuccess bool

const (
	ActivationCheckUpdateResponseEnvelopeSuccessTrue ActivationCheckUpdateResponseEnvelopeSuccess = true
)
