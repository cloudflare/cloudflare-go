// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZoneActivationCheckService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneActivationCheckService]
// method instead.
type ZoneActivationCheckService struct {
	Options []option.RequestOption
}

// NewZoneActivationCheckService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneActivationCheckService(opts ...option.RequestOption) (r *ZoneActivationCheckService) {
	r = &ZoneActivationCheckService{}
	r.Options = opts
	return
}

// Triggeres a new activation check for a PENDING Zone. This can be triggered every
// 5 min for paygo/ent customers, every hour for FREE Zones.
func (r *ZoneActivationCheckService) Trigger(ctx context.Context, body ZoneActivationCheckTriggerParams, opts ...option.RequestOption) (res *ZoneActivationCheckTriggerResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneActivationCheckTriggerResponseEnvelope
	path := fmt.Sprintf("zones/%s/activation_check", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneActivationCheckTriggerResponse struct {
	// Identifier
	ID   string                                 `json:"id"`
	JSON zoneActivationCheckTriggerResponseJSON `json:"-"`
}

// zoneActivationCheckTriggerResponseJSON contains the JSON metadata for the struct
// [ZoneActivationCheckTriggerResponse]
type zoneActivationCheckTriggerResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActivationCheckTriggerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActivationCheckTriggerResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneActivationCheckTriggerParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneActivationCheckTriggerResponseEnvelope struct {
	Errors   []ZoneActivationCheckTriggerResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneActivationCheckTriggerResponseEnvelopeMessages `json:"messages,required"`
	Result   ZoneActivationCheckTriggerResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZoneActivationCheckTriggerResponseEnvelopeSuccess `json:"success,required"`
	JSON    zoneActivationCheckTriggerResponseEnvelopeJSON    `json:"-"`
}

// zoneActivationCheckTriggerResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneActivationCheckTriggerResponseEnvelope]
type zoneActivationCheckTriggerResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActivationCheckTriggerResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActivationCheckTriggerResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneActivationCheckTriggerResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneActivationCheckTriggerResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneActivationCheckTriggerResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneActivationCheckTriggerResponseEnvelopeErrors]
type zoneActivationCheckTriggerResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActivationCheckTriggerResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActivationCheckTriggerResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneActivationCheckTriggerResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneActivationCheckTriggerResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneActivationCheckTriggerResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneActivationCheckTriggerResponseEnvelopeMessages]
type zoneActivationCheckTriggerResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActivationCheckTriggerResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneActivationCheckTriggerResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneActivationCheckTriggerResponseEnvelopeSuccess bool

const (
	ZoneActivationCheckTriggerResponseEnvelopeSuccessTrue ZoneActivationCheckTriggerResponseEnvelopeSuccess = true
)
