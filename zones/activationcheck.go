// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *ActivationCheckService) Trigger(ctx context.Context, body ActivationCheckTriggerParams, opts ...option.RequestOption) (res *ActivationCheckTriggerResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ActivationCheckTriggerResponseEnvelope
	path := fmt.Sprintf("zones/%s/activation_check", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ActivationCheckTriggerResponse struct {
	// Identifier
	ID   string                             `json:"id"`
	JSON activationCheckTriggerResponseJSON `json:"-"`
}

// activationCheckTriggerResponseJSON contains the JSON metadata for the struct
// [ActivationCheckTriggerResponse]
type activationCheckTriggerResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckTriggerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activationCheckTriggerResponseJSON) RawJSON() string {
	return r.raw
}

type ActivationCheckTriggerParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ActivationCheckTriggerResponseEnvelope struct {
	Errors   []ActivationCheckTriggerResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ActivationCheckTriggerResponseEnvelopeMessages `json:"messages,required"`
	Result   ActivationCheckTriggerResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ActivationCheckTriggerResponseEnvelopeSuccess `json:"success,required"`
	JSON    activationCheckTriggerResponseEnvelopeJSON    `json:"-"`
}

// activationCheckTriggerResponseEnvelopeJSON contains the JSON metadata for the
// struct [ActivationCheckTriggerResponseEnvelope]
type activationCheckTriggerResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckTriggerResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activationCheckTriggerResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ActivationCheckTriggerResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    activationCheckTriggerResponseEnvelopeErrorsJSON `json:"-"`
}

// activationCheckTriggerResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ActivationCheckTriggerResponseEnvelopeErrors]
type activationCheckTriggerResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckTriggerResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activationCheckTriggerResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ActivationCheckTriggerResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    activationCheckTriggerResponseEnvelopeMessagesJSON `json:"-"`
}

// activationCheckTriggerResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ActivationCheckTriggerResponseEnvelopeMessages]
type activationCheckTriggerResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckTriggerResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activationCheckTriggerResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ActivationCheckTriggerResponseEnvelopeSuccess bool

const (
	ActivationCheckTriggerResponseEnvelopeSuccessTrue ActivationCheckTriggerResponseEnvelopeSuccess = true
)

func (r ActivationCheckTriggerResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ActivationCheckTriggerResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
