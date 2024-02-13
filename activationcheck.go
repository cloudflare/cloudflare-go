// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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
func (r *ActivationCheckService) PutZonesZoneIDActivationCheck(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ActivationCheckPutZonesZoneIDActivationCheckResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelope
	path := fmt.Sprintf("zones/%s/activation_check", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ActivationCheckPutZonesZoneIDActivationCheckResponse struct {
	// Identifier
	ID   string                                                   `json:"id"`
	JSON activationCheckPutZonesZoneIDActivationCheckResponseJSON `json:"-"`
}

// activationCheckPutZonesZoneIDActivationCheckResponseJSON contains the JSON
// metadata for the struct [ActivationCheckPutZonesZoneIDActivationCheckResponse]
type activationCheckPutZonesZoneIDActivationCheckResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckPutZonesZoneIDActivationCheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelope struct {
	Errors   []ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeMessages `json:"messages,required"`
	Result   ActivationCheckPutZonesZoneIDActivationCheckResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeSuccess `json:"success,required"`
	JSON    activationCheckPutZonesZoneIDActivationCheckResponseEnvelopeJSON    `json:"-"`
}

// activationCheckPutZonesZoneIDActivationCheckResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelope]
type activationCheckPutZonesZoneIDActivationCheckResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    activationCheckPutZonesZoneIDActivationCheckResponseEnvelopeErrorsJSON `json:"-"`
}

// activationCheckPutZonesZoneIDActivationCheckResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeErrors]
type activationCheckPutZonesZoneIDActivationCheckResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    activationCheckPutZonesZoneIDActivationCheckResponseEnvelopeMessagesJSON `json:"-"`
}

// activationCheckPutZonesZoneIDActivationCheckResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeMessages]
type activationCheckPutZonesZoneIDActivationCheckResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeSuccess bool

const (
	ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeSuccessTrue ActivationCheckPutZonesZoneIDActivationCheckResponseEnvelopeSuccess = true
)
